package azuread

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azuread_access_package",
			Category:         "Identity Governance",
			ShortDescription: ``,
			Description: `

Manages an Access Package within Identity Governance in Azure Active Directory.

`,
			Keywords: []string{
				"identity",
				"governance",
				"access",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_id",
					Description: `(Required) The ID of the Catalog this access package will be created in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the access package.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name of the access package.`,
				},
				resource.Attribute{
					Name:        "hidden",
					Description: `(Optional) Whether the access package is hidden from the requestor. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource. ## Import Access Packages can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azuread_access_package.example_package 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource. ## Import Access Packages can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azuread_access_package.example_package 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_access_package_assignment_policy",
			Category:         "Identity Governance",
			ShortDescription: ``,
			Description: `

Manages an assignment policy for an access package within Identity Governance in Azure Active Directory.

`,
			Keywords: []string{
				"identity",
				"governance",
				"access",
				"package",
				"assignment",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_access_package_catalog",
			Category:         "Identity Governance",
			ShortDescription: ``,
			Description: `

Manages an access package catalog within Identity Governance in Azure Active Directory.

`,
			Keywords: []string{
				"identity",
				"governance",
				"access",
				"package",
				"catalog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the access package catalog.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name of the access package catalog.`,
				},
				resource.Attribute{
					Name:        "externally_visible",
					Description: `(Optional) Whether the access packages in this catalog can be requested by users outside the tenant.`,
				},
				resource.Attribute{
					Name:        "published",
					Description: `(Optional) Whether the access packages in this catalog are available for management. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource. ## Import An Access Package Catalog can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azuread_access_package_catalog.example 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource. ## Import An Access Package Catalog can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azuread_access_package_catalog.example 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_access_package_catalog_role_assignment",
			Category:         "Identity Governance",
			ShortDescription: ``,
			Description: `

Manages a single catalog role assignment within Azure Active Directory.

`,
			Keywords: []string{
				"identity",
				"governance",
				"access",
				"package",
				"catalog",
				"role",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_id",
					Description: `(Required) The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "principal_object_id",
					Description: `(Required) The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The object ID of the catalog role you want to assign. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_access_package_resource_catalog_association",
			Category:         "Identity Governance",
			ShortDescription: ``,
			Description: `

Manages the resources added to access package catalogs within Identity Governance in Azure Active Directory.

`,
			Keywords: []string{
				"identity",
				"governance",
				"access",
				"package",
				"resource",
				"catalog",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_id",
					Description: `(Required) The unique ID of the access package catalog. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_origin_id",
					Description: `(Required) The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_origin_system",
					Description: `(Required) The type of the resource in the origin system, such as ` + "`" + `SharePointOnline` + "`" + `, ` + "`" + `AadApplication` + "`" + ` or ` + "`" + `AadGroup` + "`" + `. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource, the ID is the concatenation of ` + "`" + `catalog_id` + "`" + ` and ` + "`" + `resource_origin_id` + "`" + ` with colon in between. ## Import The resource and catalog association can be imported using the catalog ID and the resource origin ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azuread_access_package_resource_catalog_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the Catalog ID and the Resource Origin ID in the format ` + "`" + `{CatalogID}/{ResourceOriginID}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource, the ID is the concatenation of ` + "`" + `catalog_id` + "`" + ` and ` + "`" + `resource_origin_id` + "`" + ` with colon in between. ## Import The resource and catalog association can be imported using the catalog ID and the resource origin ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azuread_access_package_resource_catalog_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the Catalog ID and the Resource Origin ID in the format ` + "`" + `{CatalogID}/{ResourceOriginID}` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_access_package_resource_package_association",
			Category:         "Identity Governance",
			ShortDescription: ``,
			Description: `

Manages the resources added to access packages within Identity Governance in Azure Active Directory.

`,
			Keywords: []string{
				"identity",
				"governance",
				"access",
				"package",
				"resource",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_package_id",
					Description: `(Required) The ID of access package this resource association is configured to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) The role of access type to the specified resource. Valid values are ` + "`" + `Member` + "`" + `, or ` + "`" + `Owner` + "`" + ` The default is ` + "`" + `Member` + "`" + `. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "catalog_resource_association_id",
					Description: `(Required) The ID of the catalog association from the ` + "`" + `azuread_access_package_resource_catalog_association` + "`" + ` resource. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource. The ID is combined by four fields with colon in between, the four fields are ` + "`" + `access_package_id` + "`" + `, this package association id, ` + "`" + `resource_origin_id` + "`" + ` and ` + "`" + `access_type` + "`" + `. ## Import The resource and catalog association can be imported using the access package ID, the resource association ID, the resource origin ID, and the access type, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azuread_access_package_resource_package_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222/33333333-3333-3333-3333-33333333/Member ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the Access Package ID, the Resource Association ID, the Resource Origin ID, and the Access Type, in the format ` + "`" + `{AccessPackageID}/{ResourceAssociationID}/{ResourceOriginID}/{AccessType}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource. The ID is combined by four fields with colon in between, the four fields are ` + "`" + `access_package_id` + "`" + `, this package association id, ` + "`" + `resource_origin_id` + "`" + ` and ` + "`" + `access_type` + "`" + `. ## Import The resource and catalog association can be imported using the access package ID, the resource association ID, the resource origin ID, and the access type, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azuread_access_package_resource_package_association.example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222/33333333-3333-3333-3333-33333333/Member ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the Access Package ID, the Resource Association ID, the Resource Origin ID, and the Access Type, in the format ` + "`" + `{AccessPackageID}/{ResourceAssociationID}/{ResourceOriginID}/{AccessType}` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_administrative_unit",
			Category:         "Administrative Units",
			ShortDescription: ``,
			Description: `

Manages an Administrative Unit within Azure Active Directory.

`,
			Keywords: []string{
				"administrative",
				"units",
				"unit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the administrative unit.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name of the administrative unit.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups. !>`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Whether the administrative unit _and_ its members are hidden or publicly viewable in the directory. Must be one of: ` + "`" + `Hiddenmembership` + "`" + ` or ` + "`" + `Public` + "`" + `. Defaults to ` + "`" + `Public` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the administrative unit. ## Import Administrative units can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_administrative_unit.example 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the administrative unit. ## Import Administrative units can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_administrative_unit.example 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_administrative_unit_member",
			Category:         "Administrative Units",
			ShortDescription: ``,
			Description: `

Manages a single administrative unit membership within Azure Active Directory.

~> **Warning** Do not use this resource at the same time as the ` + "`" + `members` + "`" + ` property of the ` + "`" + `azuread_administrative_unit` + "`" + ` resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.

`,
			Keywords: []string{
				"administrative",
				"units",
				"unit",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "administrative_unit_object_id",
					Description: `(Required) The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "member_object_id",
					Description: `(Required) The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_administrative_unit_role_member",
			Category:         "Administrative Units",
			ShortDescription: ``,
			Description: `

Manages a single directory role assignment scoped to an administrative unit within Azure Active Directory.

`,
			Keywords: []string{
				"administrative",
				"units",
				"unit",
				"role",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "administrative_unit_object_id",
					Description: `(Required) The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "member_object_id",
					Description: `(Required) The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "role_object_id",
					Description: `(Required) The object ID of the directory role you want to assign. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_app_role_assignment",
			Category:         "App Role Assignments",
			ShortDescription: ``,
			Description: `

Manages an app role assignment for a group, user or service principal. Can be used to grant admin consent for application permissions.

`,
			Keywords: []string{
				"app",
				"role",
				"assignments",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_role_id",
					Description: `(Required) The ID of the app role to be assigned, or the default role ID ` + "`" + `00000000-0000-0000-0000-000000000000` + "`" + `. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "principal_object_id",
					Description: `(Required) The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_object_id",
					Description: `(Required) The object ID of the service principal representing the resource. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "principal_display_name",
					Description: `The display name of the principal to which the app role is assigned.`,
				},
				resource.Attribute{
					Name:        "principal_type",
					Description: `The object type of the principal to which the app role is assigned.`,
				},
				resource.Attribute{
					Name:        "resource_display_name",
					Description: `The display name of the application representing the resource. ## Import App role assignments can be imported using the object ID of the service principal representing the resource and the ID of the app role assignment (note: _not_ the ID of the app role), e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_app_role_assignment.example 00000000-0000-0000-0000-000000000000/appRoleAssignment/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the Resource Service Principal Object ID and the ID of the App Role Assignment in the format ` + "`" + `{ResourcePrincipalID}/appRoleAssignment/{AppRoleAssignmentID}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "principal_display_name",
					Description: `The display name of the principal to which the app role is assigned.`,
				},
				resource.Attribute{
					Name:        "principal_type",
					Description: `The object type of the principal to which the app role is assigned.`,
				},
				resource.Attribute{
					Name:        "resource_display_name",
					Description: `The display name of the application representing the resource. ## Import App role assignments can be imported using the object ID of the service principal representing the resource and the ID of the app role assignment (note: _not_ the ID of the app role), e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_app_role_assignment.example 00000000-0000-0000-0000-000000000000/appRoleAssignment/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the Resource Service Principal Object ID and the ID of the App Role Assignment in the format ` + "`" + `{ResourcePrincipalID}/appRoleAssignment/{AppRoleAssignmentID}` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages an application registration within Azure Active Directory.

`,
			Keywords: []string{
				"applications",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api",
					Description: `(Optional) An ` + "`" + `api` + "`" + ` block as documented below, which configures API related settings for this application.`,
				},
				resource.Attribute{
					Name:        "app_role",
					Description: `(Optional) A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the application, as shown to end users.`,
				},
				resource.Attribute{
					Name:        "device_only_auth_enabled",
					Description: `(Optional) Specifies whether this application supports device authentication without a user. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the application.`,
				},
				resource.Attribute{
					Name:        "fallback_public_client_enabled",
					Description: `(Optional) Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "feature_tags",
					Description: `(Optional) A ` + "`" + `feature_tags` + "`" + ` block as described below. Cannot be used together with the ` + "`" + `tags` + "`" + ` property. ->`,
				},
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `(Optional) Configures the ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are ` + "`" + `None` + "`" + `, ` + "`" + `SecurityGroup` + "`" + `, ` + "`" + `DirectoryRole` + "`" + `, ` + "`" + `ApplicationGroup` + "`" + ` or ` + "`" + `All` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "identifier_uris",
					Description: `(Optional) A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
				},
				resource.Attribute{
					Name:        "logo_image",
					Description: `(Optional) A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.`,
				},
				resource.Attribute{
					Name:        "marketing_url",
					Description: `(Optional) URL of the application's marketing page.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) User-specified notes relevant for the management of the application.`,
				},
				resource.Attribute{
					Name:        "oauth2_post_response_required",
					Description: `(Optional) Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to ` + "`" + `false` + "`" + `, which specifies that only GET requests are allowed.`,
				},
				resource.Attribute{
					Name:        "optional_claims",
					Description: `(Optional) An ` + "`" + `optional_claims` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned. ->`,
				},
				resource.Attribute{
					Name:        "prevent_duplicate_names",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, will return an error if an existing application is found with the same name. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "privacy_statement_url",
					Description: `(Optional) URL of the application's privacy statement.`,
				},
				resource.Attribute{
					Name:        "public_client",
					Description: `(Optional) A ` + "`" + `public_client` + "`" + ` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `(Optional) A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "service_management_reference",
					Description: `(Optional) References application context information from a Service or Asset Management database.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `(Optional) The Microsoft account types that are supported for the current application. Must be one of ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `. Defaults to ` + "`" + `AzureADMyOrg` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "single_page_application",
					Description: `(Optional) A ` + "`" + `single_page_application` + "`" + ` block as documented below, which configures single-page application (SPA) related settings for this application.`,
				},
				resource.Attribute{
					Name:        "support_url",
					Description: `(Optional) URL of the application's support page.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tags to apply to the application for configuring specific behaviours of the application and linked service principals. Note that these are not provided for use by practitioners. Cannot be used together with the ` + "`" + `feature_tags` + "`" + ` block. ->`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional) Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "terms_of_service_url",
					Description: `(Optional) URL of the application's terms of service statement.`,
				},
				resource.Attribute{
					Name:        "web",
					Description: `(Optional) A ` + "`" + `web` + "`" + ` block as documented below, which configures web related settings for this application. ->`,
				},
				resource.Attribute{
					Name:        "known_client_applications",
					Description: `(Optional) A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.`,
				},
				resource.Attribute{
					Name:        "mapped_claims_enabled",
					Description: `(Optional) Allows an application to use claims mapping without specifying a custom signing key. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope",
					Description: `(Optional) One or more ` + "`" + `oauth2_permission_scope` + "`" + ` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.`,
				},
				resource.Attribute{
					Name:        "requested_access_token_version",
					Description: `(Optional) The access token version expected by this resource. Must be one of ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `, and must be ` + "`" + `2` + "`" + ` when ` + "`" + `sign_in_audience` + "`" + ` is either ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + ` Defaults to ` + "`" + `1` + "`" + `. --- ` + "`" + `oauth2_permission_scope` + "`" + ` blocks support the following:`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `(Required) Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `(Required) Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Determines if the permission scope is enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier of the delegated permission. Must be a valid UUID. ->`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to ` + "`" + `User` + "`" + `. Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `(Optional) Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `(Optional) Display name for the delegated permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. ~>`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `(Required) Specifies whether this app role definition can be assigned to users and groups by setting to ` + "`" + `User` + "`" + `, or to other applications (that are accessing this application in a standalone scenario) by setting to ` + "`" + `Application` + "`" + `, or to both.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name for the app role that appears during app role assignment and in consent experiences.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Determines if the app role is enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier of the app role. Must be a valid UUID. ->`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. ->`,
				},
				resource.Attribute{
					Name:        "custom_single_sign_on",
					Description: `(Optional) Whether this application represents a custom SAML application for linked service principals. Enabling this will assign the ` + "`" + `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` + "`" + ` tag. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enterprise",
					Description: `(Optional) Whether this application represents an Enterprise Application for linked service principals. Enabling this will assign the ` + "`" + `WindowsAzureActiveDirectoryIntegratedApp` + "`" + ` tag. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gallery",
					Description: `(Optional) Whether this application represents a gallery application for linked service principals. Enabling this will assign the ` + "`" + `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` + "`" + ` tag. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hide",
					Description: `(Optional) Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the ` + "`" + `HideApp` + "`" + ` tag. Defaults to ` + "`" + `false` + "`" + `. --- ` + "`" + `optional_claims` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `(Optional) One or more ` + "`" + `access_token` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "id_token",
					Description: `(Optional) One or more ` + "`" + `id_token` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "saml2_token",
					Description: `(Optional) One or more ` + "`" + `saml2_token` + "`" + ` blocks as documented below. --- ` + "`" + `access_token` + "`" + `, ` + "`" + `id_token` + "`" + ` and ` + "`" + `saml2_token` + "`" + ` blocks support the following:`,
				},
				resource.Attribute{
					Name:        "additional_properties",
					Description: `List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.`,
				},
				resource.Attribute{
					Name:        "essential",
					Description: `Whether the claim specified by the client is necessary to ensure a smooth authorization experience.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the optional claim.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object. --- ` + "`" + `public_client` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `(Optional) A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid ` + "`" + `https` + "`" + ` or ` + "`" + `ms-appx-web` + "`" + ` URL. --- ` + "`" + `required_resource_access` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "resource_access",
					Description: `(Required) A collection of ` + "`" + `resource_access` + "`" + ` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.`,
				},
				resource.Attribute{
					Name:        "resource_app_id",
					Description: `(Required) The unique identifier for the resource that the application requires access to. This should be the Application ID of the target application. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier for an app role or OAuth2 permission scope published by the resource application.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies whether the ` + "`" + `id` + "`" + ` property references an app role or an OAuth2 permission scope. Possible values are ` + "`" + `Role` + "`" + ` or ` + "`" + `Scope` + "`" + `. --- ` + "`" + `single_page_application` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `(Optional) A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid ` + "`" + `https` + "`" + ` URL. --- ` + "`" + `web` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `(Optional) Home page or landing page of the application.`,
				},
				resource.Attribute{
					Name:        "implicit_grant",
					Description: `(Optional) An ` + "`" + `implicit_grant` + "`" + ` block as documented above.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `(Optional) The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `(Optional) A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid ` + "`" + `http` + "`" + ` URL or a URN. --- ` + "`" + `implicit_grant` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "access_token_issuance_enabled",
					Description: `(Optional) Whether this web application can request an access token using OAuth 2.0 implicit flow.`,
				},
				resource.Attribute{
					Name:        "id_token_issuance_enabled",
					Description: `(Optional) Whether this web application can request an ID token using OAuth 2.0 implicit flow. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (also called Client ID).`,
				},
				resource.Attribute{
					Name:        "disabled_by_microsoft",
					Description: `Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. ` + "`" + `DisabledDueToViolationOfServicesAgreement` + "`" + ``,
				},
				resource.Attribute{
					Name:        "logo_url",
					Description: `CDN URL to the application's logo, as uploaded with the ` + "`" + `logo_image` + "`" + ` property.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The application's object ID.`,
				},
				resource.Attribute{
					Name:        "publisher_domain",
					Description: `The verified publisher domain for the application. ## Import Applications can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (also called Client ID).`,
				},
				resource.Attribute{
					Name:        "disabled_by_microsoft",
					Description: `Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. ` + "`" + `DisabledDueToViolationOfServicesAgreement` + "`" + ``,
				},
				resource.Attribute{
					Name:        "logo_url",
					Description: `CDN URL to the application's logo, as uploaded with the ` + "`" + `logo_image` + "`" + ` property.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The application's object ID.`,
				},
				resource.Attribute{
					Name:        "publisher_domain",
					Description: `The verified publisher domain for the application. ## Import Applications can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_certificate",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages a certificate associated with an application within Azure Active Directory. These are also referred to as client certificates during authentication.

`,
			Keywords: []string{
				"applications",
				"application",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The object ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "encoding",
					Description: `(Optional) Specifies the encoding used for the supplied certificate data. Must be one of ` + "`" + `pem` + "`" + `, ` + "`" + `base64` + "`" + ` or ` + "`" + `hex` + "`" + `. Defaults to ` + "`" + `pem` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the certificate is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created. ~> One of ` + "`" + `end_date` + "`" + ` or ` + "`" + `end_date_relative` + "`" + ` must be set. The maximum allowed duration is determined by Azure AD.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of key/certificate. Must be one of ` + "`" + `AsymmetricX509Cert` + "`" + ` or ` + "`" + `Symmetric` + "`" + `. Changing this fields forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the ` + "`" + `encoding` + "`" + ` argument. ## Attributes Reference No additional attributes are exported. ## Import Certificates can be imported using the object ID of the associated application and the key ID of the certificate credential, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application_certificate.test 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the application's object ID, the string "certificate" and the certificate's key ID in the format ` + "`" + `{ObjectId}/certificate/{CertificateKeyId}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_federated_identity_credential",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages a federated identity credential associated with an application within Azure Active Directory.

`,
			Keywords: []string{
				"applications",
				"application",
				"federated",
				"identity",
				"credential",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "audiences",
					Description: `(Required) List of audiences that can appear in the external token. This specifies what should be accepted in the ` + "`" + `aud` + "`" + ` claim of incoming tokens.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the federated identity credential.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) A unique display name for the federated identity credential. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `(Required) The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `(Required) The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "credential_id",
					Description: `A UUID used to uniquely identify this federated identity credential. ## Import Federated Identity Credentials can be imported using the object ID of the associated application and the ID of the federated identity credential, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application_federated_identity_credential.test 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the application's object ID, the string "federatedIdentityCredential" and the credential ID in the format ` + "`" + `{ObjectId}/federatedIdentityCredential/{CredentialId}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "credential_id",
					Description: `A UUID used to uniquely identify this federated identity credential. ## Import Federated Identity Credentials can be imported using the object ID of the associated application and the ID of the federated identity credential, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application_federated_identity_credential.test 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the application's object ID, the string "federatedIdentityCredential" and the credential ID in the format ` + "`" + `{ObjectId}/federatedIdentityCredential/{CredentialId}` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_password",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages a password credential associated with an application within Azure Active Directory. These are also referred to as client secrets during authentication.

`,
			Keywords: []string{
				"applications",
				"application",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The object ID of the application for which this password should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) A display name for the password. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The end date until which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the password is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "rotate_when_changed",
					Description: `(Optional) A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The start date from which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify this password credential.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The password for this application, which is generated by Azure Active Directory. ## Import This resource does not support importing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify this password credential.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The password for this application, which is generated by Azure Active Directory. ## Import This resource does not support importing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_pre_authorized",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages client applications that are pre-authorized with the specified permissions to access an application's APIs without requiring user consent.

`,
			Keywords: []string{
				"applications",
				"application",
				"pre",
				"authorized",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "authorized_application_id",
					Description: `(Optional) The application ID (client ID) of the application being authorized. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "permission_ids",
					Description: `(Required) A set of permission scope IDs required by the authorized application. ## Attributes Reference No additional attributes are exported. ## Import Pre-authorized applications can be imported using the object ID of the authorizing application and the application ID of the application being authorized, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application_pre_authorized.example 00000000-0000-0000-0000-000000000000/preAuthorizedApplication/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the authorizing application's object ID, the string "preAuthorizedApplication" and the authorized application's application ID (client ID) in the format ` + "`" + `{ObjectId}/preAuthorizedApplication/{ApplicationId}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_claims_mapping_policy",
			Category:         "Policies",
			ShortDescription: ``,
			Description: `

Manages a Claims Mapping Policy within Azure Active Directory.

`,
			Keywords: []string{
				"policies",
				"claims",
				"mapping",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) The claims mapping policy. This is a JSON formatted string, for which the [` + "`" + `jsonencode()` + "`" + `](https://www.terraform.io/language/functions/jsonencode) function can be used.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for this Claims Mapping Policy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Claims Mapping Policy. ## Import Claims Mapping Policy can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_claims_mapping_policy.my_policy 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Claims Mapping Policy. ## Import Claims Mapping Policy can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_claims_mapping_policy.my_policy 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_conditional_access_policy",
			Category:         "Conditional Access",
			ShortDescription: ``,
			Description: `

Manages a Conditional Access Policy within Azure Active Directory.

`,
			Keywords: []string{
				"conditional",
				"access",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Required) A ` + "`" + `conditions` + "`" + ` block as documented below, which specifies the rules that must be met for the policy to apply.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The friendly name for this Conditional Access Policy.`,
				},
				resource.Attribute{
					Name:        "grant_controls",
					Description: `(Required) A ` + "`" + `grant_controls` + "`" + ` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.`,
				},
				resource.Attribute{
					Name:        "session_controls",
					Description: `(Optional) A ` + "`" + `session_controls` + "`" + ` block as documented below, which specifies the session controls that are enforced after sign-in.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Required) Specifies the state of the policy object. Possible values are: ` + "`" + `enabled` + "`" + `, ` + "`" + `disabled` + "`" + ` and ` + "`" + `enabledForReportingButNotEnforced` + "`" + ` --- ` + "`" + `conditions` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Required) An ` + "`" + `applications` + "`" + ` block as documented below, which specifies applications and user actions included in and excluded from the policy.`,
				},
				resource.Attribute{
					Name:        "client_app_types",
					Description: `(Required) A list of client application types included in the policy. Possible values are: ` + "`" + `all` + "`" + `, ` + "`" + `browser` + "`" + `, ` + "`" + `mobileAppsAndDesktopClients` + "`" + `, ` + "`" + `exchangeActiveSync` + "`" + `, ` + "`" + `easSupported` + "`" + ` and ` + "`" + `other` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `(Optional) A ` + "`" + `devices` + "`" + ` block as documented below, which describes devices to be included in and excluded from the policy. A ` + "`" + `devices` + "`" + ` block can be added to an existing policy, but removing the ` + "`" + `devices` + "`" + ` block forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Required) A ` + "`" + `locations` + "`" + ` block as documented below, which specifies locations included in and excluded from the policy.`,
				},
				resource.Attribute{
					Name:        "platforms",
					Description: `(Required) A ` + "`" + `platforms` + "`" + ` block as documented below, which specifies platforms included in and excluded from the policy.`,
				},
				resource.Attribute{
					Name:        "sign_in_risk_levels",
					Description: `(Optional) A list of sign-in risk levels included in the policy. Possible values are: ` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `hidden` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `unknownFutureValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_risk_levels",
					Description: `(Optional) A list of user risk levels included in the policy. Possible values are: ` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `hidden` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `unknownFutureValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Required) A ` + "`" + `users` + "`" + ` block as documented below, which specifies users, groups, and roles included in and excluded from the policy. --- ` + "`" + `applications` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "excluded_applications",
					Description: `(Optional) A list of application IDs explicitly excluded from the policy. Can also be set to ` + "`" + `Office365` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "included_applications",
					Description: `(Optional) A list of application IDs the policy applies to, unless explicitly excluded (in ` + "`" + `excluded_applications` + "`" + `). Can also be set to ` + "`" + `All` + "`" + `, ` + "`" + `None` + "`" + ` or ` + "`" + `Office365` + "`" + `. Cannot be specified with ` + "`" + `included_user_actions` + "`" + `. One of ` + "`" + `included_applications` + "`" + ` or ` + "`" + `included_user_actions` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "included_user_actions",
					Description: `(Optional) A list of user actions to include. Supported values are ` + "`" + `urn:user:registerdevice` + "`" + ` and ` + "`" + `urn:user:registersecurityinfo` + "`" + `. Cannot be specified with ` + "`" + `included_applications` + "`" + `. One of ` + "`" + `included_applications` + "`" + ` or ` + "`" + `included_user_actions` + "`" + ` must be specified. --- ` + "`" + `devices` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A ` + "`" + `filter` + "`" + ` block as described below. A ` + "`" + `filter` + "`" + ` block can be added to an existing policy, but removing the ` + "`" + `filter` + "`" + ` block forces a new resource to be created. --- ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) Whether to include in, or exclude from, matching devices from the policy. Supported values are ` + "`" + `include` + "`" + ` or ` + "`" + `exclude` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) Condition filter to match devices. For more information, see [official documentation](https://docs.microsoft.com/en-us/azure/active-directory/conditional-access/concept-condition-filters-for-devices#supported-operators-and-device-properties-for-filters). --- ` + "`" + `users` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "excluded_groups",
					Description: `(Optional) A list of group IDs excluded from scope of policy.`,
				},
				resource.Attribute{
					Name:        "excluded_roles",
					Description: `(Optional) A list of role IDs excluded from scope of policy.`,
				},
				resource.Attribute{
					Name:        "excluded_users",
					Description: `(Optional) A list of user IDs excluded from scope of policy and/or ` + "`" + `GuestsOrExternalUsers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "included_groups",
					Description: `(Optional) A list of group IDs in scope of policy unless explicitly excluded.`,
				},
				resource.Attribute{
					Name:        "included_roles",
					Description: `(Optional) A list of role IDs in scope of policy unless explicitly excluded.`,
				},
				resource.Attribute{
					Name:        "included_users",
					Description: `(Optional) A list of user IDs in scope of policy unless explicitly excluded, or ` + "`" + `None` + "`" + ` or ` + "`" + `All` + "`" + ` or ` + "`" + `GuestsOrExternalUsers` + "`" + `. -> At least one of ` + "`" + `included_groups` + "`" + `, ` + "`" + `included_roles` + "`" + ` or ` + "`" + `included_users` + "`" + ` must be specified. --- ` + "`" + `locations` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "excluded_locations",
					Description: `(Optional) A list of location IDs excluded from scope of policy. Can also be set to ` + "`" + `AllTrusted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "included_locations",
					Description: `(Required) A list of location IDs in scope of policy unless explicitly excluded. Can also be set to ` + "`" + `All` + "`" + `, or ` + "`" + `AllTrusted` + "`" + `. --- ` + "`" + `platforms` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "excluded_platforms",
					Description: `(Optional) A list of platforms explicitly excluded from the policy. Possible values are: ` + "`" + `all` + "`" + `, ` + "`" + `android` + "`" + `, ` + "`" + `iOS` + "`" + `, ` + "`" + `linux` + "`" + `, ` + "`" + `macOS` + "`" + `, ` + "`" + `windows` + "`" + `, ` + "`" + `windowsPhone` + "`" + ` or ` + "`" + `unknownFutureValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "included_platforms",
					Description: `(Required) A list of platforms the policy applies to, unless explicitly excluded. Possible values are: ` + "`" + `all` + "`" + `, ` + "`" + `android` + "`" + `, ` + "`" + `iOS` + "`" + `, ` + "`" + `linux` + "`" + `, ` + "`" + `macOS` + "`" + `, ` + "`" + `windows` + "`" + `, ` + "`" + `windowsPhone` + "`" + ` or ` + "`" + `unknownFutureValue` + "`" + `. --- ` + "`" + `grant_controls` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "built_in_controls",
					Description: `(Required) List of built-in controls required by the policy. Possible values are: ` + "`" + `block` + "`" + `, ` + "`" + `mfa` + "`" + `, ` + "`" + `approvedApplication` + "`" + `, ` + "`" + `compliantApplication` + "`" + `, ` + "`" + `compliantDevice` + "`" + `, ` + "`" + `domainJoinedDevice` + "`" + `, ` + "`" + `passwordChange` + "`" + ` or ` + "`" + `unknownFutureValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_authentication_factors",
					Description: `(Optional) List of custom controls IDs required by the policy.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Defines the relationship of the grant controls. Possible values are: ` + "`" + `AND` + "`" + `, ` + "`" + `OR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "terms_of_use",
					Description: `(Optional) List of terms of use IDs required by the policy. --- ` + "`" + `session_controls` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "application_enforced_restrictions_enabled",
					Description: `(Optional) Whether or not application enforced restrictions are enabled. Defaults to ` + "`" + `false` + "`" + `. -> Only Office 365, Exchange Online and Sharepoint Online support application enforced restrictions.`,
				},
				resource.Attribute{
					Name:        "cloud_app_security_policy",
					Description: `(Optional) Enables cloud app security and specifies the cloud app security policy to use. Possible values are: ` + "`" + `blockDownloads` + "`" + `, ` + "`" + `mcasConfigured` + "`" + `, ` + "`" + `monitorOnly` + "`" + ` or ` + "`" + `unknownFutureValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "persistent_browser_mode",
					Description: `(Optional) Session control to define whether to persist cookies or not. Possible values are: ` + "`" + `always` + "`" + ` or ` + "`" + `never` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sign_in_frequency",
					Description: `(Optional) Number of days or hours to enforce sign-in frequency. Required when ` + "`" + `sign_in_frequency_period` + "`" + ` is specified. Due to an API issue, removing this property forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "sign_in_frequency_period",
					Description: `(Optional) The time period to enforce sign-in frequency. Possible values are: ` + "`" + `hours` + "`" + ` or ` + "`" + `days` + "`" + `. Required when ` + "`" + `sign_in_frequency_period` + "`" + ` is specified. Due to an API issue, removing this property forces a new resource to be created. --- ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Conditional Access Policy. ## Import Conditional Access Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_conditional_access_policy.my_location 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Conditional Access Policy. ## Import Conditional Access Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_conditional_access_policy.my_location 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_custom_directory_role",
			Category:         "Directory Roles",
			ShortDescription: ``,
			Description: `

Manages a Custom Directory Role within Azure Active Directory.

This resource is for managing custom directory roles. For management of built-in roles, see the [azuread_directory_role](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/directory_role) resource.

`,
			Keywords: []string{
				"directory",
				"roles",
				"custom",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the custom directory role.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name of the custom directory role.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicates whether the role is enabled for assignment.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) A collection of ` + "`" + `permissions` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional) Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) - The version of the role definition. This can be any arbitrary string between 1-128 characters. --- ` + "`" + `permissions` + "`" + ` blocks support the following:`,
				},
				resource.Attribute{
					Name:        "allowed_resource_actions",
					Description: `(Required) A set of tasks that can be performed on a resource. For more information, see the [Permissions Reference](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference) documentation. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the custom directory role. ## Import This resource does not support importing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the custom directory role. ## Import This resource does not support importing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_directory_role",
			Category:         "Directory Roles",
			ShortDescription: ``,
			Description: `

Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.

Directory Roles are built-in to Azure Active Directory and are immutable. However, by default they are not activated in a tenant (except for the Global Administrator role). This resource ensures a directory role is activated from its associated role template, and exports the object ID of the role, so that role assignments can be made for it.

Once activated, directory roles cannot be deactivated and so this resource does not perform any actions on destroy.

`,
			Keywords: []string{
				"directory",
				"roles",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of the directory role to activate. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional) The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created. ~> Either ` + "`" + `display_name` + "`" + ` or ` + "`" + `template_id` + "`" + ` must be specified. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the directory role.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the directory role. ## Import This resource does not support importing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the directory role.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the directory role. ## Import This resource does not support importing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_directory_role_assignment",
			Category:         "Directory Roles",
			ShortDescription: ``,
			Description: `

Manages a single directory role assignment within Azure Active Directory.

`,
			Keywords: []string{
				"directory",
				"roles",
				"role",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_scope_id",
					Description: `(Optional) Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with ` + "`" + `directory_scope_id` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "directory_scope_id",
					Description: `(Optional) Identifier of the directory object representing the scope of the assignment. Cannot be used with ` + "`" + `app_scope_id` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "principal_object_id",
					Description: `(Required) The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_directory_role_member",
			Category:         "Directory Roles",
			ShortDescription: ``,
			Description: `

Manages a single directory role membership (assignment) within Azure Active Directory.

-> **Deprecation Warning:** This resource has been superseded by the [azuread_directory_role_assignment](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/directory_role_assignment) resource and will be removed in version 3.0 of the AzureAD provider

`,
			Keywords: []string{
				"directory",
				"roles",
				"role",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "member_object_id",
					Description: `(Required) The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "role_object_id",
					Description: `(Required) The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_group",
			Category:         "Groups",
			ShortDescription: ``,
			Description: `

Manages a group within Azure Active Directory.

`,
			Keywords: []string{
				"groups",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "administrative_unit_ids",
					Description: `(Optional) The object IDs of administrative units in which the group is a member. If specified, new groups will be created in the scope of the first administrative unit and added to the others. If empty, new groups will be created at the tenant level. !>`,
				},
				resource.Attribute{
					Name:        "assignable_to_role",
					Description: `(Optional) Indicates whether this group can be assigned to an Azure Active Directory role. Defaults to ` + "`" + `false` + "`" + `. Can only be set to ` + "`" + `true` + "`" + ` for security-enabled groups. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "auto_subscribe_new_members",
					Description: `(Optional) Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups. ~>`,
				},
				resource.Attribute{
					Name:        "behaviors",
					Description: `(Optional) A set of behaviors for a Microsoft 365 group. Possible values are ` + "`" + `AllowOnlyMembersToPost` + "`" + `, ` + "`" + `HideGroupInOutlook` + "`" + `, ` + "`" + `SubscribeMembersToCalendarEventsDisabled` + "`" + `, ` + "`" + `SubscribeNewGroupMembers` + "`" + ` and ` + "`" + `WelcomeEmailDisabled` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the group.`,
				},
				resource.Attribute{
					Name:        "dynamic_membership",
					Description: `(Optional) A ` + "`" + `dynamic_membership` + "`" + ` block as documented below. Required when ` + "`" + `types` + "`" + ` contains ` + "`" + `DynamicMembership` + "`" + `. Cannot be used with the ` + "`" + `members` + "`" + ` property.`,
				},
				resource.Attribute{
					Name:        "external_senders_allowed",
					Description: `(Optional) Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups. ~>`,
				},
				resource.Attribute{
					Name:        "hide_from_address_lists",
					Description: `(Optional) Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups. ~>`,
				},
				resource.Attribute{
					Name:        "hide_from_outlook_clients",
					Description: `(Optional) Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups. ~>`,
				},
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `(Optional) Whether the group is a mail enabled, with a shared group mailbox. At least one of ` + "`" + `mail_enabled` + "`" + ` or ` + "`" + `security_enabled` + "`" + ` must be specified. Only Microsoft 365 groups can be mail enabled (see the ` + "`" + `types` + "`" + ` property).`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `(Optional) The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the ` + "`" + `dynamic_membership` + "`" + ` block. !>`,
				},
				resource.Attribute{
					Name:        "onpremises_group_type",
					Description: `(Optional) The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are ` + "`" + `UniversalDistributionGroup` + "`" + `, ` + "`" + `UniversalMailEnabledSecurityGroup` + "`" + `, or ` + "`" + `UniversalSecurityGroup` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A set of object IDs of principals that will be granted ownership of the group. Supported object types are users or service principals. By default, the principal being used to execute Terraform is assigned as the sole owner. Groups cannot be created with no owners or have all their owners removed. ->`,
				},
				resource.Attribute{
					Name:        "prevent_duplicate_names",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, will return an error if an existing group is found with the same name. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioning_options",
					Description: `(Optional) A set of provisioning options for a Microsoft 365 group. The only supported value is ` + "`" + `Team` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `(Optional) Whether the group is a security group for controlling access to in-app resources. At least one of ` + "`" + `security_enabled` + "`" + ` or ` + "`" + `mail_enabled` + "`" + ` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the ` + "`" + `types` + "`" + ` property).`,
				},
				resource.Attribute{
					Name:        "theme",
					Description: `(Optional) The colour theme for a Microsoft 365 group. Possible values are ` + "`" + `Blue` + "`" + `, ` + "`" + `Green` + "`" + `, ` + "`" + `Orange` + "`" + `, ` + "`" + `Pink` + "`" + `, ` + "`" + `Purple` + "`" + `, ` + "`" + `Red` + "`" + ` or ` + "`" + `Teal` + "`" + `. By default, no theme is set.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `(Optional) A set of group types to configure for the group. Supported values are ` + "`" + `DynamicMembership` + "`" + `, which denotes a group with dynamic membership, and ` + "`" + `Unified` + "`" + `, which specifies a Microsoft 365 group. Required when ` + "`" + `mail_enabled` + "`" + ` is true. Changing this forces a new resource to be created. ->`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The group join policy and group content visibility. Possible values are ` + "`" + `Private` + "`" + `, ` + "`" + `Public` + "`" + `, or ` + "`" + `Hiddenmembership` + "`" + `. Only Microsoft 365 groups can have ` + "`" + `Hiddenmembership` + "`" + ` visibility and this value must be set when the group is created. By default, security groups will receive ` + "`" + `Private` + "`" + ` visibility and Microsoft 365 groups will receive ` + "`" + `Public` + "`" + ` visibility. ->`,
				},
				resource.Attribute{
					Name:        "writeback_enabled",
					Description: `(Optional) Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used. --- ` + "`" + `dynamic_membership` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Whether rule processing is "On" (true) or "Paused" (false).`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) The rule that determines membership of this group. For more information, see official documentation on [membership rules syntax](https://docs.microsoft.com/en-gb/azure/active-directory/enterprise-users/groups-dynamic-membership). ~>`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The SMTP address for the group.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the group.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_netbios_name",
					Description: `The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_security_identifier",
					Description: `The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sync_enabled",
					Description: `Whether this group is synchronised from an on-premises directory (` + "`" + `true` + "`" + `), no longer synchronised (` + "`" + `false` + "`" + `), or has never been synchronised (` + "`" + `null` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `The preferred language for a Microsoft 365 group, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the group that direct to the same group mailbox. ## Import Groups can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mail",
					Description: `The SMTP address for the group.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the group.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_netbios_name",
					Description: `The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_security_identifier",
					Description: `The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sync_enabled",
					Description: `Whether this group is synchronised from an on-premises directory (` + "`" + `true` + "`" + `), no longer synchronised (` + "`" + `false` + "`" + `), or has never been synchronised (` + "`" + `null` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `The preferred language for a Microsoft 365 group, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the group that direct to the same group mailbox. ## Import Groups can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_group_member",
			Category:         "Groups",
			ShortDescription: ``,
			Description: `

Manages a single group membership within Azure Active Directory.

~> **Warning** Do not use this resource at the same time as the ` + "`" + `members` + "`" + ` property of the ` + "`" + `azuread_group` + "`" + ` resource for the same group. Doing so will cause a conflict and group members will be removed.

`,
			Keywords: []string{
				"groups",
				"group",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_object_id",
					Description: `(Required) The object ID of the group you want to add the member to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "member_object_id",
					Description: `(Required) The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_invitation",
			Category:         "Invitations",
			ShortDescription: ``,
			Description: `

Manages an invitation of a guest user within Azure Active Directory.

`,
			Keywords: []string{
				"invitations",
				"invitation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "message",
					Description: `(Optional) A ` + "`" + `message` + "`" + ` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.`,
				},
				resource.Attribute{
					Name:        "redirect_url",
					Description: `(Required) The URL that the user should be redirected to once the invitation is redeemed.`,
				},
				resource.Attribute{
					Name:        "user_display_name",
					Description: `(Optional) The display name of the user being invited.`,
				},
				resource.Attribute{
					Name:        "user_email_address",
					Description: `(Required) The email address of the user being invited.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `(Optional) The user type of the user being invited. Must be one of ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `. Only Global Administrators can invite users as members. Defaults to ` + "`" + `Guest` + "`" + `. --- ` + "`" + `message` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "additional_recipients",
					Description: `(Optional) Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `(Optional) Customized message body you want to send if you don't want to send the default message. Cannot be specified with ` + "`" + `language` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `(Optional) The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to ` + "`" + `en-US` + "`" + `. Cannot be specified with ` + "`" + `body` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "redeem_url",
					Description: `The URL the user can use to redeem their invitation.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Object ID of the invited user. ## Import This resource does not support importing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "redeem_url",
					Description: `The URL the user can use to redeem their invitation.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Object ID of the invited user. ## Import This resource does not support importing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_named_location",
			Category:         "Conditional Access",
			ShortDescription: ``,
			Description: `

Manages a Named Location within Azure Active Directory.

`,
			Keywords: []string{
				"conditional",
				"access",
				"named",
				"location",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) A ` + "`" + `country` + "`" + ` block as documented below, which configures a country-based named location.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The friendly name for this named location.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) An ` + "`" + `ip` + "`" + ` block as documented below, which configures an IP-based named location. -> Exactly one of ` + "`" + `ip` + "`" + ` or ` + "`" + `country` + "`" + ` must be specified. Changing between these forces a new resource to be created. --- ` + "`" + `country` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "countries_and_regions",
					Description: `(Required) List of countries and/or regions in two-letter format specified by ISO 3166-2.`,
				},
				resource.Attribute{
					Name:        "include_unknown_countries_and_regions",
					Description: `(Optional) Whether IP addresses that don't map to a country or region should be included in the named location. Defaults to ` + "`" + `false` + "`" + `. --- ` + "`" + `ip` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Required) List of IP address ranges in IPv4 CIDR format (e.g. 1.2.3.4/32) or any allowable IPv6 format from IETF RFC596.`,
				},
				resource.Attribute{
					Name:        "trusted",
					Description: `(Optional) Whether the named location is trusted. Defaults to ` + "`" + `false` + "`" + `. --- ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the named location. ## Import Named Locations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_named_location.my_location 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the named location. ## Import Named Locations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_named_location.my_location 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Manages a service principal associated with an application within Azure Active Directory.

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_enabled",
					Description: `(Optional) Whether or not the service principal account is enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alternative_names",
					Description: `(Optional) A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.`,
				},
				resource.Attribute{
					Name:        "app_role_assignment_required",
					Description: `(Optional) Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) The application ID (client ID) of the application for which to create a service principal.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the service principal provided for internal end-users.`,
				},
				resource.Attribute{
					Name:        "feature_tags",
					Description: `(Optional) A ` + "`" + `feature_tags` + "`" + ` block as described below. Cannot be used together with the ` + "`" + `tags` + "`" + ` property. ->`,
				},
				resource.Attribute{
					Name:        "login_url",
					Description: `(Optional) The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) A free text field to capture information about the service principal, typically used for operational purposes.`,
				},
				resource.Attribute{
					Name:        "notification_email_addresses",
					Description: `(Optional) A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned. ->`,
				},
				resource.Attribute{
					Name:        "preferred_single_sign_on_mode",
					Description: `(Optional) The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are ` + "`" + `oidc` + "`" + `, ` + "`" + `password` + "`" + `, ` + "`" + `saml` + "`" + ` or ` + "`" + `notSupported` + "`" + `. Omit this property or specify a blank string to unset.`,
				},
				resource.Attribute{
					Name:        "saml_single_sign_on",
					Description: `(Optional) A ` + "`" + `saml_single_sign_on` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tags to apply to the service principal for configuring specific behaviours of the service principal. Note that these are not provided for use by practitioners. Cannot be used together with the ` + "`" + `feature_tags` + "`" + ` block. ->`,
				},
				resource.Attribute{
					Name:        "use_existing",
					Description: `(Optional) When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal. ->`,
				},
				resource.Attribute{
					Name:        "custom_single_sign_on",
					Description: `(Optional) Whether this service principal represents a custom SAML application. Enabling this will assign the ` + "`" + `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` + "`" + ` tag. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enterprise",
					Description: `(Optional) Whether this service principal represents an Enterprise Application. Enabling this will assign the ` + "`" + `WindowsAzureActiveDirectoryIntegratedApp` + "`" + ` tag. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gallery",
					Description: `(Optional) Whether this service principal represents a gallery application. Enabling this will assign the ` + "`" + `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` + "`" + ` tag. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hide",
					Description: `(Optional) Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the ` + "`" + `HideApp` + "`" + ` tag. Defaults to ` + "`" + `false` + "`" + `. --- ` + "`" + `saml_single_sign_on` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "relay_state",
					Description: `(Optional) The relative URI the service provider would redirect to after completion of the single sign-on flow. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "application_tenant_id",
					Description: `The tenant ID where the associated application is registered.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page or landing page of the associated application.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `The URL that will be used by Microsoft's authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.`,
				},
				resource.Attribute{
					Name:        "saml_metadata_url",
					Description: `The URL where the service exposes SAML metadata for federation.`,
				},
				resource.Attribute{
					Name:        "service_principal_names",
					Description: `A list of identifier URI(s), copied over from the associated application.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the associated application. Possible values include ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Identifies whether the service principal represents an application or a managed identity. Possible values include ` + "`" + `Application` + "`" + ` or ` + "`" + `ManagedIdentity` + "`" + `. --- ` + "`" + `app_roles` + "`" + ` is a list of objects with the following attributes:`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are: ` + "`" + `User` + "`" + ` and ` + "`" + `Application` + "`" + `, or both.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for the app role that appears during app role assignment and in consent experiences.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the ` + "`" + `app_role` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. --- ` + "`" + `oauth2_permission_scopes` + "`" + ` is a list of objects with the following attributes:`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether the permission scope is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the delegated permission.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `Display name for the delegated permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. ## Import Service principals can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "application_tenant_id",
					Description: `The tenant ID where the associated application is registered.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page or landing page of the associated application.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `The URL that will be used by Microsoft's authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.`,
				},
				resource.Attribute{
					Name:        "saml_metadata_url",
					Description: `The URL where the service exposes SAML metadata for federation.`,
				},
				resource.Attribute{
					Name:        "service_principal_names",
					Description: `A list of identifier URI(s), copied over from the associated application.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the associated application. Possible values include ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Identifies whether the service principal represents an application or a managed identity. Possible values include ` + "`" + `Application` + "`" + ` or ` + "`" + `ManagedIdentity` + "`" + `. --- ` + "`" + `app_roles` + "`" + ` is a list of objects with the following attributes:`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are: ` + "`" + `User` + "`" + ` and ` + "`" + `Application` + "`" + `, or both.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for the app role that appears during app role assignment and in consent experiences.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the ` + "`" + `app_role` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. --- ` + "`" + `oauth2_permission_scopes` + "`" + ` is a list of objects with the following attributes:`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether the permission scope is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the delegated permission.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `Display name for the delegated permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. ## Import Service principals can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal_certificate",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Manages a certificate associated with a service principal within Azure Active Directory.

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "encoding",
					Description: `(Optional) Specifies the encoding used for the supplied certificate data. Must be one of ` + "`" + `pem` + "`" + `, ` + "`" + `base64` + "`" + ` or ` + "`" + `hex` + "`" + `. Defaults to ` + "`" + `pem` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the certificate is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created. ~> One of ` + "`" + `end_date` + "`" + ` or ` + "`" + `end_date_relative` + "`" + ` must be set. The maximum duration is determined by Azure AD.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of key/certificate. Must be one of ` + "`" + `AsymmetricX509Cert` + "`" + ` or ` + "`" + `Symmetric` + "`" + `. Changing this fields forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the ` + "`" + `encoding` + "`" + ` argument. ## Attributes Reference No additional attributes are exported. ## Import Certificates can be imported using the object ID of the associated service principal and the key ID of the certificate credential, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_certificate.test 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "certificate" and the certificate's key ID in the format ` + "`" + `{ServicePrincipalObjectId}/certificate/{CertificateKeyId}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal_claims_mapping_policy_assignment",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Manages a Claims Mapping Policy Assignment within Azure Active Directory.

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
				"claims",
				"mapping",
				"policy",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "claims_mapping_policy_id",
					Description: `(Required) The ID of the claims mapping policy to assign.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The object ID of the service principal for the policy assignment. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Claims Mapping Policy Assignment. ## Import Claims Mapping Policy can be imported using the ` + "`" + `id` + "`" + `, in the form ` + "`" + `service-principal-uuid/claimsMappingPolicy/claims-mapping-policy-uuid` + "`" + `, e.g: ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_claims_mapping_policy_assignment.app 00000000-0000-0000-0000-000000000000/claimsMappingPolicy/11111111-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Claims Mapping Policy Assignment. ## Import Claims Mapping Policy can be imported using the ` + "`" + `id` + "`" + `, in the form ` + "`" + `service-principal-uuid/claimsMappingPolicy/claims-mapping-policy-uuid` + "`" + `, e.g: ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_claims_mapping_policy_assignment.app 00000000-0000-0000-0000-000000000000/claimsMappingPolicy/11111111-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal_delegated_permission_grant",
			Category:         "Delegated Permission Grants",
			ShortDescription: ``,
			Description: `

Manages a delegated permission grant for a service principal, on behalf of a single user, or all users.

`,
			Keywords: []string{
				"delegated",
				"permission",
				"grants",
				"service",
				"principal",
				"grant",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "claim_values",
					Description: `(Required) - A set of claim values for delegated permission scopes which should be included in access tokens for the resource.`,
				},
				resource.Attribute{
					Name:        "resource_service_principal_object_id",
					Description: `(Required) The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "service_principal_object_id",
					Description: `(Required) The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "user_object_id",
					Description: `(Optional) - The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the delegated permission grant. ## Import Delegated permission grants can be imported using their ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_delegated_permission_grant.example aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the delegated permission grant. ## Import Delegated permission grants can be imported using their ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_delegated_permission_grant.example aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal_password",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Manages a password credential associated with a service principal within Azure Active Directory. See also the [azuread_application_password resource](application_password.html).

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) A display name for the password.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The end date until which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the password is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "rotate_when_changed",
					Description: `(Optional) A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The start date from which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify this password credential.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The password for this service principal, which is generated by Azure Active Directory. ## Import This resource does not support importing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify this password credential.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The password for this service principal, which is generated by Azure Active Directory. ## Import This resource does not support importing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal_token_signing_certificate",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Manages a token signing certificate associated with a service principal within Azure Active Directory.

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
				"token",
				"signing",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Specifies a friendly name for the certificate. Must start with ` + "`" + `CN=` + "`" + `. Changing this field forces a new resource to be created. ~> If not specified, it will default to ` + "`" + `CN=Microsoft Azure Federated SSO Certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify the verify certificate.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The certificate data, which is PEM encoded but does not include the header ` + "`" + `-----BEGIN CERTIFICATE-----\n` + "`" + ` or the footer ` + "`" + `\n-----END CERTIFICATE-----` + "`" + `. ## Import Token signing certificates can be imported using the object ID of the associated service principal and the key ID of the verify certificate credential, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_token_signing_certificate.example 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "tokenSigningCertificate" and the verify certificate's key ID in the format ` + "`" + `{ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify the verify certificate.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The certificate data, which is PEM encoded but does not include the header ` + "`" + `-----BEGIN CERTIFICATE-----\n` + "`" + ` or the footer ` + "`" + `\n-----END CERTIFICATE-----` + "`" + `. ## Import Token signing certificates can be imported using the object ID of the associated service principal and the key ID of the verify certificate credential, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_token_signing_certificate.example 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "tokenSigningCertificate" and the verify certificate's key ID in the format ` + "`" + `{ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_synchronization_job",
			Category:         "Synchronization",
			ShortDescription: ``,
			Description: `

Manages a synchronization job associated with a service principal (enterprise application) within Azure Active Directory.

`,
			Keywords: []string{
				"synchronization",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether or not the provisioning job is enabled. Default state is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Required) Identifier of the synchronization template this job is based on. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An ID used to uniquely identify this synchronization job.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `A ` + "`" + `schedule` + "`" + ` list as documented below. --- ` + "`" + `schedule` + "`" + ` block exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Date and time when this job will expire, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The interval between synchronization iterations ISO8601. E.g. PT40M run every 40 minutes.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the job. ## Import Synchronization jobs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_synchronization_job.example 00000000-0000-0000-0000-000000000000/job/dataBricks.f5532fc709734b1a90e8a1fa9fd03a82.8442fd39-2183-419c-8732-74b6ce866bd5 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the Service Principal Object ID and the ID of the Synchronization Job Id in the format ` + "`" + `{servicePrincipalId}/job/{jobId}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An ID used to uniquely identify this synchronization job.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `A ` + "`" + `schedule` + "`" + ` list as documented below. --- ` + "`" + `schedule` + "`" + ` block exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Date and time when this job will expire, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The interval between synchronization iterations ISO8601. E.g. PT40M run every 40 minutes.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the job. ## Import Synchronization jobs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_synchronization_job.example 00000000-0000-0000-0000-000000000000/job/dataBricks.f5532fc709734b1a90e8a1fa9fd03a82.8442fd39-2183-419c-8732-74b6ce866bd5 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the Service Principal Object ID and the ID of the Synchronization Job Id in the format ` + "`" + `{servicePrincipalId}/job/{jobId}` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_synchronization_secret",
			Category:         "Synchronization",
			ShortDescription: ``,
			Description: `

Manages synchronization secrets associated with a service principal (enterprise application) within Azure Active Directory.

`,
			Keywords: []string{
				"synchronization",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "credential",
					Description: `(Optional) One or more ` + "`" + `credential` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created. --- ` + "`" + `credential` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the secret.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the secret. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An ID used to uniquely identify this synchronization sec. ## Import This resource does not support importing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An ID used to uniquely identify this synchronization sec. ## Import This resource does not support importing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_user",
			Category:         "Users",
			ShortDescription: ``,
			Description: `

Manages a user within Azure Active Directory.

`,
			Keywords: []string{
				"users",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_enabled",
					Description: `(Optional) Whether or not the account should be enabled.`,
				},
				resource.Attribute{
					Name:        "age_group",
					Description: `(Optional) The age group of the user. Supported values are ` + "`" + `Adult` + "`" + `, ` + "`" + `NotAdult` + "`" + ` and ` + "`" + `Minor` + "`" + `. Omit this property or specify a blank string to unset.`,
				},
				resource.Attribute{
					Name:        "business_phones",
					Description: `(Optional) A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `(Optional) The city in which the user is located.`,
				},
				resource.Attribute{
					Name:        "company_name",
					Description: `(Optional) The company name which the user is associated. This property can be useful for describing the company that an external user comes from.`,
				},
				resource.Attribute{
					Name:        "consent_provided_for_minor",
					Description: `(Optional) Whether consent has been obtained for minors. Supported values are ` + "`" + `Granted` + "`" + `, ` + "`" + `Denied` + "`" + ` and ` + "`" + `NotRequired` + "`" + `. Omit this property or specify a blank string to unset.`,
				},
				resource.Attribute{
					Name:        "cost_center",
					Description: `(Optional) The cost center associated with the user.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country/region in which the user is located, e.g. ` + "`" + `US` + "`" + ` or ` + "`" + `UK` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `(Optional) The name for the department in which the user works.`,
				},
				resource.Attribute{
					Name:        "disable_password_expiration",
					Description: `(Optional) Whether the user's password is exempt from expiring. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_strong_password",
					Description: `(Optional) Whether the user is allowed weaker passwords than the default policy to be specified. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name to display in the address book for the user.`,
				},
				resource.Attribute{
					Name:        "division",
					Description: `(Optional) The name of the division in which the user works.`,
				},
				resource.Attribute{
					Name:        "employee_id",
					Description: `(Optional) The employee identifier assigned to the user by the organisation.`,
				},
				resource.Attribute{
					Name:        "employee_type",
					Description: `(Optional) Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.`,
				},
				resource.Attribute{
					Name:        "fax_number",
					Description: `(Optional) The fax number of the user.`,
				},
				resource.Attribute{
					Name:        "force_password_change",
					Description: `(Optional) Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "given_name",
					Description: `(Optional) The given name (first name) of the user.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `(Optional) The user’s job title.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `(Optional) The SMTP address for the user. This property cannot be unset once specified.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `(Optional) The mail alias for the user. Defaults to the user name part of the user principal name (UPN).`,
				},
				resource.Attribute{
					Name:        "manager_id",
					Description: `(Optional) The object ID of the user's manager.`,
				},
				resource.Attribute{
					Name:        "mobile_phone",
					Description: `(Optional) The primary cellular telephone number for the user.`,
				},
				resource.Attribute{
					Name:        "office_location",
					Description: `(Optional) The office location in the user's place of business.`,
				},
				resource.Attribute{
					Name:        "onpremises_immutable_id",
					Description: `(Optional) The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's ` + "`" + `user_principal_name` + "`" + ` property when creating a new user account.`,
				},
				resource.Attribute{
					Name:        "other_mails",
					Description: `(Optional) A list of additional email addresses for the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the user. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters. This property is required when creating a new user. ->`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `(Optional) The user's preferred language, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `(Optional) Whether or not the Outlook global address list should include this user. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The state or province in the user's address.`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address of the user's place of business.`,
				},
				resource.Attribute{
					Name:        "surname",
					Description: `(Optional) The user's surname (family name or last name).`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `(Optional) The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: ` + "`" + `NO` + "`" + `, ` + "`" + `JP` + "`" + `, and ` + "`" + `GB` + "`" + `. Cannot be reset to null once set.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `(Required) The user principal name (UPN) of the user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_type",
					Description: `Indicates whether the user account was created as a regular school or work account (` + "`" + `null` + "`" + `), an external account (` + "`" + `Invitation` + "`" + `), a local account for an Azure Active Directory B2C tenant (` + "`" + `LocalAccount` + "`" + `) or self-service sign-up using email verification (` + "`" + `EmailVerified` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "external_user_state",
					Description: `For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are ` + "`" + `PendingAcceptance` + "`" + ` or ` + "`" + `Accepted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "im_addresses",
					Description: `A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_distinguished_name",
					Description: `The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premise SAM account name of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_security_identifier",
					Description: `The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sync_enabled",
					Description: `Whether this user is synchronised from an on-premises directory (` + "`" + `true` + "`" + `), no longer synchronised (` + "`" + `false` + "`" + `), or has never been synchronised (` + "`" + `null` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the user.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the user that direct to the same mailbox.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `The user type in the directory. Possible values are ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `. ## Import Users can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_user.my_user 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_type",
					Description: `Indicates whether the user account was created as a regular school or work account (` + "`" + `null` + "`" + `), an external account (` + "`" + `Invitation` + "`" + `), a local account for an Azure Active Directory B2C tenant (` + "`" + `LocalAccount` + "`" + `) or self-service sign-up using email verification (` + "`" + `EmailVerified` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "external_user_state",
					Description: `For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are ` + "`" + `PendingAcceptance` + "`" + ` or ` + "`" + `Accepted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "im_addresses",
					Description: `A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_distinguished_name",
					Description: `The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premise SAM account name of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_security_identifier",
					Description: `The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sync_enabled",
					Description: `Whether this user is synchronised from an on-premises directory (` + "`" + `true` + "`" + `), no longer synchronised (` + "`" + `false` + "`" + `), or has never been synchronised (` + "`" + `null` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the user.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the user that direct to the same mailbox.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `The user type in the directory. Possible values are ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `. ## Import Users can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_user.my_user 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_user_flow_attribute",
			Category:         "User Flows",
			ShortDescription: ``,
			Description: `

Manages user flow attributes in an Azure Active Directory (Azure AD) tenant.

`,
			Keywords: []string{
				"user",
				"flows",
				"flow",
				"attribute",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data_type",
					Description: `(Required) The data type of the user flow attribute. Possible values are ` + "`" + `boolean` + "`" + `, ` + "`" + `dateTime` + "`" + `, ` + "`" + `int64` + "`" + `, ` + "`" + `string` + "`" + ` or ` + "`" + `stringCollection` + "`" + `. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the user flow attribute that is shown to the user at the time of sign-up.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name of the user flow attribute. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "attribute_type",
					Description: `The type of the user flow attribute. Values include ` + "`" + `builtIn` + "`" + `, ` + "`" + `custom` + "`" + ` or ` + "`" + `required` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An ID used to uniquely identify this user flow attribute. ## Import User flow attributes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_user_flow_attribute.example extension_ecc9f88db2924942b8a96f44873616fe_Hobbyjkorv ` + "`" + `` + "`" + `` + "`" + ` -> This ID can be queried using the [User Flow Attributes API](https://learn.microsoft.com/en-us/graph/api/identityuserflowattribute-list?view=graph-rest-1.0&tabs=http).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "attribute_type",
					Description: `The type of the user flow attribute. Values include ` + "`" + `builtIn` + "`" + `, ` + "`" + `custom` + "`" + ` or ` + "`" + `required` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An ID used to uniquely identify this user flow attribute. ## Import User flow attributes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_user_flow_attribute.example extension_ecc9f88db2924942b8a96f44873616fe_Hobbyjkorv ` + "`" + `` + "`" + `` + "`" + ` -> This ID can be queried using the [User Flow Attributes API](https://learn.microsoft.com/en-us/graph/api/identityuserflowattribute-list?view=graph-rest-1.0&tabs=http).`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"azuread_access_package":                                     0,
		"azuread_access_package_assignment_policy":                   1,
		"azuread_access_package_catalog":                             2,
		"azuread_access_package_catalog_role_assignment":             3,
		"azuread_access_package_resource_catalog_association":        4,
		"azuread_access_package_resource_package_association":        5,
		"azuread_administrative_unit":                                6,
		"azuread_administrative_unit_member":                         7,
		"azuread_administrative_unit_role_member":                    8,
		"azuread_app_role_assignment":                                9,
		"azuread_application":                                        10,
		"azuread_application_certificate":                            11,
		"azuread_application_federated_identity_credential":          12,
		"azuread_application_password":                               13,
		"azuread_application_pre_authorized":                         14,
		"azuread_claims_mapping_policy":                              15,
		"azuread_conditional_access_policy":                          16,
		"azuread_custom_directory_role":                              17,
		"azuread_directory_role":                                     18,
		"azuread_directory_role_assignment":                          19,
		"azuread_directory_role_member":                              20,
		"azuread_group":                                              21,
		"azuread_group_member":                                       22,
		"azuread_invitation":                                         23,
		"azuread_named_location":                                     24,
		"azuread_service_principal":                                  25,
		"azuread_service_principal_certificate":                      26,
		"azuread_service_principal_claims_mapping_policy_assignment": 27,
		"azuread_service_principal_delegated_permission_grant":       28,
		"azuread_service_principal_password":                         29,
		"azuread_service_principal_token_signing_certificate":        30,
		"azuread_synchronization_job":                                31,
		"azuread_synchronization_secret":                             32,
		"azuread_user":                                               33,
		"azuread_user_flow_attribute":                                34,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
