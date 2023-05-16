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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Service.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `(Optional) The Service's suspended state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The User's email address. Must be unique.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Optional) External ID is an alternative unique ID this user is represented by within an external service.`,
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
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the Account resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Account resource.`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `Managed By is a read only field for what service manages this user, e.g. StrongDM, Okta, Azure.`,
				},
				resource.Attribute{
					Name:        "permission_level",
					Description: `PermissionLevel is a read only field for the user's permission level e.g. admin, DBA, user. ## Import Account can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_account.example a-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Account resource.`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `Managed By is a read only field for what service manages this user, e.g. StrongDM, Okta, Azure.`,
				},
				resource.Attribute{
					Name:        "permission_level",
					Description: `PermissionLevel is a read only field for the user's permission level e.g. admin, DBA, user. ## Import Account can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_account.example a-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_account_attachment",
			Category:         "Resources",
			ShortDescription: `Provides settings for AccountAttachment. layout: “sdm” sidebar_current: “docs-sdm-resource-account-attachment"`,
			Description: `

AccountAttachments assign an account to a role.
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
					Description: `A unique identifier for the AccountAttachment resource. ## Import AccountAttachment can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_account_attachment.example aa-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the AccountAttachment resource. ## Import AccountAttachment can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_account_attachment.example aa-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "bind_address",
					Description: `(Optional) The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:listen_address_port".`,
				},
				resource.Attribute{
					Name:        "gateway_filter",
					Description: `(Optional) GatewayFilter can be used to restrict the peering between relays and gateways.`,
				},
				resource.Attribute{
					Name:        "listen_address",
					Description: `(Required) The public hostname/port tuple at which the gateway will be accessible to clients.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "gateway_filter",
					Description: `(Optional) GatewayFilter can be used to restrict the peering between relays and gateways.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the Node resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Node resource.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device is a read only device name uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location is a read only network location uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version is a read only sdm binary version uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device is a read only device name uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location is a read only network location uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version is a read only sdm binary version uploaded by the gateway process when it comes online. ## Import Node can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_node.example n-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Node resource.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device is a read only device name uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location is a read only network location uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version is a read only sdm binary version uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device is a read only device name uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location is a read only network location uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version is a read only sdm binary version uploaded by the gateway process when it comes online. ## Import Node can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_node.example n-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_remote_identity",
			Category:         "Resources",
			ShortDescription: `Provides settings for RemoteIdentity. layout: “sdm” sidebar_current: “docs-sdm-resource-remote-identity"`,
			Description: `

RemoteIdentities define the username to be used for a specific account
 when connecting to a remote resource using that group.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The account for this remote identity.`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Required) The remote identity group.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to be used as the remote identity for this account. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the RemoteIdentity resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the RemoteIdentity resource. ## Import RemoteIdentity can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_remote_identity.example i-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the RemoteIdentity resource. ## Import RemoteIdentity can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_remote_identity.example i-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_resource",
			Category:         "Resources",
			ShortDescription: `Provides settings for Resource. layout: “sdm” sidebar_current: “docs-sdm-resource-resource"`,
			Description: `

A Resource is a database, server, cluster, website, or cloud that strongDM
 delegates access to.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "cluster_name",
					Description: `(Required)`,
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
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
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
					Name:        "secret_access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "cluster_name",
					Description: `(Required)`,
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
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
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
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "cluster_name",
					Description: `(Required)`,
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
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
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
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "cluster_name",
					Description: `(Required)`,
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
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
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
					Name:        "secret_access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
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
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "region",
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
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
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
					Name:        "secret_access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
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
					Name:        "secret_access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "enable_env_variables",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "session_expiry",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_role_arn_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "secret_access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "session_expiry",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_app_id_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_app_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_tenant_id_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_tenant_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_app_id_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_app_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_tenant_id_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_tenant_id_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_private_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
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
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Hostname must contain the hostname/port pairs of all instances in the replica set separated by commas.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
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
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
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
					Name:        "access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "region",
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
					Name:        "secret_access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "keyfile",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_keyfile_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_keyfile_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_service_account_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_service_account_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
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
					Name:        "service_account_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_service_account_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_service_account_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "auth_header",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_auth_header_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_auth_header_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "host_override",
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
					Name:        "subdomain",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "host_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "url",
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
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "host_override",
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
					Name:        "subdomain",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_token_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: `(Optional)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: `(Optional)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "auth_database",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_certificate_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_client_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "region",
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
					Name:        "secret_access_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_secret_access_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "downgrade_nla_connections",
					Description: `(Optional)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `(Optional) A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_username",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "saml_metadata",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_saml_metadata_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_saml_metadata_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "key_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "key_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_private_key_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_private_key_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_password_path",
					Description: `(Required if using credentials stored in a secret store)`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if storing credentials directly strongDM)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_path",
					Description: `(Required if using credentials stored in a secret store)`,
				},
				resource.Attribute{
					Name:        "secret_store_username_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `(Optional) Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required)`,
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
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Resource.`,
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
					Name:        "port_override",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `(Optional) ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the Resource resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Resource resource.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `## Import Resource can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_resource.example rs-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Resource resource.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `## Import Resource can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_resource.example rs-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_role",
			Category:         "Resources",
			ShortDescription: `Provides settings for Role. layout: “sdm” sidebar_current: “docs-sdm-resource-role"`,
			Description: `

A Role has a list of access rules which determine which Resources the members
 of the Role have access to. An Account can be a member of multiple Roles via
 AccountAttachments.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_rules",
					Description: `(Optional) AccessRules is a list of access rules defining the resources this Role has access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the Role. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the Role resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Role resource. ## Import Role can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_role.example r-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the Role resource. ## Import Role can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_role.example r-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
 Coming soon support for HashiCorp Vault and AWS Secret Store.
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.

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
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "vault_uri",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "server_url",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
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
					Name:        "namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by the SecretStore resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the SecretStore resource. ## Import SecretStore can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_secret_store.example se-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the SecretStore resource. ## Import SecretStore can be imported using the id, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import sdm_secret_store.example se-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"sdm_account":            0,
		"sdm_account_attachment": 1,
		"sdm_node":               2,
		"sdm_remote_identity":    3,
		"sdm_resource":           4,
		"sdm_role":               5,
		"sdm_secret_store":       6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
