package http

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "http_http",
			Category:         "Data Sources",
			ShortDescription: `Retrieves the content at an HTTP or HTTPS URL.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL to request data from. This URL must respond with a ` + "`" + `200 OK` + "`" + ` response and a ` + "`" + `text/`,
				},
				resource.Attribute{
					Name:        "request_headers",
					Description: `(Optional) A map of strings representing additional HTTP headers to include in the request. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `The raw body of the HTTP response.`,
				},
				resource.Attribute{
					Name:        "response_headers",
					Description: `A map of strings representing the response HTTP headers. Duplicate headers are contatenated with ` + "`" + `, ` + "`" + ` according to [RFC2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `The raw body of the HTTP response.`,
				},
				resource.Attribute{
					Name:        "response_headers",
					Description: `A map of strings representing the response HTTP headers. Duplicate headers are contatenated with ` + "`" + `, ` + "`" + ` according to [RFC2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2)`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"http_http": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
