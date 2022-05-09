package tls

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tls_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information about the TLS certificates securing a host. Use this data source to get information, such as SHA1 fingerprint or serial number, about the TLS certificates that protects a URL.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tls_public_key",
			Category:         "Data Sources",
			ShortDescription: `Get a public key from a PEM-encoded private key. Use this data source to get the public key from a PEM (RFC 1421) https://datatracker.ietf.org/doc/html/rfc1421 or OpenSSH PEM (RFC 4716) https://datatracker.ietf.org/doc/html/rfc4716 formatted private key, for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"tls_certificate": 0,
		"tls_public_key":  1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
