package aws

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "aws_acm_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Amazon Certificate Manager (ACM) Certificate`,
			Description: `

Use this data source to get the ARN of a certificate in AWS Certificate
Manager (ACM), you can reference
it by domain without having to hard code the ARNs as input.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain of the certificate to look up. If no certificate is found with this name, an error will be returned.`,
				},
				resource.Attribute{
					Name:        "key_types",
					Description: `(Optional) A list of key algorithms to filter certificates. By default, ACM does not return all certificate types when searching. Valid values are ` + "`" + `RSA_1024` + "`" + `, ` + "`" + `RSA_2048` + "`" + `, ` + "`" + `RSA_4096` + "`" + `, ` + "`" + `EC_prime256v1` + "`" + `, ` + "`" + `EC_secp384r1` + "`" + `, and ` + "`" + `EC_secp521r1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "statuses",
					Description: `(Optional) A list of statuses on which to filter the returned list. Valid values are ` + "`" + `PENDING_VALIDATION` + "`" + `, ` + "`" + `ISSUED` + "`" + `, ` + "`" + `INACTIVE` + "`" + `, ` + "`" + `EXPIRED` + "`" + `, ` + "`" + `VALIDATION_TIMED_OUT` + "`" + `, ` + "`" + `REVOKED` + "`" + ` and ` + "`" + `FAILED` + "`" + `. If no value is specified, only certificates in the ` + "`" + `ISSUED` + "`" + ` state are returned.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `(Optional) A list of types on which to filter the returned list. Valid values are ` + "`" + `AMAZON_ISSUED` + "`" + ` and ` + "`" + `IMPORTED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If set to true, it sorts the certificates matched by previous criteria by the NotBefore field, returning only the most recent one. If set to false, it returns an error if more than one certificate is found. Defaults to false. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Set to the ARN of the found certificate, suitable for referencing in other resources that support ACM certificates.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Set to the ARN of the found certificate, suitable for referencing in other resources that support ACM certificates.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_acmpca_certificate_authority",
			Category:         "Data Sources",
			ShortDescription: `Get information on a AWS Certificate Manager Private Certificate Authority`,
			Description: `

Get information on a AWS Certificate Manager Private Certificate Authority (ACM PCA Certificate Authority).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `(Required) Amazon Resource Name (ARN) of the certificate authority. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Amazon Resource Name (ARN) of the certificate authority.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "certificate_signing_request",
					Description: `The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration",
					Description: `Nested attribute containing revocation configuration.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration",
					Description: `Nested attribute containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration.0.custom_cname",
					Description: `Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration.0.enabled",
					Description: `Boolean value that specifies whether certificate revocation lists (CRLs) are enabled.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration.0.expiration_in_days",
					Description: `Number of days until a certificate expires.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration.0.s3_bucket_name",
					Description: `Name of the S3 bucket that contains the CRL.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the certificate authority.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Specifies a key-value map of user-defined tags that are attached to the certificate authority.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the certificate authority.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Amazon Resource Name (ARN) of the certificate authority.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "certificate_signing_request",
					Description: `The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration",
					Description: `Nested attribute containing revocation configuration.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration",
					Description: `Nested attribute containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration.0.custom_cname",
					Description: `Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration.0.enabled",
					Description: `Boolean value that specifies whether certificate revocation lists (CRLs) are enabled.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration.0.expiration_in_days",
					Description: `Number of days until a certificate expires.`,
				},
				resource.Attribute{
					Name:        "revocation_configuration.0.crl_configuration.0.s3_bucket_name",
					Description: `Name of the S3 bucket that contains the CRL.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the certificate authority.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Specifies a key-value map of user-defined tags that are attached to the certificate authority.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the certificate authority.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ami",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Amazon Machine Image (AMI).`,
			Description: `

Use this data source to get the ID of a registered AMI for use in other
resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owners",
					Description: `(Required) List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, ` + "`" + `self` + "`" + ` (the current account), or an AWS owner alias (e.g. ` + "`" + `amazon` + "`" + `, ` + "`" + `aws-marketplace` + "`" + `, ` + "`" + `microsoft` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent AMI.`,
				},
				resource.Attribute{
					Name:        "executable_users",
					Description: `(Optional) Limit search to users with`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-images in the AWS CLI reference][1].`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the AMI list returned by AWS. This allows more advanced filtering not supported from the AWS API. This filtering is done locally on what AWS returns, and could have a performance impact if the result is large. It is recommended to combine this with other options to narrow down the list AWS returns. ~>`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The OS architecture of the AMI (ie: ` + "`" + `i386` + "`" + ` or ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `The block device mappings of the AMI.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.device_name",
					Description: `The physical name of the device.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.delete_on_termination",
					Description: `` + "`" + `true` + "`" + ` if the EBS volume will be deleted on termination.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.encrypted",
					Description: `` + "`" + `true` + "`" + ` if the EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.iops",
					Description: `` + "`" + `0` + "`" + ` if the EBS volume is not a provisioned IOPS image, otherwise the supported IOPS count.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.volume_size",
					Description: `The size of the volume, in GiB.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.no_device",
					Description: `Suppresses the specified device included in the block device mapping of the AMI.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.virtual_name",
					Description: `The virtual device name (for instance stores).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time the image was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the AMI that was provided during image creation.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the AMI. Should be the same as the resource ` + "`" + `id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_location",
					Description: `The location of the AMI.`,
				},
				resource.Attribute{
					Name:        "image_owner_alias",
					Description: `The AWS account alias (for example, ` + "`" + `amazon` + "`" + `, ` + "`" + `self` + "`" + `) or the AWS account ID of the AMI owner.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of image.`,
				},
				resource.Attribute{
					Name:        "kernel_id",
					Description: `The kernel associated with the image, if any. Only applicable for machine images.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the AMI that was provided during image creation.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The AWS account ID of the image owner.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The value is Windows for ` + "`" + `Windows` + "`" + ` AMIs; otherwise blank.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `Any product codes associated with the AMI.`,
				},
				resource.Attribute{
					Name:        "product_codes.#.product_code_id",
					Description: `The product code.`,
				},
				resource.Attribute{
					Name:        "product_codes.#.product_code_type",
					Description: `The type of product code.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `` + "`" + `true` + "`" + ` if the image has public launch permissions.`,
				},
				resource.Attribute{
					Name:        "ramdisk_id",
					Description: `The RAM disk associated with the image, if any. Only applicable for machine images.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The device name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device (ie: ` + "`" + `ebs` + "`" + ` or ` + "`" + `instance-store` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "root_snapshot_id",
					Description: `The snapshot id associated with the root device, if any (only applies to ` + "`" + `ebs` + "`" + ` root devices).`,
				},
				resource.Attribute{
					Name:        "sriov_net_support",
					Description: `Specifies whether enhanced networking is enabled.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the AMI. If the state is ` + "`" + `available` + "`" + `, the image is successfully registered and can be used to launch an instance.`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `Describes a state change. Fields are ` + "`" + `UNSET` + "`" + ` if not available.`,
				},
				resource.Attribute{
					Name:        "state_reason.code",
					Description: `The reason code for the state change.`,
				},
				resource.Attribute{
					Name:        "state_reason.message",
					Description: `The message for the state change.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Any tags assigned to the image.`,
				},
				resource.Attribute{
					Name:        "tags.#.key",
					Description: `The key name of the tag.`,
				},
				resource.Attribute{
					Name:        "tags.#.value",
					Description: `The value of the tag.`,
				},
				resource.Attribute{
					Name:        "virtualization_type",
					Description: `The type of virtualization of the AMI (ie: ` + "`" + `hvm` + "`" + ` or ` + "`" + `paravirtual` + "`" + `). [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-images.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "architecture",
					Description: `The OS architecture of the AMI (ie: ` + "`" + `i386` + "`" + ` or ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `The block device mappings of the AMI.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.device_name",
					Description: `The physical name of the device.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.delete_on_termination",
					Description: `` + "`" + `true` + "`" + ` if the EBS volume will be deleted on termination.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.encrypted",
					Description: `` + "`" + `true` + "`" + ` if the EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.iops",
					Description: `` + "`" + `0` + "`" + ` if the EBS volume is not a provisioned IOPS image, otherwise the supported IOPS count.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.volume_size",
					Description: `The size of the volume, in GiB.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.ebs.volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.no_device",
					Description: `Suppresses the specified device included in the block device mapping of the AMI.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings.#.virtual_name",
					Description: `The virtual device name (for instance stores).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time the image was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the AMI that was provided during image creation.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the AMI. Should be the same as the resource ` + "`" + `id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_location",
					Description: `The location of the AMI.`,
				},
				resource.Attribute{
					Name:        "image_owner_alias",
					Description: `The AWS account alias (for example, ` + "`" + `amazon` + "`" + `, ` + "`" + `self` + "`" + `) or the AWS account ID of the AMI owner.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of image.`,
				},
				resource.Attribute{
					Name:        "kernel_id",
					Description: `The kernel associated with the image, if any. Only applicable for machine images.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the AMI that was provided during image creation.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The AWS account ID of the image owner.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The value is Windows for ` + "`" + `Windows` + "`" + ` AMIs; otherwise blank.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `Any product codes associated with the AMI.`,
				},
				resource.Attribute{
					Name:        "product_codes.#.product_code_id",
					Description: `The product code.`,
				},
				resource.Attribute{
					Name:        "product_codes.#.product_code_type",
					Description: `The type of product code.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `` + "`" + `true` + "`" + ` if the image has public launch permissions.`,
				},
				resource.Attribute{
					Name:        "ramdisk_id",
					Description: `The RAM disk associated with the image, if any. Only applicable for machine images.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The device name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device (ie: ` + "`" + `ebs` + "`" + ` or ` + "`" + `instance-store` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "root_snapshot_id",
					Description: `The snapshot id associated with the root device, if any (only applies to ` + "`" + `ebs` + "`" + ` root devices).`,
				},
				resource.Attribute{
					Name:        "sriov_net_support",
					Description: `Specifies whether enhanced networking is enabled.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the AMI. If the state is ` + "`" + `available` + "`" + `, the image is successfully registered and can be used to launch an instance.`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `Describes a state change. Fields are ` + "`" + `UNSET` + "`" + ` if not available.`,
				},
				resource.Attribute{
					Name:        "state_reason.code",
					Description: `The reason code for the state change.`,
				},
				resource.Attribute{
					Name:        "state_reason.message",
					Description: `The message for the state change.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Any tags assigned to the image.`,
				},
				resource.Attribute{
					Name:        "tags.#.key",
					Description: `The key name of the tag.`,
				},
				resource.Attribute{
					Name:        "tags.#.value",
					Description: `The value of the tag.`,
				},
				resource.Attribute{
					Name:        "virtualization_type",
					Description: `The type of virtualization of the AMI (ie: ` + "`" + `hvm` + "`" + ` or ` + "`" + `paravirtual` + "`" + `). [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-images.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ami_ids",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of AMI IDs.`,
			Description: `

Use this data source to get a list of AMI IDs matching the specified criteria.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owners",
					Description: `(Required) List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, ` + "`" + `self` + "`" + ` (the current account), or an AWS owner alias (e.g. ` + "`" + `amazon` + "`" + `, ` + "`" + `aws-marketplace` + "`" + `, ` + "`" + `microsoft` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "executable_users",
					Description: `(Optional) Limit search to users with`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-images in the AWS CLI reference][1].`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the AMI list returned by AWS. This allows more advanced filtering not supported from the AWS API. This filtering is done locally on what AWS returns, and could have a performance impact if the result is large. It is recommended to combine this with other options to narrow down the list AWS returns.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_api_gateway_api_key",
			Category:         "Data Sources",
			ShortDescription: `Get information on an API Gateway API Key`,
			Description: `

Use this data source to get the name and value of a pre-existing API Key, for
example to supply credentials for a dependency microservice.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the API Key to look up. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the ID of the API Key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Set to the name of the API Key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Set to the value of the API Key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to the ID of the API Key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Set to the name of the API Key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Set to the value of the API Key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_api_gateway_resource",
			Category:         "Data Sources",
			ShortDescription: `Get information on a API Gateway Resource`,
			Description: `

Use this data source to get the id of a Resource in API Gateway. 
To fetch the Resource, you must provide the REST API id as well as the full path.  

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rest_api_id",
					Description: `(Required) The REST API id that owns the resource. If no REST API is found, an error will be returned.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The full path of the resource. If no path is found, an error will be returned. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the ID of the found Resource.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Set to the ID of the parent Resource.`,
				},
				resource.Attribute{
					Name:        "path_part",
					Description: `Set to the path relative to the parent Resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to the ID of the found Resource.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Set to the ID of the parent Resource.`,
				},
				resource.Attribute{
					Name:        "path_part",
					Description: `Set to the path relative to the parent Resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_api_gateway_rest_api",
			Category:         "Data Sources",
			ShortDescription: `Get information on a API Gateway REST API`,
			Description: `

Use this data source to get the id and root_resource_id of a REST API in
API Gateway. To fetch the REST API you must provide a name to match against. 
As there is no unique name constraint on REST APIs this data source will 
error if there is more than one match.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the ID of the found REST API.`,
				},
				resource.Attribute{
					Name:        "root_resource_id",
					Description: `Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to the ID of the found REST API.`,
				},
				resource.Attribute{
					Name:        "root_resource_id",
					Description: `Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_api_gateway_vpc_link",
			Category:         "Data Sources",
			ShortDescription: `Get information on a API Gateway VPC Link`,
			Description: `

Use this data source to get the id of a VPC Link in
API Gateway. To fetch the VPC Link you must provide a name to match against. 
As there is no unique name constraint on API Gateway VPC Links this data source will 
error if there is more than one match.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the API Gateway VPC Link to look up. If no API Gateway VPC Link is found with this name, an error will be returned. If multiple API Gateway VPC Links are found with this name, an error will be returned. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the ID of the found API Gateway VPC Link.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to the ID of the found API Gateway VPC Link.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_arn",
			Category:         "Data Sources",
			ShortDescription: `Parses an ARN into its constituent parts.`,
			Description: `

Parses an Amazon Resource Name (ARN) into its constituent parts.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `(Required) The ARN to parse. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `The partition that the resource is in.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `The [service namespace](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces) that identifies the AWS product.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the resource resides in. Note that the ARNs for some resources do not require a region, so this component might be omitted.`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `The [ID](https://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html) of the AWS account that owns the resource, without the hyphens.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `The content of this part of the ARN varies by service. It often includes an indicator of the type of resource—for example, an IAM user or Amazon RDS database —followed by a slash (/) or a colon (:), followed by the resource name itself.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "partition",
					Description: `The partition that the resource is in.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `The [service namespace](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces) that identifies the AWS product.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the resource resides in. Note that the ARNs for some resources do not require a region, so this component might be omitted.`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `The [ID](https://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html) of the AWS account that owns the resource, without the hyphens.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `The content of this part of the ARN varies by service. It often includes an indicator of the type of resource—for example, an IAM user or Amazon RDS database —followed by a slash (/) or a colon (:), followed by the resource name itself.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_autoscaling_group",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Amazon EC2 Autoscaling Group.`,
			Description: `

Use this data source to get information on an existing autoscaling group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specify the exact name of the desired autoscaling group. ## Attributes Reference ~>`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the Auto Scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Auto Scaling group.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `One or more Availability Zones for the group.`,
				},
				resource.Attribute{
					Name:        "default_cool_down",
					Description: `The amount of time, in seconds, after a scaling activity completes before another scaling activity can start.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The desired size of the group.`,
				},
				resource.Attribute{
					Name:        "health_check_grace_period",
					Description: `The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `The service to use for the health checks. The valid values are EC2 and ELB.`,
				},
				resource.Attribute{
					Name:        "launch_configuration",
					Description: `The name of the associated launch configuration.`,
				},
				resource.Attribute{
					Name:        "load_balancers",
					Description: `One or more load balancers associated with the group.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum size of the group.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum size of the group.`,
				},
				resource.Attribute{
					Name:        "placement_group",
					Description: `The name of the placement group into which to launch your instances, if any. For more information, see Placement Groups (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the Amazon Elastic Compute Cloud User Guide.`,
				},
				resource.Attribute{
					Name:        "service_linked_role_arn",
					Description: `The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current state of the group when DeleteAutoScalingGroup is in progress.`,
				},
				resource.Attribute{
					Name:        "target_group_arns",
					Description: `The Amazon Resource Names (ARN) of the target groups for your load balancer.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `The termination policies for the group.`,
				},
				resource.Attribute{
					Name:        "vpc_zone_identifier",
					Description: `VPC ID for the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the Auto Scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Auto Scaling group.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `One or more Availability Zones for the group.`,
				},
				resource.Attribute{
					Name:        "default_cool_down",
					Description: `The amount of time, in seconds, after a scaling activity completes before another scaling activity can start.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The desired size of the group.`,
				},
				resource.Attribute{
					Name:        "health_check_grace_period",
					Description: `The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `The service to use for the health checks. The valid values are EC2 and ELB.`,
				},
				resource.Attribute{
					Name:        "launch_configuration",
					Description: `The name of the associated launch configuration.`,
				},
				resource.Attribute{
					Name:        "load_balancers",
					Description: `One or more load balancers associated with the group.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum size of the group.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum size of the group.`,
				},
				resource.Attribute{
					Name:        "placement_group",
					Description: `The name of the placement group into which to launch your instances, if any. For more information, see Placement Groups (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the Amazon Elastic Compute Cloud User Guide.`,
				},
				resource.Attribute{
					Name:        "service_linked_role_arn",
					Description: `The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current state of the group when DeleteAutoScalingGroup is in progress.`,
				},
				resource.Attribute{
					Name:        "target_group_arns",
					Description: `The Amazon Resource Names (ARN) of the target groups for your load balancer.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `The termination policies for the group.`,
				},
				resource.Attribute{
					Name:        "vpc_zone_identifier",
					Description: `VPC ID for the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_autoscaling_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Autoscaling Groups within a specific region.`,
			Description: `

The Autoscaling Groups data source allows access to the list of AWS
ASGs within a specific region. This will allow you to pass a list of AutoScaling Groups to other resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags. See [related docs](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Filter.html).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are: ` + "`" + `auto-scaling-group` + "`" + `, ` + "`" + `key` + "`" + `, ` + "`" + `value` + "`" + `, and ` + "`" + `propagate-at-launch` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of the Autoscaling Groups in the current region.`,
				},
				resource.Attribute{
					Name:        "arns",
					Description: `A list of the Autoscaling Groups Arns in the current region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of the Autoscaling Groups in the current region.`,
				},
				resource.Attribute{
					Name:        "arns",
					Description: `A list of the Autoscaling Groups Arns in the current region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_availability_zone",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific availability zone`,
			Description: `

` + "`" + `aws_availability_zone` + "`" + ` provides details about a specific availability zone (AZ)
in the current region.

This can be used both to validate an availability zone given in a variable
and to split the AZ name into its component parts of an AWS region and an
AZ identifier letter. The latter may be useful e.g. for implementing a
consistent subnet numbering scheme across several regions by mapping both
the region and the subnet letter to network numbers.

This is different from the ` + "`" + `aws_availability_zones` + "`" + ` (plural) data source,
which provides a list of the available zones.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The full name of the availability zone to select.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) A specific availability zone state to require. May be any of ` + "`" + `"available"` + "`" + `, ` + "`" + `"information"` + "`" + ` or ` + "`" + `"impaired"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The zone ID of the availability zone to select. All reasonable uses of this data source will specify ` + "`" + `name` + "`" + `, since ` + "`" + `state` + "`" + ` alone would match a single AZ only in a region that itself has only one AZ. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the selected availability zone.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the selected availability zone resides. This is always the region selected on the provider, since this data source searches only within that region.`,
				},
				resource.Attribute{
					Name:        "name_suffix",
					Description: `The part of the AZ name that appears after the region name, uniquely identifying the AZ within its region.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the AZ.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The zone ID of the selected availability zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the selected availability zone.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the selected availability zone resides. This is always the region selected on the provider, since this data source searches only within that region.`,
				},
				resource.Attribute{
					Name:        "name_suffix",
					Description: `The part of the AZ name that appears after the region name, uniquely identifying the AZ within its region.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the AZ.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The zone ID of the selected availability zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_availability_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Availability Zones which can be used by an AWS account.`,
			Description: `

The Availability Zones data source allows access to the list of AWS
Availability Zones which can be accessed by an AWS account within the region
configured in the provider.

This is different from the ` + "`" + `aws_availability_zone` + "`" + ` (singular) data source,
which provides some details about a specific availability zone.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "blacklisted_names",
					Description: `(Optional) List of blacklisted Availability Zone names.`,
				},
				resource.Attribute{
					Name:        "blacklisted_zone_ids",
					Description: `(Optional) List of blacklisted Availability Zone IDs.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Allows to filter list of Availability Zones based on their current state. Can be either ` + "`" + `"available"` + "`" + `, ` + "`" + `"information"` + "`" + `, ` + "`" + `"impaired"` + "`" + ` or ` + "`" + `"unavailable"` + "`" + `. By default the list includes a complete set of Availability Zones to which the underlying AWS account has access, regardless of their state. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of the Availability Zone names available to the account.`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of the Availability Zone IDs available to the account. Note that the indexes of Availability Zone names and IDs correspond.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of the Availability Zone names available to the account.`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of the Availability Zone IDs available to the account. Note that the indexes of Availability Zone names and IDs correspond.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_batch_compute_environment",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a batch compute environment`,
			Description: `

The Batch Compute Environment data source allows access to details of a specific
compute environment within AWS Batch.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_environment_name",
					Description: `(Required) The name of the Batch Compute Environment ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the compute environment.`,
				},
				resource.Attribute{
					Name:        "ecs_cluster_arn",
					Description: `The ARN of the underlying Amazon ECS cluster used by the compute environment.`,
				},
				resource.Attribute{
					Name:        "service_role",
					Description: `The ARN of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the compute environment (for example, ` + "`" + `MANAGED` + "`" + ` or ` + "`" + `UNMANAGED` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the compute environment (for example, ` + "`" + `CREATING` + "`" + ` or ` + "`" + `VALID` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status_reason",
					Description: `A short, human-readable string to provide additional details about the current status of the compute environment.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the compute environment (for example, ` + "`" + `ENABLED` + "`" + ` or ` + "`" + `DISABLED` + "`" + `). If the state is ` + "`" + `ENABLED` + "`" + `, then the compute environment accepts jobs from a queue and can scale out automatically based on queues.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the compute environment.`,
				},
				resource.Attribute{
					Name:        "ecs_cluster_arn",
					Description: `The ARN of the underlying Amazon ECS cluster used by the compute environment.`,
				},
				resource.Attribute{
					Name:        "service_role",
					Description: `The ARN of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the compute environment (for example, ` + "`" + `MANAGED` + "`" + ` or ` + "`" + `UNMANAGED` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the compute environment (for example, ` + "`" + `CREATING` + "`" + ` or ` + "`" + `VALID` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status_reason",
					Description: `A short, human-readable string to provide additional details about the current status of the compute environment.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the compute environment (for example, ` + "`" + `ENABLED` + "`" + ` or ` + "`" + `DISABLED` + "`" + `). If the state is ` + "`" + `ENABLED` + "`" + `, then the compute environment accepts jobs from a queue and can scale out automatically based on queues.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_batch_job_queue",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a batch job queue`,
			Description: `

The Batch Job Queue data source allows access to details of a specific
job queue within AWS Batch.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the job queue. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the job queue.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the job queue (for example, ` + "`" + `CREATING` + "`" + ` or ` + "`" + `VALID` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status_reason",
					Description: `A short, human-readable string to provide additional details about the current status of the job queue.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Describes the ability of the queue to accept new jobs (for example, ` + "`" + `ENABLED` + "`" + ` or ` + "`" + `DISABLED` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the job queue. Job queues with a higher priority are evaluated first when associated with the same compute environment.`,
				},
				resource.Attribute{
					Name:        "compute_environment_order",
					Description: `The compute environments that are attached to the job queue and the order in which job placement is preferred. Compute environments are selected for job placement in ascending order.`,
				},
				resource.Attribute{
					Name:        "compute_environment_order.#.order",
					Description: `The order of the compute environment.`,
				},
				resource.Attribute{
					Name:        "compute_environment_order.#.compute_environment",
					Description: `The ARN of the compute environment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the job queue.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the job queue (for example, ` + "`" + `CREATING` + "`" + ` or ` + "`" + `VALID` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status_reason",
					Description: `A short, human-readable string to provide additional details about the current status of the job queue.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Describes the ability of the queue to accept new jobs (for example, ` + "`" + `ENABLED` + "`" + ` or ` + "`" + `DISABLED` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the job queue. Job queues with a higher priority are evaluated first when associated with the same compute environment.`,
				},
				resource.Attribute{
					Name:        "compute_environment_order",
					Description: `The compute environments that are attached to the job queue and the order in which job placement is preferred. Compute environments are selected for job placement in ascending order.`,
				},
				resource.Attribute{
					Name:        "compute_environment_order.#.order",
					Description: `The order of the compute environment.`,
				},
				resource.Attribute{
					Name:        "compute_environment_order.#.compute_environment",
					Description: `The ARN of the compute environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_billing_service_account",
			Category:         "Data Sources",
			ShortDescription: `Get AWS Billing Service Account`,
			Description: `

Use this data source to get the Account ID of the [AWS Billing and Cost Management Service Account](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-getting-started.html#step-2) for the purpose of whitelisting in S3 bucket policy.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the AWS billing service account.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the AWS billing service account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_caller_identity",
			Category:         "Data Sources",
			ShortDescription: `Get information about the identity of the caller for the provider connection to AWS.`,
			Description: `

Use this data source to get the access to the effective Account ID, User ID, and ARN in
which Terraform is authorized.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `The AWS Account ID number of the account that owns or contains the calling entity.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The AWS ARN associated with the calling entity.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The unique identifier of the calling entity.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `The AWS Account ID number of the account that owns or contains the calling entity.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The AWS ARN associated with the calling entity.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The unique identifier of the calling entity.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_canonical_user_id",
			Category:         "Data Sources",
			ShortDescription: `Provides the canonical user ID for the AWS account associated with the provider connection to AWS.`,
			Description: `

The Canonical User ID data source allows access to the [canonical user ID](http://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html)
for the effective account in which Terraform is working.  

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The canonical user ID associated with the AWS account.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The human-friendly name linked to the canonical user ID. The bucket owner's display name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The canonical user ID associated with the AWS account.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The human-friendly name linked to the canonical user ID. The bucket owner's display name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_cloudformation_export",
			Category:         "Data Sources",
			ShortDescription: `Provides metadata of a CloudFormation Export (e.g. Cross Stack References)`,
			Description: `

The CloudFormation Export data source allows access to stack
exports specified in the [Output](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html) section of the Cloudformation Template using the optional Export Property.

 -> Note: If you are trying to use a value from a Cloudformation Stack in the same Terraform run please use normal interpolation or Cloudformation Outputs. 

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the export as it appears in the console or from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value from Cloudformation export identified by the export name found from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html)`,
				},
				resource.Attribute{
					Name:        "exporting_stack_id",
					Description: `The exporting_stack_id (AWS ARNs) equivalent ` + "`" + `ExportingStackId` + "`" + ` from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `The value from Cloudformation export identified by the export name found from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html)`,
				},
				resource.Attribute{
					Name:        "exporting_stack_id",
					Description: `The exporting_stack_id (AWS ARNs) equivalent ` + "`" + `ExportingStackId` + "`" + ` from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_cloudformation_stack",
			Category:         "Data Sources",
			ShortDescription: `Provides metadata of a CloudFormation stack (e.g. outputs)`,
			Description: `

The CloudFormation Stack data source allows access to stack
outputs and other useful data including the template body.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the stack ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `A list of capabilities`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the stack`,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `Whether the rollback of the stack is disabled when stack creation fails`,
				},
				resource.Attribute{
					Name:        "notification_arns",
					Description: `A list of SNS topic ARNs to publish stack related events`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A map of outputs from the stack.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `A map of parameters that specify input parameters for the stack.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags associated with this stack.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `Structure containing the template body.`,
				},
				resource.Attribute{
					Name:        "iam_role_arn",
					Description: `The ARN of the IAM role used to create the stack.`,
				},
				resource.Attribute{
					Name:        "timeout_in_minutes",
					Description: `The amount of time that can pass before the stack status becomes ` + "`" + `CREATE_FAILED` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "capabilities",
					Description: `A list of capabilities`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the stack`,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `Whether the rollback of the stack is disabled when stack creation fails`,
				},
				resource.Attribute{
					Name:        "notification_arns",
					Description: `A list of SNS topic ARNs to publish stack related events`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A map of outputs from the stack.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `A map of parameters that specify input parameters for the stack.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags associated with this stack.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `Structure containing the template body.`,
				},
				resource.Attribute{
					Name:        "iam_role_arn",
					Description: `The ARN of the IAM role used to create the stack.`,
				},
				resource.Attribute{
					Name:        "timeout_in_minutes",
					Description: `The amount of time that can pass before the stack status becomes ` + "`" + `CREATE_FAILED` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_cloudhsm_v2_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information on a CloudHSM v2 cluster.`,
			Description: `

Use this data source to get information about a CloudHSM v2 cluster

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of Cloud HSM v2 cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_state",
					Description: `(Optional) The state of the cluster to be found. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The id of the VPC that the CloudHSM cluster resides in.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group associated with the CloudHSM cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `The IDs of subnets in which cluster operates.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates",
					Description: `The list of cluster certificates.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.cluster_certificate",
					Description: `The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.cluster_csr",
					Description: `The certificate signing request (CSR). Available only in UNINITIALIZED state.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.aws_hardware_certificate",
					Description: `The HSM hardware certificate issued (signed) by AWS CloudHSM.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.hsm_certificate",
					Description: `The HSM certificate issued (signed) by the HSM hardware.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.manufacturer_hardware_certificate",
					Description: `The HSM hardware certificate issued (signed) by the hardware manufacturer. The number of available cluster certificates may vary depending on state of the cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The id of the VPC that the CloudHSM cluster resides in.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group associated with the CloudHSM cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `The IDs of subnets in which cluster operates.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates",
					Description: `The list of cluster certificates.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.cluster_certificate",
					Description: `The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.cluster_csr",
					Description: `The certificate signing request (CSR). Available only in UNINITIALIZED state.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.aws_hardware_certificate",
					Description: `The HSM hardware certificate issued (signed) by AWS CloudHSM.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.hsm_certificate",
					Description: `The HSM certificate issued (signed) by the HSM hardware.`,
				},
				resource.Attribute{
					Name:        "cluster_certificates.0.manufacturer_hardware_certificate",
					Description: `The HSM hardware certificate issued (signed) by the hardware manufacturer. The number of available cluster certificates may vary depending on state of the cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_cloudtrail_service_account",
			Category:         "Data Sources",
			ShortDescription: `Get AWS CloudTrail Service Account ID for storing trail data in S3.`,
			Description: `

Use this data source to get the Account ID of the [AWS CloudTrail Service Account](http://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-supported-regions.html)
in a given region for the purpose of allowing CloudTrail to store trail data in S3.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Name of the region whose AWS CloudTrail account ID is desired. Defaults to the region from the AWS provider configuration. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the AWS CloudTrail service account in the selected region.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the AWS CloudTrail service account in the selected region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the AWS CloudTrail service account in the selected region.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the AWS CloudTrail service account in the selected region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_cloudwatch_log_group",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Cloudwatch Log Group.`,
			Description: `

Use this data source to get information about an AWS Cloudwatch Log Group

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Cloudwatch log group ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the Cloudwatch log group`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the Cloudwatch log group`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_codecommit_repository",
			Category:         "Data Sources",
			ShortDescription: `Provides details about CodeCommit Repository.`,
			Description: `

The CodeCommit Repository data source allows the ARN, Repository ID, Repository URL for HTTP and Repository URL for SSH to be retrieved for an CodeCommit repository.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository_name",
					Description: `(Required) The name for the repository. This needs to be less than 100 characters. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "repository_id",
					Description: `The ID of the repository`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the repository`,
				},
				resource.Attribute{
					Name:        "clone_url_http",
					Description: `The URL to use for cloning the repository over HTTPS.`,
				},
				resource.Attribute{
					Name:        "clone_url_ssh",
					Description: `The URL to use for cloning the repository over SSH.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "repository_id",
					Description: `The ID of the repository`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the repository`,
				},
				resource.Attribute{
					Name:        "clone_url_http",
					Description: `The URL to use for cloning the repository over HTTPS.`,
				},
				resource.Attribute{
					Name:        "clone_url_ssh",
					Description: `The URL to use for cloning the repository over SSH.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_cognito_user_pools",
			Category:         "Data Sources",
			ShortDescription: `Get list of cognito user pools.`,
			Description: `

Use this data source to get a list of cognito user pools.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(required) Name of the cognito user pools. Name is not a unique attribute for cognito user pool, so multiple pools might be returned with given name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `The list of cognito user pool ids.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `The list of cognito user pool ids.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_cur_report_definition",
			Category:         "Data Sources",
			ShortDescription: `Get information on an AWS Cost and Usage Report Definition.`,
			Description: `

Use this data source to get information on an AWS Cost and Usage Report Definition.

~> *NOTE:* The AWS Cost and Usage Report service is only available in ` + "`" + `us-east-1` + "`" + ` currently.

~> *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "report_name",
					Description: `(Required) The name of the report definition to match. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "time_unit",
					Description: `The frequency on which report data are measured and displayed.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Preferred compression format for report.`,
				},
				resource.Attribute{
					Name:        "compression",
					Description: `Preferred format for report.`,
				},
				resource.Attribute{
					Name:        "additional_schema_elements",
					Description: `A list of schema elements.`,
				},
				resource.Attribute{
					Name:        "s3_bucket",
					Description: `Name of customer S3 bucket.`,
				},
				resource.Attribute{
					Name:        "s3_prefix",
					Description: `Preferred report path prefix.`,
				},
				resource.Attribute{
					Name:        "s3_region",
					Description: `Region of customer S3 bucket.`,
				},
				resource.Attribute{
					Name:        "additional_artifacts",
					Description: `A list of additional artifacts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "time_unit",
					Description: `The frequency on which report data are measured and displayed.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Preferred compression format for report.`,
				},
				resource.Attribute{
					Name:        "compression",
					Description: `Preferred format for report.`,
				},
				resource.Attribute{
					Name:        "additional_schema_elements",
					Description: `A list of schema elements.`,
				},
				resource.Attribute{
					Name:        "s3_bucket",
					Description: `Name of customer S3 bucket.`,
				},
				resource.Attribute{
					Name:        "s3_prefix",
					Description: `Preferred report path prefix.`,
				},
				resource.Attribute{
					Name:        "s3_region",
					Description: `Region of customer S3 bucket.`,
				},
				resource.Attribute{
					Name:        "additional_artifacts",
					Description: `A list of additional artifacts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_customer_gateway",
			Category:         "Data Sources",
			ShortDescription: `Get an existing AWS Customer Gateway.`,
			Description: `

Get an existing AWS Customer Gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the gateway.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more [name-value pairs][dcg-filters] to filter by. [dcg-filters]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeCustomerGateways.html ## Attribute Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `(Optional) The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address of the gateway's Internet-routable external interface.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Map of key-value pairs assigned to the gateway.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of customer gateway. The only type AWS supports at this time is "ipsec.1".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `(Optional) The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address of the gateway's Internet-routable external interface.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Map of key-value pairs assigned to the gateway.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of customer gateway. The only type AWS supports at this time is "ipsec.1".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_db_cluster_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information on a DB Cluster Snapshot.`,
			Description: `

Use this data source to get information about a DB Cluster Snapshot for use when provisioning DB clusters.

~> **NOTE:** This data source does not apply to snapshots created on DB Instances. 
See the [` + "`" + `aws_db_snapshot` + "`" + ` data source](/docs/providers/aws/d/db_snapshot.html) for DB Instance snapshots.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent Snapshot.`,
				},
				resource.Attribute{
					Name:        "db_cluster_identifier",
					Description: `(Optional) Returns the list of snapshots created by the specific db_cluster`,
				},
				resource.Attribute{
					Name:        "db_cluster_snapshot_identifier",
					Description: `(Optional) Returns information on a specific snapshot_id.`,
				},
				resource.Attribute{
					Name:        "snapshot_type",
					Description: `(Optional) The type of snapshots to be returned. If you don't specify a SnapshotType value, then both automated and manual DB cluster snapshots are returned. Shared and public DB Cluster Snapshots are not included in the returned results by default. Possible values are, ` + "`" + `automated` + "`" + `, ` + "`" + `manual` + "`" + `, ` + "`" + `shared` + "`" + ` and ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "include_shared",
					Description: `(Optional) Set this value to true to include shared manual DB Cluster Snapshots from other AWS accounts that this AWS account has been given permission to copy or restore, otherwise set this value to false. The default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "include_public",
					Description: `(Optional) Set this value to true to include manual DB Cluster Snapshots that are public and can be copied or restored by any AWS account, otherwise set this value to false. The default is ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allocated_storage",
					Description: `Specifies the allocated storage size in gigabytes (GB).`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.`,
				},
				resource.Attribute{
					Name:        "db_cluster_identifier",
					Description: `Specifies the DB cluster identifier of the DB cluster that this DB cluster snapshot was created from.`,
				},
				resource.Attribute{
					Name:        "db_cluster_snapshot_arn",
					Description: `The Amazon Resource Name (ARN) for the DB Cluster Snapshot.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the database engine for this DB cluster snapshot.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Specifies the name of the database engine.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot ID.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `If storage_encrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.`,
				},
				resource.Attribute{
					Name:        "license_model",
					Description: `License model information for the restored DB cluster.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port that the DB cluster was listening on at the time of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_create_time",
					Description: `Time when the snapshot was taken, in Universal Coordinated Time (UTC).`,
				},
				resource.Attribute{
					Name:        "source_db_cluster_snapshot_identifier",
					Description: `The DB Cluster Snapshot Arn that the DB Cluster Snapshot was copied from. It only has value in case of cross customer or cross region copy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this DB Cluster Snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_encrypted",
					Description: `Specifies whether the DB cluster snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID associated with the DB cluster snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allocated_storage",
					Description: `Specifies the allocated storage size in gigabytes (GB).`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.`,
				},
				resource.Attribute{
					Name:        "db_cluster_identifier",
					Description: `Specifies the DB cluster identifier of the DB cluster that this DB cluster snapshot was created from.`,
				},
				resource.Attribute{
					Name:        "db_cluster_snapshot_arn",
					Description: `The Amazon Resource Name (ARN) for the DB Cluster Snapshot.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the database engine for this DB cluster snapshot.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Specifies the name of the database engine.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot ID.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `If storage_encrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.`,
				},
				resource.Attribute{
					Name:        "license_model",
					Description: `License model information for the restored DB cluster.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port that the DB cluster was listening on at the time of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_create_time",
					Description: `Time when the snapshot was taken, in Universal Coordinated Time (UTC).`,
				},
				resource.Attribute{
					Name:        "source_db_cluster_snapshot_identifier",
					Description: `The DB Cluster Snapshot Arn that the DB Cluster Snapshot was copied from. It only has value in case of cross customer or cross region copy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this DB Cluster Snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_encrypted",
					Description: `Specifies whether the DB cluster snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID associated with the DB cluster snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_db_event_categories",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of DB Event Categories which can be used to pass values into DB Event Subscription.`,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Optional) The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "event_categories",
					Description: `A list of the event categories.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "event_categories",
					Description: `A list of the event categories.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_db_instance",
			Category:         "Data Sources",
			ShortDescription: `Get information on an RDS Database Instance.`,
			Description: `

Use this data source to get information about an RDS instance

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_instance_identifier",
					Description: `(Required) The name of the RDS instance ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The hostname of the RDS instance. See also ` + "`" + `endpoint` + "`" + ` and ` + "`" + `port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allocated_storage",
					Description: `Specifies the allocated storage size specified in gigabytes.`,
				},
				resource.Attribute{
					Name:        "auto_minor_version_upgrade",
					Description: `Indicates that minor version patches are applied automatically.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Specifies the name of the Availability Zone the DB instance is located in.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `Specifies the number of days for which automatic DB snapshots are retained.`,
				},
				resource.Attribute{
					Name:        "db_cluster_identifier",
					Description: `If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.`,
				},
				resource.Attribute{
					Name:        "db_instance_arn",
					Description: `The Amazon Resource Name (ARN) for the DB instance.`,
				},
				resource.Attribute{
					Name:        "db_instance_class",
					Description: `Contains the name of the compute and memory capacity class of the DB instance.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `Contains the name of the initial database of this instance that was provided at create time, if one was specified when the DB instance was created. This same name is returned for the life of the DB instance.`,
				},
				resource.Attribute{
					Name:        "db_parameter_groups",
					Description: `Provides the list of DB parameter groups applied to this DB instance.`,
				},
				resource.Attribute{
					Name:        "db_security_groups",
					Description: `Provides List of DB security groups associated to this DB instance.`,
				},
				resource.Attribute{
					Name:        "db_subnet_group",
					Description: `Specifies the name of the subnet group associated with the DB instance.`,
				},
				resource.Attribute{
					Name:        "db_instance_port",
					Description: `Specifies the port that the DB instance listens on.`,
				},
				resource.Attribute{
					Name:        "enabled_cloudwatch_logs_exports",
					Description: `List of log types to export to cloudwatch.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The connection endpoint in ` + "`" + `address:port` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Provides the name of the database engine to be used for this DB instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Indicates the database engine version.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The canonical hosted zone ID of the DB instance (to be used in a Route 53 Alias record).`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `Specifies the Provisioned IOPS (I/O operations per second) value.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `If StorageEncrypted is true, the KMS key identifier for the encrypted DB instance.`,
				},
				resource.Attribute{
					Name:        "license_model",
					Description: `License model information for this DB instance.`,
				},
				resource.Attribute{
					Name:        "master_username",
					Description: `Contains the master username for the DB instance.`,
				},
				resource.Attribute{
					Name:        "monitoring_interval",
					Description: `The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.`,
				},
				resource.Attribute{
					Name:        "monitoring_role_arn",
					Description: `The ARN for the IAM role that permits RDS to send Enhanced Monitoring metrics to CloudWatch Logs.`,
				},
				resource.Attribute{
					Name:        "multi_az",
					Description: `Specifies if the DB instance is a Multi-AZ deployment.`,
				},
				resource.Attribute{
					Name:        "option_group_memberships",
					Description: `Provides the list of option group memberships for this DB instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The database port.`,
				},
				resource.Attribute{
					Name:        "preferred_backup_window",
					Description: `Specifies the daily time range during which automated backups are created.`,
				},
				resource.Attribute{
					Name:        "preferred_maintenance_window",
					Description: `Specifies the weekly time range during which system maintenance can occur in UTC.`,
				},
				resource.Attribute{
					Name:        "publicly_accessible",
					Description: `Specifies the accessibility options for the DB instance.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The RDS Resource ID of this instance.`,
				},
				resource.Attribute{
					Name:        "storage_encrypted",
					Description: `Specifies whether the DB instance is encrypted.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Specifies the storage type associated with DB instance.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The time zone of the DB instance.`,
				},
				resource.Attribute{
					Name:        "vpc_security_groups",
					Description: `Provides a list of VPC security group elements that the DB instance belongs to.`,
				},
				resource.Attribute{
					Name:        "replicate_source_db",
					Description: `The identifier of the source DB that this is a replica of.`,
				},
				resource.Attribute{
					Name:        "ca_cert_identifier",
					Description: `Specifies the identifier of the CA certificate for the DB instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `The hostname of the RDS instance. See also ` + "`" + `endpoint` + "`" + ` and ` + "`" + `port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allocated_storage",
					Description: `Specifies the allocated storage size specified in gigabytes.`,
				},
				resource.Attribute{
					Name:        "auto_minor_version_upgrade",
					Description: `Indicates that minor version patches are applied automatically.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Specifies the name of the Availability Zone the DB instance is located in.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `Specifies the number of days for which automatic DB snapshots are retained.`,
				},
				resource.Attribute{
					Name:        "db_cluster_identifier",
					Description: `If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.`,
				},
				resource.Attribute{
					Name:        "db_instance_arn",
					Description: `The Amazon Resource Name (ARN) for the DB instance.`,
				},
				resource.Attribute{
					Name:        "db_instance_class",
					Description: `Contains the name of the compute and memory capacity class of the DB instance.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `Contains the name of the initial database of this instance that was provided at create time, if one was specified when the DB instance was created. This same name is returned for the life of the DB instance.`,
				},
				resource.Attribute{
					Name:        "db_parameter_groups",
					Description: `Provides the list of DB parameter groups applied to this DB instance.`,
				},
				resource.Attribute{
					Name:        "db_security_groups",
					Description: `Provides List of DB security groups associated to this DB instance.`,
				},
				resource.Attribute{
					Name:        "db_subnet_group",
					Description: `Specifies the name of the subnet group associated with the DB instance.`,
				},
				resource.Attribute{
					Name:        "db_instance_port",
					Description: `Specifies the port that the DB instance listens on.`,
				},
				resource.Attribute{
					Name:        "enabled_cloudwatch_logs_exports",
					Description: `List of log types to export to cloudwatch.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The connection endpoint in ` + "`" + `address:port` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Provides the name of the database engine to be used for this DB instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Indicates the database engine version.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The canonical hosted zone ID of the DB instance (to be used in a Route 53 Alias record).`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `Specifies the Provisioned IOPS (I/O operations per second) value.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `If StorageEncrypted is true, the KMS key identifier for the encrypted DB instance.`,
				},
				resource.Attribute{
					Name:        "license_model",
					Description: `License model information for this DB instance.`,
				},
				resource.Attribute{
					Name:        "master_username",
					Description: `Contains the master username for the DB instance.`,
				},
				resource.Attribute{
					Name:        "monitoring_interval",
					Description: `The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.`,
				},
				resource.Attribute{
					Name:        "monitoring_role_arn",
					Description: `The ARN for the IAM role that permits RDS to send Enhanced Monitoring metrics to CloudWatch Logs.`,
				},
				resource.Attribute{
					Name:        "multi_az",
					Description: `Specifies if the DB instance is a Multi-AZ deployment.`,
				},
				resource.Attribute{
					Name:        "option_group_memberships",
					Description: `Provides the list of option group memberships for this DB instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The database port.`,
				},
				resource.Attribute{
					Name:        "preferred_backup_window",
					Description: `Specifies the daily time range during which automated backups are created.`,
				},
				resource.Attribute{
					Name:        "preferred_maintenance_window",
					Description: `Specifies the weekly time range during which system maintenance can occur in UTC.`,
				},
				resource.Attribute{
					Name:        "publicly_accessible",
					Description: `Specifies the accessibility options for the DB instance.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The RDS Resource ID of this instance.`,
				},
				resource.Attribute{
					Name:        "storage_encrypted",
					Description: `Specifies whether the DB instance is encrypted.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Specifies the storage type associated with DB instance.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The time zone of the DB instance.`,
				},
				resource.Attribute{
					Name:        "vpc_security_groups",
					Description: `Provides a list of VPC security group elements that the DB instance belongs to.`,
				},
				resource.Attribute{
					Name:        "replicate_source_db",
					Description: `The identifier of the source DB that this is a replica of.`,
				},
				resource.Attribute{
					Name:        "ca_cert_identifier",
					Description: `Specifies the identifier of the CA certificate for the DB instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_db_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information on a DB Snapshot.`,
			Description: `

Use this data source to get information about a DB Snapshot for use when provisioning DB instances

~> **NOTE:** This data source does not apply to snapshots created on Aurora DB clusters.
See the [` + "`" + `aws_db_cluster_snapshot` + "`" + ` data source](/docs/providers/aws/d/db_cluster_snapshot.html) for DB Cluster snapshots.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent Snapshot.`,
				},
				resource.Attribute{
					Name:        "db_instance_identifier",
					Description: `(Optional) Returns the list of snapshots created by the specific db_instance`,
				},
				resource.Attribute{
					Name:        "db_snapshot_identifier",
					Description: `(Optional) Returns information on a specific snapshot_id.`,
				},
				resource.Attribute{
					Name:        "snapshot_type",
					Description: `(Optional) The type of snapshots to be returned. If you don't specify a SnapshotType value, then both automated and manual snapshots are returned. Shared and public DB snapshots are not included in the returned results by default. Possible values are, ` + "`" + `automated` + "`" + `, ` + "`" + `manual` + "`" + `, ` + "`" + `shared` + "`" + ` and ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "include_shared",
					Description: `(Optional) Set this value to true to include shared manual DB snapshots from other AWS accounts that this AWS account has been given permission to copy or restore, otherwise set this value to false. The default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "include_public",
					Description: `(Optional) Set this value to true to include manual DB snapshots that are public and can be copied or restored by any AWS account, otherwise set this value to false. The default is ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot ID.`,
				},
				resource.Attribute{
					Name:        "allocated_storage",
					Description: `Specifies the allocated storage size in gigabytes (GB).`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.`,
				},
				resource.Attribute{
					Name:        "db_snapshot_arn",
					Description: `The Amazon Resource Name (ARN) for the DB snapshot.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies whether the DB snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Specifies the name of the database engine.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Specifies the version of the database engine.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "license_model",
					Description: `License model information for the restored DB instance.`,
				},
				resource.Attribute{
					Name:        "option_group_name",
					Description: `Provides the option group name for the DB snapshot.`,
				},
				resource.Attribute{
					Name:        "source_db_snapshot_identifier",
					Description: `The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.`,
				},
				resource.Attribute{
					Name:        "source_region",
					Description: `The region that the DB snapshot was created in or copied from.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of this DB snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Specifies the storage type associated with DB snapshot.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Specifies the ID of the VPC associated with the DB snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_create_time",
					Description: `Provides the time when the snapshot was taken, in Universal Coordinated Time (UTC).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot ID.`,
				},
				resource.Attribute{
					Name:        "allocated_storage",
					Description: `Specifies the allocated storage size in gigabytes (GB).`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.`,
				},
				resource.Attribute{
					Name:        "db_snapshot_arn",
					Description: `The Amazon Resource Name (ARN) for the DB snapshot.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies whether the DB snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Specifies the name of the database engine.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Specifies the version of the database engine.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "license_model",
					Description: `License model information for the restored DB instance.`,
				},
				resource.Attribute{
					Name:        "option_group_name",
					Description: `Provides the option group name for the DB snapshot.`,
				},
				resource.Attribute{
					Name:        "source_db_snapshot_identifier",
					Description: `The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.`,
				},
				resource.Attribute{
					Name:        "source_region",
					Description: `The region that the DB snapshot was created in or copied from.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of this DB snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Specifies the storage type associated with DB snapshot.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Specifies the ID of the VPC associated with the DB snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_create_time",
					Description: `Provides the time when the snapshot was taken, in Universal Coordinated Time (UTC).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_dx_gateway",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about a Direct Connect Gateway`,
			Description: `

Retrieve information about a Direct Connect Gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the gateway to retrieve. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "amazon_side_asn",
					Description: `The ASN on the Amazon side of the connection.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the gateway.`,
				},
				resource.Attribute{
					Name:        "owner_account_id",
					Description: `AWS Account ID of the gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "amazon_side_asn",
					Description: `The ASN on the Amazon side of the connection.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the gateway.`,
				},
				resource.Attribute{
					Name:        "owner_account_id",
					Description: `AWS Account ID of the gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_dynamodb_table",
			Category:         "Data Sources",
			ShortDescription: `Provides a DynamoDB table data source.`,
			Description: `

Provides information about a DynamoDB table.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the DynamoDB table. ## Attributes Reference See the [DynamoDB Table Resource](/docs/providers/aws/r/dynamodb_table.html) for details on the returned attributes - they are identical.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ebs_default_kms_key",
			Category:         "Data Sources",
			ShortDescription: `Provides metadata about the KMS key set for EBS default encryption`,
			Description: `
Use this data source to get the default EBS encryption KMS key in the current region.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_arn",
					Description: `Amazon Resource Name (ARN) of the default KMS key uses to encrypt an EBS volume in this region when no key is specified in an API call that creates the volume and encryption by default is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ebs_encryption_by_default",
			Category:         "Data Sources",
			ShortDescription: `Checks whether default EBS encryption is enabled for your AWS account in the current AWS region.`,
			Description: `

Provides a way to check whether default EBS encryption is enabled for your AWS account in the current AWS region.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether or not default EBS encryption is enabled. Returns as ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ebs_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information on an EBS Snapshot.`,
			Description: `

Use this data source to get information about an EBS Snapshot for use when provisioning EBS Volumes

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent snapshot.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) Returns the snapshots owned by the specified owner id. Multiple owners can be specified.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional) Returns information on a specific snapshot_id.`,
				},
				resource.Attribute{
					Name:        "restorable_by_user_ids",
					Description: `(Optional) One or more AWS accounts IDs that can create volumes from the snapshot.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-snapshots in the AWS CLI reference][1]. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot ID (e.g. snap-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot ID (e.g. snap-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the snapshot`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The AWS account ID of the EBS snapshot owner.`,
				},
				resource.Attribute{
					Name:        "owner_alias",
					Description: `Value from an Amazon-maintained list (` + "`" + `amazon` + "`" + `, ` + "`" + `aws-marketplace` + "`" + `, ` + "`" + `microsoft` + "`" + `) of snapshot owners.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The volume ID (e.g. vol-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the drive in GiBs.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "data_encryption_key_id",
					Description: `The data encryption key identifier for the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The snapshot state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags for the resource. [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-snapshots.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot ID (e.g. snap-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot ID (e.g. snap-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the snapshot`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The AWS account ID of the EBS snapshot owner.`,
				},
				resource.Attribute{
					Name:        "owner_alias",
					Description: `Value from an Amazon-maintained list (` + "`" + `amazon` + "`" + `, ` + "`" + `aws-marketplace` + "`" + `, ` + "`" + `microsoft` + "`" + `) of snapshot owners.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The volume ID (e.g. vol-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the drive in GiBs.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "data_encryption_key_id",
					Description: `The data encryption key identifier for the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The snapshot state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags for the resource. [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-snapshots.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ebs_snapshot_ids",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of EBS snapshot IDs.`,
			Description: `

Use this data source to get a list of EBS Snapshot IDs matching the specified
criteria.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) Returns the snapshots owned by the specified owner id. Multiple owners can be specified.`,
				},
				resource.Attribute{
					Name:        "restorable_by_user_ids",
					Description: `(Optional) One or more AWS accounts IDs that can create volumes from the snapshot.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-volumes in the AWS CLI reference][1]. ## Attributes Reference ` + "`" + `ids` + "`" + ` is set to the list of EBS snapshot IDs, sorted by creation time in descending order. [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-snapshots.html`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ebs_volume",
			Category:         "Data Sources",
			ShortDescription: `Get information on an EBS volume.`,
			Description: `

Use this data source to get information about an EBS volume for use in other
resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent Volume.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-volumes in the AWS CLI reference][1]. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The volume ID (e.g. vol-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The volume ID (e.g. vol-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ where the EBS volume exists.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the disk is encrypted.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The amount of IOPS for the disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the drive in GiBs.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot_id the EBS volume is based off.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of EBS volume.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags for the resource. [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-volumes.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The volume ID (e.g. vol-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The volume ID (e.g. vol-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ where the EBS volume exists.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the disk is encrypted.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The amount of IOPS for the disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the drive in GiBs.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot_id the EBS volume is based off.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of EBS volume.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags for the resource. [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-volumes.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ec2_transit_gateway",
			Category:         "Data Sources",
			ShortDescription: `Get information on an EC2 Transit Gateway`,
			Description: `

Get information on an EC2 Transit Gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more configuration blocks containing name-values filters. Detailed below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier of the EC2 Transit Gateway. ### filter Argument Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the filter.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) List of one or more values for the filter. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "amazon_side_asn",
					Description: `Private Autonomous System Number (ASN) for the Amazon side of a BGP session`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `EC2 Transit Gateway Amazon Resource Name (ARN)`,
				},
				resource.Attribute{
					Name:        "association_default_route_table_id",
					Description: `Identifier of the default association route table`,
				},
				resource.Attribute{
					Name:        "auto_accept_shared_attachments",
					Description: `Whether resource attachment requests are automatically accepted.`,
				},
				resource.Attribute{
					Name:        "default_route_table_association",
					Description: `Whether resource attachments are automatically associated with the default association route table.`,
				},
				resource.Attribute{
					Name:        "default_route_table_propagation",
					Description: `Whether resource attachments automatically propagate routes to the default propagation route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "dns_support",
					Description: `Whether DNS support is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway identifier`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Identifier of the AWS account that owns the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "propagation_default_route_table_id",
					Description: `Identifier of the default propagation route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "vpn_ecmp_support",
					Description: `Whether VPN Equal Cost Multipath Protocol support is enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "amazon_side_asn",
					Description: `Private Autonomous System Number (ASN) for the Amazon side of a BGP session`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `EC2 Transit Gateway Amazon Resource Name (ARN)`,
				},
				resource.Attribute{
					Name:        "association_default_route_table_id",
					Description: `Identifier of the default association route table`,
				},
				resource.Attribute{
					Name:        "auto_accept_shared_attachments",
					Description: `Whether resource attachment requests are automatically accepted.`,
				},
				resource.Attribute{
					Name:        "default_route_table_association",
					Description: `Whether resource attachments are automatically associated with the default association route table.`,
				},
				resource.Attribute{
					Name:        "default_route_table_propagation",
					Description: `Whether resource attachments automatically propagate routes to the default propagation route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "dns_support",
					Description: `Whether DNS support is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway identifier`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Identifier of the AWS account that owns the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "propagation_default_route_table_id",
					Description: `Identifier of the default propagation route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "vpn_ecmp_support",
					Description: `Whether VPN Equal Cost Multipath Protocol support is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ec2_transit_gateway_dx_gateway_attachment",
			Category:         "Data Sources",
			ShortDescription: `Get information on an EC2 Transit Gateway's attachment to a Direct Connect Gateway`,
			Description: `

Get information on an EC2 Transit Gateway's attachment to a Direct Connect Gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `(Required) Identifier of the EC2 Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "dx_gateway_id",
					Description: `(Required) Identifier of the Direct Connect Gateway. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway Attachment identifier`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway Attachment`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway Attachment identifier`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway Attachment`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ec2_transit_gateway_route_table",
			Category:         "Data Sources",
			ShortDescription: `Get information on an EC2 Transit Gateway Route Table`,
			Description: `

Get information on an EC2 Transit Gateway Route Table.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more configuration blocks containing name-values filters. Detailed below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier of the EC2 Transit Gateway Route Table. ### filter Argument Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the filter.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) List of one or more values for the filter. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "default_association_route_table",
					Description: `Boolean whether this is the default association route table for the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "default_propagation_route_table",
					Description: `Boolean whether this is the default propagation route table for the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway Route Table identifier`,
				},
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `EC2 Transit Gateway identifier`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway Route Table`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "default_association_route_table",
					Description: `Boolean whether this is the default association route table for the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "default_propagation_route_table",
					Description: `Boolean whether this is the default propagation route table for the EC2 Transit Gateway`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway Route Table identifier`,
				},
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `EC2 Transit Gateway identifier`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway Route Table`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ec2_transit_gateway_vpc_attachment",
			Category:         "Data Sources",
			ShortDescription: `Get information on an EC2 Transit Gateway VPC Attachment`,
			Description: `

Get information on an EC2 Transit Gateway VPC Attachment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more configuration blocks containing name-values filters. Detailed below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier of the EC2 Transit Gateway VPC Attachment. ### filter Argument Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the filter.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) List of one or more values for the filter. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dns_support",
					Description: `Whether DNS support is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway VPC Attachment identifier`,
				},
				resource.Attribute{
					Name:        "ipv6_support",
					Description: `Whether IPv6 support is enabled.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `Identifiers of EC2 Subnets.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `EC2 Transit Gateway identifier`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway VPC Attachment`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Identifier of EC2 VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_owner_id",
					Description: `Identifier of the AWS account that owns the EC2 VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dns_support",
					Description: `Whether DNS support is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway VPC Attachment identifier`,
				},
				resource.Attribute{
					Name:        "ipv6_support",
					Description: `Whether IPv6 support is enabled.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `Identifiers of EC2 Subnets.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `EC2 Transit Gateway identifier`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway VPC Attachment`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Identifier of EC2 VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_owner_id",
					Description: `Identifier of the AWS account that owns the EC2 VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ec2_transit_gateway_vpn_attachment",
			Category:         "Data Sources",
			ShortDescription: `Get information on an EC2 Transit Gateway VPN Attachment`,
			Description: `

Get information on an EC2 Transit Gateway VPN Attachment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `(Required) Identifier of the EC2 Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `(Required) Identifier of the EC2 VPN Connection. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway VPN Attachment identifier`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway VPN Attachment`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `EC2 Transit Gateway VPN Attachment identifier`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value tags for the EC2 Transit Gateway VPN Attachment`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ecr_image",
			Category:         "Data Sources",
			ShortDescription: `Provides details about an ECR Image`,
			Description: `

The ECR Image data source allows the details of an image with a particular tag or digest to be retrieved.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "registry_id",
					Description: `(Optional) The ID of the Registry where the repository resides.`,
				},
				resource.Attribute{
					Name:        "repository_name",
					Description: `(Required) The name of the ECR Repository.`,
				},
				resource.Attribute{
					Name:        "image_digest",
					Description: `(Optional) The sha256 digest of the image manifest. At least one of ` + "`" + `image_digest` + "`" + ` or ` + "`" + `image_tag` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "image_tag",
					Description: `(Optional) The tag associated with this image. At least one of ` + "`" + `image_digest` + "`" + ` or ` + "`" + `image_tag` + "`" + ` must be specified. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "image_pushed_at",
					Description: `The date and time, expressed as a unix timestamp, at which the current image was pushed to the repository.`,
				},
				resource.Attribute{
					Name:        "image_size_in_bytes",
					Description: `The size, in bytes, of the image in the repository.`,
				},
				resource.Attribute{
					Name:        "image_tags",
					Description: `The list of tags associated with this image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image_pushed_at",
					Description: `The date and time, expressed as a unix timestamp, at which the current image was pushed to the repository.`,
				},
				resource.Attribute{
					Name:        "image_size_in_bytes",
					Description: `The size, in bytes, of the image in the repository.`,
				},
				resource.Attribute{
					Name:        "image_tags",
					Description: `The list of tags associated with this image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ecr_repository",
			Category:         "Data Sources",
			ShortDescription: `Provides details about an ECR Repository`,
			Description: `

The ECR Repository data source allows the ARN, Repository URI and Registry ID to be retrieved for an ECR repository.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ECR Repository. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Full ARN of the repository.`,
				},
				resource.Attribute{
					Name:        "registry_id",
					Description: `The registry ID where the repository was created.`,
				},
				resource.Attribute{
					Name:        "repository_url",
					Description: `The URL of the repository (in the form ` + "`" + `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Full ARN of the repository.`,
				},
				resource.Attribute{
					Name:        "registry_id",
					Description: `The registry ID where the repository was created.`,
				},
				resource.Attribute{
					Name:        "repository_url",
					Description: `The URL of the repository (in the form ` + "`" + `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ecs_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides details about an ecs cluster`,
			Description: `

The ECS Cluster data source allows access to details of a specific
cluster within an AWS ECS service.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the ECS Cluster ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "pending_tasks_count",
					Description: `The number of pending tasks for the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "running_tasks_count",
					Description: `The number of running tasks for the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "registered_container_instances_count",
					Description: `The number of registered container instances for the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "setting",
					Description: `The settings associated with the ECS Cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "pending_tasks_count",
					Description: `The number of pending tasks for the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "running_tasks_count",
					Description: `The number of running tasks for the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "registered_container_instances_count",
					Description: `The number of registered container instances for the ECS Cluster`,
				},
				resource.Attribute{
					Name:        "setting",
					Description: `The settings associated with the ECS Cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ecs_container_definition",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a single container within an ecs task definition`,
			Description: `

The ECS container definition data source allows access to details of
a specific container within an AWS ECS service.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "task_definition",
					Description: `(Required) The ARN of the task definition which contains the container`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `(Required) The name of the container definition ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The docker image in use, including the digest`,
				},
				resource.Attribute{
					Name:        "image_digest",
					Description: `The digest of the docker image in use`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The CPU limit for this container definition`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The memory limit for this container definition`,
				},
				resource.Attribute{
					Name:        "memory_reservation",
					Description: `The soft limit (in MiB) of memory to reserve for the container. When system memory is under contention, Docker attempts to keep the container memory to this soft limit`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The environment in use`,
				},
				resource.Attribute{
					Name:        "disable_networking",
					Description: `Indicator if networking is disabled`,
				},
				resource.Attribute{
					Name:        "docker_labels",
					Description: `Set docker labels`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image",
					Description: `The docker image in use, including the digest`,
				},
				resource.Attribute{
					Name:        "image_digest",
					Description: `The digest of the docker image in use`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The CPU limit for this container definition`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The memory limit for this container definition`,
				},
				resource.Attribute{
					Name:        "memory_reservation",
					Description: `The soft limit (in MiB) of memory to reserve for the container. When system memory is under contention, Docker attempts to keep the container memory to this soft limit`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The environment in use`,
				},
				resource.Attribute{
					Name:        "disable_networking",
					Description: `Indicator if networking is disabled`,
				},
				resource.Attribute{
					Name:        "docker_labels",
					Description: `Set docker labels`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ecs_service",
			Category:         "Data Sources",
			ShortDescription: `Provides details about an ecs service`,
			Description: `

The ECS Service data source allows access to details of a specific
Service within a AWS ECS Cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the ECS Service`,
				},
				resource.Attribute{
					Name:        "cluster_arn",
					Description: `(Required) The arn of the ECS Cluster ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the ECS Service`,
				},
				resource.Attribute{
					Name:        "desired_count",
					Description: `The number of tasks for the ECS Service`,
				},
				resource.Attribute{
					Name:        "launch_type",
					Description: `The launch type for the ECS Service`,
				},
				resource.Attribute{
					Name:        "scheduling_strategy",
					Description: `The scheduling strategy for the ECS Service`,
				},
				resource.Attribute{
					Name:        "task_definition",
					Description: `The family for the latest ACTIVE revision`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the ECS Service`,
				},
				resource.Attribute{
					Name:        "desired_count",
					Description: `The number of tasks for the ECS Service`,
				},
				resource.Attribute{
					Name:        "launch_type",
					Description: `The launch type for the ECS Service`,
				},
				resource.Attribute{
					Name:        "scheduling_strategy",
					Description: `The scheduling strategy for the ECS Service`,
				},
				resource.Attribute{
					Name:        "task_definition",
					Description: `The family for the latest ACTIVE revision`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ecs_task_definition",
			Category:         "Data Sources",
			ShortDescription: `Provides details about an ecs task definition`,
			Description: `

The ECS task definition data source allows access to details of
a specific AWS ECS task definition.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "task_definition",
					Description: `(Required) The family for the latest ACTIVE revision, family and revision (family:revision) for a specific revision in the family, the ARN of the task definition to access to. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The family of this task definition`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `The Docker networking mode to use for the containers in this task.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of this task definition`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this task definition`,
				},
				resource.Attribute{
					Name:        "task_role_arn",
					Description: `The ARN of the IAM role that containers in this task can assume`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "family",
					Description: `The family of this task definition`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `The Docker networking mode to use for the containers in this task.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of this task definition`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this task definition`,
				},
				resource.Attribute{
					Name:        "task_role_arn",
					Description: `The ARN of the IAM role that containers in this task can assume`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_efs_file_system",
			Category:         "Data Sources",
			ShortDescription: `Provides an Elastic File System (EFS) data source.`,
			Description: `

Provides information about an Elastic File System (EFS).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_system_id",
					Description: `(Optional) The ID that identifies the file system (e.g. fs-ccfc0d65).`,
				},
				resource.Attribute{
					Name:        "creation_token",
					Description: `(Optional) Restricts the list to the file system with this creation token. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name of the file system.`,
				},
				resource.Attribute{
					Name:        "performance_mode",
					Description: `The PerformanceMode of the file system.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The list of tags assigned to the file system.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether EFS is encrypted.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name of the file system.`,
				},
				resource.Attribute{
					Name:        "performance_mode",
					Description: `The PerformanceMode of the file system.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The list of tags assigned to the file system.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether EFS is encrypted.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_efs_mount_target",
			Category:         "Data Sources",
			ShortDescription: `Provides an Elastic File System Mount Target (EFS) data source.`,
			Description: `

Provides information about an Elastic File System Mount Target (EFS).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mount_target_id",
					Description: `(Required) ID of the mount target that you want to have described ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "file_system_arn",
					Description: `Amazon Resource Name of the file system for which the mount target is intended.`,
				},
				resource.Attribute{
					Name:        "file_system_id",
					Description: `ID of the file system for which the mount target is intended.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the mount target's subnet.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Address at which the file system may be mounted via the mount target.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `List of VPC security group IDs attached to the mount target.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The ID of the network interface that Amazon EFS created when it created the mount target.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "file_system_arn",
					Description: `Amazon Resource Name of the file system for which the mount target is intended.`,
				},
				resource.Attribute{
					Name:        "file_system_id",
					Description: `ID of the file system for which the mount target is intended.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the mount target's subnet.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Address at which the file system may be mounted via the mount target.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `List of VPC security group IDs attached to the mount target.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The ID of the network interface that Amazon EFS created when it created the mount target.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_eip",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Elastic IP`,
			Description: `

` + "`" + `aws_eip` + "`" + ` provides details about a specific Elastic IP.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAddresses.html).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The allocation id of the specific VPC EIP to retrieve. If a classic EIP is required, do NOT set ` + "`" + `id` + "`" + `, only set ` + "`" + `public_ip` + "`" + ``,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) The public IP of the specific EIP to retrieve.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired Elastic IP ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "association_id",
					Description: `The ID representing the association of the address with an instance in a VPC.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Indicates whether the address is for use in EC2-Classic (standard) or in a VPC (vpc).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `If VPC Elastic IP, the allocation identifier. If EC2-Classic Elastic IP, the public IP address.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance that the address is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The ID of the network interface.`,
				},
				resource.Attribute{
					Name:        "network_interface_owner_id",
					Description: `The ID of the AWS account that owns the network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address associated with the Elastic IP address.`,
				},
				resource.Attribute{
					Name:        "private_dns",
					Description: `The Private DNS associated with the Elastic IP address.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of Elastic IP.`,
				},
				resource.Attribute{
					Name:        "public_dns",
					Description: `Public DNS associated with the Elastic IP address.`,
				},
				resource.Attribute{
					Name:        "public_ipv4_pool",
					Description: `The ID of an address pool.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value map of tags associated with Elastic IP. ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "association_id",
					Description: `The ID representing the association of the address with an instance in a VPC.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Indicates whether the address is for use in EC2-Classic (standard) or in a VPC (vpc).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `If VPC Elastic IP, the allocation identifier. If EC2-Classic Elastic IP, the public IP address.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance that the address is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The ID of the network interface.`,
				},
				resource.Attribute{
					Name:        "network_interface_owner_id",
					Description: `The ID of the AWS account that owns the network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address associated with the Elastic IP address.`,
				},
				resource.Attribute{
					Name:        "private_dns",
					Description: `The Private DNS associated with the Elastic IP address.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of Elastic IP.`,
				},
				resource.Attribute{
					Name:        "public_dns",
					Description: `Public DNS associated with the Elastic IP address.`,
				},
				resource.Attribute{
					Name:        "public_ipv4_pool",
					Description: `The ID of an address pool.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key-value map of tags associated with Elastic IP. ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_eks_cluster",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about an EKS Cluster`,
			Description: `

Retrieve information about an EKS Cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the cluster`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the cluster.`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `Nested attribute containing ` + "`" + `certificate-authority-data` + "`" + ` for your cluster.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `The base64 encoded certificate data required to communicate with your cluster. Add this to the ` + "`" + `certificate-authority-data` + "`" + ` section of the ` + "`" + `kubeconfig` + "`" + ` file for your cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The Unix epoch time stamp in seconds for when the cluster was created.`,
				},
				resource.Attribute{
					Name:        "enabled_cluster_log_types",
					Description: `The enabled control plane logs.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The endpoint for your Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. For an example using this information to enable IAM Roles for Service Accounts, see the [` + "`" + `aws_eks_cluster` + "`" + ` resource documentation](/docs/providers/aws/r/eks_cluster.html).`,
				},
				resource.Attribute{
					Name:        "oidc",
					Description: `Nested attribute containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `Issuer URL for the OpenID Connect identity provider.`,
				},
				resource.Attribute{
					Name:        "platform_version",
					Description: `The platform version for the cluster.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EKS cluster. One of ` + "`" + `CREATING` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Kubernetes server version for the cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_config",
					Description: `Nested attribute containing VPC configuration for the cluster.`,
				},
				resource.Attribute{
					Name:        "endpoint_private_access",
					Description: `Indicates whether or not the Amazon EKS private API server endpoint is enabled.`,
				},
				resource.Attribute{
					Name:        "endpoint_public_access",
					Description: `Indicates whether or not the Amazon EKS public API server endpoint is enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the cluster`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the cluster.`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `Nested attribute containing ` + "`" + `certificate-authority-data` + "`" + ` for your cluster.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `The base64 encoded certificate data required to communicate with your cluster. Add this to the ` + "`" + `certificate-authority-data` + "`" + ` section of the ` + "`" + `kubeconfig` + "`" + ` file for your cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The Unix epoch time stamp in seconds for when the cluster was created.`,
				},
				resource.Attribute{
					Name:        "enabled_cluster_log_types",
					Description: `The enabled control plane logs.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The endpoint for your Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. For an example using this information to enable IAM Roles for Service Accounts, see the [` + "`" + `aws_eks_cluster` + "`" + ` resource documentation](/docs/providers/aws/r/eks_cluster.html).`,
				},
				resource.Attribute{
					Name:        "oidc",
					Description: `Nested attribute containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `Issuer URL for the OpenID Connect identity provider.`,
				},
				resource.Attribute{
					Name:        "platform_version",
					Description: `The platform version for the cluster.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EKS cluster. One of ` + "`" + `CREATING` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Kubernetes server version for the cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_config",
					Description: `Nested attribute containing VPC configuration for the cluster.`,
				},
				resource.Attribute{
					Name:        "endpoint_private_access",
					Description: `Indicates whether or not the Amazon EKS private API server endpoint is enabled.`,
				},
				resource.Attribute{
					Name:        "endpoint_public_access",
					Description: `Indicates whether or not the Amazon EKS public API server endpoint is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_eks_cluster_auth",
			Category:         "Data Sources",
			ShortDescription: `Get an authentication token to communicate with an EKS Cluster`,
			Description: `

Get an authentication token to communicate with an EKS cluster.

Uses IAM credentials from the AWS provider to generate a temporary token that is compatible with
[AWS IAM Authenticator](https://github.com/kubernetes-sigs/aws-iam-authenticator) authentication.
This can be used to authenticate to an EKS cluster or to a cluster that has the AWS IAM Authenticator
server configured.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The token to use to authenticate with the cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `The token to use to authenticate with the cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_elastic_beanstalk_application",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about an Elastic Beanstalk Application`,
			Description: `

Retrieve information about an Elastic Beanstalk Application.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the application`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the application.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Short description of the application Application version lifecycle (` + "`" + `appversion_lifecycle` + "`" + `) supports the nested attribute containing.`,
				},
				resource.Attribute{
					Name:        "service_role",
					Description: `The ARN of an IAM service role under which the application version is deleted. Elastic Beanstalk must have permission to assume this role.`,
				},
				resource.Attribute{
					Name:        "max_count",
					Description: `The maximum number of application versions to retain.`,
				},
				resource.Attribute{
					Name:        "max_age_in_days",
					Description: `The number of days to retain an application version.`,
				},
				resource.Attribute{
					Name:        "delete_source_from_s3",
					Description: `Specifies whether delete a version's source bundle from S3 when the application version is deleted.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the application`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the application.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Short description of the application Application version lifecycle (` + "`" + `appversion_lifecycle` + "`" + `) supports the nested attribute containing.`,
				},
				resource.Attribute{
					Name:        "service_role",
					Description: `The ARN of an IAM service role under which the application version is deleted. Elastic Beanstalk must have permission to assume this role.`,
				},
				resource.Attribute{
					Name:        "max_count",
					Description: `The maximum number of application versions to retain.`,
				},
				resource.Attribute{
					Name:        "max_age_in_days",
					Description: `The number of days to retain an application version.`,
				},
				resource.Attribute{
					Name:        "delete_source_from_s3",
					Description: `Specifies whether delete a version's source bundle from S3 when the application version is deleted.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_elastic_beanstalk_hosted_zone",
			Category:         "Data Sources",
			ShortDescription: `Get an elastic beanstalk hosted zone.`,
			Description: `

Use this data source to get the ID of an [elastic beanstalk hosted zone](http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region you'd like the zone for. By default, fetches the current region. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the hosted zone.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the hosted zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the hosted zone.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the hosted zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_elastic_beanstalk_solution_stack",
			Category:         "Data Sources",
			ShortDescription: `Get an elastic beanstalk solution stack.`,
			Description: `

Use this data source to get the name of a elastic beanstalk solution stack.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent solution stack.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `A regex string to apply to the solution stack list returned by AWS. See [Elastic Beanstalk Supported Platforms][beanstalk-platforms] from AWS documentation for reference solution stack names. ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the solution stack. [beanstalk-platforms]: http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html "AWS Elastic Beanstalk Supported Platforms documentation"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the solution stack. [beanstalk-platforms]: http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html "AWS Elastic Beanstalk Supported Platforms documentation"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_elasticache_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information on an ElastiCache Cluster resource.`,
			Description: `

Use this data source to get information about an Elasticache Cluster

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "replication_group_id",
					Description: `The replication group to which this cache cluster belongs.`,
				},
				resource.Attribute{
					Name:        "snapshot_window",
					Description: `The daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of the cache cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_retention_limit",
					Description: `The number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The Availability Zone for the cache cluster.`,
				},
				resource.Attribute{
					Name:        "configuration_endpoint",
					Description: `(Memcached only) The configuration endpoint to allow host discovery.`,
				},
				resource.Attribute{
					Name:        "cluster_address",
					Description: `(Memcached only) The DNS name of the cache cluster without the port appended.`,
				},
				resource.Attribute{
					Name:        "cache_nodes",
					Description: `List of node objects including ` + "`" + `id` + "`" + `, ` + "`" + `address` + "`" + `, ` + "`" + `port` + "`" + ` and ` + "`" + `availability_zone` + "`" + `. Referenceable e.g. as ` + "`" + `${data.aws_elasticache_cluster.bar.cache_nodes.0.address}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned to the resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "replication_group_id",
					Description: `The replication group to which this cache cluster belongs.`,
				},
				resource.Attribute{
					Name:        "snapshot_window",
					Description: `The daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of the cache cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_retention_limit",
					Description: `The number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The Availability Zone for the cache cluster.`,
				},
				resource.Attribute{
					Name:        "configuration_endpoint",
					Description: `(Memcached only) The configuration endpoint to allow host discovery.`,
				},
				resource.Attribute{
					Name:        "cluster_address",
					Description: `(Memcached only) The DNS name of the cache cluster without the port appended.`,
				},
				resource.Attribute{
					Name:        "cache_nodes",
					Description: `List of node objects including ` + "`" + `id` + "`" + `, ` + "`" + `address` + "`" + `, ` + "`" + `port` + "`" + ` and ` + "`" + `availability_zone` + "`" + `. Referenceable e.g. as ` + "`" + `${data.aws_elasticache_cluster.bar.cache_nodes.0.address}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned to the resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_elasticache_replication_group",
			Category:         "Data Sources",
			ShortDescription: `Get information on an ElastiCache Replication Group resource.`,
			Description: `

Use this data source to get information about an Elasticache Replication Group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "replication_group_id",
					Description: `The identifier for the replication group.`,
				},
				resource.Attribute{
					Name:        "replication_group_description",
					Description: `The description of the replication group.`,
				},
				resource.Attribute{
					Name:        "auth_token_enabled",
					Description: `A flag that enables using an AuthToken (password) when issuing Redis commands.`,
				},
				resource.Attribute{
					Name:        "automatic_failover_enabled",
					Description: `A flag whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails.`,
				},
				resource.Attribute{
					Name:        "member_clusters",
					Description: `The identifiers of all the nodes that are part of this replication group.`,
				},
				resource.Attribute{
					Name:        "snapshot_window",
					Description: `The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).`,
				},
				resource.Attribute{
					Name:        "snapshot_retention_limit",
					Description: `The number of days for which ElastiCache retains automatic cache cluster snapshots before deleting them.`,
				},
				resource.Attribute{
					Name:        "configuration_endpoint_address",
					Description: `The configuration endpoint address to allow host discovery.`,
				},
				resource.Attribute{
					Name:        "primary_endpoint_address",
					Description: `The endpoint of the primary node in this node group (shard).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "replication_group_id",
					Description: `The identifier for the replication group.`,
				},
				resource.Attribute{
					Name:        "replication_group_description",
					Description: `The description of the replication group.`,
				},
				resource.Attribute{
					Name:        "auth_token_enabled",
					Description: `A flag that enables using an AuthToken (password) when issuing Redis commands.`,
				},
				resource.Attribute{
					Name:        "automatic_failover_enabled",
					Description: `A flag whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails.`,
				},
				resource.Attribute{
					Name:        "member_clusters",
					Description: `The identifiers of all the nodes that are part of this replication group.`,
				},
				resource.Attribute{
					Name:        "snapshot_window",
					Description: `The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).`,
				},
				resource.Attribute{
					Name:        "snapshot_retention_limit",
					Description: `The number of days for which ElastiCache retains automatic cache cluster snapshots before deleting them.`,
				},
				resource.Attribute{
					Name:        "configuration_endpoint_address",
					Description: `The configuration endpoint address to allow host discovery.`,
				},
				resource.Attribute{
					Name:        "primary_endpoint_address",
					Description: `The endpoint of the primary node in this node group (shard).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_elasticsearch_domain",
			Category:         "Data Sources",
			ShortDescription: `Get information on an ElasticSearch Domain resource.`,
			Description: `

Use this data source to get information about an Elasticsearch Domain

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "advanced_options",
					Description: `Key-value string pairs to specify advanced configuration options.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `Cluster configuration of the domain.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of data nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "dedicated_master_enabled",
					Description: `Indicates whether dedicated master nodes are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "dedicated_master_type",
					Description: `Instance type of the dedicated master nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "dedicated_master_count",
					Description: `Number of dedicated master nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "zone_awareness_enabled",
					Description: `Indicates whether zone awareness is enabled.`,
				},
				resource.Attribute{
					Name:        "zone_awareness_config",
					Description: `Configuration block containing zone awareness settings.`,
				},
				resource.Attribute{
					Name:        "availability_zone_count",
					Description: `Number of availability zones used.`,
				},
				resource.Attribute{
					Name:        "cognito_options",
					Description: `Domain Amazon Cognito Authentication options for Kibana.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether Amazon Cognito Authentication is enabled.`,
				},
				resource.Attribute{
					Name:        "user_pool_id",
					Description: `The Cognito User pool used by the domain.`,
				},
				resource.Attribute{
					Name:        "identity_pool_id",
					Description: `The Cognito Identity pool used by the domain.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `The IAM Role with the AmazonESCognitoAccess policy attached.`,
				},
				resource.Attribute{
					Name:        "ebs_options",
					Description: `EBS Options for the instances in the domain.`,
				},
				resource.Attribute{
					Name:        "ebs_enabled",
					Description: `Whether EBS volumes are attached to data nodes in the domain.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of EBS volumes attached to data nodes.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of EBS volumes attached to data nodes (in GB).`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The baseline input/output (I/O) performance of EBS volumes attached to data nodes.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest",
					Description: `Domain encryption at rest related options.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether encryption at rest is enabled in the domain.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The KMS key id used to encrypt data at rest.`,
				},
				resource.Attribute{
					Name:        "kibana_endpoint",
					Description: `Domain-specific endpoint used to access the Kibana application.`,
				},
				resource.Attribute{
					Name:        "log_publishing_options",
					Description: `Domain log publishing related options.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `The type of Elasticsearch log being published.`,
				},
				resource.Attribute{
					Name:        "cloudwatch_log_group_arn",
					Description: `The CloudWatch Log Group where the logs are published.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether log publishing is enabled.`,
				},
				resource.Attribute{
					Name:        "node_to_node_encryption",
					Description: `Domain in transit encryption related options.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether node to node encryption is enabled.`,
				},
				resource.Attribute{
					Name:        "automated_snapshot_start_hour",
					Description: `Hour during which the service takes an automated daily snapshot of the indices in the domain.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned to the domain.`,
				},
				resource.Attribute{
					Name:        "vpc_options",
					Description: `VPC Options for private Elasticsearch domains.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The availability zones used by the domain.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `The security groups used by the domain.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `The subnets used by the domain.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC used by the domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "advanced_options",
					Description: `Key-value string pairs to specify advanced configuration options.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `Cluster configuration of the domain.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of data nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "dedicated_master_enabled",
					Description: `Indicates whether dedicated master nodes are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "dedicated_master_type",
					Description: `Instance type of the dedicated master nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "dedicated_master_count",
					Description: `Number of dedicated master nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "zone_awareness_enabled",
					Description: `Indicates whether zone awareness is enabled.`,
				},
				resource.Attribute{
					Name:        "zone_awareness_config",
					Description: `Configuration block containing zone awareness settings.`,
				},
				resource.Attribute{
					Name:        "availability_zone_count",
					Description: `Number of availability zones used.`,
				},
				resource.Attribute{
					Name:        "cognito_options",
					Description: `Domain Amazon Cognito Authentication options for Kibana.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether Amazon Cognito Authentication is enabled.`,
				},
				resource.Attribute{
					Name:        "user_pool_id",
					Description: `The Cognito User pool used by the domain.`,
				},
				resource.Attribute{
					Name:        "identity_pool_id",
					Description: `The Cognito Identity pool used by the domain.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `The IAM Role with the AmazonESCognitoAccess policy attached.`,
				},
				resource.Attribute{
					Name:        "ebs_options",
					Description: `EBS Options for the instances in the domain.`,
				},
				resource.Attribute{
					Name:        "ebs_enabled",
					Description: `Whether EBS volumes are attached to data nodes in the domain.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of EBS volumes attached to data nodes.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of EBS volumes attached to data nodes (in GB).`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The baseline input/output (I/O) performance of EBS volumes attached to data nodes.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest",
					Description: `Domain encryption at rest related options.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether encryption at rest is enabled in the domain.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The KMS key id used to encrypt data at rest.`,
				},
				resource.Attribute{
					Name:        "kibana_endpoint",
					Description: `Domain-specific endpoint used to access the Kibana application.`,
				},
				resource.Attribute{
					Name:        "log_publishing_options",
					Description: `Domain log publishing related options.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `The type of Elasticsearch log being published.`,
				},
				resource.Attribute{
					Name:        "cloudwatch_log_group_arn",
					Description: `The CloudWatch Log Group where the logs are published.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether log publishing is enabled.`,
				},
				resource.Attribute{
					Name:        "node_to_node_encryption",
					Description: `Domain in transit encryption related options.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether node to node encryption is enabled.`,
				},
				resource.Attribute{
					Name:        "automated_snapshot_start_hour",
					Description: `Hour during which the service takes an automated daily snapshot of the indices in the domain.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned to the domain.`,
				},
				resource.Attribute{
					Name:        "vpc_options",
					Description: `VPC Options for private Elasticsearch domains.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The availability zones used by the domain.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `The security groups used by the domain.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `The subnets used by the domain.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC used by the domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_elb",
			Category:         "Data Sources",
			ShortDescription: `Provides a classic Elastic Load Balancer data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the load balancer. ## Attributes Reference See the [ELB Resource](/docs/providers/aws/r/elb.html) for details on the returned attributes - they are identical.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_elb_hosted_zone_id",
			Category:         "Data Sources",
			ShortDescription: `Get AWS Elastic Load Balancing Hosted Zone Id`,
			Description: `

Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
in a given region for the purpose of using in an AWS Route53 Alias.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Name of the region whose AWS ELB HostedZoneId is desired. Defaults to the region from the AWS provider configuration. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the AWS ELB HostedZoneId in the selected region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the AWS ELB HostedZoneId in the selected region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_elb_service_account",
			Category:         "Data Sources",
			ShortDescription: `Get AWS Elastic Load Balancing Service Account`,
			Description: `

Use this data source to get the Account ID of the [AWS Elastic Load Balancing Service Account](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy)
in a given region for the purpose of whitelisting in S3 bucket policy.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Name of the region whose AWS ELB account ID is desired. Defaults to the region from the AWS provider configuration. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the AWS ELB service account in the selected region.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the AWS ELB service account in the selected region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the AWS ELB service account in the selected region.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the AWS ELB service account in the selected region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_glue_script",
			Category:         "Data Sources",
			ShortDescription: `Generate Glue script from Directed Acyclic Graph`,
			Description: `

Use this data source to generate a Glue script from a Directed Acyclic Graph (DAG).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dag_edge",
					Description: `(Required) A list of the edges in the DAG. Defined below.`,
				},
				resource.Attribute{
					Name:        "dag_node",
					Description: `(Required) A list of the nodes in the DAG. Defined below.`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `(Optional) The programming language of the resulting code from the DAG. Defaults to ` + "`" + `PYTHON` + "`" + `. Valid values are ` + "`" + `PYTHON` + "`" + ` and ` + "`" + `SCALA` + "`" + `. ### dag_edge Argument Reference`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The ID of the node at which the edge starts.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The ID of the node at which the edge ends.`,
				},
				resource.Attribute{
					Name:        "target_parameter",
					Description: `(Optional) The target of the edge. ### dag_node Argument Reference`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Required) Nested configuration an argument or property of a node. Defined below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) A node identifier that is unique within the node's graph.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Required) The type of node this is.`,
				},
				resource.Attribute{
					Name:        "line_number",
					Description: `(Optional) The line number of the node. #### args Argument Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the argument or property.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the argument or property.`,
				},
				resource.Attribute{
					Name:        "param",
					Description: `(Optional) Boolean if the value is used as a parameter. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "python_script",
					Description: `The Python script generated from the DAG when the ` + "`" + `language` + "`" + ` argument is set to ` + "`" + `PYTHON` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scala_code",
					Description: `The Scala code generated from the DAG when the ` + "`" + `language` + "`" + ` argument is set to ` + "`" + `SCALA` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "python_script",
					Description: `The Python script generated from the DAG when the ` + "`" + `language` + "`" + ` argument is set to ` + "`" + `PYTHON` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scala_code",
					Description: `The Scala code generated from the DAG when the ` + "`" + `language` + "`" + ` argument is set to ` + "`" + `SCALA` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_iam_account_alias",
			Category:         "Data Sources",
			ShortDescription: `Provides the account alias for the AWS account associated with the provider connection to AWS.`,
			Description: `

The IAM Account Alias data source allows access to the account alias
for the effective account in which Terraform is working.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_alias",
					Description: `The alias associated with the AWS account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_alias",
					Description: `The alias associated with the AWS account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_iam_group",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Amazon IAM group`,
			Description: `

This data source can be used to fetch information about a specific
IAM group. By using this data source, you can reference IAM group
properties without having to hard code ARNs as input.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) The friendly IAM group name to match. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) specifying the group.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The stable and unique string identifying the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) specifying the group.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The stable and unique string identifying the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_iam_instance_profile",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Amazon IAM Instance Profile`,
			Description: `

This data source can be used to fetch information about a specific
IAM instance profile. By using this data source, you can reference IAM
instance profile properties without having to hard code ARNs as input.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The friendly IAM instance profile name to match. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) specifying the instance profile.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The string representation of the date the instance profile was created.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the instance profile.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `The role arn associated with this instance profile.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The role id associated with this instance profile.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `The role name associated with this instance profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) specifying the instance profile.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The string representation of the date the instance profile was created.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the instance profile.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `The role arn associated with this instance profile.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The role id associated with this instance profile.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `The role name associated with this instance profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_iam_policy",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Amazon IAM policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `(Required) ARN of the IAM policy. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the IAM policy.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) specifying the policy.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the policy.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The policy document of the policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the IAM policy.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) specifying the policy.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the policy.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The policy document of the policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_iam_policy_document",
			Category:         "Data Sources",
			ShortDescription: `Generates an IAM policy document in JSON format`,
			Description: `

Generates an IAM policy document in JSON format.

This is a data source which can be used to construct a JSON representation of
an IAM policy document, for use with resources which expect policy documents,
such as the ` + "`" + `aws_iam_policy` + "`" + ` resource.

-> For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](/docs/providers/aws/guides/iam-policy-documents.html).

` + "`" + `` + "`" + `` + "`" + `hcl
data "aws_iam_policy_document" "example" {
  statement {
    sid = "1"

    actions = [
      "s3:ListAllMyBuckets",
      "s3:GetBucketLocation",
    ]

    resources = [
      "arn:aws:s3:::*",
    ]
  }

  statement {
    actions = [
      "s3:ListBucket",
    ]

    resources = [
      "arn:aws:s3:::${var.s3_bucket_name}",
    ]

    condition {
      test     = "StringLike"
      variable = "s3:prefix"

      values = [
        "",
        "home/",
        "home/&{aws:username}/",
      ]
    }
  }

  statement {
    actions = [
      "s3:*",
    ]

    resources = [
      "arn:aws:s3:::${var.s3_bucket_name}/home/&{aws:username}",
      "arn:aws:s3:::${var.s3_bucket_name}/home/&{aws:username}/*",
    ]
  }
}

resource "aws_iam_policy" "example" {
  name   = "example_policy"
  path   = "/"
  policy = "${data.aws_iam_policy_document.example.json}"
}
` + "`" + `` + "`" + `` + "`" + `

Using this data source to generate policy documents is *optional*. It is also
valid to use literal JSON strings within your configuration, or to use the
` + "`" + `file` + "`" + ` interpolation function to read a raw JSON policy document from a file.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `The above arguments serialized as a standard JSON policy document. ## Example with Multiple Principals Showing how you can use this as an assume role policy as well as showing how you can specify multiple principal blocks with different types. ` + "`" + `` + "`" + `` + "`" + `hcl data "aws_iam_policy_document" "event_stream_bucket_role_assume_role_policy" { statement { actions = ["sts:AssumeRole"] principals { type = "Service" identifiers = ["firehose.amazonaws.com"] } principals { type = "AWS" identifiers = ["${var.trusted_role_arn}"] } } } ` + "`" + `` + "`" + `` + "`" + ` ## Example with Source and Override Showing how you can use ` + "`" + `source_json` + "`" + ` and ` + "`" + `override_json` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl data "aws_iam_policy_document" "source" { statement { actions = ["ec2:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `The above arguments serialized as a standard JSON policy document. ## Example with Multiple Principals Showing how you can use this as an assume role policy as well as showing how you can specify multiple principal blocks with different types. ` + "`" + `` + "`" + `` + "`" + `hcl data "aws_iam_policy_document" "event_stream_bucket_role_assume_role_policy" { statement { actions = ["sts:AssumeRole"] principals { type = "Service" identifiers = ["firehose.amazonaws.com"] } principals { type = "AWS" identifiers = ["${var.trusted_role_arn}"] } } } ` + "`" + `` + "`" + `` + "`" + ` ## Example with Source and Override Showing how you can use ` + "`" + `source_json` + "`" + ` and ` + "`" + `override_json` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl data "aws_iam_policy_document" "source" { statement { actions = ["ec2:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_iam_role",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Amazon IAM role`,
			Description: `

This data source can be used to fetch information about a specific
IAM role. By using this data source, you can reference IAM role
properties without having to hard code ARNs as input.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The friendly IAM role name to match. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The friendly IAM role name to match.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) specifying the role.`,
				},
				resource.Attribute{
					Name:        "assume_role_policy",
					Description: `The policy document associated with the role.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the role in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the role.`,
				},
				resource.Attribute{
					Name:        "max_session_duration",
					Description: `Maximum session duration.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the role.`,
				},
				resource.Attribute{
					Name:        "permissions_boundary",
					Description: `The ARN of the policy that is used to set the permissions boundary for the role.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The stable and unique string identifying the role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The friendly IAM role name to match.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) specifying the role.`,
				},
				resource.Attribute{
					Name:        "assume_role_policy",
					Description: `The policy document associated with the role.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the role in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the role.`,
				},
				resource.Attribute{
					Name:        "max_session_duration",
					Description: `Maximum session duration.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the role.`,
				},
				resource.Attribute{
					Name:        "permissions_boundary",
					Description: `The ARN of the policy that is used to set the permissions boundary for the role.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The stable and unique string identifying the role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_iam_server_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information about a server certificate`,
			Description: `

Use this data source to lookup information about IAM Server Certificates.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_prefix",
					Description: `prefix of cert to filter by`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `prefix of path to filter by`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `exact name of the cert to lookup`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `sort results by expiration date. returns the certificate with expiration date in furthest in the future. ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_iam_user",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Amazon IAM user`,
			Description: `

This data source can be used to fetch information about a specific
IAM user. By using this data source, you can reference IAM user
properties without having to hard code ARNs or unique IDs as input.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) The friendly IAM user name to match. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) assigned by AWS for this user.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path in which this user was created.`,
				},
				resource.Attribute{
					Name:        "permissions_boundary",
					Description: `The ARN of the policy that is used to set the permissions boundary for the user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The unique ID assigned by AWS for this user.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The name associated to this User`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) assigned by AWS for this user.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path in which this user was created.`,
				},
				resource.Attribute{
					Name:        "permissions_boundary",
					Description: `The ARN of the policy that is used to set the permissions boundary for the user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The unique ID assigned by AWS for this user.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The name associated to this User`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_inspector_rules_packages",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of AWS Inspector Rules packages which can be used by AWS Inspector.`,
			Description: `

The AWS Inspector Rules Packages data source allows access to the list of AWS
Inspector Rules Packages which can be used by AWS Inspector within the region
configured in the provider.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arns",
					Description: `A list of the AWS Inspector Rules Packages arns available in the AWS region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_instance",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Amazon EC2 Instance.`,
			Description: `

Use this data source to get the ID of an Amazon EC2 Instance for use in other
resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Specify the exact Instance ID with which to populate the data source.`,
				},
				resource.Attribute{
					Name:        "instance_tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired Instance.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out [describe-instances in the AWS CLI reference][1].`,
				},
				resource.Attribute{
					Name:        "get_password_data",
					Description: `(Optional) If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the ` + "`" + `password_data` + "`" + ` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.`,
				},
				resource.Attribute{
					Name:        "get_user_data",
					Description: `(Optional) Retrieve Base64 encoded User Data contents into the ` + "`" + `user_data_base64` + "`" + ` attribute. A SHA-1 hash of the User Data contents will always be present in the ` + "`" + `user_data` + "`" + ` attribute. Defaults to ` + "`" + `false` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "ami",
					Description: `The ID of the AMI used to launch the instance.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the instance.`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `Whether or not the Instance is associated with a public IP address or not (Boolean).`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the Instance.`,
				},
				resource.Attribute{
					Name:        "ebs_block_device",
					Description: `The EBS block device mappings of the Instance.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `If the EBS volume will be deleted on termination.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The physical name of the device.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `If the EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `` + "`" + `0` + "`" + ` If the EBS volume is not a provisioned IOPS image, otherwise the supported IOPS count.`,
				},
				resource.Attribute{
					Name:        "kms_key_arn",
					Description: `Amazon Resource Name (ARN) of KMS Key, if EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in GiB.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `Whether the Instance is EBS optimized or not (Boolean).`,
				},
				resource.Attribute{
					Name:        "ephemeral_block_device",
					Description: `The ephemeral block device mappings of the Instance.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The physical name of the device.`,
				},
				resource.Attribute{
					Name:        "no_device",
					Description: `Whether the specified device included in the device mapping was suppressed or not (Boolean).`,
				},
				resource.Attribute{
					Name:        "virtual_name",
					Description: `The virtual device name.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `The name of the instance profile associated with the Instance.`,
				},
				resource.Attribute{
					Name:        "ipv6_addresses",
					Description: `The IPv6 addresses associated to the Instance, if applicable.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `The state of the instance. One of: ` + "`" + `pending` + "`" + `, ` + "`" + `running` + "`" + `, ` + "`" + `shutting-down` + "`" + `, ` + "`" + `terminated` + "`" + `, ` + "`" + `stopping` + "`" + `, ` + "`" + `stopped` + "`" + `. See [Instance Lifecycle](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html) for more information.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the Instance.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The key name of the Instance.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Whether detailed monitoring is enabled or disabled for the Instance (Boolean).`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The ID of the network interface that was created with the Instance.`,
				},
				resource.Attribute{
					Name:        "password_data",
					Description: `Base-64 encoded encrypted password data for the instance. Useful for getting the administrator password for instances running Microsoft Windows. This attribute is only exported if ` + "`" + `get_password_data` + "`" + ` is true. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.`,
				},
				resource.Attribute{
					Name:        "placement_group",
					Description: `The placement group of the Instance.`,
				},
				resource.Attribute{
					Name:        "private_dns",
					Description: `The private DNS name assigned to the Instance. Can only be used inside the Amazon EC2, and only available if you've enabled DNS hostnames for your VPC.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the Instance.`,
				},
				resource.Attribute{
					Name:        "public_dns",
					Description: `The public DNS name assigned to the Instance. For EC2-VPC, this is only available if you've enabled DNS hostnames for your VPC.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address assigned to the Instance, if applicable.`,
				},
				resource.Attribute{
					Name:        "root_block_device",
					Description: `The root block device mappings of the Instance`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `If the root block device will be deleted on termination.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `If the EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `` + "`" + `0` + "`" + ` If the volume is not a provisioned IOPS image, otherwise the supported IOPS count.`,
				},
				resource.Attribute{
					Name:        "kms_key_arn",
					Description: `Amazon Resource Name (ARN) of KMS Key, if EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in GiB.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `The associated security groups.`,
				},
				resource.Attribute{
					Name:        "source_dest_check",
					Description: `Whether the network interface performs source/destination checking (Boolean).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The VPC subnet ID.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `SHA-1 hash of User Data supplied to the Instance.`,
				},
				resource.Attribute{
					Name:        "user_data_base64",
					Description: `Base64 encoded contents of User Data supplied to the Instance. Valid UTF-8 contents can be decoded with the [` + "`" + `base64decode` + "`" + ` function](/docs/configuration/functions/base64decode.html). This attribute is only exported if ` + "`" + `get_user_data` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Instance.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the instance: ` + "`" + `dedicated` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `The Id of the dedicated host the instance will be assigned to.`,
				},
				resource.Attribute{
					Name:        "vpc_security_group_ids",
					Description: `The associated security groups in a non-default VPC.`,
				},
				resource.Attribute{
					Name:        "credit_specification",
					Description: `The credit specification of the Instance. [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ami",
					Description: `The ID of the AMI used to launch the instance.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the instance.`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `Whether or not the Instance is associated with a public IP address or not (Boolean).`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the Instance.`,
				},
				resource.Attribute{
					Name:        "ebs_block_device",
					Description: `The EBS block device mappings of the Instance.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `If the EBS volume will be deleted on termination.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The physical name of the device.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `If the EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `` + "`" + `0` + "`" + ` If the EBS volume is not a provisioned IOPS image, otherwise the supported IOPS count.`,
				},
				resource.Attribute{
					Name:        "kms_key_arn",
					Description: `Amazon Resource Name (ARN) of KMS Key, if EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in GiB.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `Whether the Instance is EBS optimized or not (Boolean).`,
				},
				resource.Attribute{
					Name:        "ephemeral_block_device",
					Description: `The ephemeral block device mappings of the Instance.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The physical name of the device.`,
				},
				resource.Attribute{
					Name:        "no_device",
					Description: `Whether the specified device included in the device mapping was suppressed or not (Boolean).`,
				},
				resource.Attribute{
					Name:        "virtual_name",
					Description: `The virtual device name.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `The name of the instance profile associated with the Instance.`,
				},
				resource.Attribute{
					Name:        "ipv6_addresses",
					Description: `The IPv6 addresses associated to the Instance, if applicable.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `The state of the instance. One of: ` + "`" + `pending` + "`" + `, ` + "`" + `running` + "`" + `, ` + "`" + `shutting-down` + "`" + `, ` + "`" + `terminated` + "`" + `, ` + "`" + `stopping` + "`" + `, ` + "`" + `stopped` + "`" + `. See [Instance Lifecycle](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html) for more information.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the Instance.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The key name of the Instance.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Whether detailed monitoring is enabled or disabled for the Instance (Boolean).`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The ID of the network interface that was created with the Instance.`,
				},
				resource.Attribute{
					Name:        "password_data",
					Description: `Base-64 encoded encrypted password data for the instance. Useful for getting the administrator password for instances running Microsoft Windows. This attribute is only exported if ` + "`" + `get_password_data` + "`" + ` is true. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.`,
				},
				resource.Attribute{
					Name:        "placement_group",
					Description: `The placement group of the Instance.`,
				},
				resource.Attribute{
					Name:        "private_dns",
					Description: `The private DNS name assigned to the Instance. Can only be used inside the Amazon EC2, and only available if you've enabled DNS hostnames for your VPC.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the Instance.`,
				},
				resource.Attribute{
					Name:        "public_dns",
					Description: `The public DNS name assigned to the Instance. For EC2-VPC, this is only available if you've enabled DNS hostnames for your VPC.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address assigned to the Instance, if applicable.`,
				},
				resource.Attribute{
					Name:        "root_block_device",
					Description: `The root block device mappings of the Instance`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `If the root block device will be deleted on termination.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `If the EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `` + "`" + `0` + "`" + ` If the volume is not a provisioned IOPS image, otherwise the supported IOPS count.`,
				},
				resource.Attribute{
					Name:        "kms_key_arn",
					Description: `Amazon Resource Name (ARN) of KMS Key, if EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in GiB.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `The associated security groups.`,
				},
				resource.Attribute{
					Name:        "source_dest_check",
					Description: `Whether the network interface performs source/destination checking (Boolean).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The VPC subnet ID.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `SHA-1 hash of User Data supplied to the Instance.`,
				},
				resource.Attribute{
					Name:        "user_data_base64",
					Description: `Base64 encoded contents of User Data supplied to the Instance. Valid UTF-8 contents can be decoded with the [` + "`" + `base64decode` + "`" + ` function](/docs/configuration/functions/base64decode.html). This attribute is only exported if ` + "`" + `get_user_data` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Instance.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the instance: ` + "`" + `dedicated` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `The Id of the dedicated host the instance will be assigned to.`,
				},
				resource.Attribute{
					Name:        "vpc_security_group_ids",
					Description: `The associated security groups in a non-default VPC.`,
				},
				resource.Attribute{
					Name:        "credit_specification",
					Description: `The credit specification of the Instance. [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_instances",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Amazon EC2 instances.`,
			Description: `

Use this data source to get IDs or IPs of Amazon EC2 instances to be referenced elsewhere,
e.g. to allow easier migration from another management solution
or to make it easier for an operator to connect through bastion host(s).

-> **Note:** It's a best practice to expose instance details via [outputs](https://www.terraform.io/docs/configuration/outputs.html)
and [remote state](https://www.terraform.io/docs/state/remote.html) and
**use [` + "`" + `terraform_remote_state` + "`" + `](https://www.terraform.io/docs/providers/terraform/d/remote_state.html)
data source instead** if you manage referenced instances via Terraform.

~> **Note:** It's strongly discouraged to use this data source for querying ephemeral
instances (e.g. managed via autoscaling group), as the output may change at any time
and you'd need to re-run ` + "`" + `apply` + "`" + ` every time an instance comes up or dies.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on desired instances.`,
				},
				resource.Attribute{
					Name:        "instance_state_names",
					Description: `(Optional) A list of instance states that should be applicable to the desired instances. The permitted values are: ` + "`" + `pending, running, shutting-down, stopped, stopping, terminated` + "`" + `. The default value is ` + "`" + `running` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out [describe-instances in the AWS CLI reference][1]. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `IDs of instances found through the filter`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Private IP addresses of instances found through the filter`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Public IP addresses of instances found through the filter [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `IDs of instances found through the filter`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Private IP addresses of instances found through the filter`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Public IP addresses of instances found through the filter [1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_internet_gateway",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Internet Gateway`,
			Description: `

` + "`" + `aws_internet_gateway` + "`" + ` provides details about a specific Internet Gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_gateway_id",
					Description: `(Optional) The id of the specific Internet Gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired Internet Gateway.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInternetGateways.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. An Internet Gateway will be selected if any one of the given values matches. ## Attributes Reference All of the argument attributes except ` + "`" + `filter` + "`" + ` block are also exported as result attributes. This data source will complete the data by populating any fields that are not included in the configuration with the data for the selected Internet Gateway. ` + "`" + `attachments` + "`" + ` are also exported with the following attributes, when there are relevants: Each attachement supports the following:`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the internet gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the attachment between the gateway and the VPC. Present only if a VPC is attached`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of an attached VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the internet gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the attachment between the gateway and the VPC. Present only if a VPC is attached`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of an attached VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_iot_endpoint",
			Category:         "Data Sources",
			ShortDescription: `Get the unique IoT endpoint`,
			Description: `

Returns a unique endpoint specific to the AWS account making the call.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `(Optional) Endpoint type. Valid values: ` + "`" + `iot:CredentialProvider` + "`" + `, ` + "`" + `iot:Data` + "`" + `, ` + "`" + `iot:Data-ATS` + "`" + `, ` + "`" + `iot:Job` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "endpoint_address",
					Description: `The endpoint based on ` + "`" + `endpoint_type` + "`" + `:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_address",
					Description: `The endpoint based on ` + "`" + `endpoint_type` + "`" + `:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get information on AWS IP ranges.`,
			Description: `

Use this data source to get the IP ranges of various AWS products and services. For more information about the contents of this data source and required JSON syntax if referencing a custom URL, see the [AWS IP Address Ranges documention][1].

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) Filter IP ranges by regions (or include all regions, if omitted). Valid items are ` + "`" + `global` + "`" + ` (for ` + "`" + `cloudfront` + "`" + `) as well as all AWS regions (e.g. ` + "`" + `eu-central-1` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) Filter IP ranges by services. Valid items are ` + "`" + `amazon` + "`" + ` (for amazon.com), ` + "`" + `cloudfront` + "`" + `, ` + "`" + `codebuild` + "`" + `, ` + "`" + `ec2` + "`" + `, ` + "`" + `route53` + "`" + `, ` + "`" + `route53_healthchecks` + "`" + ` and ` + "`" + `S3` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) Custom URL for source JSON file. Syntax must match [AWS IP Address Ranges documention][1]. Defaults to ` + "`" + `https://ip-ranges.amazonaws.com/ip-ranges.json` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The lexically ordered list of CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_blocks",
					Description: `The lexically ordered list of IPv6 CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The publication time of the IP ranges (e.g. ` + "`" + `2016-08-03-23-46-05` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "sync_token",
					Description: `The publication time of the IP ranges, in Unix epoch time format (e.g. ` + "`" + `1470267965` + "`" + `). [1]: https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The lexically ordered list of CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_blocks",
					Description: `The lexically ordered list of IPv6 CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The publication time of the IP ranges (e.g. ` + "`" + `2016-08-03-23-46-05` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "sync_token",
					Description: `The publication time of the IP ranges, in Unix epoch time format (e.g. ` + "`" + `1470267965` + "`" + `). [1]: https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_kinesis_stream",
			Category:         "Data Sources",
			ShortDescription: `Provides a Kinesis Stream data source.`,
			Description: `

Use this data source to get information about a Kinesis Stream for use in other
resources.

For more details, see the [Amazon Kinesis Documentation][1].

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Kinesis Stream. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the Amazon Resource Name (ARN) of the Kinesis Stream. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the Kinesis Stream (same as id).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Kinesis Stream.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `The approximate UNIX timestamp that the stream was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the stream. The stream status is one of CREATING, DELETING, ACTIVE, or UPDATING.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Length of time (in hours) data records are accessible after they are added to the stream.`,
				},
				resource.Attribute{
					Name:        "open_shards",
					Description: `The list of shard ids in the OPEN state. See [Shard State][2] for more.`,
				},
				resource.Attribute{
					Name:        "closed_shards",
					Description: `The list of shard ids in the CLOSED state. See [Shard State][2] for more.`,
				},
				resource.Attribute{
					Name:        "shard_level_metrics",
					Description: `A list of shard-level CloudWatch metrics which are enabled for the stream. See [Monitoring with CloudWatch][3] for more.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the stream. [1]: https://aws.amazon.com/documentation/kinesis/ [2]: https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing [3]: https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the Kinesis Stream (same as id).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Kinesis Stream.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `The approximate UNIX timestamp that the stream was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the stream. The stream status is one of CREATING, DELETING, ACTIVE, or UPDATING.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Length of time (in hours) data records are accessible after they are added to the stream.`,
				},
				resource.Attribute{
					Name:        "open_shards",
					Description: `The list of shard ids in the OPEN state. See [Shard State][2] for more.`,
				},
				resource.Attribute{
					Name:        "closed_shards",
					Description: `The list of shard ids in the CLOSED state. See [Shard State][2] for more.`,
				},
				resource.Attribute{
					Name:        "shard_level_metrics",
					Description: `A list of shard-level CloudWatch metrics which are enabled for the stream. See [Monitoring with CloudWatch][3] for more.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the stream. [1]: https://aws.amazon.com/documentation/kinesis/ [2]: https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing [3]: https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_kms_alias",
			Category:         "Data Sources",
			ShortDescription: `Get information on a AWS Key Management Service (KMS) Alias`,
			Description: `

Use this data source to get the ARN of a KMS key alias.
By using this data source, you can reference key alias
without having to hard code the ARN as input.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name(ARN) of the key alias.`,
				},
				resource.Attribute{
					Name:        "target_key_id",
					Description: `Key identifier pointed to by the alias.`,
				},
				resource.Attribute{
					Name:        "target_key_arn",
					Description: `ARN pointed to by the alias.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name(ARN) of the key alias.`,
				},
				resource.Attribute{
					Name:        "target_key_id",
					Description: `Key identifier pointed to by the alias.`,
				},
				resource.Attribute{
					Name:        "target_key_arn",
					Description: `ARN pointed to by the alias.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_kms_ciphertext",
			Category:         "Data Sources",
			ShortDescription: `Provides ciphertext encrypted using a KMS key`,
			Description: `

The KMS ciphertext data source allows you to encrypt plaintext into ciphertext
by using an AWS KMS customer master key. The value returned by this data source
changes every apply. For a stable ciphertext value, see the [` + "`" + `aws_kms_ciphertext` + "`" + `
resource](/docs/providers/aws/r/kms_ciphertext.html).

~> **Note:** All arguments including the plaintext be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `(Required) Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required) Globally unique key ID for the customer master key.`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Optional) An optional mapping that makes up the encryption context. ## Attributes Reference All of the argument attributes are also exported as result attributes.`,
				},
				resource.Attribute{
					Name:        "ciphertext_blob",
					Description: `Base64 encoded ciphertext`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ciphertext_blob",
					Description: `Base64 encoded ciphertext`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_kms_key",
			Category:         "Data Sources",
			ShortDescription: `Get information on a AWS Key Management Service (KMS) Key`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required) Key identifier which can be one of the following format:`,
				},
				resource.Attribute{
					Name:        "grant_tokens",
					Description: `(Optional) List of grant tokens ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_kms_secrets",
			Category:         "Data Sources",
			ShortDescription: `Decrypt multiple secrets from data encrypted with the AWS KMS service`,
			Description: `

Decrypt multiple secrets from data encrypted with the AWS KMS service.

~> **NOTE**: Using this data provider will allow you to conceal secret data within your resource definitions but does not take care of protecting that data in all Terraform logging and state output. Please take care to secure your secret data beyond just the Terraform configuration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "secret",
					Description: `(Required) One or more encrypted payload definitions from the KMS service. See the Secret Definitions below. ### Secret Definitions Each ` + "`" + `secret` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to export this secret under in the attributes.`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `(Required) Base64 encoded payload, as returned from a KMS encrypt operation.`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Optional) An optional mapping that makes up the Encryption Context for the secret.`,
				},
				resource.Attribute{
					Name:        "plaintext",
					Description: `Map containing each ` + "`" + `secret` + "`" + ` ` + "`" + `name` + "`" + ` as the key with its decrypted plaintext value`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `Map containing each ` + "`" + `secret` + "`" + ` ` + "`" + `name` + "`" + ` as the key with its decrypted plaintext value`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_lambda_function",
			Category:         "Data Sources",
			ShortDescription: `Provides a Lambda Function data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required) Name of the lambda function.`,
				},
				resource.Attribute{
					Name:        "qualifier",
					Description: `(Optional) Alias name or version number of the lambda function. e.g. ` + "`" + `$LATEST` + "`" + `, ` + "`" + `my-alias` + "`" + `, or ` + "`" + `1` + "`" + ` ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Unqualified (no ` + "`" + `:QUALIFIER` + "`" + ` or ` + "`" + `:VERSION` + "`" + ` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also ` + "`" + `qualified_arn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dead_letter_config",
					Description: `Configure the function's`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of what your Lambda Function does.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The Lambda environment's configuration settings.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `The function entrypoint in your code.`,
				},
				resource.Attribute{
					Name:        "invoke_arn",
					Description: `The ARN to be used for invoking Lambda Function from API Gateway.`,
				},
				resource.Attribute{
					Name:        "kms_key_arn",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified.`,
				},
				resource.Attribute{
					Name:        "layers",
					Description: `A list of Lambda Layer ARNs attached to your Lambda Function.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Amount of memory in MB your Lambda Function can use at runtime.`,
				},
				resource.Attribute{
					Name:        "qualified_arn",
					Description: `Qualified (` + "`" + `:QUALIFIER` + "`" + ` or ` + "`" + `:VERSION` + "`" + ` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also ` + "`" + `arn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reserved_concurrent_executions",
					Description: `The amount of reserved concurrent executions for this lambda function or ` + "`" + `-1` + "`" + ` if unreserved.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `IAM role attached to the Lambda Function.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `The runtime environment for the Lambda function..`,
				},
				resource.Attribute{
					Name:        "source_code_hash",
					Description: `Base64-encoded representation of raw SHA-256 sum of the zip file.`,
				},
				resource.Attribute{
					Name:        "source_code_size",
					Description: `The size in bytes of the function .zip file.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The function execution time at which Lambda should terminate the function.`,
				},
				resource.Attribute{
					Name:        "tracing_config",
					Description: `Tracing settings of the function.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the Lambda function.`,
				},
				resource.Attribute{
					Name:        "vpc_config",
					Description: `VPC configuration associated with your Lambda function.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Unqualified (no ` + "`" + `:QUALIFIER` + "`" + ` or ` + "`" + `:VERSION` + "`" + ` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also ` + "`" + `qualified_arn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dead_letter_config",
					Description: `Configure the function's`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of what your Lambda Function does.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The Lambda environment's configuration settings.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `The function entrypoint in your code.`,
				},
				resource.Attribute{
					Name:        "invoke_arn",
					Description: `The ARN to be used for invoking Lambda Function from API Gateway.`,
				},
				resource.Attribute{
					Name:        "kms_key_arn",
					Description: `The ARN for the KMS encryption key.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified.`,
				},
				resource.Attribute{
					Name:        "layers",
					Description: `A list of Lambda Layer ARNs attached to your Lambda Function.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Amount of memory in MB your Lambda Function can use at runtime.`,
				},
				resource.Attribute{
					Name:        "qualified_arn",
					Description: `Qualified (` + "`" + `:QUALIFIER` + "`" + ` or ` + "`" + `:VERSION` + "`" + ` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also ` + "`" + `arn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reserved_concurrent_executions",
					Description: `The amount of reserved concurrent executions for this lambda function or ` + "`" + `-1` + "`" + ` if unreserved.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `IAM role attached to the Lambda Function.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `The runtime environment for the Lambda function..`,
				},
				resource.Attribute{
					Name:        "source_code_hash",
					Description: `Base64-encoded representation of raw SHA-256 sum of the zip file.`,
				},
				resource.Attribute{
					Name:        "source_code_size",
					Description: `The size in bytes of the function .zip file.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The function execution time at which Lambda should terminate the function.`,
				},
				resource.Attribute{
					Name:        "tracing_config",
					Description: `Tracing settings of the function.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the Lambda function.`,
				},
				resource.Attribute{
					Name:        "vpc_config",
					Description: `VPC configuration associated with your Lambda function.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_lambda_invocation",
			Category:         "Data Sources",
			ShortDescription: `Invoke AWS Lambda Function as data source`,
			Description: `

Use this data source to invoke custom lambda functions as data source.
The lambda function is invoked with [RequestResponse](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax)
invocation type.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required) The name of the lambda function.`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `(Required) A string in JSON format that is passed as payload to the lambda function.`,
				},
				resource.Attribute{
					Name:        "qualifier",
					Description: `(Optional) The qualifier (a.k.a version) of the lambda function. Defaults to ` + "`" + `$LATEST` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `String result of the lambda function invocation.`,
				},
				resource.Attribute{
					Name:        "result_map",
					Description: `This field is set only if result is a map of primitive types, where the map is string keys and string values. In Terraform 0.12 and later, use the [` + "`" + `jsondecode()` + "`" + ` function](/docs/configuration/functions/jsondecode.html) with the ` + "`" + `result` + "`" + ` attribute instead to convert the result to all supported native Terraform types.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `String result of the lambda function invocation.`,
				},
				resource.Attribute{
					Name:        "result_map",
					Description: `This field is set only if result is a map of primitive types, where the map is string keys and string values. In Terraform 0.12 and later, use the [` + "`" + `jsondecode()` + "`" + ` function](/docs/configuration/functions/jsondecode.html) with the ` + "`" + `result` + "`" + ` attribute instead to convert the result to all supported native Terraform types.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_lambda_layer_version",
			Category:         "Data Sources",
			ShortDescription: `Provides a Lambda Layer Version data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer_name",
					Description: `(Required) Name of the lambda layer.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Specific layer version. Conflicts with ` + "`" + `compatible_runtime` + "`" + `. If omitted, the latest available layer version will be used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the specific Lambda Layer version.`,
				},
				resource.Attribute{
					Name:        "license_info",
					Description: `License info associated with the specific Lambda Layer version.`,
				},
				resource.Attribute{
					Name:        "compatible_runtimes",
					Description: `A list of [Runtimes][1] the specific Lambda Layer version is compatible with.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the Lambda Layer with version.`,
				},
				resource.Attribute{
					Name:        "layer_arn",
					Description: `The Amazon Resource Name (ARN) of the Lambda Layer without version.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `The date this resource was created.`,
				},
				resource.Attribute{
					Name:        "source_code_hash",
					Description: `Base64-encoded representation of raw SHA-256 sum of the zip file.`,
				},
				resource.Attribute{
					Name:        "source_code_size",
					Description: `The size in bytes of the function .zip file.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `This Lamba Layer version. [1]: https://docs.aws.amazon.com/lambda/latest/dg/API_GetLayerVersion.html#SSS-GetLayerVersion-response-CompatibleRuntimes`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the specific Lambda Layer version.`,
				},
				resource.Attribute{
					Name:        "license_info",
					Description: `License info associated with the specific Lambda Layer version.`,
				},
				resource.Attribute{
					Name:        "compatible_runtimes",
					Description: `A list of [Runtimes][1] the specific Lambda Layer version is compatible with.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the Lambda Layer with version.`,
				},
				resource.Attribute{
					Name:        "layer_arn",
					Description: `The Amazon Resource Name (ARN) of the Lambda Layer without version.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `The date this resource was created.`,
				},
				resource.Attribute{
					Name:        "source_code_hash",
					Description: `Base64-encoded representation of raw SHA-256 sum of the zip file.`,
				},
				resource.Attribute{
					Name:        "source_code_size",
					Description: `The size in bytes of the function .zip file.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `This Lamba Layer version. [1]: https://docs.aws.amazon.com/lambda/latest/dg/API_GetLayerVersion.html#SSS-GetLayerVersion-response-CompatibleRuntimes`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_launch_configuration",
			Category:         "Data Sources",
			ShortDescription: `Provides a Launch Configuration data source.`,
			Description: `

Provides information about a Launch Configuration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the launch configuration. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the launch configuration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Name of the launch configuration.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The EC2 Image ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The Instance Type of the instance to launch.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `The IAM Instance Profile to associate with launched instances.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The Key Name that should be used for the instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A list of associated Security Group IDS.`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `Whether a Public IP address is associated with the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_classic_link_id",
					Description: `The ID of a ClassicLink-enabled VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_classic_link_security_groups",
					Description: `The IDs of one or more Security Groups for the specified ClassicLink-enabled VPC.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The User Data of the instance.`,
				},
				resource.Attribute{
					Name:        "enable_monitoring",
					Description: `Whether Detailed Monitoring is Enabled.`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `Whether the launched EC2 instance will be EBS-optimized.`,
				},
				resource.Attribute{
					Name:        "root_block_device",
					Description: `The Root Block Device of the instance.`,
				},
				resource.Attribute{
					Name:        "ebs_block_device",
					Description: `The EBS Block Devices attached to the instance.`,
				},
				resource.Attribute{
					Name:        "ephemeral_block_device",
					Description: `The Ephemeral volumes on the instance.`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `The Price to use for reserving Spot instances.`,
				},
				resource.Attribute{
					Name:        "placement_tenancy",
					Description: `The Tenancy of the instance. ` + "`" + `root_block_device` + "`" + ` is exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `Whether the EBS Volume will be deleted on instance termination.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the volume is Encrypted.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The provisioned IOPs of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The Size of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The Type of the volume. ` + "`" + `ebs_block_device` + "`" + ` is exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `Whether the EBS Volume will be deleted on instance termination.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The Name of the device.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The provisioned IOPs of the volume.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The Snapshot ID of the mount.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The Size of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The Type of the volume.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the volume is Encrypted. ` + "`" + `ephemeral_block_device` + "`" + ` is exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The Name of the device.`,
				},
				resource.Attribute{
					Name:        "virtual_name",
					Description: `The Virtual Name of the device.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the launch configuration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Name of the launch configuration.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The EC2 Image ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The Instance Type of the instance to launch.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `The IAM Instance Profile to associate with launched instances.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The Key Name that should be used for the instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A list of associated Security Group IDS.`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `Whether a Public IP address is associated with the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_classic_link_id",
					Description: `The ID of a ClassicLink-enabled VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_classic_link_security_groups",
					Description: `The IDs of one or more Security Groups for the specified ClassicLink-enabled VPC.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The User Data of the instance.`,
				},
				resource.Attribute{
					Name:        "enable_monitoring",
					Description: `Whether Detailed Monitoring is Enabled.`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `Whether the launched EC2 instance will be EBS-optimized.`,
				},
				resource.Attribute{
					Name:        "root_block_device",
					Description: `The Root Block Device of the instance.`,
				},
				resource.Attribute{
					Name:        "ebs_block_device",
					Description: `The EBS Block Devices attached to the instance.`,
				},
				resource.Attribute{
					Name:        "ephemeral_block_device",
					Description: `The Ephemeral volumes on the instance.`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `The Price to use for reserving Spot instances.`,
				},
				resource.Attribute{
					Name:        "placement_tenancy",
					Description: `The Tenancy of the instance. ` + "`" + `root_block_device` + "`" + ` is exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `Whether the EBS Volume will be deleted on instance termination.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the volume is Encrypted.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The provisioned IOPs of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The Size of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The Type of the volume. ` + "`" + `ebs_block_device` + "`" + ` is exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `Whether the EBS Volume will be deleted on instance termination.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The Name of the device.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The provisioned IOPs of the volume.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The Snapshot ID of the mount.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The Size of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The Type of the volume.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the volume is Encrypted. ` + "`" + `ephemeral_block_device` + "`" + ` is exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The Name of the device.`,
				},
				resource.Attribute{
					Name:        "virtual_name",
					Description: `The Virtual Name of the device.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_launch_template",
			Category:         "Data Sources",
			ShortDescription: `Provides a Launch Template data source.`,
			Description: `

Provides information about a Launch Template.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the launch template. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of the launch template.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the launch template.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `The default version of the launch template.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The latest version of the launch template.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the launch template.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `Specify volumes to attach to the instance besides the volumes specified by the AMI.`,
				},
				resource.Attribute{
					Name:        "credit_specification",
					Description: `Customize the credit specification of the instance. See [Credit Specification](#credit-specification) below for more details.`,
				},
				resource.Attribute{
					Name:        "disable_api_termination",
					Description: `If ` + "`" + `true` + "`" + `, enables [EC2 Instance Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `If ` + "`" + `true` + "`" + `, the launched EC2 instance will be EBS-optimized.`,
				},
				resource.Attribute{
					Name:        "elastic_gpu_specifications",
					Description: `The elastic GPU to attach to the instance. See [Elastic GPU](#elastic-gpu) below for more details.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `The IAM Instance Profile to launch the instance with. See [Instance Profile](#instance-profile) below for more details.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The AMI from which to launch the instance.`,
				},
				resource.Attribute{
					Name:        "instance_initiated_shutdown_behavior",
					Description: `Shutdown behavior for the instance. Can be ` + "`" + `stop` + "`" + ` or ` + "`" + `terminate` + "`" + `. (Default: ` + "`" + `stop` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "instance_market_options",
					Description: `The market (purchasing) option for the instance. below for details.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the instance.`,
				},
				resource.Attribute{
					Name:        "kernel_id",
					Description: `The kernel ID.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The key name to use for the instance.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `The monitoring option for the instance.`,
				},
				resource.Attribute{
					Name:        "network_interfaces",
					Description: `Customize network interfaces to be attached at instance boot time. See [Network Interfaces](#network-interfaces) below for more details.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `The placement of the instance.`,
				},
				resource.Attribute{
					Name:        "ram_disk_id",
					Description: `The ID of the RAM disk.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `A list of security group names to associate with. If you are creating Instances in a VPC, use ` + "`" + `vpc_security_group_ids` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "vpc_security_group_ids",
					Description: `A list of security group IDs to associate with.`,
				},
				resource.Attribute{
					Name:        "tag_specifications",
					Description: `The tags to apply to the resources during launch.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the launch template.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded user data to provide when launching the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of the launch template.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the launch template.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `The default version of the launch template.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The latest version of the launch template.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the launch template.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `Specify volumes to attach to the instance besides the volumes specified by the AMI.`,
				},
				resource.Attribute{
					Name:        "credit_specification",
					Description: `Customize the credit specification of the instance. See [Credit Specification](#credit-specification) below for more details.`,
				},
				resource.Attribute{
					Name:        "disable_api_termination",
					Description: `If ` + "`" + `true` + "`" + `, enables [EC2 Instance Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `If ` + "`" + `true` + "`" + `, the launched EC2 instance will be EBS-optimized.`,
				},
				resource.Attribute{
					Name:        "elastic_gpu_specifications",
					Description: `The elastic GPU to attach to the instance. See [Elastic GPU](#elastic-gpu) below for more details.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `The IAM Instance Profile to launch the instance with. See [Instance Profile](#instance-profile) below for more details.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The AMI from which to launch the instance.`,
				},
				resource.Attribute{
					Name:        "instance_initiated_shutdown_behavior",
					Description: `Shutdown behavior for the instance. Can be ` + "`" + `stop` + "`" + ` or ` + "`" + `terminate` + "`" + `. (Default: ` + "`" + `stop` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "instance_market_options",
					Description: `The market (purchasing) option for the instance. below for details.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the instance.`,
				},
				resource.Attribute{
					Name:        "kernel_id",
					Description: `The kernel ID.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The key name to use for the instance.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `The monitoring option for the instance.`,
				},
				resource.Attribute{
					Name:        "network_interfaces",
					Description: `Customize network interfaces to be attached at instance boot time. See [Network Interfaces](#network-interfaces) below for more details.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `The placement of the instance.`,
				},
				resource.Attribute{
					Name:        "ram_disk_id",
					Description: `The ID of the RAM disk.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `A list of security group names to associate with. If you are creating Instances in a VPC, use ` + "`" + `vpc_security_group_ids` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "vpc_security_group_ids",
					Description: `A list of security group IDs to associate with.`,
				},
				resource.Attribute{
					Name:        "tag_specifications",
					Description: `The tags to apply to the resources during launch.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the launch template.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded user data to provide when launching the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_lb",
			Category:         "Data Sources",
			ShortDescription: `Provides a Load Balancer data source.`,
			Description: `

~> **Note:** ` + "`" + `aws_alb` + "`" + ` is known as ` + "`" + `aws_lb` + "`" + `. The functionality is identical.

Provides information about a Load Balancer.

This data source can prove useful when a module accepts an LB as an input
variable and needs to, for example, determine the security groups associated
with it, etc.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `(Optional) The full ARN of the load balancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The unique name of the load balancer. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_lb_listener",
			Category:         "Data Sources",
			ShortDescription: `Provides a Load Balancer Listener data source.`,
			Description: `

~> **Note:** ` + "`" + `aws_alb_listener` + "`" + ` is known as ` + "`" + `aws_lb_listener` + "`" + `. The functionality is identical.

Provides information about a Load Balancer Listener.

This data source can prove useful when a module accepts an LB Listener as an
input variable and needs to know the LB it is attached to, or other
information specific to the listener in question.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `(Optional) The arn of the listener. Required if ` + "`" + `load_balancer_arn` + "`" + ` and ` + "`" + `port` + "`" + ` is not set.`,
				},
				resource.Attribute{
					Name:        "load_balancer_arn",
					Description: `(Optional) The arn of the load balancer. Required if ` + "`" + `arn` + "`" + ` is not set.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port of the listener. Required if ` + "`" + `arn` + "`" + ` is not set. ## Attributes Reference See the [LB Listener Resource](/docs/providers/aws/r/lb_listener.html) for details on the returned attributes - they are identical.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_lb_target_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a Load Balancer Target Group data source.`,
			Description: `

~> **Note:** ` + "`" + `aws_alb_target_group` + "`" + ` is known as ` + "`" + `aws_lb_target_group` + "`" + `. The functionality is identical.

Provides information about a Load Balancer Target Group.

This data source can prove useful when a module accepts an LB Target Group as an
input variable and needs to know its attributes. It can also be used to get the ARN of
an LB Target Group for use in other resources, given LB Target Group name.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `(Optional) The full ARN of the target group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The unique name of the target group. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_mq_broker",
			Category:         "Data Sources",
			ShortDescription: `Provides a MQ Broker data source.`,
			Description: `

Provides information about a MQ Broker.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "broker_id",
					Description: `(Optional) The unique id of the mq broker.`,
				},
				resource.Attribute{
					Name:        "broker_name",
					Description: `(Optional) The unique name of the mq broker. ## Attributes Reference See the [` + "`" + `aws_mq_broker` + "`" + ` resource](/docs/providers/aws/r/mq_broker.html) for details on the returned attributes. They are identical except for user password, which is not returned when describing broker.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_msk_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Amazon MSK Cluster`,
			Description: `

Get information on an Amazon MSK Cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Name of the cluster. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of the MSK cluster.`,
				},
				resource.Attribute{
					Name:        "bootstrap_brokers",
					Description: `A comma separated list of one or more hostname:port pairs of Kafka brokers suitable to boostrap connectivity to the Kafka cluster.`,
				},
				resource.Attribute{
					Name:        "bootstrap_brokers_tls",
					Description: `A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster.`,
				},
				resource.Attribute{
					Name:        "kafka_version",
					Description: `Apache Kafka version.`,
				},
				resource.Attribute{
					Name:        "number_of_broker_nodes",
					Description: `Number of broker nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Map of key-value pairs assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "zookeeper_connect_string",
					Description: `A comma separated list of one or more IP:port pairs to use to connect to the Apache Zookeeper cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of the MSK cluster.`,
				},
				resource.Attribute{
					Name:        "bootstrap_brokers",
					Description: `A comma separated list of one or more hostname:port pairs of Kafka brokers suitable to boostrap connectivity to the Kafka cluster.`,
				},
				resource.Attribute{
					Name:        "bootstrap_brokers_tls",
					Description: `A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster.`,
				},
				resource.Attribute{
					Name:        "kafka_version",
					Description: `Apache Kafka version.`,
				},
				resource.Attribute{
					Name:        "number_of_broker_nodes",
					Description: `Number of broker nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Map of key-value pairs assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "zookeeper_connect_string",
					Description: `A comma separated list of one or more IP:port pairs to use to connect to the Apache Zookeeper cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_msk_configuration",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Amazon MSK Configuration`,
			Description: `

Get information on an Amazon MSK Configuration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the configuration. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of the configuration.`,
				},
				resource.Attribute{
					Name:        "latest_revision",
					Description: `Latest revision of the configuration.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the configuration.`,
				},
				resource.Attribute{
					Name:        "kafka_versions",
					Description: `List of Apache Kafka versions which can use this configuration.`,
				},
				resource.Attribute{
					Name:        "server_properties",
					Description: `Contents of the server.properties file.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of the configuration.`,
				},
				resource.Attribute{
					Name:        "latest_revision",
					Description: `Latest revision of the configuration.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the configuration.`,
				},
				resource.Attribute{
					Name:        "kafka_versions",
					Description: `List of Apache Kafka versions which can use this configuration.`,
				},
				resource.Attribute{
					Name:        "server_properties",
					Description: `Contents of the server.properties file.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_nat_gateway",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Nat Gateway`,
			Description: `

Provides details about a specific Nat Gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the specific Nat Gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The id of subnet that the Nat Gateway resides in.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The id of the VPC that the Nat Gateway resides in.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The state of the NAT gateway (pending | failed | available | deleting | deleted ).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeNatGateways.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. An Nat Gateway will be selected if any one of the given values matches. ## Attributes Reference All of the argument attributes except ` + "`" + `filter` + "`" + ` block are also exported as result attributes. This data source will complete the data by populating any fields that are not included in the configuration with the data for the selected Nat Gateway. ` + "`" + `addresses` + "`" + ` are also exported with the following attributes, when they are relevant: Each attachement supports the following:`,
				},
				resource.Attribute{
					Name:        "allocation_id",
					Description: `The Id of the EIP allocated to the selected Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The Id of the ENI allocated to the selected Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private Ip address of the selected Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public Ip (EIP) address of the selected Nat Gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_id",
					Description: `The Id of the EIP allocated to the selected Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The Id of the ENI allocated to the selected Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private Ip address of the selected Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public Ip (EIP) address of the selected Nat Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_network_acls",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of network ACL ids for a VPC`,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The VPC ID that you want to filter from.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired network ACLs.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeNetworkAcls.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. A VPC will be selected if any one of the given values matches. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the network ACL ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the network ACL ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_network_interface",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Network Interface resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "association",
					Description: `The association information for an Elastic IP address (IPv4) associated with the network interface. See supported fields below.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The Availability Zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the network interface.`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `The type of interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_addresses",
					Description: `List of IPv6 addresses to assign to the ENI.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The AWS account ID of the owner of the network interface.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The private DNS name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IPv4 address of the network interface within the subnet.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IPv4 addresses associated with the network interface.`,
				},
				resource.Attribute{
					Name:        "requester_id",
					Description: `The ID of the entity that launched the instance on your behalf.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `The list of security groups for the network interface.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Any tags assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC. ### ` + "`" + `association` + "`" + ``,
				},
				resource.Attribute{
					Name:        "allocation_id",
					Description: `The allocation ID.`,
				},
				resource.Attribute{
					Name:        "association_id",
					Description: `The association ID.`,
				},
				resource.Attribute{
					Name:        "ip_owner_id",
					Description: `The ID of the Elastic IP address owner.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The public DNS name.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The address of the Elastic IP address bound to the network interface. ## Import Elastic Network Interfaces can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aws_network_interface.test eni-12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "association",
					Description: `The association information for an Elastic IP address (IPv4) associated with the network interface. See supported fields below.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The Availability Zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the network interface.`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `The type of interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_addresses",
					Description: `List of IPv6 addresses to assign to the ENI.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The AWS account ID of the owner of the network interface.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The private DNS name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IPv4 address of the network interface within the subnet.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IPv4 addresses associated with the network interface.`,
				},
				resource.Attribute{
					Name:        "requester_id",
					Description: `The ID of the entity that launched the instance on your behalf.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `The list of security groups for the network interface.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Any tags assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC. ### ` + "`" + `association` + "`" + ``,
				},
				resource.Attribute{
					Name:        "allocation_id",
					Description: `The allocation ID.`,
				},
				resource.Attribute{
					Name:        "association_id",
					Description: `The association ID.`,
				},
				resource.Attribute{
					Name:        "ip_owner_id",
					Description: `The ID of the Elastic IP address owner.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The public DNS name.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The address of the Elastic IP address bound to the network interface. ## Import Elastic Network Interfaces can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aws_network_interface.test eni-12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_network_interfaces",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of network interface ids`,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired network interfaces.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeNetworkInterfaces.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the network interface ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the network interface ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_organizations_organization",
			Category:         "Data Sources",
			ShortDescription: `Get information about the organization that the user's account belongs to`,
			Description: `

Get information about the organization that the user's account belongs to

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the organization.`,
				},
				resource.Attribute{
					Name:        "feature_set",
					Description: `The FeatureSet of the organization.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the organization.`,
				},
				resource.Attribute{
					Name:        "master_account_arn",
					Description: `The Amazon Resource Name (ARN) of the account that is designated as the master account for the organization.`,
				},
				resource.Attribute{
					Name:        "master_account_email",
					Description: `The email address that is associated with the AWS account that is designated as the master account for the organization.`,
				},
				resource.Attribute{
					Name:        "master_account_id",
					Description: `The unique identifier (ID) of the master account of an organization. ### Master Account Attributes Reference If the account is the master account for the organization, the following attributes are also exported:`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `List of organization accounts including the master account. For a list excluding the master account, see the ` + "`" + `non_master_accounts` + "`" + ` attribute. All elements have these attributes:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `ARN of the account`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email of the account`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Identifier of the account`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the account`,
				},
				resource.Attribute{
					Name:        "aws_service_access_principals",
					Description: `A list of AWS service principal names that have integration enabled with your organization. Organization must have ` + "`" + `feature_set` + "`" + ` set to ` + "`" + `ALL` + "`" + `. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).`,
				},
				resource.Attribute{
					Name:        "enabled_policy_types",
					Description: `A list of Organizations policy types that are enabled in the Organization Root. Organization must have ` + "`" + `feature_set` + "`" + ` set to ` + "`" + `ALL` + "`" + `. For additional information about valid policy types (e.g. ` + "`" + `SERVICE_CONTROL_POLICY` + "`" + `), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).`,
				},
				resource.Attribute{
					Name:        "non_master_accounts",
					Description: `List of organization accounts excluding the master account. For a list including the master account, see the ` + "`" + `accounts` + "`" + ` attribute. All elements have these attributes:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `ARN of the account`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email of the account`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Identifier of the account`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the account`,
				},
				resource.Attribute{
					Name:        "roots",
					Description: `List of organization roots. All elements have these attributes:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `ARN of the root`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Identifier of the root`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the root`,
				},
				resource.Attribute{
					Name:        "policy_types",
					Description: `List of policy types enabled for this root. All elements have these attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the policy type`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the policy type as it relates to the associated root`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the organization.`,
				},
				resource.Attribute{
					Name:        "feature_set",
					Description: `The FeatureSet of the organization.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the organization.`,
				},
				resource.Attribute{
					Name:        "master_account_arn",
					Description: `The Amazon Resource Name (ARN) of the account that is designated as the master account for the organization.`,
				},
				resource.Attribute{
					Name:        "master_account_email",
					Description: `The email address that is associated with the AWS account that is designated as the master account for the organization.`,
				},
				resource.Attribute{
					Name:        "master_account_id",
					Description: `The unique identifier (ID) of the master account of an organization. ### Master Account Attributes Reference If the account is the master account for the organization, the following attributes are also exported:`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `List of organization accounts including the master account. For a list excluding the master account, see the ` + "`" + `non_master_accounts` + "`" + ` attribute. All elements have these attributes:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `ARN of the account`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email of the account`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Identifier of the account`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the account`,
				},
				resource.Attribute{
					Name:        "aws_service_access_principals",
					Description: `A list of AWS service principal names that have integration enabled with your organization. Organization must have ` + "`" + `feature_set` + "`" + ` set to ` + "`" + `ALL` + "`" + `. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).`,
				},
				resource.Attribute{
					Name:        "enabled_policy_types",
					Description: `A list of Organizations policy types that are enabled in the Organization Root. Organization must have ` + "`" + `feature_set` + "`" + ` set to ` + "`" + `ALL` + "`" + `. For additional information about valid policy types (e.g. ` + "`" + `SERVICE_CONTROL_POLICY` + "`" + `), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).`,
				},
				resource.Attribute{
					Name:        "non_master_accounts",
					Description: `List of organization accounts excluding the master account. For a list including the master account, see the ` + "`" + `accounts` + "`" + ` attribute. All elements have these attributes:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `ARN of the account`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email of the account`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Identifier of the account`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the account`,
				},
				resource.Attribute{
					Name:        "roots",
					Description: `List of organization roots. All elements have these attributes:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `ARN of the root`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Identifier of the root`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the root`,
				},
				resource.Attribute{
					Name:        "policy_types",
					Description: `List of policy types enabled for this root. All elements have these attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the policy type`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the policy type as it relates to the associated root`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_partition",
			Category:         "Data Sources",
			ShortDescription: `Get AWS partition identifier`,
			Description: `

Use this data source to lookup information about the current AWS partition in
which Terraform is working.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_prefix_list",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific prefix list`,
			Description: `

` + "`" + `aws_prefix_list` + "`" + ` provides details about a specific prefix list (PL)
in the current region.

This can be used both to validate a prefix list given in a variable
and to obtain the CIDR blocks (IP address ranges) for the associated
AWS service. The latter may be useful e.g. for adding network ACL
rules.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "prefix_list_id",
					Description: `(Optional) The ID of the prefix list to select.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the prefix list to select. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the selected prefix list.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the selected prefix list.`,
				},
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The list of CIDR blocks for the AWS service associated with the prefix list.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the selected prefix list.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the selected prefix list.`,
				},
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The list of CIDR blocks for the AWS service associated with the prefix list.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_pricing_product",
			Category:         "Data Sources",
			ShortDescription: `Get information regarding the pricing of an Amazon product`,
			Description: `

Use this data source to get the pricing information of all products in AWS.
This data source is only available in a us-east-1 or ap-south-1 provider.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_code",
					Description: `(Required) The code of the service. Available service codes can be fetched using the DescribeServices pricing API call.`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Required) A list of filters. Passed directly to the API (see GetProducts API reference). These filters must describe a single product, this resource will fail if more than one product is returned by the API. ### filters`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `Set to the product returned from the API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `Set to the product returned from the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ram_resource_share",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about a RAM Resource Share`,
			Description: `

` + "`" + `aws_ram_resource_share` + "`" + ` Retrieve information about a RAM Resource Share.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the resource share to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags. See [related docs] (https://docs.aws.amazon.com/ram/latest/APIReference/API_TagFilter.html).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag key to filter on.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the tag key. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the resource share.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Amazon Resource Name (ARN) of the resource share.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The Status of the RAM share.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The Tags attached to the RAM share`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the resource share.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Amazon Resource Name (ARN) of the resource share.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The Status of the RAM share.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The Tags attached to the RAM share`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_rds_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides a RDS cluster data source.`,
			Description: `

Provides information about a RDS cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_identifier",
					Description: `(Required) The cluster identifier of the RDS cluster. ## Attributes Reference See the [RDS Cluster Resource](/docs/providers/aws/r/rds_cluster.html) for details on the returned attributes - they are identical.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_redshift_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific redshift cluster`,
			Description: `

Provides details about a specific redshift cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_identifier",
					Description: `(Required) The cluster identifier ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allow_version_upgrade",
					Description: `Whether major version upgrades can be applied during maintenance period`,
				},
				resource.Attribute{
					Name:        "automated_snapshot_retention_period",
					Description: `The backup retention period`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the cluster`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `The name of the S3 bucket where the log files are to be stored`,
				},
				resource.Attribute{
					Name:        "cluster_identifier",
					Description: `The cluster identifier`,
				},
				resource.Attribute{
					Name:        "cluster_parameter_group_name",
					Description: `The name of the parameter group to be associated with this cluster`,
				},
				resource.Attribute{
					Name:        "cluster_public_key",
					Description: `The public key for the cluster`,
				},
				resource.Attribute{
					Name:        "cluster_revision_number",
					Description: `The cluster revision number`,
				},
				resource.Attribute{
					Name:        "cluster_security_groups",
					Description: `The security groups associated with the cluster`,
				},
				resource.Attribute{
					Name:        "cluster_subnet_group_name",
					Description: `The name of a cluster subnet group to be associated with this cluster`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The cluster type`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the default database in the cluster`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `The Elastic IP of the cluster`,
				},
				resource.Attribute{
					Name:        "enable_logging",
					Description: `Whether cluster logging is enabled`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the cluster data is encrypted`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The cluster endpoint`,
				},
				resource.Attribute{
					Name:        "enhanced_vpc_routing",
					Description: `Whether enhanced VPC routing is enabled`,
				},
				resource.Attribute{
					Name:        "iam_roles",
					Description: `The IAM roles associated to the cluster`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The KMS encryption key associated to the cluster`,
				},
				resource.Attribute{
					Name:        "master_username",
					Description: `Username for the master DB user`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `The cluster node type`,
				},
				resource.Attribute{
					Name:        "number_of_nodes",
					Description: `The number of nodes in the cluster`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port the cluster responds on`,
				},
				resource.Attribute{
					Name:        "preferred_maintenance_window",
					Description: `The maintenance window`,
				},
				resource.Attribute{
					Name:        "publicly_accessible",
					Description: `Whether the cluster is publicly accessible`,
				},
				resource.Attribute{
					Name:        "s3_key_prefix",
					Description: `The folder inside the S3 bucket where the log files are stored`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags associated to the cluster`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC Id associated with the cluster`,
				},
				resource.Attribute{
					Name:        "vpc_security_group_ids",
					Description: `The VPC security group Ids associated with the cluster`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_version_upgrade",
					Description: `Whether major version upgrades can be applied during maintenance period`,
				},
				resource.Attribute{
					Name:        "automated_snapshot_retention_period",
					Description: `The backup retention period`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the cluster`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `The name of the S3 bucket where the log files are to be stored`,
				},
				resource.Attribute{
					Name:        "cluster_identifier",
					Description: `The cluster identifier`,
				},
				resource.Attribute{
					Name:        "cluster_parameter_group_name",
					Description: `The name of the parameter group to be associated with this cluster`,
				},
				resource.Attribute{
					Name:        "cluster_public_key",
					Description: `The public key for the cluster`,
				},
				resource.Attribute{
					Name:        "cluster_revision_number",
					Description: `The cluster revision number`,
				},
				resource.Attribute{
					Name:        "cluster_security_groups",
					Description: `The security groups associated with the cluster`,
				},
				resource.Attribute{
					Name:        "cluster_subnet_group_name",
					Description: `The name of a cluster subnet group to be associated with this cluster`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The cluster type`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the default database in the cluster`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `The Elastic IP of the cluster`,
				},
				resource.Attribute{
					Name:        "enable_logging",
					Description: `Whether cluster logging is enabled`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the cluster data is encrypted`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The cluster endpoint`,
				},
				resource.Attribute{
					Name:        "enhanced_vpc_routing",
					Description: `Whether enhanced VPC routing is enabled`,
				},
				resource.Attribute{
					Name:        "iam_roles",
					Description: `The IAM roles associated to the cluster`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The KMS encryption key associated to the cluster`,
				},
				resource.Attribute{
					Name:        "master_username",
					Description: `Username for the master DB user`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `The cluster node type`,
				},
				resource.Attribute{
					Name:        "number_of_nodes",
					Description: `The number of nodes in the cluster`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port the cluster responds on`,
				},
				resource.Attribute{
					Name:        "preferred_maintenance_window",
					Description: `The maintenance window`,
				},
				resource.Attribute{
					Name:        "publicly_accessible",
					Description: `Whether the cluster is publicly accessible`,
				},
				resource.Attribute{
					Name:        "s3_key_prefix",
					Description: `The folder inside the S3 bucket where the log files are stored`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags associated to the cluster`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC Id associated with the cluster`,
				},
				resource.Attribute{
					Name:        "vpc_security_group_ids",
					Description: `The VPC security group Ids associated with the cluster`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_redshift_service_account",
			Category:         "Data Sources",
			ShortDescription: `Get AWS Redshift Service Account for storing audit data in S3.`,
			Description: `

Use this data source to get the Account ID of the [AWS Redshift Service Account](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
in a given region for the purpose of allowing Redshift to store audit data in S3.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Name of the region whose AWS Redshift account ID is desired. Defaults to the region from the AWS provider configuration. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the AWS Redshift service account in the selected region.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the AWS Redshift service account in the selected region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the AWS Redshift service account in the selected region.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the AWS Redshift service account in the selected region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_region",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific service region`,
			Description: `

` + "`" + `aws_region` + "`" + ` provides details about a specific AWS region.

As well as validating a given region name this resource can be used to
discover the name of the region configured within the provider. The latter
can be useful in a child module which is inheriting an AWS provider
configuration from its parent module.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The full name of the region to select.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) The EC2 endpoint of the region to select. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the selected region.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The EC2 endpoint for the selected region.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The region's description in this format: "Location (Region name)".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the selected region.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The EC2 endpoint for the selected region.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The region's description in this format: "Location (Region name)".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_route",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Route`,
			Description: `

` + "`" + `aws_route` + "`" + ` provides details about a specific Route.

This resource can prove useful when finding the resource
associated with a CIDR. For example, finding the peering
connection associated with a CIDR value.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) The id of the specific Route Table containing the Route entry.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `(Optional) The CIDR block of the Route belonging to the Route Table.`,
				},
				resource.Attribute{
					Name:        "destination_ipv6_cidr_block",
					Description: `(Optional) The IPv6 CIDR block of the Route belonging to the Route Table.`,
				},
				resource.Attribute{
					Name:        "egress_only_gateway_id",
					Description: `(Optional) The Egress Only Gateway ID of the Route belonging to the Route Table.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `(Optional) The Gateway ID of the Route belonging to the Route Table.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) The Instance ID of the Route belonging to the Route Table.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `(Optional) The NAT Gateway ID of the Route belonging to the Route Table.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `(Optional) The EC2 Transit Gateway ID of the Route belonging to the Route Table.`,
				},
				resource.Attribute{
					Name:        "vpc_peering_connection_id",
					Description: `(Optional) The VPC Peering Connection ID of the Route belonging to the Route Table.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `(Optional) The Network Interface ID of the Route belonging to the Route Table. ## Attributes Reference All of the argument attributes are also exported as result attributes when there is data available. For example, the ` + "`" + `vpc_peering_connection_id` + "`" + ` field will be empty when the route is attached to a Network Interface.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_route53_delegation_set",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Route 53 Delegation Set`,
			Description: `

` + "`" + `aws_route53_delegation_set` + "`" + ` provides details about a specific Route 53 Delegation Set.

This data source allows to find a list of name servers associated with a specific delegation set.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The Hosted Zone id of the desired delegation set. The following attribute is additionally exported:`,
				},
				resource.Attribute{
					Name:        "caller_reference",
					Description: `Caller Reference of the delegation set.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `The list of DNS name servers for the delegation set.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_route53_resolver_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Route53 Resolver rule`,
			Description: `

` + "`" + `aws_route53_resolver_rule` + "`" + ` provides details about a specific Route53 Resolver rule.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) The domain name the desired resolver rule forwards DNS queries for. Conflicts with ` + "`" + `resolver_rule_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The friendly name of the desired resolver rule. Conflicts with ` + "`" + `resolver_rule_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Optional) The rule type of the desired resolver rule. Valid values are ` + "`" + `FORWARD` + "`" + `, ` + "`" + `SYSTEM` + "`" + ` and ` + "`" + `RECURSIVE` + "`" + `. Conflicts with ` + "`" + `resolver_rule_id` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resolver rule.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN (Amazon Resource Name) for the resolver rule.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.`,
				},
				resource.Attribute{
					Name:        "share_status",
					Description: `Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account. Values are ` + "`" + `NOT_SHARED` + "`" + `, ` + "`" + `SHARED_BY_ME` + "`" + ` or ` + "`" + `SHARED_WITH_ME` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resolver rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resolver rule.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN (Amazon Resource Name) for the resolver rule.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.`,
				},
				resource.Attribute{
					Name:        "share_status",
					Description: `Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account. Values are ` + "`" + `NOT_SHARED` + "`" + `, ` + "`" + `SHARED_BY_ME` + "`" + ` or ` + "`" + `SHARED_WITH_ME` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resolver rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_route53_resolver_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a set of Route53 Resolver rules`,
			Description: `

` + "`" + `aws_route53_resolver_rules` + "`" + ` provides details about a set of Route53 Resolver rules.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resolver_rule_ids",
					Description: `The IDs of the matched resolver rules.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resolver_rule_ids",
					Description: `The IDs of the matched resolver rules.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_route53_zone",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Route 53 Hosted Zone`,
			Description: `

` + "`" + `aws_route53_zone` + "`" + ` provides details about a specific Route 53 Hosted Zone.

This data source allows to find a Hosted Zone ID given Hosted Zone name and certain search criteria.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The Hosted Zone id of the desired Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Hosted Zone name of the desired Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "private_zone",
					Description: `(Optional) Used with ` + "`" + `name` + "`" + ` field to get a private Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Used with ` + "`" + `name` + "`" + ` field to get a private Hosted Zone associated with the vpc_id (in this case, private_zone is not mandatory).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Used with ` + "`" + `name` + "`" + ` field. A mapping of tags, each pair of which must exactly match a pair on the desired Hosted Zone. ## Attributes Reference All of the argument attributes are also exported as result attributes. This data source will complete the data by populating any fields that are not included in the configuration with the data for the selected Hosted Zone. The following attribute is additionally exported:`,
				},
				resource.Attribute{
					Name:        "caller_reference",
					Description: `Caller Reference of the Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `The comment field of the Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `The list of DNS name servers for the Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "resource_record_set_count",
					Description: `The number of Record Set in the Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "linked_service_principal",
					Description: `The service that created the Hosted Zone (e.g. ` + "`" + `servicediscovery.amazonaws.com` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "linked_service_description",
					Description: `The description provided by the service that created the Hosted Zone (e.g. ` + "`" + `arn:aws:servicediscovery:us-east-1:1234567890:namespace/ns-xxxxxxxxxxxxxxxx` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "caller_reference",
					Description: `Caller Reference of the Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `The comment field of the Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `The list of DNS name servers for the Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "resource_record_set_count",
					Description: `The number of Record Set in the Hosted Zone.`,
				},
				resource.Attribute{
					Name:        "linked_service_principal",
					Description: `The service that created the Hosted Zone (e.g. ` + "`" + `servicediscovery.amazonaws.com` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "linked_service_description",
					Description: `The description provided by the service that created the Hosted Zone (e.g. ` + "`" + `arn:aws:servicediscovery:us-east-1:1234567890:namespace/ns-xxxxxxxxxxxxxxxx` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_route_table",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Route Table`,
			Description: `

` + "`" + `aws_route_table` + "`" + ` provides details about a specific Route Table.

This resource can prove useful when a module accepts a Subnet id as
an input variable and needs to, for example, add a route in
the Route Table.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) The id of the specific Route Table to retrieve.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired Route Table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The id of the VPC that the desired Route Table belongs to.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The id of a Subnet which is connected to the Route Table (not be exported if not given in parameter). More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeRouteTables.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. A Route Table will be selected if any one of the given values matches. ## Attributes Reference All of the argument attributes except ` + "`" + `filter` + "`" + ` and ` + "`" + `subnet_id` + "`" + ` blocks are also exported as result attributes. This data source will complete the data by populating any fields that are not included in the configuration with the data for the selected Route Table. In addition the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the route table ` + "`" + `routes` + "`" + ` are also exported with the following attributes, when there are relevants: Each route supports the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the route.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_block",
					Description: `The IPv6 CIDR block of the route.`,
				},
				resource.Attribute{
					Name:        "egress_only_gateway_id",
					Description: `The ID of the Egress Only Internet Gateway.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The Internet Gateway ID.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The NAT Gateway ID.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The EC2 instance ID.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `The EC2 Transit Gateway ID.`,
				},
				resource.Attribute{
					Name:        "vpc_peering_connection_id",
					Description: `The VPC Peering ID.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The ID of the elastic network interface (eni) to use. ` + "`" + `associations` + "`" + ` are also exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "route_table_association_id",
					Description: `The Association ID .`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The Subnet ID.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If the Association due to the Main Route Table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the route table ` + "`" + `routes` + "`" + ` are also exported with the following attributes, when there are relevants: Each route supports the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the route.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_block",
					Description: `The IPv6 CIDR block of the route.`,
				},
				resource.Attribute{
					Name:        "egress_only_gateway_id",
					Description: `The ID of the Egress Only Internet Gateway.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The Internet Gateway ID.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The NAT Gateway ID.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The EC2 instance ID.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `The EC2 Transit Gateway ID.`,
				},
				resource.Attribute{
					Name:        "vpc_peering_connection_id",
					Description: `The VPC Peering ID.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The ID of the elastic network interface (eni) to use. ` + "`" + `associations` + "`" + ` are also exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "route_table_association_id",
					Description: `The Association ID .`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The Subnet ID.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If the Association due to the Main Route Table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_route_tables",
			Category:         "Data Sources",
			ShortDescription: `Get information on Amazon route tables.`,
			Description: `

This resource can be useful for getting back a list of route table ids to be referenced elsewhere.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The VPC ID that you want to filter from.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired route tables. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeRouteTables.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. A Route Table will be selected if any one of the given values matches. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the route table ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the route table ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_s3_bucket",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific S3 bucket`,
			Description: `

Provides details about a specific S3 bucket.

This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
Distribution.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the bucket. Will be of format ` + "`" + `arn:aws:s3:::bucketname` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.s3.amazonaws.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_regional_domain_name",
					Description: `The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The AWS region this bucket resides in.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the bucket. Will be of format ` + "`" + `arn:aws:s3:::bucketname` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.s3.amazonaws.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_regional_domain_name",
					Description: `The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The AWS region this bucket resides in.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_s3_bucket_object",
			Category:         "Data Sources",
			ShortDescription: `Provides metadata and optionally content of an S3 object`,
			Description: `

The S3 object data source allows access to the metadata and
_optionally_ (see below) content of an object stored inside S3 bucket.

~> **Note:** The content of an object (` + "`" + `body` + "`" + ` field) is available only for objects which have a human-readable ` + "`" + `Content-Type` + "`" + ` (` + "`" + `text/*` + "`" + ` and ` + "`" + `application/json` + "`" + `). This is to prevent printing unsafe characters and potentially downloading large amount of data which would be thrown away in favour of metadata.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to read the object from`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The full path to the object inside the bucket`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) Specific version ID of the object returned (defaults to latest version) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `[Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for ` + "`" + `Standard` + "`" + ` storage class objects.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `[Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for ` + "`" + `Standard` + "`" + ` storage class objects.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_s3_bucket_objects",
			Category:         "Data Sources",
			ShortDescription: `Returns keys and metadata of S3 objects`,
			Description: `

~> **NOTE on ` + "`" + `max_keys` + "`" + `:** Retrieving very large numbers of keys can adversely affect Terraform's performance.

The bucket-objects data source returns keys (i.e., file names) and other metadata about objects in an S3 bucket.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Lists object keys in this S3 bucket`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Limits results to object keys with this prefix (Default: none)`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) A character used to group keys (Default: none)`,
				},
				resource.Attribute{
					Name:        "encoding_type",
					Description: `(Optional) Encodes keys using this method (Default: none; besides none, only "url" can be used)`,
				},
				resource.Attribute{
					Name:        "max_keys",
					Description: `(Optional) Maximum object keys to return (Default: 1000)`,
				},
				resource.Attribute{
					Name:        "start_after",
					Description: `(Optional) Returns key names lexicographically after a specific object key in your bucket (Default: none; S3 lists object keys in UTF-8 character encoding in lexicographical order)`,
				},
				resource.Attribute{
					Name:        "fetch_owner",
					Description: `(Optional) Boolean specifying whether to populate the owner list (Default: false) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `List of strings representing object keys`,
				},
				resource.Attribute{
					Name:        "common_prefixes",
					Description: `List of any keys between ` + "`" + `prefix` + "`" + ` and the next occurrence of ` + "`" + `delimiter` + "`" + ` (i.e., similar to subdirectories of the ` + "`" + `prefix` + "`" + ` "directory"); the list is only returned when you specify ` + "`" + `delimiter` + "`" + ``,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `List of strings representing object owner IDs (see ` + "`" + `fetch_owner` + "`" + ` above)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keys",
					Description: `List of strings representing object keys`,
				},
				resource.Attribute{
					Name:        "common_prefixes",
					Description: `List of any keys between ` + "`" + `prefix` + "`" + ` and the next occurrence of ` + "`" + `delimiter` + "`" + ` (i.e., similar to subdirectories of the ` + "`" + `prefix` + "`" + ` "directory"); the list is only returned when you specify ` + "`" + `delimiter` + "`" + ``,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `List of strings representing object owner IDs (see ` + "`" + `fetch_owner` + "`" + ` above)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_secretsmanager_secret",
			Category:         "Data Sources",
			ShortDescription: `Retrieve metadata information about a Secrets Manager secret`,
			Description: `

Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the [` + "`" + `aws_secretsmanager_secret_version` + "`" + ` data source](/docs/providers/aws/d/secretsmanager_secret_version.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `(Optional) The Amazon Resource Name (ARN) of the secret to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the secret to retrieve. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the secret.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the secret.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The Key Management Service (KMS) Customer Master Key (CMK) associated with the secret.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Amazon Resource Name (ARN) of the secret.`,
				},
				resource.Attribute{
					Name:        "rotation_enabled",
					Description: `Whether rotation is enabled or not.`,
				},
				resource.Attribute{
					Name:        "rotation_lambda_arn",
					Description: `Rotation Lambda function Amazon Resource Name (ARN) if rotation is enabled.`,
				},
				resource.Attribute{
					Name:        "rotation_rules",
					Description: `Rotation rules if rotation is enabled.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the secret.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The resource-based policy document that's attached to the secret.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the secret.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the secret.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The Key Management Service (KMS) Customer Master Key (CMK) associated with the secret.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Amazon Resource Name (ARN) of the secret.`,
				},
				resource.Attribute{
					Name:        "rotation_enabled",
					Description: `Whether rotation is enabled or not.`,
				},
				resource.Attribute{
					Name:        "rotation_lambda_arn",
					Description: `Rotation Lambda function Amazon Resource Name (ARN) if rotation is enabled.`,
				},
				resource.Attribute{
					Name:        "rotation_rules",
					Description: `Rotation rules if rotation is enabled.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the secret.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The resource-based policy document that's attached to the secret.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_secretsmanager_secret_version",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about a Secrets Manager secret version including its secret value`,
			Description: `

Retrieve information about a Secrets Manager secret version, including its secret value. To retrieve secret metadata, see the [` + "`" + `aws_secretsmanager_secret` + "`" + ` data source](/docs/providers/aws/d/secretsmanager_secret.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_id",
					Description: `(Required) Specifies the secret containing the version that you want to retrieve. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) Specifies the unique identifier of the version of the secret that you want to retrieve. Overrides ` + "`" + `version_stage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version_stage",
					Description: `(Optional) Specifies the secret version that you want to retrieve by the staging label attached to the version. Defaults to ` + "`" + `AWSCURRENT` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the secret.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of this version of the secret.`,
				},
				resource.Attribute{
					Name:        "secret_string",
					Description: `The decrypted part of the protected secret information that was originally provided as a string.`,
				},
				resource.Attribute{
					Name:        "secret_binary",
					Description: `The decrypted part of the protected secret information that was originally provided as a binary. Base64 encoded.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The unique identifier of this version of the secret.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the secret.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of this version of the secret.`,
				},
				resource.Attribute{
					Name:        "secret_string",
					Description: `The decrypted part of the protected secret information that was originally provided as a string.`,
				},
				resource.Attribute{
					Name:        "secret_binary",
					Description: `The decrypted part of the protected secret information that was originally provided as a binary. Base64 encoded.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The unique identifier of this version of the secret.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_security_group",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Security Group`,
			Description: `

` + "`" + `aws_security_group` + "`" + ` provides details about a specific Security Group.

This resource can prove useful when a module accepts a Security Group id as
an input variable and needs to, for example, determine the id of the
VPC that the security group belongs to.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the specific security group to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name that the desired security group must have.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired security group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The id of the VPC that the desired security group belongs to. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroups.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. A Security Group will be selected if any one of the given values matches. ## Attributes Reference All of the argument attributes except ` + "`" + `filter` + "`" + ` blocks are also exported as result attributes. This data source will complete the data by populating any fields that are not included in the configuration with the data for the selected Security Group. The following fields are also exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The computed ARN of the security group. ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The computed ARN of the security group. ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_security_groups",
			Category:         "Data Sources",
			ShortDescription: `Get information about a set of Security Groups.`,
			Description: `

Use this data source to get IDs and VPC membership of Security Groups that are created
outside of Terraform.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match for desired security groups.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out [describe-security-groups in the AWS CLI reference][1]. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `IDs of the matches security groups.`,
				},
				resource.Attribute{
					Name:        "vpc_ids",
					Description: `The VPC IDs of the matched security groups. The data source's tag or filter`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `IDs of the matches security groups.`,
				},
				resource.Attribute{
					Name:        "vpc_ids",
					Description: `The VPC IDs of the matched security groups. The data source's tag or filter`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_servicequotas_service",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about a Service Quotas Service`,
			Description: `

Retrieve information about a Service Quotas Service.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) Service name to lookup within Service Quotas. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Code of the service.`,
				},
				resource.Attribute{
					Name:        "service_code",
					Description: `Code of the service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Code of the service.`,
				},
				resource.Attribute{
					Name:        "service_code",
					Description: `Code of the service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_servicequotas_service_quota",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about a Service Quota`,
			Description: `

Retrieve information about a Service Quota.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_code",
					Description: `(Required) Service code for the quota. Available values can be found with the [` + "`" + `aws_servicequotas_service` + "`" + ` data source](/docs/providers/aws/d/servicequotas_service.html) or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).`,
				},
				resource.Attribute{
					Name:        "quota_code",
					Description: `(Optional) Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).`,
				},
				resource.Attribute{
					Name:        "quota_name",
					Description: `(Optional) Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "adjustable",
					Description: `Whether the service quota is adjustable.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of the service quota.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `Default value of the service quota.`,
				},
				resource.Attribute{
					Name:        "global_quota",
					Description: `Whether the service quota is global for the AWS account.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Amazon Resource Name (ARN) of the service quota.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Name of the service.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Current value of the service quota.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "adjustable",
					Description: `Whether the service quota is adjustable.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of the service quota.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `Default value of the service quota.`,
				},
				resource.Attribute{
					Name:        "global_quota",
					Description: `Whether the service quota is global for the AWS account.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Amazon Resource Name (ARN) of the service quota.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Name of the service.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Current value of the service quota.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_sns_topic",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Amazon Simple Notification Service (SNS) Topic`,
			Description: `

Use this data source to get the ARN of a topic in AWS Simple Notification
Service (SNS). By using this data source, you can reference SNS topics
without having to hard code the ARNs as input.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The friendly name of the topic to match. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Set to the ARN of the found topic, suitable for referencing in other resources that support SNS topics.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Set to the ARN of the found topic, suitable for referencing in other resources that support SNS topics.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_sqs_queue",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Amazon Simple Queue Service (SQS) Queue`,
			Description: `

Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
By using this data source, you can reference SQS queues without having to hardcode
the ARNs as input.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the queue to match. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the queue.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the queue.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The Amazon Resource Name (ARN) of the queue.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the queue.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ssm_document",
			Category:         "Data Sources",
			ShortDescription: `Provides a SSM Document datasource`,
			Description: `

Gets the contents of the specified Systems Manager document.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Systems Manager document.`,
				},
				resource.Attribute{
					Name:        "document_format",
					Description: `(Optional) Returns the document in the specified format. The document format can be either JSON or YAML. JSON is the default format.`,
				},
				resource.Attribute{
					Name:        "document_version",
					Description: `(Optional) The document version for which you want information. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the document.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `The contents of the document.`,
				},
				resource.Attribute{
					Name:        "document_type",
					Description: `The type of the document.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the document.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `The contents of the document.`,
				},
				resource.Attribute{
					Name:        "document_type",
					Description: `The type of the document.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_ssm_parameter",
			Category:         "Data Sources",
			ShortDescription: `Provides a SSM Parameter datasource`,
			Description: `

Provides an SSM Parameter data source.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the parameter.`,
				},
				resource.Attribute{
					Name:        "with_decryption",
					Description: `(Optional) Whether to return decrypted ` + "`" + `SecureString` + "`" + ` value. Defaults to ` + "`" + `true` + "`" + `. In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the parameter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the parameter.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the parameter. Valid types are ` + "`" + `String` + "`" + `, ` + "`" + `StringList` + "`" + ` and ` + "`" + `SecureString` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the parameter.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the parameter.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_storagegateway_local_disk",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about a Storage Gateway local disk`,
			Description: `

Retrieve information about a Storage Gateway local disk. The disk identifier is useful for adding the disk as a cache or upload buffer to a gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_arn",
					Description: `(Required) The Amazon Resource Name (ARN) of the gateway.`,
				},
				resource.Attribute{
					Name:        "disk_node",
					Description: `(Optional) The device node of the local disk to retrieve. For example, ` + "`" + `/dev/sdb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_path",
					Description: `(Optional) The device path of the local disk to retrieve. For example, ` + "`" + `/dev/xvdb` + "`" + ` or ` + "`" + `/dev/nvme1n1` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `The disk identifier. e.g. ` + "`" + `pci-0000:03:00.0-scsi-0:0:0:0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The disk identifier. e.g. ` + "`" + `pci-0000:03:00.0-scsi-0:0:0:0` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disk_id",
					Description: `The disk identifier. e.g. ` + "`" + `pci-0000:03:00.0-scsi-0:0:0:0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The disk identifier. e.g. ` + "`" + `pci-0000:03:00.0-scsi-0:0:0:0` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_subnet",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC subnet`,
			Description: `

` + "`" + `aws_subnet` + "`" + ` provides details about a specific VPC subnet.

This resource can prove useful when a module accepts a subnet id as
an input variable and needs to, for example, determine the id of the
VPC that the subnet belongs to.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone where the subnet must reside.`,
				},
				resource.Attribute{
					Name:        "availability_zone_id",
					Description: `(Optional) The ID of the Availability Zone for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The cidr block of the desired subnet.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_block",
					Description: `(Optional) The Ipv6 cidr block of the desired subnet`,
				},
				resource.Attribute{
					Name:        "default_for_az",
					Description: `(Optional) Boolean constraint for whether the desired subnet must be the default subnet for its associated availability zone.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the specific subnet to retrieve.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The state that the desired subnet must have.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The id of the VPC that the desired subnet belongs to. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSubnets.html). For example, if matching against tag ` + "`" + `Name` + "`" + `, use: ` + "`" + `` + "`" + `` + "`" + `hcl data "aws_subnet" "selected" { filter { name = "tag:Name" values = [""] # insert value here } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. A subnet will be selected if any one of the given values matches. ## Attributes Reference All of the argument attributes except ` + "`" + `filter` + "`" + ` blocks are also exported as result attributes. This data source will complete the data by populating any fields that are not included in the configuration with the data for the selected subnet. In addition the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the subnet.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the subnet.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_subnet_ids",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of subnet Ids for a VPC`,
			Description: `

` + "`" + `aws_subnet_ids` + "`" + ` provides a list of ids for a vpc_id

This resource can be useful for getting back a list of subnet ids for a vpc.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC ID that you want to filter from.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired subnets. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSubnets.html). For example, if matching against tag ` + "`" + `Name` + "`" + `, use: ` + "`" + `` + "`" + `` + "`" + `hcl data "aws_subnet_ids" "selected" { filter { name = "tag:Name" values = [""] # insert values here } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. Subnet IDs will be selected if any one of the given values match. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A set of all the subnet ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A set of all the subnet ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_transfer_server",
			Category:         "Data Sources",
			ShortDescription: `Get information on an AWS Transfer Server resource`,
			Description: `

Use this data source to get the ARN of an AWS Transfer Server for use in other
resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) ID for an SFTP server. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of Transfer Server`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The endpoint of the Transfer Server (e.g. ` + "`" + `s-12345678.server.transfer.REGION.amazonaws.com` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "identity_provider_type",
					Description: `The mode of authentication enabled for this service. The default value is ` + "`" + `SERVICE_MANAGED` + "`" + `, which allows you to store and access SFTP user credentials within the service. ` + "`" + `API_GATEWAY` + "`" + ` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.`,
				},
				resource.Attribute{
					Name:        "invocation_role",
					Description: `Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an ` + "`" + `identity_provider_type` + "`" + ` of ` + "`" + `API_GATEWAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging_role",
					Description: `Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the service endpoint used to authenticate users with an ` + "`" + `identity_provider_type` + "`" + ` of ` + "`" + `API_GATEWAY` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of Transfer Server`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The endpoint of the Transfer Server (e.g. ` + "`" + `s-12345678.server.transfer.REGION.amazonaws.com` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "identity_provider_type",
					Description: `The mode of authentication enabled for this service. The default value is ` + "`" + `SERVICE_MANAGED` + "`" + `, which allows you to store and access SFTP user credentials within the service. ` + "`" + `API_GATEWAY` + "`" + ` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.`,
				},
				resource.Attribute{
					Name:        "invocation_role",
					Description: `Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an ` + "`" + `identity_provider_type` + "`" + ` of ` + "`" + `API_GATEWAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging_role",
					Description: `Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the service endpoint used to authenticate users with an ` + "`" + `identity_provider_type` + "`" + ` of ` + "`" + `API_GATEWAY` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_vpc",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC`,
			Description: `

` + "`" + `aws_vpc` + "`" + ` provides details about a specific VPC.

This resource can prove useful when a module accepts a vpc id as
an input variable and needs to, for example, determine the CIDR block of that
VPC.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The cidr block of the desired VPC.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_id",
					Description: `(Optional) The DHCP options id of the desired VPC.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) Boolean constraint on whether the desired VPC is the default VPC for the region.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The current state of the desired VPC. Can be either ` + "`" + `"pending"` + "`" + ` or ` + "`" + `"available"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired VPC. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. A VPC will be selected if any one of the given values matches. ## Attributes Reference All of the argument attributes except ` + "`" + `filter` + "`" + ` blocks are also exported as result attributes. This data source will complete the data by populating any fields that are not included in the configuration with the data for the selected VPC. The following attribute is additionally exported:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of VPC`,
				},
				resource.Attribute{
					Name:        "enable_dns_support",
					Description: `Whether or not the VPC has DNS support`,
				},
				resource.Attribute{
					Name:        "enable_dns_hostnames",
					Description: `Whether or not the VPC has DNS hostname support`,
				},
				resource.Attribute{
					Name:        "instance_tenancy",
					Description: `The allowed tenancy of instances launched into the selected VPC. May be any of ` + "`" + `"default"` + "`" + `, ` + "`" + `"dedicated"` + "`" + `, or ` + "`" + `"host"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_association_id",
					Description: `The association ID for the IPv6 CIDR block.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_block",
					Description: `The IPv6 CIDR block.`,
				},
				resource.Attribute{
					Name:        "main_route_table_id",
					Description: `The ID of the main route table associated with this VPC.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the VPC. ` + "`" + `cidr_block_associations` + "`" + ` is also exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "association_id",
					Description: `The association ID for the the IPv4 CIDR block.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block for the association.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The State of the association.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "arn",
					Description: `Amazon Resource Name (ARN) of VPC`,
				},
				resource.Attribute{
					Name:        "enable_dns_support",
					Description: `Whether or not the VPC has DNS support`,
				},
				resource.Attribute{
					Name:        "enable_dns_hostnames",
					Description: `Whether or not the VPC has DNS hostname support`,
				},
				resource.Attribute{
					Name:        "instance_tenancy",
					Description: `The allowed tenancy of instances launched into the selected VPC. May be any of ` + "`" + `"default"` + "`" + `, ` + "`" + `"dedicated"` + "`" + `, or ` + "`" + `"host"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_association_id",
					Description: `The association ID for the IPv6 CIDR block.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_block",
					Description: `The IPv6 CIDR block.`,
				},
				resource.Attribute{
					Name:        "main_route_table_id",
					Description: `The ID of the main route table associated with this VPC.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the VPC. ` + "`" + `cidr_block_associations` + "`" + ` is also exported with the following attributes:`,
				},
				resource.Attribute{
					Name:        "association_id",
					Description: `The association ID for the the IPv4 CIDR block.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block for the association.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The State of the association.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_vpc_dhcp_options",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about an EC2 DHCP Options configuration`,
			Description: `

Retrieve information about an EC2 DHCP Options configuration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_id",
					Description: `(Optional) The EC2 DHCP Options ID.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) List of custom filters as described below. ### filter For more information about filtering, see the [EC2 API documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeDhcpOptions.html).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values for filtering. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "dhcp_options_id",
					Description: `EC2 DHCP Options ID`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The suffix domain name to used when resolving non Fully Qualified Domain Names. e.g. the ` + "`" + `search` + "`" + ` value in the ` + "`" + `/etc/resolv.conf` + "`" + ` file.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `List of name servers.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 DHCP Options ID`,
				},
				resource.Attribute{
					Name:        "netbios_name_servers",
					Description: `List of NETBIOS name servers.`,
				},
				resource.Attribute{
					Name:        "netbios_node_type",
					Description: `The NetBIOS node type (1, 2, 4, or 8). For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `List of NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the DHCP options set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_id",
					Description: `EC2 DHCP Options ID`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The suffix domain name to used when resolving non Fully Qualified Domain Names. e.g. the ` + "`" + `search` + "`" + ` value in the ` + "`" + `/etc/resolv.conf` + "`" + ` file.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `List of name servers.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `EC2 DHCP Options ID`,
				},
				resource.Attribute{
					Name:        "netbios_name_servers",
					Description: `List of NETBIOS name servers.`,
				},
				resource.Attribute{
					Name:        "netbios_node_type",
					Description: `The NetBIOS node type (1, 2, 4, or 8). For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `List of NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the DHCP options set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_vpc_endpoint",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC endpoint.`,
			Description: `

The VPC Endpoint data source provides details about
a specific VPC endpoint.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific VPC Endpoint to retrieve.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The AWS service name of the specific VPC Endpoint to retrieve.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The state of the specific VPC Endpoint to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of the VPC in which the specific VPC Endpoint is used. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type ` + "`" + `Gateway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns_entry",
					Description: `The DNS entries for the VPC Endpoint. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `. DNS blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface_ids",
					Description: `One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The policy document associated with the VPC Endpoint. Applicable for endpoints of type ` + "`" + `Gateway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prefix_list_id",
					Description: `The prefix list ID of the exposed AWS service. Applicable for endpoints of type ` + "`" + `Gateway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_dns_enabled",
					Description: `Whether or not the VPC is associated with a private hosted zone - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "requester_managed",
					Description: `Whether or not the VPC Endpoint is being managed by its service - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `One or more route tables associated with the VPC Endpoint. Applicable for endpoints of type ` + "`" + `Gateway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `One or more security groups associated with the network interfaces. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `One or more subnets in which the VPC Endpoint is located. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "vpc_endpoint_type",
					Description: `The VPC Endpoint type, ` + "`" + `Gateway` + "`" + ` or ` + "`" + `Interface` + "`" + `. DNS blocks (for ` + "`" + `dns_entry` + "`" + `) support the following attributes:`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The ID of the private hosted zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type ` + "`" + `Gateway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns_entry",
					Description: `The DNS entries for the VPC Endpoint. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `. DNS blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface_ids",
					Description: `One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the AWS account that owns the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The policy document associated with the VPC Endpoint. Applicable for endpoints of type ` + "`" + `Gateway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prefix_list_id",
					Description: `The prefix list ID of the exposed AWS service. Applicable for endpoints of type ` + "`" + `Gateway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_dns_enabled",
					Description: `Whether or not the VPC is associated with a private hosted zone - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "requester_managed",
					Description: `Whether or not the VPC Endpoint is being managed by its service - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `One or more route tables associated with the VPC Endpoint. Applicable for endpoints of type ` + "`" + `Gateway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `One or more security groups associated with the network interfaces. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `One or more subnets in which the VPC Endpoint is located. Applicable for endpoints of type ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "vpc_endpoint_type",
					Description: `The VPC Endpoint type, ` + "`" + `Gateway` + "`" + ` or ` + "`" + `Interface` + "`" + `. DNS blocks (for ` + "`" + `dns_entry` + "`" + `) support the following attributes:`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The ID of the private hosted zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_vpc_endpoint_service",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific service that can be specified when creating a VPC endpoint.`,
			Description: `

The VPC Endpoint Service data source details about a specific service that
can be specified when creating a VPC endpoint within the region configured in the provider.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) The common name of an AWS service (e.g. ` + "`" + `s3` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The service name that can be specified when creating a VPC endpoint. ~>`,
				},
				resource.Attribute{
					Name:        "acceptance_required",
					Description: `Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The Availability Zones in which the service is available.`,
				},
				resource.Attribute{
					Name:        "base_endpoint_dns_names",
					Description: `The DNS names for the service.`,
				},
				resource.Attribute{
					Name:        "manages_vpc_endpoints",
					Description: `Whether or not the service manages its VPC endpoints - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The AWS account ID of the service owner or ` + "`" + `amazon` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The private DNS name for the service.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The service type, ` + "`" + `Gateway` + "`" + ` or ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "vpc_endpoint_policy_supported",
					Description: `Whether or not the service supports endpoint policies - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "acceptance_required",
					Description: `Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The Availability Zones in which the service is available.`,
				},
				resource.Attribute{
					Name:        "base_endpoint_dns_names",
					Description: `The DNS names for the service.`,
				},
				resource.Attribute{
					Name:        "manages_vpc_endpoints",
					Description: `Whether or not the service manages its VPC endpoints - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The AWS account ID of the service owner or ` + "`" + `amazon` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The private DNS name for the service.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The service type, ` + "`" + `Gateway` + "`" + ` or ` + "`" + `Interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "vpc_endpoint_policy_supported",
					Description: `Whether or not the service supports endpoint policies - ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_vpc_peering_connection",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC peering connection.`,
			Description: `

The VPC Peering Connection data source provides details about
a specific VPC peering connection.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of the requester VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `(Optional) The AWS account ID of the owner of the requester VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The CIDR block of the requester VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region of the requester VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `(Optional) The ID of the accepter VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "peer_owner_id",
					Description: `(Optional) The AWS account ID of the owner of the accepter VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "peer_cidr_block",
					Description: `(Optional) The CIDR block of the accepter VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "peer_region",
					Description: `(Optional) The region of the accepter VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired VPC Peering Connection. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcPeeringConnections.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. A VPC Peering Connection will be selected if any one of the given values matches. ## Attributes Reference All of the argument attributes except ` + "`" + `filter` + "`" + ` are also exported as result attributes.`,
				},
				resource.Attribute{
					Name:        "accepter",
					Description: `A configuration block that describes [VPC Peering Connection] (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.`,
				},
				resource.Attribute{
					Name:        "requester",
					Description: `A configuration block that describes [VPC Peering Connection] (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC. #### Accepter and Requester Attributes Reference`,
				},
				resource.Attribute{
					Name:        "allow_remote_vpc_dns_resolution",
					Description: `Indicates whether a local VPC can resolve public DNS hostnames to private IP addresses when queried from instances in a peer VPC.`,
				},
				resource.Attribute{
					Name:        "allow_classic_link_to_remote_vpc",
					Description: `Indicates whether a local ClassicLink connection can communicate with the peer VPC over the VPC peering connection.`,
				},
				resource.Attribute{
					Name:        "allow_vpc_to_remote_classic_link",
					Description: `Indicates whether a local VPC can communicate with a ClassicLink connection in the peer VPC over the VPC peering connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accepter",
					Description: `A configuration block that describes [VPC Peering Connection] (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.`,
				},
				resource.Attribute{
					Name:        "requester",
					Description: `A configuration block that describes [VPC Peering Connection] (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC. #### Accepter and Requester Attributes Reference`,
				},
				resource.Attribute{
					Name:        "allow_remote_vpc_dns_resolution",
					Description: `Indicates whether a local VPC can resolve public DNS hostnames to private IP addresses when queried from instances in a peer VPC.`,
				},
				resource.Attribute{
					Name:        "allow_classic_link_to_remote_vpc",
					Description: `Indicates whether a local ClassicLink connection can communicate with the peer VPC over the VPC peering connection.`,
				},
				resource.Attribute{
					Name:        "allow_vpc_to_remote_classic_link",
					Description: `Indicates whether a local VPC can communicate with a ClassicLink connection in the peer VPC over the VPC peering connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_vpcs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPC Ids in a region`,
			Description: `

This resource can be useful for getting back a list of VPC Ids for a region.

The following example retrieves a list of VPC Ids with a custom tag of ` + "`" + `service` + "`" + ` set to a value of "production".

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired vpcs.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. A VPC will be selected if any one of the given values matches. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the VPC Ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the VPC Ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_vpn_gateway",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPN gateway.`,
			Description: `

The VPN Gateway data source provides details about
a specific VPN gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific VPN Gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The state of the specific VPN Gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The Availability Zone of the specific VPN Gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "attached_vpc_id",
					Description: `(Optional) The ID of a VPC attached to the specific VPN Gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags, each pair of which must exactly match a pair on the desired VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "amazon_side_asn",
					Description: `(Optional) The Autonomous System Number (ASN) for the Amazon side of the specific VPN Gateway to retrieve. More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by, as defined by [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpnGateways.html).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. A VPN Gateway will be selected if any one of the given values matches. ## Attributes Reference All of the argument attributes are also exported as result attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_waf_ipset",
			Category:         "Data Sources",
			ShortDescription: `Retrieves an AWS WAF IP set id.`,
			Description: `

` + "`" + `aws_waf_ipset` + "`" + ` Retrieves a WAF IP Set Resource Id.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the WAF IP set. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF IP set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF IP set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_waf_rule",
			Category:         "Data Sources",
			ShortDescription: `Retrieves an AWS WAF rule id.`,
			Description: `

` + "`" + `aws_waf_rule` + "`" + ` Retrieves a WAF Rule Resource Id.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the WAF rule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_waf_web_acl",
			Category:         "Data Sources",
			ShortDescription: `Retrieves a WAF Web ACL id.`,
			Description: `

` + "`" + `aws_waf_rule` + "`" + ` Retrieves a WAF Web ACL Resource Id.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the WAF Web ACL. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF WebACL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF WebACL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_wafregional_ipset",
			Category:         "Data Sources",
			ShortDescription: `Retrieves an AWS WAF Regional IP set id.`,
			Description: `

` + "`" + `aws_wafregional_ipset` + "`" + ` Retrieves a WAF Regional IP Set Resource Id.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the WAF Regional IP set. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF Regional IP set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF Regional IP set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_wafregional_rule",
			Category:         "Data Sources",
			ShortDescription: `Retrieves an AWS WAF Regional rule id.`,
			Description: `

` + "`" + `aws_wafregional_rule` + "`" + ` Retrieves a WAF Regional Rule Resource Id.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the WAF rule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF Regional rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF Regional rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_wafregional_web_acl",
			Category:         "Data Sources",
			ShortDescription: `Retrieves a WAF Regional Web ACL id.`,
			Description: `

` + "`" + `aws_wafregional_web_acl` + "`" + ` Retrieves a WAF Regional Web ACL Resource Id.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the WAF Web ACL. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF Regional WebACL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF Regional WebACL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aws_workspaces_bundle",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Workspaces Bundle.`,
			Description: `

Use this data source to get information about a Workspaces Bundle.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the compute type. ### ` + "`" + `root_storage` + "`" + ``,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The size of the root volume. ### ` + "`" + `user_storage` + "`" + ``,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The size of the user storage.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the compute type. ### ` + "`" + `root_storage` + "`" + ``,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The size of the root volume. ### ` + "`" + `user_storage` + "`" + ``,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The size of the user storage.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"aws_acm_certificate":                           0,
		"aws_acmpca_certificate_authority":              1,
		"aws_ami":                                       2,
		"aws_ami_ids":                                   3,
		"aws_api_gateway_api_key":                       4,
		"aws_api_gateway_resource":                      5,
		"aws_api_gateway_rest_api":                      6,
		"aws_api_gateway_vpc_link":                      7,
		"aws_arn":                                       8,
		"aws_autoscaling_group":                         9,
		"aws_autoscaling_groups":                        10,
		"aws_availability_zone":                         11,
		"aws_availability_zones":                        12,
		"aws_batch_compute_environment":                 13,
		"aws_batch_job_queue":                           14,
		"aws_billing_service_account":                   15,
		"aws_caller_identity":                           16,
		"aws_canonical_user_id":                         17,
		"aws_cloudformation_export":                     18,
		"aws_cloudformation_stack":                      19,
		"aws_cloudhsm_v2_cluster":                       20,
		"aws_cloudtrail_service_account":                21,
		"aws_cloudwatch_log_group":                      22,
		"aws_codecommit_repository":                     23,
		"aws_cognito_user_pools":                        24,
		"aws_cur_report_definition":                     25,
		"aws_customer_gateway":                          26,
		"aws_db_cluster_snapshot":                       27,
		"aws_db_event_categories":                       28,
		"aws_db_instance":                               29,
		"aws_db_snapshot":                               30,
		"aws_dx_gateway":                                31,
		"aws_dynamodb_table":                            32,
		"aws_ebs_default_kms_key":                       33,
		"aws_ebs_encryption_by_default":                 34,
		"aws_ebs_snapshot":                              35,
		"aws_ebs_snapshot_ids":                          36,
		"aws_ebs_volume":                                37,
		"aws_ec2_transit_gateway":                       38,
		"aws_ec2_transit_gateway_dx_gateway_attachment": 39,
		"aws_ec2_transit_gateway_route_table":           40,
		"aws_ec2_transit_gateway_vpc_attachment":        41,
		"aws_ec2_transit_gateway_vpn_attachment":        42,
		"aws_ecr_image":                                 43,
		"aws_ecr_repository":                            44,
		"aws_ecs_cluster":                               45,
		"aws_ecs_container_definition":                  46,
		"aws_ecs_service":                               47,
		"aws_ecs_task_definition":                       48,
		"aws_efs_file_system":                           49,
		"aws_efs_mount_target":                          50,
		"aws_eip":                                       51,
		"aws_eks_cluster":                               52,
		"aws_eks_cluster_auth":                          53,
		"aws_elastic_beanstalk_application":             54,
		"aws_elastic_beanstalk_hosted_zone":             55,
		"aws_elastic_beanstalk_solution_stack":          56,
		"aws_elasticache_cluster":                       57,
		"aws_elasticache_replication_group":             58,
		"aws_elasticsearch_domain":                      59,
		"aws_elb":                                       60,
		"aws_elb_hosted_zone_id":                        61,
		"aws_elb_service_account":                       62,
		"aws_glue_script":                               63,
		"aws_iam_account_alias":                         64,
		"aws_iam_group":                                 65,
		"aws_iam_instance_profile":                      66,
		"aws_iam_policy":                                67,
		"aws_iam_policy_document":                       68,
		"aws_iam_role":                                  69,
		"aws_iam_server_certificate":                    70,
		"aws_iam_user":                                  71,
		"aws_inspector_rules_packages":                  72,
		"aws_instance":                                  73,
		"aws_instances":                                 74,
		"aws_internet_gateway":                          75,
		"aws_iot_endpoint":                              76,
		"aws_ip_ranges":                                 77,
		"aws_kinesis_stream":                            78,
		"aws_kms_alias":                                 79,
		"aws_kms_ciphertext":                            80,
		"aws_kms_key":                                   81,
		"aws_kms_secrets":                               82,
		"aws_lambda_function":                           83,
		"aws_lambda_invocation":                         84,
		"aws_lambda_layer_version":                      85,
		"aws_launch_configuration":                      86,
		"aws_launch_template":                           87,
		"aws_lb":                                        88,
		"aws_lb_listener":                               89,
		"aws_lb_target_group":                           90,
		"aws_mq_broker":                                 91,
		"aws_msk_cluster":                               92,
		"aws_msk_configuration":                         93,
		"aws_nat_gateway":                               94,
		"aws_network_acls":                              95,
		"aws_network_interface":                         96,
		"aws_network_interfaces":                        97,
		"aws_organizations_organization":                98,
		"aws_partition":                                 99,
		"aws_prefix_list":                               100,
		"aws_pricing_product":                           101,
		"aws_ram_resource_share":                        102,
		"aws_rds_cluster":                               103,
		"aws_redshift_cluster":                          104,
		"aws_redshift_service_account":                  105,
		"aws_region":                                    106,
		"aws_route":                                     107,
		"aws_route53_delegation_set":                    108,
		"aws_route53_resolver_rule":                     109,
		"aws_route53_resolver_rules":                    110,
		"aws_route53_zone":                              111,
		"aws_route_table":                               112,
		"aws_route_tables":                              113,
		"aws_s3_bucket":                                 114,
		"aws_s3_bucket_object":                          115,
		"aws_s3_bucket_objects":                         116,
		"aws_secretsmanager_secret":                     117,
		"aws_secretsmanager_secret_version":             118,
		"aws_security_group":                            119,
		"aws_security_groups":                           120,
		"aws_servicequotas_service":                     121,
		"aws_servicequotas_service_quota":               122,
		"aws_sns_topic":                                 123,
		"aws_sqs_queue":                                 124,
		"aws_ssm_document":                              125,
		"aws_ssm_parameter":                             126,
		"aws_storagegateway_local_disk":                 127,
		"aws_subnet":                                    128,
		"aws_subnet_ids":                                129,
		"aws_transfer_server":                           130,
		"aws_vpc":                                       131,
		"aws_vpc_dhcp_options":                          132,
		"aws_vpc_endpoint":                              133,
		"aws_vpc_endpoint_service":                      134,
		"aws_vpc_peering_connection":                    135,
		"aws_vpcs":                                      136,
		"aws_vpn_gateway":                               137,
		"aws_waf_ipset":                                 138,
		"aws_waf_rule":                                  139,
		"aws_waf_web_acl":                               140,
		"aws_wafregional_ipset":                         141,
		"aws_wafregional_rule":                          142,
		"aws_wafregional_web_acl":                       143,
		"aws_workspaces_bundle":                         144,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
