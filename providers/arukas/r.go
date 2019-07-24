package arukas

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "arukas_container",
			Category:         "Resources",
			ShortDescription: `Manages Arukas Containers`,
			Description:      ``,
			Keywords: []string{
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the container.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required, string) The ID of the image to back this container.It must be a public image on DockerHub.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `(Optional, int) The count of the instance. It must be between ` + "`" + `1` + "`" + ` and ` + "`" + `10` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Optional, string) The plan of the Arukas. It must be ` + "`" + `free` + "`" + ` or ` + "`" + `hobby` + "`" + ` or ` + "`" + `standard-1` + "`" + ` or ` + "`" + `standard-2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional,string) The subdomain part of the endpoint assigned by Arukas. If it is not set, Arukas will do automatic assignment.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required , block) See [Ports](#ports) below for details.`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `(Required , block) See [Environments](#environments) below for details.`,
				},
				resource.Attribute{
					Name:        "cmd",
					Description: `(Optional , string) The command of the container. <a id="ports"></a> ### Ports ` + "`" + `ports` + "`" + ` is a block within the configuration that can be repeated to specify the port mappings of the container. Each ` + "`" + `ports` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, string) Protocol that can be used over this port, defaults to ` + "`" + `tcp` + "`" + `,It must be ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `(Optional, int) Port within the container,defaults to ` + "`" + `80` + "`" + `, It must be between ` + "`" + `1` + "`" + ` to ` + "`" + `65535` + "`" + `. <a id="environments"></a> ### Environments ` + "`" + `environments` + "`" + ` is a block within the configuration that can be repeated to specify the environment variables. Each ` + "`" + `environments` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, string) Key of environment variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, string) Value of environment variable. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Arukas application to which the Arukas Service and the container belongs.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the Arukas Service to which the container belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The ID of the image to back this container.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The count of the instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The name of region`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The name of plan`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The subdomain part of the endpoint assigned by Arukas.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `See [Ports](#ports) below for details.`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `See [Environments](#environments) below for details.`,
				},
				resource.Attribute{
					Name:        "cmd",
					Description: `The command of the container.`,
				},
				resource.Attribute{
					Name:        "port_mappings",
					Description: `See [PortMappings](#port_mappings) below for details.`,
				},
				resource.Attribute{
					Name:        "endpoint_full_url",
					Description: `The URL of endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint_full_hostname",
					Description: `The Hostname of endpoint. <a id="port_mappings"></a> ### PortMappings ` + "`" + `port_mappings` + "`" + ` is a block within the configuration that the port mappings of the container. Each ` + "`" + `port_mappings` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The name of the host actually running the container.`,
				},
				resource.Attribute{
					Name:        "ipaddress",
					Description: `The IP address of the host actually running the container.`,
				},
				resource.Attribute{
					Name:        "container_port",
					Description: `Port within the container.`,
				},
				resource.Attribute{
					Name:        "service_port",
					Description: `The actual port mapped to the port in the container.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Arukas application to which the Arukas Service and the container belongs.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the Arukas Service to which the container belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The ID of the image to back this container.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The count of the instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The name of region`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The name of plan`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The subdomain part of the endpoint assigned by Arukas.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `See [Ports](#ports) below for details.`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `See [Environments](#environments) below for details.`,
				},
				resource.Attribute{
					Name:        "cmd",
					Description: `The command of the container.`,
				},
				resource.Attribute{
					Name:        "port_mappings",
					Description: `See [PortMappings](#port_mappings) below for details.`,
				},
				resource.Attribute{
					Name:        "endpoint_full_url",
					Description: `The URL of endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint_full_hostname",
					Description: `The Hostname of endpoint. <a id="port_mappings"></a> ### PortMappings ` + "`" + `port_mappings` + "`" + ` is a block within the configuration that the port mappings of the container. Each ` + "`" + `port_mappings` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The name of the host actually running the container.`,
				},
				resource.Attribute{
					Name:        "ipaddress",
					Description: `The IP address of the host actually running the container.`,
				},
				resource.Attribute{
					Name:        "container_port",
					Description: `Port within the container.`,
				},
				resource.Attribute{
					Name:        "service_port",
					Description: `The actual port mapped to the port in the container.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"arukas_container": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
