package runscope

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "runscope_bucket",
			Category:         "Resources",
			ShortDescription: `Provides a Runscope bucket resource.`,
			Description:      ``,
			Keywords: []string{
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(String, Required) The name of this bucket.`,
				},
				resource.Attribute{
					Name:        "team_uuid",
					Description: `(String, Required) Unique identifier for the team this bucket is being created for. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this bucket.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this bucket.`,
				},
				resource.Attribute{
					Name:        "team_uuid",
					Description: `Unique identifier for the team this bucket belongs to. ## Import Buckets can be imported using the bucket ` + "`" + `key` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import runscope_bucket.example t2f4bkvnggcx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of this bucket.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this bucket.`,
				},
				resource.Attribute{
					Name:        "team_uuid",
					Description: `Unique identifier for the team this bucket belongs to. ## Import Buckets can be imported using the bucket ` + "`" + `key` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import runscope_bucket.example t2f4bkvnggcx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "runscope_environment",
			Category:         "Resources",
			ShortDescription: `Provides a Runscope environment resource.`,
			Description:      ``,
			Keywords: []string{
				"environment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_id",
					Description: `(Required) The id of the bucket to associate this environment with.`,
				},
				resource.Attribute{
					Name:        "test_id",
					Description: `(Optional) The id of the test to associate this environment with. If given, creates a test specific environment, otherwise creates a shared environment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of environment.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Optional) The [script](https://www.runscope.com/docs/api-testing/scripts#initial-script) to to run to setup the environment`,
				},
				resource.Attribute{
					Name:        "preserve_cookies",
					Description: `(Optional) If this is set to true, tests using this enviornment will manage cookies between steps.`,
				},
				resource.Attribute{
					Name:        "initial_variables",
					Description: `(Optional) Map of keys and values being used for variables when the test begins.`,
				},
				resource.Attribute{
					Name:        "integrations",
					Description: `(Optional) A list of integration ids to enable for test runs using this environment.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) A list of [Runscope regions](https://www.runscope.com/docs/regions) to execute test runs in when using this environment.`,
				},
				resource.Attribute{
					Name:        "remote_agents",
					Description: `(Optional) A list of [Remote Agents](https://www.runscope.com/docs/api/agents) to execute test runs in when using this environment. Remote Agents documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the remote agent`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) The uuid of the remote agent Emails (` + "`" + `emails` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "notify_all",
					Description: `(Required) Send an email to all team members according to the ` + "`" + `notify_on` + "`" + ` rules.`,
				},
				resource.Attribute{
					Name:        "notify_on",
					Description: `(Required) Upon completion of a test run Runscope will send email notifications, allowed values: ` + "`" + `all` + "`" + `, ` + "`" + `failures` + "`" + `, ` + "`" + `threshold` + "`" + ` or ` + "`" + `switch` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the person.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique identifier for this person's account.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The email address for this account. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the environment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "runscope_schedule",
			Category:         "Resources",
			ShortDescription: `Provides a Runscope schedule resource.`,
			Description:      ``,
			Keywords: []string{
				"schedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_id",
					Description: `(Required) The id of the bucket to associate this schedule with.`,
				},
				resource.Attribute{
					Name:        "test_id",
					Description: `(Required) The id of the test to associate this schedule with.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The id of the environment to use when running the test. If given, creates a test specific schedule, otherwise creates a shared schedule.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Required) The schedule's interval, must be one of:`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) A human-friendly description for the schedule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "runscope_step",
			Category:         "Resources",
			ShortDescription: `Provides a Runscope step resource.`,
			Description:      ``,
			Keywords: []string{
				"step",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_id",
					Description: `(Required) The id of the bucket to associate this step with.`,
				},
				resource.Attribute{
					Name:        "test_id",
					Description: `(Required) The id of the test to associate this step with.`,
				},
				resource.Attribute{
					Name:        "step_type",
					Description: `(Required) The type of step.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) The HTTP method for this request step.`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `(Optional) A list of variables to extract out of the HTTP response from this request. Variables documented below.`,
				},
				resource.Attribute{
					Name:        "assertions",
					Description: `(Optional) A list of assertions to apply to the HTTP response from this request. Assertions documented below.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) A list of headers to apply to the request. Headers documented below.`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `(Optional) A string to use as the body of the request.`,
				},
				resource.Attribute{
					Name:        "auth",
					Description: `(Optional) The credentials used to authenticate the request`,
				},
				resource.Attribute{
					Name:        "before_script",
					Description: `(Optional) Runs a script before the request is made`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Optional) Runs a script after the request is made Variables (` + "`" + `variables` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the variable to define.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `(Required) The name of the source property. i.e. header name or json path`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The variable source, for list of allowed values see: https://www.runscope.com/docs/api/steps#assertions Assertions (` + "`" + `assertions` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The assertion source, for list of allowed values see: https://www.runscope.com/docs/api/steps#assertions`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `(Optional) The name of the source property. i.e. header name or json path`,
				},
				resource.Attribute{
					Name:        "comparison",
					Description: `(Required) The assertion comparison to make i.e. ` + "`" + `equals` + "`" + `, for list of allowed values see: https://www.runscope.com/docs/api/steps#assertions`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value the ` + "`" + `comparison` + "`" + ` will use`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Required) The name of the header`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The name header value ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the step.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the step.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "runscope_test",
			Category:         "Resources",
			ShortDescription: `Provides a Runscope test resource.`,
			Description:      ``,
			Keywords: []string{
				"test",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(String, Required) The name of this test.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String, Optional) Human-readable description of the new test. is being created for. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the test.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this test.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the new test.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the test.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this test.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the new test.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"runscope_bucket":      0,
		"runscope_environment": 1,
		"runscope_schedule":    2,
		"runscope_step":        3,
		"runscope_test":        4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
