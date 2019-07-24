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
			ShortDescription: `Creates a PEM-encoded certificate request.`,
			Description:      ``,
			Keywords: []string{
				"cert",
				"request",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_algorithm",
					Description: `(Required) The name of the algorithm for the key provided in ` + "`" + `private_key_pem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_key_pem",
					Description: `(Required) PEM-encoded private key data. This can be read from a separate file using the ` + "`" + `` + "`" + `file` + "`" + `` + "`" + ` interpolation function. Only an irreversable secure hash of the private key will be stored in the Terraform state.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `(Required) The subject for which a certificate is being requested. This is a nested configuration block whose structure is described below.`,
				},
				resource.Attribute{
					Name:        "dns_names",
					Description: `(Optional) List of DNS names for which a certificate is being requested.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Optional) List of IP addresses for which a certificate is being requested. The nested ` + "`" + `subject` + "`" + ` block accepts the following arguments, all optional, with their meaning corresponding to the similarly-named attributes defined in [RFC5280](https://tools.ietf.org/html/rfc5280#section-4.1.2.4):`,
				},
				resource.Attribute{
					Name:        "cert_request_pem",
					Description: `The certificate request data in PEM format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_request_pem",
					Description: `The certificate request data in PEM format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tls_locally_signed_cert",
			Category:         "Resources",
			ShortDescription: `Creates a locally-signed TLS certificate in PEM format.`,
			Description:      ``,
			Keywords: []string{
				"locally",
				"signed",
				"cert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_request_pem",
					Description: `(Required) PEM-encoded request certificate data.`,
				},
				resource.Attribute{
					Name:        "ca_key_algorithm",
					Description: `(Required) The name of the algorithm for the key provided in ` + "`" + `ca_private_key_pem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_private_key_pem",
					Description: `(Required) PEM-encoded private key data for the CA. This can be read from a separate file using the ` + "`" + `` + "`" + `file` + "`" + `` + "`" + ` interpolation function.`,
				},
				resource.Attribute{
					Name:        "ca_cert_pem",
					Description: `(Required) PEM-encoded certificate data for the CA.`,
				},
				resource.Attribute{
					Name:        "validity_period_hours",
					Description: `(Required) The number of hours after initial issuing that the certificate will become invalid.`,
				},
				resource.Attribute{
					Name:        "allowed_uses",
					Description: `(Required) List of keywords each describing a use that is permitted for the issued certificate. The valid keywords are listed below.`,
				},
				resource.Attribute{
					Name:        "early_renewal_hours",
					Description: `(Optional) If set, the resource will consider the certificate to have expired the given number of hours before its actual expiry time. This can be useful to deploy an updated certificate in advance of the expiration of the current certificate. Note however that the old certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate revocation. Note also that this advance update can only be performed should the Terraform configuration be applied during the early renewal period.`,
				},
				resource.Attribute{
					Name:        "is_ca_certificate",
					Description: `(Optional) Boolean controlling whether the CA flag will be set in the generated certificate. Defaults to ` + "`" + `false` + "`" + `, meaning that the certificate does not represent a certificate authority. The ` + "`" + `allowed_uses` + "`" + ` list accepts the following keywords, combining the set of flags defined by both [Key Usage](https://tools.ietf.org/html/rfc5280#section-4.2.1.3) and [Extended Key Usage](https://tools.ietf.org/html/rfc5280#section-4.2.1.12) in [RFC5280](https://tools.ietf.org/html/rfc5280):`,
				},
				resource.Attribute{
					Name:        "cert_pem",
					Description: `The certificate data in PEM format.`,
				},
				resource.Attribute{
					Name:        "validity_start_time",
					Description: `The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.`,
				},
				resource.Attribute{
					Name:        "validity_end_time",
					Description: `The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp. ## Automatic Renewal This resource considers its instances to have been deleted after either their validity periods ends or the early renewal period is reached. At this time, applying the Terraform configuration will cause a new certificate to be generated for the instance. Therefore in a development environment with frequent deployments it may be convenient to set a relatively-short expiration time and use early renewal to automatically provision a new certificate when the current one is about to expire. The creation of a new certificate may of course cause dependent resources to be updated or replaced, depending on the lifecycle rules applying to those resources.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_pem",
					Description: `The certificate data in PEM format.`,
				},
				resource.Attribute{
					Name:        "validity_start_time",
					Description: `The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.`,
				},
				resource.Attribute{
					Name:        "validity_end_time",
					Description: `The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp. ## Automatic Renewal This resource considers its instances to have been deleted after either their validity periods ends or the early renewal period is reached. At this time, applying the Terraform configuration will cause a new certificate to be generated for the instance. Therefore in a development environment with frequent deployments it may be convenient to set a relatively-short expiration time and use early renewal to automatically provision a new certificate when the current one is about to expire. The creation of a new certificate may of course cause dependent resources to be updated or replaced, depending on the lifecycle rules applying to those resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tls_private_key",
			Category:         "Resources",
			ShortDescription: `Creates a PEM-encoded private key.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) The name of the algorithm to use for the key. Currently-supported values are "RSA" and "ECDSA".`,
				},
				resource.Attribute{
					Name:        "rsa_bits",
					Description: `(Optional) When ` + "`" + `algorithm` + "`" + ` is "RSA", the size of the generated RSA key in bits. Defaults to 2048.`,
				},
				resource.Attribute{
					Name:        "ecdsa_curve",
					Description: `(Optional) When ` + "`" + `algorithm` + "`" + ` is "ECDSA", the name of the elliptic curve to use. May be any one of "P224", "P256", "P384" or "P521", with "P224" as the default. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The algorithm that was selected for the key.`,
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
					Description: `The md5 hash of the public key data in OpenSSH MD5 hash format, e.g. ` + "`" + `aa:bb:cc:...` + "`" + `. Only available if the selected private key format is compatible, as per the rules for ` + "`" + `public_key_openssh` + "`" + `. ## Generating a New Key Since a private key is a logical resource that lives only in the Terraform state, it will persist until it is explicitly destroyed by the user. In order to force the generation of a new key within an existing state, the private key instance can be "tainted": ` + "`" + `` + "`" + `` + "`" + ` terraform taint tls_private_key.example ` + "`" + `` + "`" + `` + "`" + ` A new key will then be generated on the next ` + "`" + `` + "`" + `terraform apply` + "`" + `` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "algorithm",
					Description: `The algorithm that was selected for the key.`,
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
					Description: `The md5 hash of the public key data in OpenSSH MD5 hash format, e.g. ` + "`" + `aa:bb:cc:...` + "`" + `. Only available if the selected private key format is compatible, as per the rules for ` + "`" + `public_key_openssh` + "`" + `. ## Generating a New Key Since a private key is a logical resource that lives only in the Terraform state, it will persist until it is explicitly destroyed by the user. In order to force the generation of a new key within an existing state, the private key instance can be "tainted": ` + "`" + `` + "`" + `` + "`" + ` terraform taint tls_private_key.example ` + "`" + `` + "`" + `` + "`" + ` A new key will then be generated on the next ` + "`" + `` + "`" + `terraform apply` + "`" + `` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tls_self_signed_cert",
			Category:         "Resources",
			ShortDescription: `Creates a self-signed TLS certificate in PEM format.`,
			Description:      ``,
			Keywords: []string{
				"self",
				"signed",
				"cert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_algorithm",
					Description: `(Required) The name of the algorithm for the key provided in ` + "`" + `private_key_pem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_key_pem",
					Description: `(Required) PEM-encoded private key data. This can be read from a separate file using the ` + "`" + `` + "`" + `file` + "`" + `` + "`" + ` interpolation function. If the certificate is being generated to be used for a throwaway development environment or other non-critical application, the ` + "`" + `tls_private_key` + "`" + ` resource can be used to generate a TLS private key from within Terraform. Only an irreversable secure hash of the private key will be stored in the Terraform state.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `(Required) The subject for which a certificate is being requested. This is a nested configuration block whose structure matches the corresponding block for [` + "`" + `tls_cert_request` + "`" + `](cert_request.html).`,
				},
				resource.Attribute{
					Name:        "validity_period_hours",
					Description: `(Required) The number of hours after initial issuing that the certificate will become invalid.`,
				},
				resource.Attribute{
					Name:        "allowed_uses",
					Description: `(Required) List of keywords each describing a use that is permitted for the issued certificate. The valid keywords are listed below.`,
				},
				resource.Attribute{
					Name:        "dns_names",
					Description: `(Optional) List of DNS names for which a certificate is being requested.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Optional) List of IP addresses for which a certificate is being requested.`,
				},
				resource.Attribute{
					Name:        "early_renewal_hours",
					Description: `(Optional) If set, the resource will consider the certificate to have expired the given number of hours before its actual expiry time. This can be useful to deploy an updated certificate in advance of the expiration of the current certificate. Note however that the old certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate revocation. Note also that this advance update can only be performed should the Terraform configuration be applied during the early renewal period.`,
				},
				resource.Attribute{
					Name:        "is_ca_certificate",
					Description: `(Optional) Boolean controlling whether the CA flag will be set in the generated certificate. Defaults to ` + "`" + `false` + "`" + `, meaning that the certificate does not represent a certificate authority. The ` + "`" + `allowed_uses` + "`" + ` list accepts the following keywords, combining the set of flags defined by both [Key Usage](https://tools.ietf.org/html/rfc5280#section-4.2.1.3) and [Extended Key Usage](https://tools.ietf.org/html/rfc5280#section-4.2.1.12) in [RFC5280](https://tools.ietf.org/html/rfc5280):`,
				},
				resource.Attribute{
					Name:        "cert_pem",
					Description: `The certificate data in PEM format.`,
				},
				resource.Attribute{
					Name:        "validity_start_time",
					Description: `The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.`,
				},
				resource.Attribute{
					Name:        "validity_end_time",
					Description: `The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp. ## Automatic Renewal This resource considers its instances to have been deleted after either their validity periods ends or the early renewal period is reached. At this time, applying the Terraform configuration will cause a new certificate to be generated for the instance. Therefore in a development environment with frequent deployments it may be convenient to set a relatively-short expiration time and use early renewal to automatically provision a new certificate when the current one is about to expire. The creation of a new certificate may of course cause dependent resources to be updated or replaced, depending on the lifecycle rules applying to those resources.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_pem",
					Description: `The certificate data in PEM format.`,
				},
				resource.Attribute{
					Name:        "validity_start_time",
					Description: `The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.`,
				},
				resource.Attribute{
					Name:        "validity_end_time",
					Description: `The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp. ## Automatic Renewal This resource considers its instances to have been deleted after either their validity periods ends or the early renewal period is reached. At this time, applying the Terraform configuration will cause a new certificate to be generated for the instance. Therefore in a development environment with frequent deployments it may be convenient to set a relatively-short expiration time and use early renewal to automatically provision a new certificate when the current one is about to expire. The creation of a new certificate may of course cause dependent resources to be updated or replaced, depending on the lifecycle rules applying to those resources.`,
				},
			},
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
