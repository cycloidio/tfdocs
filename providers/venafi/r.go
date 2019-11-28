package venafi

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "venafi_certificate",
			Category:         "Resources",
			ShortDescription: `Provides access to TLS key and certificate data in Venafi. This can be used to define a Venafi certificate.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required, string) The common name of the certificate.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional, string) Key encryption algorithm, either RSA or ECDSA. Defaults to "RSA".`,
				},
				resource.Attribute{
					Name:        "rsa_bits",
					Description: `(Optional, integer) Number of bits to use when generating an RSA key. Applies when algorithm=RSA. Defaults to 2048.`,
				},
				resource.Attribute{
					Name:        "ecsa_curve",
					Description: `(Optional, string) Elliptic curve to use when generating an ECDSA key pair. Applies when algorithm=ECDSA. Defaults to "P521".`,
				},
				resource.Attribute{
					Name:        "san_dns",
					Description: `(Optional, set of strings) List of DNS names to use as alternative subjects of the certificate.`,
				},
				resource.Attribute{
					Name:        "san_email",
					Description: `(Optional, set of strings) List of email addresses to use as alternative subjects of the certificate.`,
				},
				resource.Attribute{
					Name:        "san_ip",
					Description: `(Optional, set of strings) List of IP addresses to use as alternative subjects of the certificate.`,
				},
				resource.Attribute{
					Name:        "key_password",
					Description: `(Optional, string) The password used to encrypt the private key.`,
				},
				resource.Attribute{
					Name:        "expiration_window",
					Description: `(Optional, integer) Number of hours before certificate expiry to request a new certificate. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "private_key_pem",
					Description: `The private key in PEM format.`,
				},
				resource.Attribute{
					Name:        "chain",
					Description: `The trust chain of X509 certificate authority certificates in PEM format concatenated together.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The X509 certificate in PEM format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "private_key_pem",
					Description: `The private key in PEM format.`,
				},
				resource.Attribute{
					Name:        "chain",
					Description: `The trust chain of X509 certificate authority certificates in PEM format concatenated together.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The X509 certificate in PEM format.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"venafi_certificate": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
