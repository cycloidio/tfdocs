package time

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "time_offset",
			Category:         "Resources",
			ShortDescription: `Manages a offset time resource.`,
			Description: `

Manages an offset time resource, which keeps an UTC timestamp stored in the Terraform state that is offset from a locally sourced base timestamp. This prevents perpetual differences caused by using the [` + "`" + `timestamp()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/timestamp.html).

-> Further manipulation of incoming or outgoing values can be accomplished with the [` + "`" + `formatdate()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/formatdate.html) and the [` + "`" + `timeadd()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/timeadd.html).

`,
			Keywords: []string{
				"offset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "base_rfc3339",
					Description: `(Optional) Configure the base timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (` + "`" + `YYYY-MM-DDTHH:MM:SSZ` + "`" + `). Defaults to the current time.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. See [the main provider documentation](../index.html) for more information.`,
				},
				resource.Attribute{
					Name:        "offset_days",
					Description: `(Optional) Number of days to offset the base timestamp. Conflicts with other ` + "`" + `offset_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "offset_hours",
					Description: `(Optional) Number of hours to offset the base timestamp. Conflicts with other ` + "`" + `offset_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "offset_minutes",
					Description: `(Optional) Number of minutes to offset the base timestamp. Conflicts with other ` + "`" + `offset_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "offset_months",
					Description: `(Optional) Number of months to offset the base timestamp. Conflicts with other ` + "`" + `offset_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "offset_seconds",
					Description: `(Optional) Number of seconds to offset the base timestamp. Conflicts with other ` + "`" + `offset_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "offset_years",
					Description: `(Optional) Number of years to offset the base timestamp. Conflicts with other ` + "`" + `offset_` + "`" + ` arguments. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `Number day of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Number hour of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UTC RFC3339 format of the base timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `Number minute of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `Number month of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "rfc3339",
					Description: `UTC RFC3339 format of the offset timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "second",
					Description: `Number second of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "unix",
					Description: `Number of seconds since epoch time, e.g. ` + "`" + `1581489373` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "year",
					Description: `Number year of offset timestamp. ## Import This resource can be imported using the base UTC RFC3339 timestamp and offset years, months, days, hours, minutes, and seconds, separated by commas (` + "`" + `,` + "`" + `), e.g. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_offset.example 2020-02-12T06:36:13Z,0,0,7,0,0,0 ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `triggers` + "`" + ` argument cannot be imported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "day",
					Description: `Number day of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Number hour of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UTC RFC3339 format of the base timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `Number minute of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `Number month of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "rfc3339",
					Description: `UTC RFC3339 format of the offset timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "second",
					Description: `Number second of offset timestamp.`,
				},
				resource.Attribute{
					Name:        "unix",
					Description: `Number of seconds since epoch time, e.g. ` + "`" + `1581489373` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "year",
					Description: `Number year of offset timestamp. ## Import This resource can be imported using the base UTC RFC3339 timestamp and offset years, months, days, hours, minutes, and seconds, separated by commas (` + "`" + `,` + "`" + `), e.g. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_offset.example 2020-02-12T06:36:13Z,0,0,7,0,0,0 ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `triggers` + "`" + ` argument cannot be imported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "time_rotating",
			Category:         "Resources",
			ShortDescription: `Manages a rotating time resource.`,
			Description: `

Manages a rotating time resource, which keeps a rotating UTC timestamp stored in the Terraform state and proposes resource recreation when the locally sourced current time is beyond the rotation time. This rotation only occurs when Terraform is executed, meaning there will be drift between the rotation timestamp and actual rotation. The new rotation timestamp offset includes this drift. This prevents perpetual differences caused by using the [` + "`" + `timestamp()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/timestamp.html) by only forcing a new value on the set cadence.

-> Further manipulation of incoming or outgoing values can be accomplished with the [` + "`" + `formatdate()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/formatdate.html) and the [` + "`" + `timeadd()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/timeadd.html).

`,
			Keywords: []string{
				"rotating",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rfc3339",
					Description: `(Optional) Configure the base timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (` + "`" + `YYYY-MM-DDTHH:MM:SSZ` + "`" + `). Defaults to the current time.`,
				},
				resource.Attribute{
					Name:        "rotation_days",
					Description: `(Optional) Number of days to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other ` + "`" + `rotation_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "rotation_hours",
					Description: `(Optional) Number of hours to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other ` + "`" + `rotation_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "rotation_minutes",
					Description: `(Optional) Number of minutes to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other ` + "`" + `rotation_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "rotation_months",
					Description: `(Optional) Number of months to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other ` + "`" + `rotation_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "rotation_rfc3339",
					Description: `(Optional) Configure the rotation timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (` + "`" + `YYYY-MM-DDTHH:MM:SSZ` + "`" + `). When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other ` + "`" + `rotation_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "rotation_years",
					Description: `(Optional) Number of years to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other ` + "`" + `rotation_` + "`" + ` arguments.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. These conditions recreate the resource in addition to other rotation arguments. See [the main provider documentation](../index.html) for more information. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `Number day of timestamp.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Number hour of timestamp.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UTC RFC3339 format of the base timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `Number minute of timestamp.`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `Number month of timestamp.`,
				},
				resource.Attribute{
					Name:        "second",
					Description: `Number second of timestamp.`,
				},
				resource.Attribute{
					Name:        "unix",
					Description: `Number of seconds since epoch time, e.g. ` + "`" + `1581489373` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "year",
					Description: `Number year of timestamp. ## Import This resource can be imported using the base UTC RFC3339 value and rotation years, months, days, hours, and minutes, separated by commas (` + "`" + `,` + "`" + `), e.g. for 30 days ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_rotation.example 2020-02-12T06:36:13Z,0,0,30,0,0 ` + "`" + `` + "`" + `` + "`" + ` Otherwise, to import with the rotation RFC3339 value, the base UTC RFC3339 value and rotation UTC RFC3339 value, separated by commas (` + "`" + `,` + "`" + `), e.g. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_rotation.example 2020-02-12T06:36:13Z,2020-02-13T06:36:13Z ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `triggers` + "`" + ` argument cannot be imported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "day",
					Description: `Number day of timestamp.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Number hour of timestamp.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UTC RFC3339 format of the base timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `Number minute of timestamp.`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `Number month of timestamp.`,
				},
				resource.Attribute{
					Name:        "second",
					Description: `Number second of timestamp.`,
				},
				resource.Attribute{
					Name:        "unix",
					Description: `Number of seconds since epoch time, e.g. ` + "`" + `1581489373` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "year",
					Description: `Number year of timestamp. ## Import This resource can be imported using the base UTC RFC3339 value and rotation years, months, days, hours, and minutes, separated by commas (` + "`" + `,` + "`" + `), e.g. for 30 days ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_rotation.example 2020-02-12T06:36:13Z,0,0,30,0,0 ` + "`" + `` + "`" + `` + "`" + ` Otherwise, to import with the rotation RFC3339 value, the base UTC RFC3339 value and rotation UTC RFC3339 value, separated by commas (` + "`" + `,` + "`" + `), e.g. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_rotation.example 2020-02-12T06:36:13Z,2020-02-13T06:36:13Z ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `triggers` + "`" + ` argument cannot be imported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "time_sleep",
			Category:         "Resources",
			ShortDescription: `Manages a static time resource.`,
			Description: `

Manages a resource that delays creation and/or destruction, typically for further resources. This prevents cross-platform compatibility and destroy-time issues with using the [` + "`" + `local-exec` + "`" + ` provisioner](https://www.terraform.io/docs/provisioners/local-exec.html).

-> In many cases, this resource should be considered a workaround for issues that should be reported and handled in downstream Terraform Provider logic. Downstream resources can usually introduce or adjust retries in their code to handle time delay issues for all Terraform configurations or upstream resources can be improved to better wait for a resource to be fully ready and available.

`,
			Keywords: []string{
				"sleep",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "create_duration",
					Description: `(Optional) [Time duration][1] to delay resource creation. For example, ` + "`" + `30s` + "`" + ` for 30 seconds or ` + "`" + `5m` + "`" + ` for 5 minutes. Updating this value by itself will not trigger a delay.`,
				},
				resource.Attribute{
					Name:        "destroy_duration",
					Description: `(Optional) [Time duration][1] to delay resource destroy. For example, ` + "`" + `30s` + "`" + ` for 30 seconds or ` + "`" + `5m` + "`" + ` for 5 minutes. Updating this value by itself will not trigger a delay. This value or any updates to it must be successfully applied into the Terraform state before destroying this resource to take effect.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) Arbitrary map of values that, when changed, will run any creation or destroy delays again. See [the main provider documentation](../index.html) for more information. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UTC RFC3339 timestamp of the creation or import, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `. ## Import This resource can be imported with the ` + "`" + `create_duration` + "`" + ` and ` + "`" + `destroy_duration` + "`" + `, separated by a comma (` + "`" + `,` + "`" + `). e.g. For 30 seconds create duration with no destroy duration: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_sleep.example 30s, ` + "`" + `` + "`" + `` + "`" + ` e.g. For 30 seconds destroy duration with no create duration: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_sleep.example ,30s ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `triggers` + "`" + ` argument cannot be imported. [1]: https://golang.org/pkg/time/#ParseDuration`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UTC RFC3339 timestamp of the creation or import, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `. ## Import This resource can be imported with the ` + "`" + `create_duration` + "`" + ` and ` + "`" + `destroy_duration` + "`" + `, separated by a comma (` + "`" + `,` + "`" + `). e.g. For 30 seconds create duration with no destroy duration: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_sleep.example 30s, ` + "`" + `` + "`" + `` + "`" + ` e.g. For 30 seconds destroy duration with no create duration: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_sleep.example ,30s ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `triggers` + "`" + ` argument cannot be imported. [1]: https://golang.org/pkg/time/#ParseDuration`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "time_static",
			Category:         "Resources",
			ShortDescription: `Manages a static time resource.`,
			Description: `

Manages a static time resource, which keeps a locally sourced UTC timestamp stored in the Terraform state. This prevents perpetual differences caused by using the [` + "`" + `timestamp()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/timestamp.html).

-> Further manipulation of incoming or outgoing values can be accomplished with the [` + "`" + `formatdate()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/formatdate.html) and the [` + "`" + `timeadd()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/timeadd.html).

`,
			Keywords: []string{
				"static",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. See [the main provider documentation](../index.html) for more information.`,
				},
				resource.Attribute{
					Name:        "rfc3339",
					Description: `(Optional) Configure the base timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (` + "`" + `YYYY-MM-DDTHH:MM:SSZ` + "`" + `). Defaults to the current time. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `Number day of timestamp.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Number hour of timestamp.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UTC RFC3339 timestamp format, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `Number minute of timestamp.`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `Number month of timestamp.`,
				},
				resource.Attribute{
					Name:        "rfc3339",
					Description: `UTC RFC3339 format of timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "second",
					Description: `Number second of timestamp.`,
				},
				resource.Attribute{
					Name:        "unix",
					Description: `Number of seconds since epoch time, e.g. ` + "`" + `1581489373` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "year",
					Description: `Number year of timestamp. ## Import This resource can be imported using the UTC RFC3339 value, e.g. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_static.example 2020-02-12T06:36:13Z ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `triggers` + "`" + ` argument cannot be imported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "day",
					Description: `Number day of timestamp.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Number hour of timestamp.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UTC RFC3339 timestamp format, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `Number minute of timestamp.`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `Number month of timestamp.`,
				},
				resource.Attribute{
					Name:        "rfc3339",
					Description: `UTC RFC3339 format of timestamp, e.g. ` + "`" + `2020-02-12T06:36:13Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "second",
					Description: `Number second of timestamp.`,
				},
				resource.Attribute{
					Name:        "unix",
					Description: `Number of seconds since epoch time, e.g. ` + "`" + `1581489373` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "year",
					Description: `Number year of timestamp. ## Import This resource can be imported using the UTC RFC3339 value, e.g. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import time_static.example 2020-02-12T06:36:13Z ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `triggers` + "`" + ` argument cannot be imported.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"time_offset":   0,
		"time_rotating": 1,
		"time_sleep":    2,
		"time_static":   3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
