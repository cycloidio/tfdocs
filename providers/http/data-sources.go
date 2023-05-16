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
			ShortDescription: `The http data source makes an HTTP GET request to the given URL and exports information about the response. The given URL may be either an http or https URL. At present this resource can only retrieve data from URLs that respond with text/* or application/json content types, and expects the result to be UTF-8 encoded regardless of the returned content type header. ~> Important Although https URLs can be used, there is currently no mechanism to authenticate the remote server except for general verification of the server certificate's chain of trust. Data retrieved from servers not under your control should be treated as untrustworthy. By default, there are no retries. Configuring the retry block will result in retries if an error is returned by the client (e.g., connection errors) or if a 5xx-range (except 501) status code is received. For further details see go-retryablehttp https://pkg.go.dev/github.com/hashicorp/go-retryablehttp.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
