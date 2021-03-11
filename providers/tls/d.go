package tls

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tls_public_key",
			Category:         "Data Sources",
			ShortDescription: `Get a public key from a PEM-encoded private key.`,
			Description: `

Use this data source to get the public key from a PEM-encoded private key for use in other
resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "private_key_pem",
					Description: `(Required) The private key to use. Currently-supported key types are "RSA" or "ECDSA". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "private_key_pem",
					Description: `The private key data in PEM format.`,
				},
				resource.Attribute{
					Name:        "public_key_pem",
					Description: `The public key data in PEM format.`,
				},
				resource.Attribute{
					Name:        "public_key_openssh",
					Description: `The public key data in OpenSSH ` + "`" + `authorized_keys` + "`" + ` format, if the selected private key format is compatible. All RSA keys are supported, and ECDSA keys with curves "P256", "P384" and "P521" are supported. This attribute is empty if an incompatible ECDSA curve is selected.`,
				},
				resource.Attribute{
					Name:        "public_key_fingerprint_md5",
					Description: `The md5 hash of the public key data in OpenSSH MD5 hash format, e.g. ` + "`" + `aa:bb:cc:...` + "`" + `. Only available if the selected private key format is compatible, as per the rules for ` + "`" + `public_key_openssh` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "private_key_pem",
					Description: `The private key data in PEM format.`,
				},
				resource.Attribute{
					Name:        "public_key_pem",
					Description: `The public key data in PEM format.`,
				},
				resource.Attribute{
					Name:        "public_key_openssh",
					Description: `The public key data in OpenSSH ` + "`" + `authorized_keys` + "`" + ` format, if the selected private key format is compatible. All RSA keys are supported, and ECDSA keys with curves "P256", "P384" and "P521" are supported. This attribute is empty if an incompatible ECDSA curve is selected.`,
				},
				resource.Attribute{
					Name:        "public_key_fingerprint_md5",
					Description: `The md5 hash of the public key data in OpenSSH MD5 hash format, e.g. ` + "`" + `aa:bb:cc:...` + "`" + `. Only available if the selected private key format is compatible, as per the rules for ` + "`" + `public_key_openssh` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tls_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information about the TLS certificates securing a host.`,
			Description: `

Use this data source to get information, such as SHA1 fingerprint or serial number, about the TLS certificates that
protect an HTTPS website. Note that the certificate chain isn't verified.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL of the website to get the certificates from.`,
				},
				resource.Attribute{
					Name:        "verify_chain",
					Description: `(Optional) Whether to verify the certificate chain while parsing it or not ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `The certificates protecting the site, with the root of the chain first.`,
				},
				resource.Attribute{
					Name:        "certificates.#.not_after",
					Description: `The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.`,
				},
				resource.Attribute{
					Name:        "certificates.#.not_before",
					Description: `The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.`,
				},
				resource.Attribute{
					Name:        "certificates.#.is_ca",
					Description: `` + "`" + `true` + "`" + ` if this certificate is a ca certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.issuer",
					Description: `Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).`,
				},
				resource.Attribute{
					Name:        "certificates.#.public_key_algorithm",
					Description: `The algorithm used to create the certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.serial_number",
					Description: `Number that uniquely identifies the certificate with the CA's system. The ` + "`" + `format` + "`" + ` function can be used to convert this base 10 number into other bases, such as hex.`,
				},
				resource.Attribute{
					Name:        "certificates.#.sha1_fingerprint",
					Description: `The SHA1 fingerprint of the public key of the certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.signature_algorithm",
					Description: `The algorithm used to sign the certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.subject",
					Description: `The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).`,
				},
				resource.Attribute{
					Name:        "certificates.#.version",
					Description: `The version the certificate is in.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificates",
					Description: `The certificates protecting the site, with the root of the chain first.`,
				},
				resource.Attribute{
					Name:        "certificates.#.not_after",
					Description: `The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.`,
				},
				resource.Attribute{
					Name:        "certificates.#.not_before",
					Description: `The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.`,
				},
				resource.Attribute{
					Name:        "certificates.#.is_ca",
					Description: `` + "`" + `true` + "`" + ` if this certificate is a ca certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.issuer",
					Description: `Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).`,
				},
				resource.Attribute{
					Name:        "certificates.#.public_key_algorithm",
					Description: `The algorithm used to create the certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.serial_number",
					Description: `Number that uniquely identifies the certificate with the CA's system. The ` + "`" + `format` + "`" + ` function can be used to convert this base 10 number into other bases, such as hex.`,
				},
				resource.Attribute{
					Name:        "certificates.#.sha1_fingerprint",
					Description: `The SHA1 fingerprint of the public key of the certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.signature_algorithm",
					Description: `The algorithm used to sign the certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.subject",
					Description: `The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).`,
				},
				resource.Attribute{
					Name:        "certificates.#.version",
					Description: `The version the certificate is in.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"tls_public_key":  0,
		"tls_certificate": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
