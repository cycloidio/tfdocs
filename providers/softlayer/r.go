package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ssh_key",
			Category:         "Resources",
			ShortDescription: `Manages SoftLayer SSH Keys.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"key",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A descriptive name used to identify an SSH key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public SSH key.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) A small note about an SSH key to use at your discretion. The ` + "`" + `name` + "`" + ` and ` + "`" + `notes` + "`" + ` fields are editable. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the new SSH key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `sequence of bytes to authenticate or lookup a longer SSH key.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the new SSH key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `sequence of bytes to authenticate or lookup a longer SSH key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "virtual_guest",
			Category:         "Resources",
			ShortDescription: `Manages SoftLayer Virtual Guests.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"guest",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual guest.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual guest.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ssh_key":       0,
		"virtual_guest": 1,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
