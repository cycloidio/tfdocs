package statuscake

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "statuscake_test",
			Category:         "Resources",
			ShortDescription: `The statuscake_test resource allows StatusCake tests to be managed by Terraform.`,
			Description:      ``,
			Keywords: []string{
				"test",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "website_name",
					Description: `(Required) This is the name of the test and the website to be monitored.`,
				},
				resource.Attribute{
					Name:        "website_url",
					Description: `(Required) The URL of the website to be monitored`,
				},
				resource.Attribute{
					Name:        "check_rate",
					Description: `(Optional) Test check rate in seconds. Defaults to 300`,
				},
				resource.Attribute{
					Name:        "contact_group",
					Description: `(Optional) Set test contact groups, must be array of strings.`,
				},
				resource.Attribute{
					Name:        "test_type",
					Description: `(Required) The type of Test. Either HTTP, TCP, PING, or DNS`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `(Optional) Whether or not the test is paused. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout of the test in seconds.`,
				},
				resource.Attribute{
					Name:        "confirmations",
					Description: `(Optional) The number of confirmation servers to use in order to detect downtime. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port to use when specifying a TCP test.`,
				},
				resource.Attribute{
					Name:        "trigger_rate",
					Description: `(Optional) The number of minutes to wait before sending an alert. Default is ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_header",
					Description: `(Optional) Custom HTTP header, must be supplied as JSON.`,
				},
				resource.Attribute{
					Name:        "user_agent",
					Description: `(Optional) Test with a custom user agent set.`,
				},
				resource.Attribute{
					Name:        "node_locations",
					Description: `(Optional) Set test node locations, must be array of strings.`,
				},
				resource.Attribute{
					Name:        "ping_url",
					Description: `(Optional) A URL to ping if a site goes down.`,
				},
				resource.Attribute{
					Name:        "basic_user",
					Description: `(Optional) A Basic Auth User account to use to login`,
				},
				resource.Attribute{
					Name:        "basic_pass",
					Description: `(Optional) If BasicUser is set then this should be the password for the BasicUser.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(Optional) Set 1 to enable public reporting, 0 to disable.`,
				},
				resource.Attribute{
					Name:        "logo_image",
					Description: `(Optional) A URL to a image to use for public reporting.`,
				},
				resource.Attribute{
					Name:        "branding",
					Description: `(Optional) Set to 0 to use branding (default) or 1 to disable public reporting branding).`,
				},
				resource.Attribute{
					Name:        "website_host",
					Description: `(Optional) Used internally, when possible please add.`,
				},
				resource.Attribute{
					Name:        "virus",
					Description: `(Optional) Enable virus checking or not. 1 to enable`,
				},
				resource.Attribute{
					Name:        "find_string",
					Description: `(Optional) A string that should either be found or not found.`,
				},
				resource.Attribute{
					Name:        "do_not_find",
					Description: `(Optional) If the above string should be found to trigger a alert. 1 = will trigger if find_string found.`,
				},
				resource.Attribute{
					Name:        "real_browser",
					Description: `(Optional) Use 1 to TURN OFF real browser testing.`,
				},
				resource.Attribute{
					Name:        "test_tags",
					Description: `(Optional) Set test tags, must be array of strings.`,
				},
				resource.Attribute{
					Name:        "status_codes",
					Description: `(Optional) Comma Separated List of StatusCodes to Trigger Error on. Defaults are "204, 205, 206, 303, 400, 401, 403, 404, 405, 406, 408, 410, 413, 444, 429, 494, 495, 496, 499, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 510, 511, 521, 522, 523, 524, 520, 598, 599".`,
				},
				resource.Attribute{
					Name:        "use_jar",
					Description: `(Optional) Set to true to enable the Cookie Jar. Required for some redirects. Default is false.`,
				},
				resource.Attribute{
					Name:        "post_raw",
					Description: `(Optional) Use to populate the RAW POST data field on the test.`,
				},
				resource.Attribute{
					Name:        "final_endpoint",
					Description: `(Optional) Use to specify the expected Final URL in the testing process.`,
				},
				resource.Attribute{
					Name:        "enable_ssl_alert",
					Description: `(Optional) HTTP Tests only. If enabled, tests will send warnings if the SSL certificate is about to expire. Paid users only. Default is false`,
				},
				resource.Attribute{
					Name:        "follow_redirect",
					Description: `(Optional) Use to specify whether redirects should be followed, set to true to enable. Default is false. ## Attributes Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "test_id",
					Description: `A unique identifier for the test. ## Import StatusCake test can be imported using the test id, e.g. ` + "`" + `` + "`" + `` + "`" + ` tf import statuscake_test.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "test_id",
					Description: `A unique identifier for the test. ## Import StatusCake test can be imported using the test id, e.g. ` + "`" + `` + "`" + `` + "`" + ` tf import statuscake_test.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"statuscake_test": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
