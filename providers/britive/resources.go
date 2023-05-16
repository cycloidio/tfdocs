package britive

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "britive_permission",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the permission.`,
				},
				resource.Attribute{
					Name:        "consumer",
					Description: `(Required) A component/entity that will use the policy engine for access decisions.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) A resource in the definition of the corresponding consumer, or '`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `(Required) A set of pre-defined actions for each consumer. ## Attribute Reference In addition to the above arguments, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with format ` + "`" + `permissions/{{permissionID}}` + "`" + ` ## Import You can import a permission using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_permission.new permissions/{{name}} terraform import britive_permission.new {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with format ` + "`" + `permissions/{{permissionID}}` + "`" + ` ## Import You can import a permission using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_permission.new permissions/{{name}} terraform import britive_permission.new {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "britive_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the policy.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Type of access the policy provides. This can have two values "Allow"/"Deny". Defaults to "Allow".`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Set of members under this policy. This is a JSON formatted string. Includes the usernames of ` + "`" + `serviceIdentities` + "`" + `, ` + "`" + `tags` + "`" + `, ` + "`" + `tokens` + "`" + ` and ` + "`" + `users` + "`" + ``,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions associated to the policy. Either a role/permission is to be assigned to a policy.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) Roles associated to the policy. Either a role/permission is to be assigned to a policy.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) Set of conditions applied to this policy. This is a JSON formatted string. Includes the username for ` + "`" + `tags` + "`" + ` and ` + "`" + `userIds` + "`" + ` under ` + "`" + `approvers` + "`" + `. The ` + "`" + `approval` + "`" + ` block also includes the ` + "`" + `notificationMedium` + "`" + `. The ` + "`" + `timeToApprove` + "`" + ` is provided in minutes, ` + "`" + `validFor` + "`" + ` can be provided in days or minutes, depending on ` + "`" + `isValidForInDays` + "`" + ` boolean value being set to true or false respectively. The condition based on ` + "`" + `ipAddress` + "`" + ` should be specified as comma separated IP addresses in CIDR or dotted decimal format. The ` + "`" + `timeOfAccess` + "`" + ` can be scheduled based on date, days, both or ` + "`" + `null` + "`" + `. The ` + "`" + `dateSchedule` + "`" + ` should contain the ` + "`" + `fromDate` + "`" + `, ` + "`" + `toDate` + "`" + ` in format of "YYYY-MM-DD HH:MM:SS" and ` + "`" + `timezone` + "`" + ` as a string from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. If ` + "`" + `dateSchedule` + "`" + ` is not required, it has to be set to ` + "`" + `null` + "`" + `. The ` + "`" + `daysSchedule` + "`" + ` should contain the ` + "`" + `fromTime` + "`" + `, ` + "`" + `toTime` + "`" + ` in format of "HH:MM:SS", ` + "`" + `timezone` + "`" + ` as a string from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones and ` + "`" + `days` + "`" + ` as a list of strings. If ` + "`" + `daysSchedule` + "`" + ` is not required, it has to be set to ` + "`" + `null` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `(Optional) Indicates if a policy is active. Boolean value accepts true/false. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "is_draft",
					Description: `(Optional) Indicates if a policy is a draft. Boolean value accepts true/false. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "is_read_only",
					Description: `(Optional) Indicates if a policy is read only. Boolean value accepts true/false. Defaults to false. ## Attribute Reference In addition to the above arguments, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the policy with format ` + "`" + `policies/{{name}}` + "`" + ` ## Import You can import a policy using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_policy.new policies/{{name}} terraform import britive_policy.new {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the policy with format ` + "`" + `policies/{{name}}` + "`" + ` ## Import You can import a policy using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_policy.new policies/{{name}} terraform import britive_policy.new {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "britive_profile",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_container_id",
					Description: `(Required, Forces new resource) The identity of the Britive application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Britive profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the Britive profile.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) The status of the Britive profile. By default, the Britive profile is enabled. To disable a Britive profile, set ` + "`" + `disabled = true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expiration_duration",
					Description: `(Required) The expiration time for the Britive profile. For example, ` + "`" + `25m0s` + "`" + ``,
				},
				resource.Attribute{
					Name:        "extendable",
					Description: `(Optional) The Boolean flag that indicates whether profile expiry is extendable or not. The default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_prior_to_expiration",
					Description: `(Optional) The Britive profile expiry notification as a time value. For example, ` + "`" + `10m0s` + "`" + ``,
				},
				resource.Attribute{
					Name:        "extension_duration",
					Description: `(Optional) The Britive profile expiry extension duration as a time value. For example: ` + "`" + `12m30s` + "`" + ``,
				},
				resource.Attribute{
					Name:        "extension_limit",
					Description: `(Optional) The Britive profile expiry extension limit. For example: ` + "`" + `2` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination_url",
					Description: `(Optional) The console URL where the user will be redirected upon checking out the profile. For example: ` + "`" + `https://console.aws.amazon.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "associations",
					Description: `(Required) The list of associations for the Britive profile. The format of ` + "`" + `associations` + "`" + ` is documented below. ### ` + "`" + `associations` + "`" + ` block supports`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of association, should be one of [Environment, EnvironmentGroup, ApplicationResource].`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The association value.`,
				},
				resource.Attribute{
					Name:        "parent_name",
					Description: `(Optional) The parent name of the resource. Required only if the association type is ApplicationResource. ## Attribute Reference In addition to the above arguments, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The identity of the Britive profile. ## Import You can import a profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_profile.new apps/{{app_name}}/paps/{{name}} terraform import britive_profile.new {{app_name}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The identity of the Britive profile. ## Import You can import a profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_profile.new apps/{{app_name}}/paps/{{name}} terraform import britive_profile.new {{app_name}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "britive_profile_permission",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Required, Forces new resource) The identifier of the profile.`,
				},
				resource.Attribute{
					Name:        "permission_name",
					Description: `(Required, Forces new resource) The name of permission.`,
				},
				resource.Attribute{
					Name:        "permission_type",
					Description: `(Required, Forces new resource) The type of permission. ## Attribute Reference In addition to the above arguments, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with the format ` + "`" + `paps/{{profileID}}/permissions/{{permission_name}}/type/{{permission_type}}` + "`" + ` ## Import You can import a Britive profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_profile_permission.new apps/{{app_name}}/paps/{{profile_name}}/permissions/{{permission_name}}/type/{{permission_type}} terraform import britive_profile_permission.new {{app_name}}/{{profile_name}}/{{permission_name}}/{{permission_type}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with the format ` + "`" + `paps/{{profileID}}/permissions/{{permission_name}}/type/{{permission_type}}` + "`" + ` ## Import You can import a Britive profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_profile_permission.new apps/{{app_name}}/paps/{{profile_name}}/permissions/{{permission_name}}/type/{{permission_type}} terraform import britive_profile_permission.new {{app_name}}/{{profile_name}}/{{permission_name}}/{{permission_type}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "britive_profile_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Required, Forces new resource) The identifier of the profile.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) The name of the profile-policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the profile-policy.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Set of members under this policy. This is a JSON formatted string. Includes the usernames of ` + "`" + `serviceIdentities` + "`" + `, ` + "`" + `tags` + "`" + ` and ` + "`" + `users` + "`" + ``,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) Set of conditions applied to this policy. This is a JSON formatted string. Includes the username for ` + "`" + `tags` + "`" + ` and ` + "`" + `userIds` + "`" + ` under ` + "`" + `approvers` + "`" + `. The ` + "`" + `approval` + "`" + ` block also includes the ` + "`" + `notificationMedium` + "`" + ` and ` + "`" + `timeToApprove` + "`" + ` in minutes, ` + "`" + `validFor` + "`" + ` can be provided in days or minutes, depending on ` + "`" + `isValidForInDays` + "`" + ` boolean value being set to true or false respectively. The condition based on ` + "`" + `ipAddress` + "`" + ` should be specified as comma separated IP addresses in CIDR or dotted decimal format. The ` + "`" + `timeOfAccess` + "`" + ` can be scheduled based on date, days, both or ` + "`" + `null` + "`" + `. The ` + "`" + `dateSchedule` + "`" + ` should contain the ` + "`" + `fromDate` + "`" + `, ` + "`" + `toDate` + "`" + ` in format of "YYYY-MM-DD HH:MM:SS" and ` + "`" + `timezone` + "`" + ` as a string from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. If ` + "`" + `dateSchedule` + "`" + ` is not required, it has to be set to ` + "`" + `null` + "`" + `. The ` + "`" + `daysSchedule` + "`" + ` should contain the ` + "`" + `fromTime` + "`" + `, ` + "`" + `toTime` + "`" + ` in format of "HH:MM:SS", ` + "`" + `timezone` + "`" + ` as a string from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones and ` + "`" + `days` + "`" + ` as a list of strings. If ` + "`" + `daysSchedule` + "`" + ` is not required, it has to be set to ` + "`" + `null` + "`" + ``,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Type of access the policy provides. This can have two values "Allow"/"Deny". Defaults to "Allow".`,
				},
				resource.Attribute{
					Name:        "consumer",
					Description: `(Optional) A component/entity that will use the policy engine for access decisions. Defaults to "papservice". Do not provide any other value.`,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `(Optional) Indicates if a policy is active. Boolean value accepts true/false. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "is_draft",
					Description: `(Optional) Indicates if a policy is a draft. Boolean value accepts true/false. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "is_read_only",
					Description: `(Optional) Indicates if a policy is read only. Boolean value accepts true/false. Defaults to false. ## Attribute Reference In addition to the above arguments, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the policy for the profile with format ` + "`" + `paps/{{profile_id}}/policies/{{policy_name}}` + "`" + ` ## Import You can import a policy for the profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_profile_policy.new paps/{{profile_id}}/policies/{{policy_name}} terraform import britive_profile_policy.new {{profile_id}}/{{policy_name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the policy for the profile with format ` + "`" + `paps/{{profile_id}}/policies/{{policy_name}}` + "`" + ` ## Import You can import a policy for the profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_profile_policy.new paps/{{profile_id}}/policies/{{policy_name}} terraform import britive_profile_policy.new {{profile_id}}/{{policy_name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "britive_profile_session_attribute",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Required, Forces new resource) The identifier of the profile.`,
				},
				resource.Attribute{
					Name:        "attribute_type",
					Description: `(Optional, Forces new resource) The type of attribute, should be one of [Static, Identity]. The default value is ` + "`" + `Identity` + "`" + ``,
				},
				resource.Attribute{
					Name:        "attribute_name",
					Description: `(Optional, Required when ` + "`" + `attribute_type` + "`" + ` is Identity, Forces new resource) The name of attribute.`,
				},
				resource.Attribute{
					Name:        "attribute_value",
					Description: `(Optional, Required when ` + "`" + `attribute_type` + "`" + ` is Static) The value of attribute.`,
				},
				resource.Attribute{
					Name:        "mapping_name",
					Description: `(Required) The name for attribute mapping.`,
				},
				resource.Attribute{
					Name:        "transitive",
					Description: `(Optional) The Boolean flag that indicates whether the attribute is transitive or not. The default value is ` + "`" + `false` + "`" + ` ## Attribute Reference In addition to the above arguments, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with the format ` + "`" + `paps/{{profileID}}/session-attributes/{{sessionAttributeID}}` + "`" + ` ## Import You can import a Britive profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_profile_session_attribute.new apps/{{app_name}}/paps/{{profile_name}}/session-attributes/type/{{attribute_type}}/mapping-name/{{mapping_name}} terraform import britive_profile_session_attribute.new {{app_name}}/{{profile_name}}/{{attribute_type}}/{{mapping_name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with the format ` + "`" + `paps/{{profileID}}/session-attributes/{{sessionAttributeID}}` + "`" + ` ## Import You can import a Britive profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_profile_session_attribute.new apps/{{app_name}}/paps/{{profile_name}}/session-attributes/type/{{attribute_type}}/mapping-name/{{mapping_name}} terraform import britive_profile_session_attribute.new {{app_name}}/{{profile_name}}/{{attribute_type}}/{{mapping_name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "britive_role",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the role.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) Names of the permissions associated to the role. This is a JSON formatted string. Mimimum of 1 permission is required to create a role. ` + "`" + `name` + "`" + ` corresponds to the name of the permission ## Attribute Reference In addition to the above arguments, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with format ` + "`" + `roles/{{roleID}}` + "`" + ` ## Import You can import a role using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_role.new roles/{{name}} terraform import britive_role.new {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with format ` + "`" + `roles/{{roleID}}` + "`" + ` ## Import You can import a role using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_role.new roles/{{name}} terraform import britive_role.new {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "britive_tag",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of Britive tag.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the Britive tag.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `(Required) The unique identity of the identity provider associated with the Britive tag.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) The status of the Britive tag. By default, the Britive tag is enabled. To disable a Britive tag, set ` + "`" + `disabled = true` + "`" + `. ## Attribute Reference In addition to the above arguments, the following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The identity of the Britive tag.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `The boolean attribute that indicates whether the tag is external or not. ## Import You can import the Britive tag using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_tag.new tags/{{tag_name}} terraform import britive_tag.new {{tag_name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The identity of the Britive tag.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `The boolean attribute that indicates whether the tag is external or not. ## Import You can import the Britive tag using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_tag.new tags/{{tag_name}} terraform import britive_tag.new {{tag_name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "britive_tag_member",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tag_id",
					Description: `(Required, Forces new resource) The identifier of the Britive tag.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required, Forces new resource) The username of the user added to the Britive tag. ## Attribute Reference In addition to the above arguments, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with format ` + "`" + `tags/{{tagID}}/users/{{userID}}` + "`" + ` ## Import You can import the Britive tag member using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_tag_member.new tags/{{tag_name}}/users/{{username}} terraform import britive_tag_member.new {{tag_name}}/{{username}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with format ` + "`" + `tags/{{tagID}}/users/{{userID}}` + "`" + ` ## Import You can import the Britive tag member using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_tag_member.new tags/{{tag_name}}/users/{{username}} terraform import britive_tag_member.new {{tag_name}}/{{username}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"britive_permission":                0,
		"britive_policy":                    1,
		"britive_profile":                   2,
		"britive_profile_permission":        3,
		"britive_profile_policy":            4,
		"britive_profile_session_attribute": 5,
		"britive_role":                      6,
		"britive_tag":                       7,
		"britive_tag_member":                8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
