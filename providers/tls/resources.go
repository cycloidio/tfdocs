package tls

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tls_cert_request",
			Category:         "Resources",
			ShortDescription: `Creates a Certificate Signing Request (CSR) in PEM (RFC 1421) https://datatracker.ietf.org/doc/html/rfc1421 format. PEM is the typical format used to request a certificate from a Certificate Authority (CA). This resource is intended to be used in conjunction with a Terraform provider for a particular certificate authority in order to provision a new certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tls_locally_signed_cert",
			Category:         "Resources",
			ShortDescription: `Creates a TLS certificate in PEM (RFC 1421) https://datatracker.ietf.org/doc/html/rfc1421 format using a Certificate Signing Request (CSR) and signs it with a provided (local) Certificate Authority (CA).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tls_private_key",
			Category:         "Resources",
			ShortDescription: `Creates a PEM (and OpenSSH) formatted private key. Generates a secure private key and encodes it in PEM (RFC 1421) https://datatracker.ietf.org/doc/html/rfc1421 and OpenSSH PEM (RFC 4716) https://datatracker.ietf.org/doc/html/rfc4716 formats. This resource is primarily intended for easily bootstrapping throwaway development environments.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tls_self_signed_cert",
			Category:         "Resources",
			ShortDescription: `Creates a self-signed TLS certificate in PEM (RFC 1421) https://datatracker.ietf.org/doc/html/rfc1421 format.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"tls_cert_request":        0,
		"tls_locally_signed_cert": 1,
		"tls_private_key":         2,
		"tls_self_signed_cert":    3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
