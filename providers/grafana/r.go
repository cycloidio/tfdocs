package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "grafana_alert_notification",
			Category:         "Resources",
			ShortDescription: `The grafana_alert_notification resource allows a Grafana Alert Notification channel to be created.`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"notification",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the alert notification channel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the alert notification channel.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Is this the default channel for all your alerts.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Optional) Additional settings, for full reference lookup [Grafana HTTP API documentation](http://docs.grafana.org/http_api/alerting).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_dashboard",
			Category:         "Resources",
			ShortDescription: `The grafana_dashboard resource allows a Grafana dashboard to be created.`,
			Description:      ``,
			Keywords: []string{
				"dashboard",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "config_json",
					Description: `(Required) The JSON configuration for the dashboard. ## Attributes Reference The resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `A URL "slug" for this dashboard, generated by Grafana by removing certain characters from the dashboard name given as part of the ` + "`" + `config_json` + "`" + ` argument. This can be used to generate the URL for a dashboard. ## Import Existing organizations can be imported using the dashboard "slug" which can be obtained from the url of the dashboard in Grafana (e.g. ` + "`" + `https://grafana.mydomain.com/d/abcdef-gh/dashboard-slug` + "`" + `) ` + "`" + `` + "`" + `` + "`" + ` $ terraform import grafana_dashboard.dashboard_name {dashboard_slug} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "slug",
					Description: `A URL "slug" for this dashboard, generated by Grafana by removing certain characters from the dashboard name given as part of the ` + "`" + `config_json` + "`" + ` argument. This can be used to generate the URL for a dashboard. ## Import Existing organizations can be imported using the dashboard "slug" which can be obtained from the url of the dashboard in Grafana (e.g. ` + "`" + `https://grafana.mydomain.com/d/abcdef-gh/dashboard-slug` + "`" + `) ` + "`" + `` + "`" + `` + "`" + ` $ terraform import grafana_dashboard.dashboard_name {dashboard_slug} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_data_source",
			Category:         "Resources",
			ShortDescription: `The grafana_data_source resource allows a Grafana data source to be created.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The data source type. Must be one of the data source keywords supported by the Grafana server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the data source within the Grafana server.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) The URL for the data source. The type of URL required varies depending on the chosen data source type.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) If true, the data source will be the default source used by the Grafana server. Only one data source on a server can be the default.`,
				},
				resource.Attribute{
					Name:        "basic_auth_enabled",
					Description: `(Optional) - If true, HTTP basic authentication will be used to make requests.`,
				},
				resource.Attribute{
					Name:        "basic_auth_username",
					Description: `(Required if ` + "`" + `basic_auth_enabled` + "`" + ` is true) The username to use for basic auth.`,
				},
				resource.Attribute{
					Name:        "basic_auth_password",
					Description: `(Required if ` + "`" + `basic_auth_enabled` + "`" + ` is true) The password to use for basic auth.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required by some data source types) The username to use to authenticate to the data source.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required by some data source types) The password to use to authenticate to the data source.`,
				},
				resource.Attribute{
					Name:        "json_data",
					Description: `(Required by some data source types) The default region and authentication type to access the data source. ` + "`" + `json_data` + "`" + ` is documented in more detail below.`,
				},
				resource.Attribute{
					Name:        "secure_json_data",
					Description: `(Required by some data source types) The access and secret keys required to access the data source. ` + "`" + `secure_json_data` + "`" + ` is documented in more detail below.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required by some data source types) The name of the database to use on the selected data source server.`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) The method by which the browser-based Grafana application will access the data source. The default is "proxy", which means that the application will make requests via a proxy endpoint on the Grafana server. JSON Data (` + "`" + `json_data` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Required by some data source types) The authentication type type used to access the data source.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required by some data source types) The default region for the data source.`,
				},
				resource.Attribute{
					Name:        "custom_metrics_namespaces",
					Description: `(Optional, for the CloudWatch data source type) A comma-separated list of custom namespaces to be queried by the CloudWatch data source.`,
				},
				resource.Attribute{
					Name:        "assume_role_arn",
					Description: `(Optional, for the CloudWatch data source type) The role ARN to be assumed by Grafana when using the CloudWatch data source. Secure JSON Data (` + "`" + `secure_json_data` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required by some data source types) The access key required to access the data source.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required by some data source types) The secret key required to access the data source. ## Attributes Reference The resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The opaque unique id assigned to the data source by the Grafana server.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The opaque unique id assigned to the data source by the Grafana server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_folder",
			Category:         "Resources",
			ShortDescription: `The grafana_folder resource allows a Grafana folder to be created.`,
			Description:      ``,
			Keywords: []string{
				"folder",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The title of the folder. ## Attributes Reference The resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The internal id of the folder in Grafana (only guaranteed to be unique within this Grafana instance). The ` + "`" + `id` + "`" + ` is used by the ` + "`" + `grafana_dashboard` + "`" + ` resource to place a dashboard within a folder.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `An external id of the folder in Grafana (stable when folders are migrated between Grafana instances). The ` + "`" + `uid` + "`" + ` is required by several Grafana Folder APIs. ## Import Folders cannot be imported.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The internal id of the folder in Grafana (only guaranteed to be unique within this Grafana instance). The ` + "`" + `id` + "`" + ` is used by the ` + "`" + `grafana_dashboard` + "`" + ` resource to place a dashboard within a folder.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `An external id of the folder in Grafana (stable when folders are migrated between Grafana instances). The ` + "`" + `uid` + "`" + ` is required by several Grafana Folder APIs. ## Import Folders cannot be imported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_organization",
			Category:         "Resources",
			ShortDescription: `The grafana_organization resource allows a Grafana organization to be created.`,
			Description:      ``,
			Keywords: []string{
				"organization",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The display name for the Grafana organization created.`,
				},
				resource.Attribute{
					Name:        "admin_user",
					Description: `(Optional) The login name of the configured [default admin user](http://docs.grafana.org/installation/configuration/#admin-user) for the Grafana installation. If unset, this value defaults to ` + "`" + `admin` + "`" + `, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and this parameter keeps Terraform from removing it from organizations.`,
				},
				resource.Attribute{
					Name:        "create_users",
					Description: `(Optional) Whether or not to create Grafana users specified in the organization's membership if they don't already exist in Grafana. If unspecified, this parameter defaults to ` + "`" + `true` + "`" + `, creating placeholder users with the ` + "`" + `name` + "`" + `, ` + "`" + `login` + "`" + `, and ` + "`" + `email` + "`" + ` set to the email of the user, and a random password. Setting this option to ` + "`" + `false` + "`" + ` will cause an error to be thrown for any users that do not already exist in Grafana. This option is particularly useful when integrating Grafana with external authentication services such as [` + "`" + `auth.github` + "`" + `](http://docs.grafana.org/installation/configuration/#auth-github) and [` + "`" + `auth.google` + "`" + `](http://docs.grafana.org/installation/configuration/#auth-google).`,
				},
				resource.Attribute{
					Name:        "admins",
					Description: `(Optional) A list of email addresses corresponding to users who should be given ` + "`" + `admin` + "`" + ` access to the organization. Note: users specified here must already exist in Grafana unless 'create_users' is set to true.`,
				},
				resource.Attribute{
					Name:        "editors",
					Description: `(Optional) A list of email addresses corresponding to users who should be given ` + "`" + `editor` + "`" + ` access to the organization. Note: users specified here must already exist in Grafana unless 'create_users' is set to true.`,
				},
				resource.Attribute{
					Name:        "viewers",
					Description: `(Optional) A list of email addresses corresponding to users who should be given ` + "`" + `viewer` + "`" + ` access to the organization. Note: users specified here must already exist in Grafana unless 'create_users' is set to true. A user can only be listed under one role-group for an organization, listing the same user under multiple roles will cause an error to be thrown. Note - Users specified for each role-group (` + "`" + `admins` + "`" + `, ` + "`" + `editors` + "`" + `, ` + "`" + `viewers` + "`" + `) should be listed in ascending alphabetical order (A-Z). By defining users in alphabetical order, Terraform is prevented from detecting unnecessary changes when comparing the list of defined users in the resource to the (ordered) list returned by the Grafana API. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The organization id assigned to this organization by Grafana. ## Import Existing organizations can be imported using the organization id obtained from the Grafana Web UI under 'Server Admin'. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import grafana_organization.org_name {org_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "org_id",
					Description: `The organization id assigned to this organization by Grafana. ## Import Existing organizations can be imported using the organization id obtained from the Grafana Web UI under 'Server Admin'. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import grafana_organization.org_name {org_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"grafana_alert_notification": 0,
		"grafana_dashboard":          1,
		"grafana_data_source":        2,
		"grafana_folder":             3,
		"grafana_organization":       4,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
