package britive

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

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
			Type:             "britive_profile_identity",
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
					Name:        "username",
					Description: `(Required, Forces new resource) The name of the identity.`,
				},
				resource.Attribute{
					Name:        "access_period",
					Description: `(Optional) The access period of the identity in a profile. The format of an ` + "`" + `access_period` + "`" + ` is documented below. ### ` + "`" + `access_period` + "`" + ` block supports`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The start of the access period for the associated identity.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) The end of the access period for the associated identity. ## Attribute Reference In addition to the above argument, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with the format ` + "`" + `paps/{{profileID}}/users/{{userID}}` + "`" + ` ## Import You can import a Britive profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_profile_identity.new apps/{{app_name}}/paps/{{profile_name}}/users/{{username}} terraform import britive_profile_identity.new {{app_name}}/{{profile_name}}/{{username}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with the format ` + "`" + `paps/{{profileID}}/users/{{userID}}` + "`" + ` ## Import You can import a Britive profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `sh terraform import britive_profile_identity.new apps/{{app_name}}/paps/{{profile_name}}/users/{{username}} terraform import britive_profile_identity.new {{app_name}}/{{profile_name}}/{{username}} ` + "`" + `` + "`" + `` + "`" + ``,
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
			Type:             "britive_profile_tag",
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
					Name:        "tag_name",
					Description: `(Required, Forces new resource) The name of the tag.`,
				},
				resource.Attribute{
					Name:        "access_period",
					Description: `(Optional) The access period of the tag in the Britive profile. The format of ` + "`" + `access_period` + "`" + ` is documented below. ### ` + "`" + `access_period` + "`" + ` block supports`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The start of the access period for the associated tag.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) The end of the access period for the associated tag. ## Attribute Reference In addition to the above arguments, the following attribute is exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with format ` + "`" + `paps/{{profileID}}/tags/{{tagID}}` + "`" + ` ## Import You can import a Britive profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_profile_tag.new apps/{{app_name}}/paps/{{profile_name}}/tags/{{tag_name}} terraform import britive_profile_tag.new {{app_name}}/{{profile_name}}/{{tag_name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier of the resource with format ` + "`" + `paps/{{profileID}}/tags/{{tagID}}` + "`" + ` ## Import You can import a Britive profile using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + `SH terraform import britive_profile_tag.new apps/{{app_name}}/paps/{{profile_name}}/tags/{{tag_name}} terraform import britive_profile_tag.new {{app_name}}/{{profile_name}}/{{tag_name}} ` + "`" + `` + "`" + `` + "`" + ``,
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

		"britive_profile":                   0,
		"britive_profile_identity":          1,
		"britive_profile_permission":        2,
		"britive_profile_session_attribute": 3,
		"britive_profile_tag":               4,
		"britive_tag":                       5,
		"britive_tag_member":                6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
