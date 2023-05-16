package sdm

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "sdm_account",
			Category:         "Data Sources",
			ShortDescription: `Query for existing Accounts instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-account"`,
			Description: `

Accounts are users that have access to strongDM. There are two types of accounts:
 1. **Users:** humans who are authenticated through username and password or SSO.
 2. **Service Accounts:** machines that are authenticated using a service token.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://www.strongdm.com/docs/automation/getting-started/filters for more information.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The User's email address. Must be unique.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Optional) External ID is an alternative unique ID this user is represented by within an external service.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Optional) The User's first name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the User.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Optional) The User's last name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the Service.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `(Optional) The User's suspended state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a Accounts data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_account ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `A single element list containing a map, where each key lists one of the following objects:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Service.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `The Service's suspended state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The User's email address. Must be unique.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External ID is an alternative unique ID this user is represented by within an external service.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The User's first name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the User.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The User's last name.`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `Managed By is a read only field for what service manages this user, e.g. StrongDM, Okta, Azure.`,
				},
				resource.Attribute{
					Name:        "permission_level",
					Description: `PermissionLevel is a read only field for the user's permission level e.g. admin, DBA, user.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `The User's suspended state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_account ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `A single element list containing a map, where each key lists one of the following objects:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Service.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `The Service's suspended state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The User's email address. Must be unique.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External ID is an alternative unique ID this user is represented by within an external service.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The User's first name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the User.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The User's last name.`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `Managed By is a read only field for what service manages this user, e.g. StrongDM, Okta, Azure.`,
				},
				resource.Attribute{
					Name:        "permission_level",
					Description: `PermissionLevel is a read only field for the user's permission level e.g. admin, DBA, user.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `The User's suspended state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_account_attachment",
			Category:         "Data Sources",
			ShortDescription: `Query for existing AccountAttachments instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-account-attachment"`,
			Description: `

AccountAttachments assign an account to a role.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The id of the account of this AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Optional) The id of the attached role of this AccountAttachment. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a AccountAttachments data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_account_attachment ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "account_attachments",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id of the account of this AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The id of the attached role of this AccountAttachment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_account_attachment ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "account_attachments",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id of the account of this AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The id of the attached role of this AccountAttachment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_node",
			Category:         "Data Sources",
			ShortDescription: `Query for existing Nodes instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-node"`,
			Description: `

Nodes make up the strongDM network, and allow your users to connect securely to your resources.
 There are two types of nodes:
 1. **Relay:** creates connectivity to your datasources, while maintaining the egress-only nature of your firewall
 1. **Gateways:** a relay that also listens for connections from strongDM clients
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://www.strongdm.com/docs/automation/getting-started/filters for more information.`,
				},
				resource.Attribute{
					Name:        "bind_address",
					Description: `(Optional) The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:listen_address_port".`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the Relay.`,
				},
				resource.Attribute{
					Name:        "listen_address",
					Description: `(Optional) The public hostname/port tuple at which the gateway will be accessible to clients.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a Nodes data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_node ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A single element list containing a map, where each key lists one of the following objects:`,
				},
				resource.Attribute{
					Name:        "bind_address",
					Description: `The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:listen_address_port".`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device is a read only device name uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "gateway_filter",
					Description: `GatewayFilter can be used to restrict the peering between relays and gateways.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Gateway.`,
				},
				resource.Attribute{
					Name:        "listen_address",
					Description: `The public hostname/port tuple at which the gateway will be accessible to clients.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location is a read only network location uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
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
					Name:        "gateway_filter",
					Description: `GatewayFilter can be used to restrict the peering between relays and gateways.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Relay.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location is a read only network location uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version is a read only sdm binary version uploaded by the gateway process when it comes online.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_node ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A single element list containing a map, where each key lists one of the following objects:`,
				},
				resource.Attribute{
					Name:        "bind_address",
					Description: `The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:listen_address_port".`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device is a read only device name uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "gateway_filter",
					Description: `GatewayFilter can be used to restrict the peering between relays and gateways.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Gateway.`,
				},
				resource.Attribute{
					Name:        "listen_address",
					Description: `The public hostname/port tuple at which the gateway will be accessible to clients.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location is a read only network location uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
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
					Name:        "gateway_filter",
					Description: `GatewayFilter can be used to restrict the peering between relays and gateways.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Relay.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location is a read only network location uploaded by the gateway process when it comes online.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version is a read only sdm binary version uploaded by the gateway process when it comes online.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_remote_identity",
			Category:         "Data Sources",
			ShortDescription: `Query for existing RemoteIdentities instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-remote-identity"`,
			Description: `

RemoteIdentities define the username to be used for a specific account
 when connecting to a remote resource using that group.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account for this remote identity.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the RemoteIdentity.`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `(Optional) The remote identity group.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username to be used as the remote identity for this account. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a RemoteIdentities data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_remote_identity ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "remote_identities",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account for this remote identity.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the RemoteIdentity.`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `The remote identity group.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username to be used as the remote identity for this account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_remote_identity ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "remote_identities",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account for this remote identity.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the RemoteIdentity.`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: `The remote identity group.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username to be used as the remote identity for this account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_remote_identity_group",
			Category:         "Data Sources",
			ShortDescription: `Query for existing RemoteIdentityGroups instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-remote-identity-group"`,
			Description: `

A RemoteIdentityGroup defines a group of remote identities.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the RemoteIdentityGroup.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the RemoteIdentityGroup. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a RemoteIdentityGroups data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_remote_identity_group ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "remote_identity_groups",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the RemoteIdentityGroup.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the RemoteIdentityGroup.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_remote_identity_group ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "remote_identity_groups",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the RemoteIdentityGroup.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the RemoteIdentityGroup.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_resource",
			Category:         "Data Sources",
			ShortDescription: `Query for existing Resources instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-resource"`,
			Description: `

A Resource is a database, server, cluster, website, or cloud that strongDM
 delegates access to.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://www.strongdm.com/docs/automation/getting-started/filters for more information.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a Resources data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_resource ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `A single element list containing a map, where each key lists one of the following objects:`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "enable_env_variables",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "session_expiry",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "session_expiry",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "project",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname must contain the hostname/port pairs of all instances in the replica set separated by commas.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "keyfile",
					Description: ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_header",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "host_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "host_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "host_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "downgrade_nla_connections",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "schema",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "saml_metadata",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "schema",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_forwarding",
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
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_forwarding",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_forwarding",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_resource ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `A single element list containing a map, where each key lists one of the following objects:`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "enable_env_variables",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "session_expiry",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "session_expiry",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "project",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname must contain the hostname/port pairs of all instances in the replica set separated by commas.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "keyfile",
					Description: ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_header",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "host_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "host_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "host_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
					Description: `The path used to check the health of your connection. Defaults to ` + "`" + `default` + "`" + `. This field is required, and is only marked as optional for backwards compatibility.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: ``,
				},
				resource.Attribute{
					Name:        "role_external_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "downgrade_nla_connections",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tls_required",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "schema",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "saml_metadata",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "override_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "schema",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_forwarding",
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
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_forwarding",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_group_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "remote_identity_healthcheck_username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_forwarding",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "bind_interface",
					Description: `Bind interface`,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "egress_filter",
					Description: `A filter applied to the routing logic to pin datasource to nodes.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `Subdomain is the local DNS address. (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_role",
			Category:         "Data Sources",
			ShortDescription: `Query for existing Roles instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-role"`,
			Description: `

A Role has a list of access rules which determine which Resources the members
 of the Role have access to. An Account can be a member of multiple Roles via
 AccountAttachments.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the Role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the Role.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a Roles data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_role ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_rules",
					Description: `AccessRules is a list of access rules defining the resources this Role has access to.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Role.`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `Managed By is a read only field for what service manages this role, e.g. StrongDM, Okta, Azure.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Role.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_role ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_rules",
					Description: `AccessRules is a list of access rules defining the resources this Role has access to.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the Role.`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `Managed By is a read only field for what service manages this role, e.g. StrongDM, Okta, Azure.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Role.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_secret_store",
			Category:         "Data Sources",
			ShortDescription: `Query for existing SecretStores instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-secret-store"`,
			Description: `

A SecretStore is a server where resource secrets (passwords, keys) are stored.
 Coming soon support for HashiCorp Vault and AWS Secret Store.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://www.strongdm.com/docs/automation/getting-started/filters for more information.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags is a map of key, value pairs. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a SecretStores data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_secret_store ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "secret_stores",
					Description: `A single element list containing a map, where each key lists one of the following objects:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "vault_uri",
					Description: ``,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "server_url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: ``,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "ca_cert_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_cert_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: ``,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: ``,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_secret_store ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "secret_stores",
					Description: `A single element list containing a map, where each key lists one of the following objects:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "vault_uri",
					Description: ``,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "server_url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: ``,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "ca_cert_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_cert_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "client_key_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: ``,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the SecretStore.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: ``,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_ssh_ca_pubkey",
			Category:         "Data Sources",
			ShortDescription: `Query for the SSH Certificate Authority Public Key. layout: “sdm” sidebar_current: “docs-sdm-datasource-ssh-ca-pub-key"`,
			Description: `

The SSH CA Pubkey is a public key used for setting up SSH resources.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `the SSH Certificate Authority public key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `the SSH Certificate Authority public key.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"sdm_account":               0,
		"sdm_account_attachment":    1,
		"sdm_node":                  2,
		"sdm_remote_identity":       3,
		"sdm_remote_identity_group": 4,
		"sdm_resource":              5,
		"sdm_role":                  6,
		"sdm_secret_store":          7,
		"sdm_ssh_ca_pubkey":         8,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
