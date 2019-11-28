package genymotion

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "genymotion_cloud",
			Category:         "Resources",
			ShortDescription: `Create a Genymotion Cloud Android virtual device`,
			Description:      ``,
			Keywords: []string{
				"cloud",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Android instance.`,
				},
				resource.Attribute{
					Name:        "recipe_uuid",
					Description: `(Required) Recipe UUID is the identifier used when starting an instance. It can be retrieved using ` + "`" + `gmsaas recipes list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "adbconnect",
					Description: `(Optional) If is true, it will connect the instance to ADB. Defaults to "true".`,
				},
				resource.Attribute{
					Name:        "adb_serial_port",
					Description: `(Optional) If the --adb_serial_port <PORT> option is set, the instance will be connected to ADB on localhost:<PORT>.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"genymotion_cloud": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
