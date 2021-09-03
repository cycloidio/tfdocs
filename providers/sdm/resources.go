package sdm

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "sdm_account",
			Category:         "Resources",
			ShortDescription: `Provides settings for Account. layout: “sdm” sidebar_current: “docs-sdm-resource-account"`,
			Description: `

Accounts are users that have access to strongDM. There are two types of accounts:
 1. **Users:** humans who are authenticated through username and password or SSO.
 2. **Service Accounts:** machines that are authenticated using a service token.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The User's email address. Must be unique.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Required) The User's first name.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Required) The User's last name.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `(Optional) The User's suspended state.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Service.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `(Optional) The Service's suspended state. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the Account resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Account resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Account resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_account_attachment",
			Category:         "Resources",
			ShortDescription: `Provides settings for AccountAttachment. layout: “sdm” sidebar_current: “docs-sdm-resource-account-attachment"`,
			Description: `

AccountAttachments assign an account to a role or composite role.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The id of the account of this AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The id of the attached role of this AccountAttachment. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the AccountAttachment resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the AccountAttachment resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the AccountAttachment resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_account_grant",
			Category:         "Resources",
			ShortDescription: `Provides settings for AccountGrant. layout: “sdm” sidebar_current: “docs-sdm-resource-account-grant"`,
			Description: `

AccountGrants connect a resource directly to an account, giving the account the permission to connect to that resource.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The id of the composite role of this AccountGrant.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The id of the attached role of this AccountGrant. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the AccountGrant resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the AccountGrant resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the AccountGrant resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_node",
			Category:         "Resources",
			ShortDescription: `Provides settings for Node. layout: “sdm” sidebar_current: “docs-sdm-resource-node"`,
			Description: `

Nodes make up the strongDM network, and allow your users to connect securely to your resources.
 There are two types of nodes:
 1. **Relay:** creates connectivity to your datasources, while maintaining the egress-only nature of your firewall
 1. **Gateways:** a relay that also listens for connections from strongDM clients
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "gateway_filter",
					Description: `(Optional) GatewayFilter can be used to restrict the peering between relays and gateways.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "listen_address",
					Description: `(Required) The public hostname/port tuple at which the gateway will be accessible to clients.`,
				},
				resource.Attribute{
					Name:        "bind_address",
					Description: `(Optional) The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:<listen_address_port>".`,
				},
				resource.Attribute{
					Name:        "gateway_filter",
					Description: `(Optional) GatewayFilter can be used to restrict the peering between relays and gateways. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the Node resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Node resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Node resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_resource",
			Category:         "Resources",
			ShortDescription: `Provides settings for Resource. layout: “sdm” sidebar_current: “docs-sdm-resource-resource"`,
			Description: `

A Resource is a database or server for which strongDM manages access.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_private_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_private_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "auth_header",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_auth_header_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_auth_header_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_external_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_service_account_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_service_account_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_service_account_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_service_account_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_certificate_authority_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_forwarding",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_forwarding",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_private_key_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_private_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port_forwarding",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the Resource resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Resource resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Resource resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_role",
			Category:         "Resources",
			ShortDescription: `Provides settings for Role. layout: “sdm” sidebar_current: “docs-sdm-resource-role"`,
			Description: `

A Role is a collection of access grants, and typically corresponds to a team, Active Directory OU, or other organizational unit. Users are granted access to resources by assigning them to roles.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Role.`,
				},
				resource.Attribute{
					Name:        "access_rules",
					Description: `(Optional) AccessRules JSON encoded access rules data.`,
				},
				resource.Attribute{
					Name:        "composite",
					Description: `(Optional) True if the Role is a composite role. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the Role resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Role resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Role resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_role_attachment",
			Category:         "Resources",
			ShortDescription: `Provides settings for RoleAttachment. layout: “sdm” sidebar_current: “docs-sdm-resource-role-attachment"`,
			Description: `

A RoleAttachment assigns a role to a composite role.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "composite_role_id",
					Description: `(Required) The id of the composite role of this RoleAttachment.`,
				},
				resource.Attribute{
					Name:        "attached_role_id",
					Description: `(Required) The id of the attached role of this RoleAttachment. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the RoleAttachment resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the RoleAttachment resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the RoleAttachment resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_role_grant",
			Category:         "Resources",
			ShortDescription: `Provides settings for RoleGrant. layout: “sdm” sidebar_current: “docs-sdm-resource-role-grant"`,
			Description: `

A RoleGrant connects a resource to a role, granting members of the role access to that resource.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The id of the resource of this RoleGrant.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The id of the attached role of this RoleGrant. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the RoleGrant resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the RoleGrant resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the RoleGrant resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_secret_store",
			Category:         "Resources",
			ShortDescription: `Provides settings for SecretStore. layout: “sdm” sidebar_current: “docs-sdm-resource-secret-store"`,
			Description: `

A SecretStore is a server where resource secrets (passwords, keys) are stored. 
 Coming soon support for HashiCorp Vault and AWS Secret Store. Contact support@strongdm.com to request access to the beta.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "ca_cert_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_cert_path",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "client_key_path",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Required) ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the SecretStore resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the SecretStore resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the SecretStore resource.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"sdm_account":            0,
		"sdm_account_attachment": 1,
		"sdm_account_grant":      2,
		"sdm_node":               3,
		"sdm_resource":           4,
		"sdm_role":               5,
		"sdm_role_attachment":    6,
		"sdm_role_grant":         7,
		"sdm_secret_store":       8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
