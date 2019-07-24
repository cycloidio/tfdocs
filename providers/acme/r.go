package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "acme_certificate",
			Category:         "Resources",
			ShortDescription: `Provides a resource to manage certificates on an ACME CA.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "common_name",
					Description: `The certificate's common name, the primary domain that the certificate will be recognized for. Required when not specifying a CSR.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `The certificate's subject alternative names, domains that this certificate will also be recognized for. Only valid when not specifying a CSR.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `The key type for the certificate's private key. Can be one of: ` + "`" + `P256` + "`" + ` and ` + "`" + `P384` + "`" + ` (for ECDSA keys of respective length) or ` + "`" + `2048` + "`" + `, ` + "`" + `4096` + "`" + `, and ` + "`" + `8192` + "`" + ` (for RSA keys of respective length). Required when not specifying a CSR. The default is ` + "`" + `2048` + "`" + ` (RSA key of 2048 bits).`,
				},
				resource.Attribute{
					Name:        "certificate_request_pem",
					Description: `A pre-created certificate request, such as one from [` + "`" + `tls_cert_request` + "`" + `][tls-cert-request], or one from an external source, in PEM format. Either this, or the in-resource request options (` + "`" + `common_name` + "`" + `, ` + "`" + `key_type` + "`" + `, and optionally ` + "`" + `subject_alternative_names` + "`" + `) need to be specified.`,
				},
				resource.Attribute{
					Name:        "certificate_p12_password",
					Description: `(Optional) Password to be used when generating the PFX file stored in [` + "`" + `certificate_p12` + "`" + `](#certificate_p12). Defaults to an empty string. ### Using DNS challenges As the usage model of Terraform generally sees it as being run on a different server than a certificate would normally be placed on, the ` + "`" + `acme_certifiate` + "`" + ` resource only supports DNS challenges. This method authenticates certificate domains by requiring the requester to place a TXT record on the FQDNs in the certificate. The ACME provider responds to DNS challenges automatically by utilizing one of the supported DNS challenge providers. Most providers take credentials as environment variables, but if you would rather use configuration for this purpose, you can by specifying ` + "`" + `config` + "`" + ` blocks within a [` + "`" + `dns_challenge` + "`" + `](#dns_challenge) block, along with the ` + "`" + `provider` + "`" + ` parameter. For a full list of providers, click [here][dns-challenge-providers]. [dns-challenge-providers]: /docs/providers/acme/dns_providers/index.html Example with the [Route 53 provider][route53-dns-provider]: [route53-dns-provider]: /docs/providers/acme/dns_providers/route53.html ` + "`" + `` + "`" + `` + "`" + `hcl resource "acme_certificate" "certificate" { #... dns_challenge { provider = "route53" config = { AWS_ACCESS_KEY_ID = "${var.aws_access_key}" AWS_SECRET_ACCESS_KEY = "${var.aws_secret_key}" AWS_DEFAULT_REGION = "us-east-1" } } #... } ` + "`" + `` + "`" + `` + "`" + ` #### Using Variable Files for Provider Arguments Most provider arguments can be suffixed with ` + "`" + `_FILE` + "`" + ` to specify that you wish to store that value in a local file. This can be useful if local storage for these values is desired over configuration as variables or within the environment. Building on the above [Route 53 provider][route53-dns-provider] example, the following example uses local files to get the access key ID and secret access key. ` + "`" + `` + "`" + `` + "`" + `hcl resource "acme_certificate" "certificate" { #... dns_challenge { provider = "route53" config = { AWS_ACCESS_KEY_ID_FILE = "/data/secrets/aws_access_key_id" AWS_SECRET_ACCESS_KEY_FILE = "/data/secrets/aws_secret_access_key" AWS_DEFAULT_REGION = "us-east-1" } } #... } ` + "`" + `` + "`" + `` + "`" + ` #### Manually specifying recursive nameservers for propagation checks The ACME provider will normally use your system-configured DNS resolvers to check for propagation of the TXT records before proceeding with the certificate request. In split horizon scenarios, this check may never succeed, as the machine running Terraform may not have visibility into these public DNS records. To override this default behavior, supply the ` + "`" + `recursive_nameservers` + "`" + ` to use as a list in ` + "`" + `host:port` + "`" + ` form within the ` + "`" + `dns_challenge` + "`" + ` block: ` + "`" + `` + "`" + `` + "`" + `hcl resource "acme_certificate" "certificate" { #... recursive_nameservers = ["8.8.8.8:53"] dns_challenge { provider = "route53" } #... } ` + "`" + `` + "`" + `` + "`" + ` #### Using multiple primary DNS providers The ACME provider will allow you to configure multiple DNS challenges in the event that you have more than one primary DNS provider. ` + "`" + `` + "`" + `` + "`" + `hcl resource "acme_certificate" "certificate" { #... dns_challenge { provider = "azure" } dns_challenge { provider = "gcloud" } dns_challenge { provider = "route53" } #... } ` + "`" + `` + "`" + `` + "`" + ` Some considerations need to be kept in mind when using multiple providers:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The full URL of the certificate within the ACME CA.`,
				},
				resource.Attribute{
					Name:        "certificate_url",
					Description: `The full URL of the certificate within the ACME CA. Same as ` + "`" + `id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_domain",
					Description: `The common name of the certificate.`,
				},
				resource.Attribute{
					Name:        "private_key_pem",
					Description: `The certificate's private key, in PEM format, if the certificate was generated from scratch and not with [` + "`" + `certificate_request_pem` + "`" + `](#certificate_request_pem). If ` + "`" + `certificate_request_pem` + "`" + ` was used, this will be blank.`,
				},
				resource.Attribute{
					Name:        "certificate_pem",
					Description: `The certificate in PEM format.`,
				},
				resource.Attribute{
					Name:        "issuer_pem",
					Description: `The intermediate certificate of the issuer.`,
				},
				resource.Attribute{
					Name:        "certificate_p12",
					Description: `The certificate, intermediate, and the private key archived as a PFX file (PKCS12 format, generally used by Microsoft products). The data is base64 encoded (including padding), and its password is configurable via the [` + "`" + `certificate_p12_password` + "`" + `](#certificate_p12_password) argument. This field is empty if creating a certificate from a CSR.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The full URL of the certificate within the ACME CA.`,
				},
				resource.Attribute{
					Name:        "certificate_url",
					Description: `The full URL of the certificate within the ACME CA. Same as ` + "`" + `id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_domain",
					Description: `The common name of the certificate.`,
				},
				resource.Attribute{
					Name:        "private_key_pem",
					Description: `The certificate's private key, in PEM format, if the certificate was generated from scratch and not with [` + "`" + `certificate_request_pem` + "`" + `](#certificate_request_pem). If ` + "`" + `certificate_request_pem` + "`" + ` was used, this will be blank.`,
				},
				resource.Attribute{
					Name:        "certificate_pem",
					Description: `The certificate in PEM format.`,
				},
				resource.Attribute{
					Name:        "issuer_pem",
					Description: `The intermediate certificate of the issuer.`,
				},
				resource.Attribute{
					Name:        "certificate_p12",
					Description: `The certificate, intermediate, and the private key archived as a PFX file (PKCS12 format, generally used by Microsoft products). The data is base64 encoded (including padding), and its password is configurable via the [` + "`" + `certificate_p12_password` + "`" + `](#certificate_p12_password) argument. This field is empty if creating a certificate from a CSR.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "acme_registration",
			Category:         "Resources",
			ShortDescription: `Provides a resource to manage accounts on an ACME CA.`,
			Description:      ``,
			Keywords: []string{
				"registration",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"acme_certificate":  0,
		"acme_registration": 1,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
