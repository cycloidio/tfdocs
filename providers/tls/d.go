package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
	}

	dataSourcesMap = map[string]Resource{

		"tls_public_key": 0,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
