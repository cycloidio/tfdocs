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
					Description: `(Optional) a filter to query only one subtype. See Attribute Reference for all subtypes.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The User's email address. Must be unique.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Optional) The User's first name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the Service.`,
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
					Description: `(Optional) The Service's suspended state. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a Accounts data source:`,
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
					Description: `Unique identifier of the User.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The User's email address. Must be unique.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The User's first name.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The User's last name.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `The User's suspended state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
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
					Description: `Unique identifier of the User.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The User's email address. Must be unique.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The User's first name.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The User's last name.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `The User's suspended state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_account_attachment",
			Category:         "Data Sources",
			ShortDescription: `Query for existing AccountAttachments instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-account-attachment"`,
			Description: `

AccountAttachments assign an account to a role or composite role.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The id of the account of this AccountAttachment.`,
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
					Name:        "id",
					Description: `Unique identifier of the AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id of the account of this AccountAttachment.`,
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
					Name:        "id",
					Description: `Unique identifier of the AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id of the account of this AccountAttachment.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The id of the attached role of this AccountAttachment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_account_grant",
			Category:         "Data Sources",
			ShortDescription: `Query for existing AccountGrants instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-account-grant"`,
			Description: `

AccountGrants connect a resource directly to an account, giving the account the permission to connect to that resource.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the AccountGrant.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) The id of the composite role of this AccountGrant.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The id of the attached role of this AccountGrant. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a AccountGrants data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_account_grant ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "account_grants",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the AccountGrant.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The id of the composite role of this AccountGrant.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id of the attached role of this AccountGrant.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_account_grant ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "account_grants",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the AccountGrant.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The id of the composite role of this AccountGrant.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id of the attached role of this AccountGrant.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_index",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Optional) a filter to query only one subtype. See Attribute Reference for all subtypes.`,
				},
				resource.Attribute{
					Name:        "bind_address",
					Description: `(Optional) The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:<listen_address_port>".`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the Gateway.`,
				},
				resource.Attribute{
					Name:        "listen_address",
					Description: `(Optional) The public hostname/port tuple at which the gateway will be accessible to clients.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a Nodes data source:`,
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
					Name:        "id",
					Description: `Unique identifier of the Relay.`,
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
					Name:        "id",
					Description: `Unique identifier of the Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "listen_address",
					Description: `The public hostname/port tuple at which the gateway will be accessible to clients.`,
				},
				resource.Attribute{
					Name:        "bind_address",
					Description: `The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:<listen_address_port>".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
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
					Name:        "id",
					Description: `Unique identifier of the Relay.`,
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
					Name:        "id",
					Description: `Unique identifier of the Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.`,
				},
				resource.Attribute{
					Name:        "listen_address",
					Description: `The public hostname/port tuple at which the gateway will be accessible to clients.`,
				},
				resource.Attribute{
					Name:        "bind_address",
					Description: `The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:<listen_address_port>".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_resource",
			Category:         "Data Sources",
			ShortDescription: `Query for existing Resources instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-resource"`,
			Description: `

A Resource is a database or server for which strongDM manages access.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) a filter to query only one subtype. See Attribute Reference for all subtypes.`,
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_region",
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
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
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_header",
					Description: ``,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
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
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
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
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
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
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
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
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
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
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "schema",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
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
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_forwarding",
					Description: ``,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "allow_deprecated_key_exchanges",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_region",
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
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
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_header",
					Description: ``,
				},
				resource.Attribute{
					Name:        "headers_blacklist",
					Description: ``,
				},
				resource.Attribute{
					Name:        "default_path",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subdomain",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
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
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
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
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: ``,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: ``,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
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
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "token",
					Description: ``,
				},
				resource.Attribute{
					Name:        "healthcheck_namespace",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auth_database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
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
					Name:        "replica_set",
					Description: ``,
				},
				resource.Attribute{
					Name:        "connect_to_replica",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
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
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
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
					Name:        "id",
					Description: `Unique identifier of the Resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
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
					Name:        "tls_required",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
					Description: ``,
				},
				resource.Attribute{
					Name:        "schema",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
					Description: ``,
				},
				resource.Attribute{
					Name:        "database",
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
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "override_database",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_forwarding",
					Description: ``,
				},
				resource.Attribute{
					Name:        "allow_deprecated_key_exchanges",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "allow_deprecated_key_exchanges",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port_override",
					Description: ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: ``,
				},
				resource.Attribute{
					Name:        "password",
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
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
				resource.Attribute{
					Name:        "secret_store_id",
					Description: `ID of the secret store containing credentials for this resource, if any.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: ``,
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
					Name:        "port",
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

A Role is a collection of access grants, and typically corresponds to a team, Active Directory OU, or other organizational unit. Users are granted access to resources by assigning them to roles.
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
					Name:        "composite",
					Description: `(Optional) True if the Role is a composite role. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a Roles data source:`,
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
					Name:        "id",
					Description: `Unique identifier of the Role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Role.`,
				},
				resource.Attribute{
					Name:        "composite",
					Description: `True if the Role is a composite role.`,
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
					Name:        "id",
					Description: `Unique identifier of the Role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique human-readable name of the Role.`,
				},
				resource.Attribute{
					Name:        "composite",
					Description: `True if the Role is a composite role.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags is a map of key, value pairs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_role_attachment",
			Category:         "Data Sources",
			ShortDescription: `Query for existing RoleAttachments instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-role-attachment"`,
			Description: `

A RoleAttachment assigns a role to a composite role.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the RoleAttachment.`,
				},
				resource.Attribute{
					Name:        "composite_role_id",
					Description: `(Optional) The id of the composite role of this RoleAttachment.`,
				},
				resource.Attribute{
					Name:        "attached_role_id",
					Description: `(Optional) The id of the attached role of this RoleAttachment. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a RoleAttachments data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_role_attachment ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "role_attachments",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the RoleAttachment.`,
				},
				resource.Attribute{
					Name:        "composite_role_id",
					Description: `The id of the composite role of this RoleAttachment.`,
				},
				resource.Attribute{
					Name:        "attached_role_id",
					Description: `The id of the attached role of this RoleAttachment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_role_attachment ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "role_attachments",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the RoleAttachment.`,
				},
				resource.Attribute{
					Name:        "composite_role_id",
					Description: `The id of the composite role of this RoleAttachment.`,
				},
				resource.Attribute{
					Name:        "attached_role_id",
					Description: `The id of the attached role of this RoleAttachment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdm_role_grant",
			Category:         "Data Sources",
			ShortDescription: `Query for existing RoleGrants instances. layout: “sdm” sidebar_current: “docs-sdm-datasource-role-grant"`,
			Description: `

A RoleGrant connects a resource to a role, granting members of the role access to that resource.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier of the RoleGrant.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) The id of the resource of this RoleGrant.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Optional) The id of the attached role of this RoleGrant. ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a RoleGrants data source:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_role_grant ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "role_grants",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the RoleGrant.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The id of the resource of this RoleGrant.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The id of the attached role of this RoleGrant.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `a generated id representing this request, unrelated to input id and sdm_role_grant ids.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `a list of strings of ids of data sources that match the given arguments.`,
				},
				resource.Attribute{
					Name:        "role_grants",
					Description: `A list where each element has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the RoleGrant.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The id of the resource of this RoleGrant.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The id of the attached role of this RoleGrant.`,
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
 Coming soon support for HashiCorp Vault and AWS Secret Store. Contact support@strongdm.com to request access to the beta.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) a filter to query only one subtype. See Attribute Reference for all subtypes.`,
				},
				resource.Attribute{
					Name:        "ca_cert_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_cert_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_key_path",
					Description: `(Optional)`,
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
					Name:        "region",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Optional) ## Attribute Reference In addition to provided arguments above, the following attributes are returned by a SecretStores data source:`,
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
					Name:        "server_address",
					Description: ``,
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
					Name:        "server_address",
					Description: ``,
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

		"sdm_account":            0,
		"sdm_account_attachment": 1,
		"sdm_account_grant":      2,
		"sdm_index":              3,
		"sdm_node":               4,
		"sdm_resource":           5,
		"sdm_role":               6,
		"sdm_role_attachment":    7,
		"sdm_role_grant":         8,
		"sdm_secret_store":       9,
		"sdm_ssh_ca_pubkey":      10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
