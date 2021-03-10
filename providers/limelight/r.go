package limelight

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "limelight_delivery",
			Category:         "Resources",
			ShortDescription: `A resource that can be used to configure Content Delivery.`,
			Description:      ``,
			Keywords: []string{
				"delivery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "shortname",
					Description: `(Required) The account name (shortname).`,
				},
				resource.Attribute{
					Name:        "service_profile",
					Description: `(Optional) Service profile to use. Defaults to ` + "`" + `LLNW-Generic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "published_hostname",
					Description: `(Required) Published hostname for the content.`,
				},
				resource.Attribute{
					Name:        "published_path",
					Description: `(Required) Published path for the content.`,
				},
				resource.Attribute{
					Name:        "source_hostname",
					Description: `(Required) Source (origin) hostname for the content.`,
				},
				resource.Attribute{
					Name:        "source_path",
					Description: `(Required) Source path on the origin for the content.`,
				},
				resource.Attribute{
					Name:        "protocol_set",
					Description: `(Required) Protocol configuration for the delivery as child blocks (max of 2):`,
				},
				resource.Attribute{
					Name:        "published_protocol",
					Description: `(Required) Published protocol to use (e.g. ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_protocol",
					Description: `(Required) Source protocol to use (e.g. ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) Source port to use. Defaults to ` + "`" + `80` + "`" + ` for http and ` + "`" + `443` + "`" + ` for https.`,
				},
				resource.Attribute{
					Name:        "option",
					Description: `(Optional) Protocol options to use specified as child blocks:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Option name.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Required) List of string parameters for the option. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The delivery ID.`,
				},
				resource.Attribute{
					Name:        "version_number",
					Description: `The delivery version. ## Importing An existing Delivery configuration can be [imported](https://www.terraform.io/docs/import/index.html) into this resource, via the following command: ` + "`" + `` + "`" + `` + "`" + ` terraform import limelight_delivery.example_website UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the Delivery named ` + "`" + `example_website` + "`" + ` with the ID ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The delivery ID.`,
				},
				resource.Attribute{
					Name:        "version_number",
					Description: `The delivery version. ## Importing An existing Delivery configuration can be [imported](https://www.terraform.io/docs/import/index.html) into this resource, via the following command: ` + "`" + `` + "`" + `` + "`" + ` terraform import limelight_delivery.example_website UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the Delivery named ` + "`" + `example_website` + "`" + ` with the ID ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "limelight_edgefunction",
			Category:         "Resources",
			ShortDescription: `A resource that can be used to manage EdgeFunctions.`,
			Description:      ``,
			Keywords: []string{
				"edgefunction",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "shortname",
					Description: `(Required) The account name (shortname).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the EdgeFunction.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the EdgeFunction.`,
				},
				resource.Attribute{
					Name:        "function_archive",
					Description: `(Required) Path to the function archive (zip file).`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `(Required) Handler that's run when the EdgeFunction is invoked.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Required) The runtime for the EdgeFunction.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) The memory allocated to the EdgeFunction. Defaults to ` + "`" + `256` + "`" + `. CPU is allocated proportional to memory.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout for the EdgeFunction execution in milliseconds. Defaults to ` + "`" + `5000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "can_debug",
					Description: `(Optional) Boolean flag to enable debug IO. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "function_sha256",
					Description: `(Required) The SHA256 value of the ` + "`" + `function_archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reserved_concurrency",
					Description: `(Optional) Sets the reserved concurrency for the EdgeFunction. Defaults to ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "environment_variable",
					Description: `(Optional) Zero or more environment variables for the EdgeFunction as child blocks:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The environment variable name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The environment variable value. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "revision_id",
					Description: `Revision number of the EdgeFunction. ## Importing An existing EdgeFunction can be [imported](https://www.terraform.io/docs/import/index.html) into this resource, via the following command: ` + "`" + `` + "`" + `` + "`" + ` terraform import limelight_edgefunction.my_func FUNCTION_ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the EdgeFunction named ` + "`" + `my_func` + "`" + ` with the ID ` + "`" + `FUNCTION_ID` + "`" + ` where ` + "`" + `FUNCTION_ID` + "`" + ` is of the form ` + "`" + `<shortname>:<function_name>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "revision_id",
					Description: `Revision number of the EdgeFunction. ## Importing An existing EdgeFunction can be [imported](https://www.terraform.io/docs/import/index.html) into this resource, via the following command: ` + "`" + `` + "`" + `` + "`" + ` terraform import limelight_edgefunction.my_func FUNCTION_ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the EdgeFunction named ` + "`" + `my_func` + "`" + ` with the ID ` + "`" + `FUNCTION_ID` + "`" + ` where ` + "`" + `FUNCTION_ID` + "`" + ` is of the form ` + "`" + `<shortname>:<function_name>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "limelight_edgefunction_alias",
			Category:         "Resources",
			ShortDescription: `A resource that can be used to manage EdgeFunction Aliases.`,
			Description:      ``,
			Keywords: []string{
				"edgefunction",
				"alias",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "shortname",
					Description: `(Required) The account name (shortname).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the EdgeFunction alias.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the EdgeFunction alias.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required) The EdgeFunction's name to create the alias for.`,
				},
				resource.Attribute{
					Name:        "function_version",
					Description: `(Required) The EdgeFunction's version to create the alias for. If a version other than ` + "`" + `$LATEST` + "`" + ` is used, this version must already exist. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "revision_id",
					Description: `Revision number of the EdgeFunction alias. ## Importing An existing EdgeFunction Alias can be [imported](https://www.terraform.io/docs/import/index.html) into this resource, via the following command: ` + "`" + `` + "`" + `` + "`" + ` terraform import limelight_edgefunction_alias.my_alias ALIAS_UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the EdgeFunction Alias named ` + "`" + `my_alias` + "`" + ` with the ID ` + "`" + `ALIAS_UUID` + "`" + ` where ` + "`" + `ALIAS_UUID` + "`" + ` is of the form ` + "`" + `<shortname>:<function_name>:<alias_name>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "revision_id",
					Description: `Revision number of the EdgeFunction alias. ## Importing An existing EdgeFunction Alias can be [imported](https://www.terraform.io/docs/import/index.html) into this resource, via the following command: ` + "`" + `` + "`" + `` + "`" + ` terraform import limelight_edgefunction_alias.my_alias ALIAS_UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the EdgeFunction Alias named ` + "`" + `my_alias` + "`" + ` with the ID ` + "`" + `ALIAS_UUID` + "`" + ` where ` + "`" + `ALIAS_UUID` + "`" + ` is of the form ` + "`" + `<shortname>:<function_name>:<alias_name>` + "`" + `.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"limelight_delivery":           0,
		"limelight_edgefunction":       1,
		"limelight_edgefunction_alias": 2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
