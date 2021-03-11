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
					Name:        "custom_fields",
					Description: `(Optional, map) Collection of Custom Field name-value pairs to assign to the certificate.`,
				},
				resource.Attribute{
					Name:        "valid_days",
					Description: `(Optional, integer) Desired number of days for which the new certificate will be valid.`,
				},
				resource.Attribute{
					Name:        "issuer_hint",
					Description: `(Optional, string) Used with valid_days to indicate the target issuer when using Trust Protection Platform. Relevant values are: "DigiCert", "Entrust", and "Microsoft".`,
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
				resource.Attribute{
					Name:        "pkcs12",
					Description: `A base64-encoded PKCS#12 keystore secured by the ` + "`" + `key_password` + "`" + `. Useful when working with resources like [azurerm_key_vault_certificate](https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate.html). ## Certificate Renewal The ` + "`" + `venafi_certificate` + "`" + ` resource handles certificate renewals as long as a ` + "`" + `terraform apply` + "`" + ` is done within the ` + "`" + `expiration_window` + "`" + ` period. Keep in mind that the ` + "`" + `expiration_window` + "`" + ` in the Terraform configuration needs to align with the renewal window of the issuing CA to achieve the desired result.`,
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
				resource.Attribute{
					Name:        "pkcs12",
					Description: `A base64-encoded PKCS#12 keystore secured by the ` + "`" + `key_password` + "`" + `. Useful when working with resources like [azurerm_key_vault_certificate](https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate.html). ## Certificate Renewal The ` + "`" + `venafi_certificate` + "`" + ` resource handles certificate renewals as long as a ` + "`" + `terraform apply` + "`" + ` is done within the ` + "`" + `expiration_window` + "`" + ` period. Keep in mind that the ` + "`" + `expiration_window` + "`" + ` in the Terraform configuration needs to align with the renewal window of the issuing CA to achieve the desired result.`,
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
