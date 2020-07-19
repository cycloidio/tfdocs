package rubrik

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "assign_sla",
			Category:         "Resources",
			ShortDescription: `Assign a Rubrik object to a specified SLA Domain.`,
			Description:      ``,
			Keywords: []string{
				"assign",
				"sla",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "object_name",
					Description: `(Required) The name of the Rubrik object you wish to assign to an SLA Domain.`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `(Required) The Rubrik object type you want to assign to the SLA Domain. Currently, ` + "`" + `vmware` + "`" + ` and ` + "`" + `ahv` + "`" + ` are the only supported values.`,
				},
				resource.Attribute{
					Name:        "sla_name",
					Description: `(Required) The name of the SLA Domain you wish to assign an object to. To exclude the object from all SLA assignments use ` + "`" + `do not protect` + "`" + ` as the ` + "`" + `sla_name` + "`" + `. To assign the selected object to the SLA of the next higher level object use ` + "`" + `clear` + "`" + ` as the ` + "`" + `sla_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The number of seconds to wait to establish a connection the Rubrik cluster before returning a timeout error. Default is 15. Valid object_type choices:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An ID unique to Terraform for this port group. The convention is a prefix, the host system ID, and the port group name. An example would be ` + "`" + `vmware-exampleVMName-assigned-sla-Gold` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sla_domain",
					Description: `The name of the SLA Domain assigned to the ` + "`" + `object_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `The type of the Rubrik object you assigned an SLA Domain to.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `The name of the Rubrik object you assigned an SLA Domain to. ## Destroy Behavior On ` + "`" + `terraform destroy` + "`" + `, this resource will set the ` + "`" + `sla_domain` + "`" + ` to ` + "`" + `clear` + "`" + ` which will assign the select ` + "`" + `object_name` + "`" + ` to the SLA of the next higher level object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An ID unique to Terraform for this port group. The convention is a prefix, the host system ID, and the port group name. An example would be ` + "`" + `vmware-exampleVMName-assigned-sla-Gold` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sla_domain",
					Description: `The name of the SLA Domain assigned to the ` + "`" + `object_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `The type of the Rubrik object you assigned an SLA Domain to.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `The name of the Rubrik object you assigned an SLA Domain to. ## Destroy Behavior On ` + "`" + `terraform destroy` + "`" + `, this resource will set the ` + "`" + `sla_domain` + "`" + ` to ` + "`" + `clear` + "`" + ` which will assign the select ` + "`" + `object_name` + "`" + ` to the SLA of the next higher level object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configure_timezone",
			Category:         "Resources",
			ShortDescription: `Set the timezone on a Rubrik CDM cluster.`,
			Description:      ``,
			Keywords: []string{
				"configure",
				"timezone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "timezone",
					Description: `(Required) The timezone used by the Rubrik cluster which uses the specified time zone for time values in the web UI, all reports, SLA Domain settings, and all other time related operations. Valid choices are America/Anchorage, America/Araguaina, America/Barbados, America/Chicago, America/Denver, America/Los_Angeles America/Mexico_City, America/New_York, America/Noronha, America/Phoenix, America/Toronto, America/Vancouver, Asia/Bangkok, Asia/Dhaka, Asia/Dubai, Asia/Hong_Kong, Asia/Karachi,Asia/Kathmandu, Asia/Kolkata, Asia/Magadan, Asia/Singapore, Asia/Tokyo, Atlantic/Cape_Verde, Australia/Perth, Australia/Sydney, Europe/Amsterdam, Europe/Athens, Europe/London, Europe/Moscow, Pacific/Auckland, Pacific/Honolulu, Pacific/Midway, or UTC.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The number of seconds to wait to establish a connection the Rubrik cluster before returning a timeout error. Default is 15.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"assign_sla":         0,
		"configure_timezone": 1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
