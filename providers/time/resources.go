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
			ShortDescription: `Manages an offset time resource, which keeps an UTC timestamp stored in the Terraform state that is offset from a locally sourced base timestamp. This prevents perpetual differences caused by using the timestamp() function https://www.terraform.io/docs/configuration/functions/timestamp.html.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "time_rotating",
			Category:         "Resources",
			ShortDescription: `Manages a rotating time resource, which keeps a rotating UTC timestamp stored in the Terraform state and proposes resource recreation when the locally sourced current time is beyond the rotation time. This rotation only occurs when Terraform is executed, meaning there will be drift between the rotation timestamp and actual rotation. The new rotation timestamp offset includes this drift. This prevents perpetual differences caused by using the timestamp() function https://www.terraform.io/docs/configuration/functions/timestamp.html by only forcing a new value on the set cadence.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "time_sleep",
			Category:         "Resources",
			ShortDescription: `Manages a resource that delays creation and/or destruction, typically for further resources. This prevents cross-platform compatibility and destroy-time issues with using the local-exec provisioner https://www.terraform.io/docs/provisioners/local-exec.html.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "time_static",
			Category:         "Resources",
			ShortDescription: `Manages a static time resource, which keeps a locally sourced UTC timestamp stored in the Terraform state. This prevents perpetual differences caused by using the timestamp() function https://www.terraform.io/docs/configuration/functions/timestamp.html.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
