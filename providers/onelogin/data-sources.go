package onelogin

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_user",
			Category:         "Data Sources",
			ShortDescription: `Returns User resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `The user's username.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user's ID. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The user's id`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The user's email.`,
				},
				resource.Attribute{
					Name:        "firstname",
					Description: `The user's first name`,
				},
				resource.Attribute{
					Name:        "lastname",
					Description: `The user's last name`,
				},
				resource.Attribute{
					Name:        "distinguished_name",
					Description: `The user's distinguished name`,
				},
				resource.Attribute{
					Name:        "samaccountname",
					Description: `The user's samaccount name`,
				},
				resource.Attribute{
					Name:        "userprincipalname",
					Description: `The user's user principal name`,
				},
				resource.Attribute{
					Name:        "member_of",
					Description: `The user's member_of`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `The user's phone number`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The user's title`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `The user's company`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `The user's department`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `A comment about the user`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The user's state. Must be one of ` + "`" + `0: Unapproved` + "`" + ` ` + "`" + `1: Approved` + "`" + ` ` + "`" + `2: Rejected` + "`" + ` ` + "`" + `3: Unlicensed` + "`" + ``,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The user's status. Must be one of ` + "`" + `0: Unactivated` + "`" + ` ` + "`" + `1: Active` + "`" + ` ` + "`" + `2: Suspended` + "`" + ` ` + "`" + `3: Locked` + "`" + ` ` + "`" + `4: Password expired` + "`" + ` ` + "`" + `5: Awaiting password reset` + "`" + ` ` + "`" + `7: Password Pending` + "`" + ` ` + "`" + `8: Security questions required` + "`" + ``,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The user's group_id`,
				},
				resource.Attribute{
					Name:        "directory_id",
					Description: `The user's directory_id`,
				},
				resource.Attribute{
					Name:        "trusted_idp_id",
					Description: `The user's trusted_idp_id`,
				},
				resource.Attribute{
					Name:        "manager_ad_id",
					Description: `The user's manager_ad_id`,
				},
				resource.Attribute{
					Name:        "manager_user_id",
					Description: `The user's manager_user_id`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `The user's external_id`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The user's id`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The user's email.`,
				},
				resource.Attribute{
					Name:        "firstname",
					Description: `The user's first name`,
				},
				resource.Attribute{
					Name:        "lastname",
					Description: `The user's last name`,
				},
				resource.Attribute{
					Name:        "distinguished_name",
					Description: `The user's distinguished name`,
				},
				resource.Attribute{
					Name:        "samaccountname",
					Description: `The user's samaccount name`,
				},
				resource.Attribute{
					Name:        "userprincipalname",
					Description: `The user's user principal name`,
				},
				resource.Attribute{
					Name:        "member_of",
					Description: `The user's member_of`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `The user's phone number`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The user's title`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `The user's company`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `The user's department`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `A comment about the user`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The user's state. Must be one of ` + "`" + `0: Unapproved` + "`" + ` ` + "`" + `1: Approved` + "`" + ` ` + "`" + `2: Rejected` + "`" + ` ` + "`" + `3: Unlicensed` + "`" + ``,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The user's status. Must be one of ` + "`" + `0: Unactivated` + "`" + ` ` + "`" + `1: Active` + "`" + ` ` + "`" + `2: Suspended` + "`" + ` ` + "`" + `3: Locked` + "`" + ` ` + "`" + `4: Password expired` + "`" + ` ` + "`" + `5: Awaiting password reset` + "`" + ` ` + "`" + `7: Password Pending` + "`" + ` ` + "`" + `8: Security questions required` + "`" + ``,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The user's group_id`,
				},
				resource.Attribute{
					Name:        "directory_id",
					Description: `The user's directory_id`,
				},
				resource.Attribute{
					Name:        "trusted_idp_id",
					Description: `The user's trusted_idp_id`,
				},
				resource.Attribute{
					Name:        "manager_ad_id",
					Description: `The user's manager_ad_id`,
				},
				resource.Attribute{
					Name:        "manager_user_id",
					Description: `The user's manager_user_id`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `The user's external_id`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_users",
			Category:         "Data Sources",
			ShortDescription: `Returns User IDs matching the given attributes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `The user's username.`,
				},
				resource.Attribute{
					Name:        "firstname",
					Description: `The user's first name`,
				},
				resource.Attribute{
					Name:        "lastname",
					Description: `The user's last name`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The user's email.`,
				},
				resource.Attribute{
					Name:        "samaccountname",
					Description: `The user's samaccount name`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `The user's external_id`,
				},
				resource.Attribute{
					Name:        "directory_id",
					Description: `The user's directory_id ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `List of user's id`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `List of user's id`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"onelogin_onelogin_user":  0,
		"onelogin_onelogin_users": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
