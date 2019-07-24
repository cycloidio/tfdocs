package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vault_approle_auth_backend_login",
			Category:         "Resources",
			ShortDescription: `Log into Vault using the AppRole auth backend.`,
			Description:      ``,
			Keywords: []string{
				"approle",
				"auth",
				"backend",
				"login",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The ID of the role to log in with.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `(Optional) The secret ID of the role to log in with. Required unless ` + "`" + `bind_secret_id` + "`" + ` is set to false on the role.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `The unique path of the Vault backend to log in with. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies applied to the token.`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `Whether the token is renewable or not.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `How long the token is valid for, in seconds.`,
				},
				resource.Attribute{
					Name:        "lease_started",
					Description: `The date and time the lease started, in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for the token.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The Vault token created.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the token.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies applied to the token.`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `Whether the token is renewable or not.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `How long the token is valid for, in seconds.`,
				},
				resource.Attribute{
					Name:        "lease_started",
					Description: `The date and time the lease started, in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for the token.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The Vault token created.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the token.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_approle_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages AppRole auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"approle",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Optional) The RoleID of this role. If not specified, one will be auto-generated.`,
				},
				resource.Attribute{
					Name:        "bind_secret_id",
					Description: `(Optional) Whether or not to require ` + "`" + `secret_id` + "`" + ` to be presented when logging in using this AppRole. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bound_cidr_list",
					Description: `(Optional) If set, specifies blocks of IP addresses which can perform the login operation.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) An array of strings specifying the policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "secret_id_num_uses",
					Description: `(Optional) The number of times any particular SecretID can be used to fetch a token from this AppRole, after which the SecretID will expire. A value of zero will allow unlimited uses.`,
				},
				resource.Attribute{
					Name:        "secret_id_ttl",
					Description: `(Optional) The number of seconds after which any SecretID expires.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The number of times issued tokens can be used. A value of 0 means unlimited uses.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The TTL period of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum allowed lifetime of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. The maximum allowed lifetime of token issued using this role. Specified as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name of the auth backend to configure. Defaults to ` + "`" + `approle` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import AppRole authentication backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_approle_auth_backend_role.example auth/approle/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_approle_auth_backend_role_secret_id",
			Category:         "Resources",
			ShortDescription: `Manages AppRole auth backend role SecretIDs in Vault.`,
			Description:      ``,
			Keywords: []string{
				"approle",
				"auth",
				"backend",
				"role",
				"secret",
				"id",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role to create the SecretID for.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A JSON-encoded string containing metadata in key-value pairs to be set on tokens issued with this SecretID.`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Optional) If set, specifies blocks of IP addresses which can perform the login operation using this SecretID.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `(Optional) The SecretID to be created. If set, uses "Push" mode. Defaults to Vault auto-generating SecretIDs. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The unique ID for this SecretID that can be safely logged.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "accessor",
					Description: `The unique ID for this SecretID that can be safely logged.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_audit",
			Category:         "Resources",
			ShortDescription: `Writes audit backends for Vault`,
			Description:      ``,
			Keywords: []string{
				"audit",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the audit device, such as 'file'.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(optional) The path to mount the audit device. This defaults to the type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description of the audit device.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Required) Configuration options to pass to the audit device itself. For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html) ## Attributes Reference No additional attributes are exported by this resource. ## Import Audit devices can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_audit.test syslog ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_auth_backend",
			Category:         "Resources",
			ShortDescription: `Writes auth methods for Vault`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The name of the auth method type`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to mount the auth method — this defaults to the name of the type`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the auth method`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `(Optional) Speficies whether to show this mount in the UI-specific listing endpoint.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Specifies if the auth method is local only. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth method ## Import Auth methods can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth method ## Import Auth methods can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_cert",
			Category:         "Resources",
			ShortDescription: `Manages a certificate for an AWS Auth Backend in Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"cert",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cert_name",
					Description: `(Required) The name of the certificate.`,
				},
				resource.Attribute{
					Name:        "aws_public_cert",
					Description: `(Required) The Base64 encoded AWS Public key required to verify PKCS7 signature of the EC2 instance metadata. You can find this key in the [AWS documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Either "pkcs7" or "identity", indicating the type of document which can be verified using the given certificate. Defaults to "pkcs7".`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path the AWS auth backend being configured was mounted at. Defaults to ` + "`" + `aws` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend certificates can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/config/certificate/` + "`" + `, and the ` + "`" + `cert_name` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_cert.example auth/aws/config/certificate/my-cert ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_client",
			Category:         "Resources",
			ShortDescription: `Configures the client used by an AWS Auth Backend in Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"client",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path the AWS auth backend being configured was mounted at. Defaults to ` + "`" + `aws` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) The AWS access key that Vault should use for the auth backend.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The AWS secret key that Vault should use for the auth backend.`,
				},
				resource.Attribute{
					Name:        "ec2_endpoint",
					Description: `(Optional) Override the URL Vault uses when making EC2 API calls.`,
				},
				resource.Attribute{
					Name:        "iam_endpoint",
					Description: `(Optional) Override the URL Vault uses when making IAM API calls.`,
				},
				resource.Attribute{
					Name:        "sts_endpoint",
					Description: `(Optional) Override the URL Vault uses when making STS API calls.`,
				},
				resource.Attribute{
					Name:        "iam_server_id_header_value",
					Description: `(Optional) The value to require in the ` + "`" + `X-Vault-AWS-IAM-Server-ID` + "`" + ` header as part of ` + "`" + `GetCallerIdentity` + "`" + ` requests that are used in the IAM auth method. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend clients can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, and ` + "`" + `/config/client` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_client.example auth/aws/config/client ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_identity_whitelist",
			Category:         "Resources",
			ShortDescription: `Configures the periodic tidying operation of the whitelisted identity entries.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"identity",
				"whitelist",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path of the AWS backend being configured.`,
				},
				resource.Attribute{
					Name:        "safety_buffer",
					Description: `(Optional) The amount of extra time, in minutes, that must have passed beyond the roletag expiration, before it is removed from the backend storage.`,
				},
				resource.Attribute{
					Name:        "disable_periodic_tidy",
					Description: `(Optional) If set to true, disables the periodic tidying of the identity-whitelist entries. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend identity whitelists can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, and ` + "`" + `/config/tidy/identity-whitelist` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_identity_whitelist.example auth/aws/config/tidy/identity-whitelist ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_login",
			Category:         "Resources",
			ShortDescription: `Manages Vault tokens acquired using the AWS auth backend.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"login",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name of the AWS auth backend. Defaults to 'aws'.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The name of the AWS auth backend role to create tokens against.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `(Optional) The base64-encoded EC2 instance identity document to authenticate with. Can be retrieved from the EC2 metadata server.`,
				},
				resource.Attribute{
					Name:        "signature",
					Description: `(Optional) The base64-encoded SHA256 RSA signature of the instance identity document to authenticate with, with all newline characters removed. Can be retrieved from the EC2 metadata server.`,
				},
				resource.Attribute{
					Name:        "pkcs7",
					Description: `(Optional) The PKCS#7 signature of the identity document to authenticate with, with all newline characters removed. Can be retrieved from the EC2 metadata server.`,
				},
				resource.Attribute{
					Name:        "nonce",
					Description: `(Optional) The unique nonce to be used for login requests. Can be set to a user-specified value, or will contain the server-generated value once a token is issued. EC2 instances can only acquire a single token until the whitelist is tidied again unless they keep track of this nonce.`,
				},
				resource.Attribute{
					Name:        "iam_http_request_method",
					Description: `(Optional) The HTTP method used in the signed IAM request.`,
				},
				resource.Attribute{
					Name:        "iam_request_url",
					Description: `(Optional) The base64-encoded HTTP URL used in the signed request.`,
				},
				resource.Attribute{
					Name:        "iam_request_body",
					Description: `(Optional) The base64-encoded body of the signed request.`,
				},
				resource.Attribute{
					Name:        "iam_request_headers",
					Description: `(Optional) The base64-encoded, JSON serialized representation of the GetCallerIdentity HTTP request headers. ## Attributes Reference In addition to the fields above, the following attributes are also exposed:`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration in seconds the token will be valid, relative to the time in ` + "`" + `lease_start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `The approximate time at which the token was created, using the clock of the system where Terraform was running.`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `Set to true if the token can be extended through renewal.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of information returned by the Vault server about the authentication used to generate this token.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `The authentication type used to generate this token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The Vault policies assigned to this token.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The token's accessor.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The token returned by Vault.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration in seconds the token will be valid, relative to the time in ` + "`" + `lease_start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `The approximate time at which the token was created, using the clock of the system where Terraform was running.`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `Set to true if the token can be extended through renewal.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of information returned by the Vault server about the authentication used to generate this token.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `The authentication type used to generate this token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The Vault policies assigned to this token.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The token's accessor.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The token returned by Vault.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages AWS auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) The auth type permitted for this role. Valid choices are ` + "`" + `ec2` + "`" + ` and ` + "`" + `iam` + "`" + `. Defaults to ` + "`" + `iam` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bound_ami_ids",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they should be using the AMI ID specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_account_ids",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they should be using the account ID specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_regions",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that the region in their identity document must match the one specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_vpc_ids",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they be associated with the VPC ID that matches the value specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_subnet_ids",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they be associated with the subnet ID that matches the value specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_iam_role_arns",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they must match the IAM role ARN specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_iam_instance_profile_arns",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they must be associated with an IAM instance profile ARN which has a prefix that matches the value specified by this field. The value is prefix-matched as though it were a glob ending in ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "role_tag",
					Description: `(Optional) If set, enable role tags for this role. The value set for this field should be the key of the tag on the EC2 instance. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_iam_principal_arns",
					Description: `(Optional) If set, defines the IAM principal that must be authenticated when ` + "`" + `auth_type` + "`" + ` is set to ` + "`" + `iam` + "`" + `. Wildcards are supported at the end of the ARN.`,
				},
				resource.Attribute{
					Name:        "inferred_entity_type",
					Description: `(Optional) If set, instructs Vault to turn on inferencing. The only valid value is ` + "`" + `ec2_instance` + "`" + `, which instructs Vault to infer that the role comes from an EC2 instance in an IAM instance profile. This only applies when ` + "`" + `auth_type` + "`" + ` is set to ` + "`" + `iam` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "inferred_aws_region",
					Description: `(Optional) When ` + "`" + `inferred_entity_type` + "`" + ` is set, this is the region to search for the inferred entities. Required if ` + "`" + `inferred_entity_type` + "`" + ` is set. This only applies when ` + "`" + `auth_type` + "`" + ` is set to ` + "`" + `iam` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resolve_aws_unique_ids",
					Description: `(Optional, Forces new resource) If set to ` + "`" + `true` + "`" + `, the ` + "`" + `bound_iam_principal_arns` + "`" + ` are resolved to [AWS Unique IDs](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-unique-ids) for the bound principal ARN. This field is ignored when a ` + "`" + `bound_iam_principal_arn` + "`" + ` ends in a wildcard. Resolving to unique IDs more closely mimics the behavior of AWS services in that if an IAM user or role is deleted and a new one is recreated with the same name, those new users or roles won't get access to roles in Vault that were permissioned to the prior principals of the same name. Defaults to ` + "`" + `true` + "`" + `. Once set to ` + "`" + `true` + "`" + `, this cannot be changed to ` + "`" + `false` + "`" + ` without recreating the role.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL period of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum allowed lifetime of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. The maximum allowed lifetime of token issued using this role. Specified as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) An array of strings specifying the policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "allow_instance_migration",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, allows migration of the underlying instance where the client resides.`,
				},
				resource.Attribute{
					Name:        "disallow_reauthentication",
					Description: `(Optional) IF set to ` + "`" + `true` + "`" + `, only allows a single token to be granted per instance ID. This can only be set when ` + "`" + `auth_type` + "`" + ` is set to ` + "`" + `ec2` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend roles can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/role/` + "`" + `, and the ` + "`" + `role` + "`" + ` name e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_role.example auth/aws/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_role_tag",
			Category:         "Resources",
			ShortDescription: `Reads role tags from a Vault AWS auth backend.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"role",
				"tag",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the AWS auth backend role to read role tags from, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path to the AWS auth backend to read role tags from, with no leading or trailing ` + "`" + `/` + "`" + `s. Defaults to "aws".`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) The policies to be associated with the tag. Must be a subset of the policies associated with the role.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum TTL of the tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.`,
				},
				resource.Attribute{
					Name:        "allow_instance_migration",
					Description: `(Optional) If set, allows migration of the underlying instances where the client resides. Use with caution.`,
				},
				resource.Attribute{
					Name:        "disallow_reauthentication",
					Description: `(Optional) If set, only allows a single token to be granted per instance ID. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `The key of the role tag.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `The value to set the role key.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "tag_key",
					Description: `The key of the role tag.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `The value to set the role key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_roletag_blacklist",
			Category:         "Resources",
			ShortDescription: `Configures the periodic tidying operation of the blacklisted role tag entries.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"roletag",
				"blacklist",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the AWS auth backend being configured was mounted at.`,
				},
				resource.Attribute{
					Name:        "safety_buffer",
					Description: `(Oprtional) The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage. Defaults to 259,200 seconds, or 72 hours.`,
				},
				resource.Attribute{
					Name:        "disable_periodic_tidy",
					Description: `(Optional) If set to true, disables the periodic tidying of the roletag blacklist entries. Defaults to false. ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_sts_role",
			Category:         "Resources",
			ShortDescription: `Configures an STS role in the Vault AWS Auth backend.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"sts",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The AWS account ID to configure the STS role for.`,
				},
				resource.Attribute{
					Name:        "sts_role",
					Description: `(Optional) The STS role to assume when verifying requests made by EC2 instances in the account specified by ` + "`" + `account_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path the AWS auth backend being configured was mounted at. Defaults to ` + "`" + `aws` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend STS roles can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/config/sts/` + "`" + `, and the ` + "`" + `account_id` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_sts_role.example auth/aws/config/sts/1234567890 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an AWS secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"secret",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) The AWS Access Key ID this backend should use to issue new credentials.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) The AWS Secret Key this backend should use to issue new credentials. ~>`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The AWS region for API calls. Defaults to ` + "`" + `us-east-1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `aws` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default TTL for credentials issued by this backend.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS secret backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_secret_backend.aws aws ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Creates a role on an AWS Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the AWS secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to identify this role within the backend. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "policy_document",
					Description: `(Optional) The JSON-formatted policy to associate with this role. Either ` + "`" + `policy_document` + "`" + ` or ` + "`" + `policy_arns` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "policy_arns",
					Description: `(Optional) The ARN for a pre-existing policy to associate with this role. Either ` + "`" + `policy_document` + "`" + ` or ` + "`" + `policy_arns` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "role_arns",
					Description: `(Optional) Specifies the ARNs of the AWS roles this Vault role is allowed to assume. Required when ` + "`" + `credential_type` + "`" + ` is ` + "`" + `assumed_role` + "`" + ` and prohibited otherwise.`,
				},
				resource.Attribute{
					Name:        "credential_type",
					Description: `(Required) Specifies the type of credential to be used when retrieving credentials from the role. Must be one of ` + "`" + `iam_user` + "`" + `, ` + "`" + `assumed_role` + "`" + `, or ` + "`" + `federation_token` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_secret_backend_role.role aws/roles/deploy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_azure_auth_backend_config",
			Category:         "Resources",
			ShortDescription: `Configures the Azure Auth Backend in Vault.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"auth",
				"backend",
				"config",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The tenant id for the Azure Active Directory organization.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) The configured URL for the application registered in Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path the Azure auth backend being configured was mounted at. Defaults to ` + "`" + `azure` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are required.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) The client secret for credentials to query the Azure APIs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud. Defaults to ` + "`" + `AzurePublicCloud` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import Azure auth backends can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, and ` + "`" + `/config` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_azure_auth_backend_config.example auth/azure/config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_azure_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages Azure auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "bound_service_principal_ids",
					Description: `(Optional) If set, defines a constraint on the service principals that can perform the login operation that they should be posess the ids specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_group_ids",
					Description: `(Optional) If set, defines a constraint on the groups that can perform the login operation that they should be using the group ID specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_locations",
					Description: `(Optional) If set, defines a constraint on the virtual machines that can perform the login operation that the location in their identity document must match the one specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_subscription_ids",
					Description: `(Optional) If set, defines a constraint on the subscriptions that can perform the login operation to ones which matches the value specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_resource_groups",
					Description: `(Optional) If set, defines a constraint on the virtual machiness that can perform the login operation that they be associated with the resource group that matches the value specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_scale_sets",
					Description: `(Optional) If set, defines a constraint on the virtual machines that can perform the login operation that they must match the scale set specified by this field.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL period of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum allowed lifetime of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. The maximum allowed lifetime of token issued using this role. Specified as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) An array of strings specifying the policies to be set on tokens issued using this role. ## Attributes Reference No additional attributes are exported by this resource. ## Import Azure auth backend roles can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/role/` + "`" + `, and the ` + "`" + `role` + "`" + ` name e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_azure_auth_backend_role.example auth/azure/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_cert_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Managing roles in an Cert auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"cert",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the role`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) CA certificate used to validate client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_names",
					Description: `(Optional) Allowed subject names for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_common_names",
					Description: `(Optional) Allowed the common names for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_dns_sans",
					Description: `(Optional) Allowed alternative dns names for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_email_sans",
					Description: `(Optional) Allowed emails for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_uri_sans",
					Description: `(Optional) Allowed URIs for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_organization_units",
					Description: `(Optional) Allowed organization units for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "required_extensions",
					Description: `(Optional) TLS extensions required on client certificates`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Default TTL of tokens issued by the backend`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) Maximum TTL of tokens issued by the backend`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Duration in seconds for token. If set, the issued token is a periodic token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Policies to grant on the issued token`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name to display on tokens issued under this role.`,
				},
				resource.Attribute{
					Name:        "bound_cidrs",
					Description: `(Optional) Restriction usage of the certificates to client IPs falling within the range of the specified CIDRs`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Path to the mounted Cert auth backend For more details on the usage of each argument consult the [Vault Cert API documentation](https://www.vaultproject.io/api/auth/cert/index.html). ## Attribute Reference No additional attributes are exposed by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_consul_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates a Consul secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"consul",
				"secret",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The Consul management token this backend should use to issue new tokens. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique location this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `consul` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) Specifies the URL scheme to use. Defaults to ` + "`" + `http` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default TTL for credentials issued by this backend.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend. ## Attributes Reference No additional attributes are exported by this resource. ## Import Consul secret backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_consul_secret_backend.example consul ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_database_secret_backend_connection",
			Category:         "Resources",
			ShortDescription: `Configures a database secret backend connection for Vault.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"secret",
				"backend",
				"connection",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to give the database connection.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The unique name of the Vault mount to configure.`,
				},
				resource.Attribute{
					Name:        "verify_connection",
					Description: `(Optional) Whether the connection should be verified on initial configuration or not.`,
				},
				resource.Attribute{
					Name:        "allowed_roles",
					Description: `(Optional) A list of roles that are allowed to use this connection.`,
				},
				resource.Attribute{
					Name:        "cassandra",
					Description: `(Optional) A nested block containing configuration options for Cassandra connections.`,
				},
				resource.Attribute{
					Name:        "mongodb",
					Description: `(Optional) A nested block containing configuration options for MongoDB connections.`,
				},
				resource.Attribute{
					Name:        "hana",
					Description: `(Optional) A nested block containing configuration options for SAP HanaDB connections.`,
				},
				resource.Attribute{
					Name:        "mssql",
					Description: `(Optional) A nested block containing configuration options for MSSQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql",
					Description: `(Optional) A nested block containing configuration options for MySQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql_rds",
					Description: `(Optional) A nested block containing configuration options for RDS MySQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql_aurora",
					Description: `(Optional) A nested block containing configuration options for Aurora MySQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql_legacy",
					Description: `(Optional) A nested block containing configuration options for legacy MySQL connections.`,
				},
				resource.Attribute{
					Name:        "postgresql",
					Description: `(Optional) A nested block containing configuration options for PostgreSQL connections.`,
				},
				resource.Attribute{
					Name:        "oracle",
					Description: `(Optional) A nested block containing configuration options for Oracle connections. Exactly one of the nested blocks of configuration options must be supplied. ### Cassandra Configuration Options`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `(Required) The hosts to connect to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to authenticate with.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to authenticate with.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The default port to connect to if no port is specified as part of the host.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Whether to use TLS when connecting to Cassandra.`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Whether to skip verification of the server certificate when using TLS.`,
				},
				resource.Attribute{
					Name:        "pem_bundle",
					Description: `(Optional) Concatenated PEM blocks configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "pem_json",
					Description: `(Optional) A JSON structure configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "protocol_version",
					Description: `(Optional) The CQL protocol version to use.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) The number of seconds to use as a connection timeout. ### MongoDB Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload) for an example. ### SAP HanaDB Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api/secret/databases/hanadb.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for. ### MSSQL Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api/secret/databases/mssql.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for. ### MySQL Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api/secret/databases/mysql-maria.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for. ### PostgreSQL Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api/secret/databases/postgresql.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for. ### Oracle Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for. ## Attributes Reference No additional attributes are exported by this resource. ## Import Database secret backend connections can be imported using the ` + "`" + `backend` + "`" + `, ` + "`" + `/config/` + "`" + `, and the ` + "`" + `name` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_database_secret_backend_connection.example postgres/config/postgres ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_database_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Configures a database secret backend role for Vault.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to give the role.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The unique name of the Vault mount to configure.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `(Required) The unique name of the database connection to use for the role.`,
				},
				resource.Attribute{
					Name:        "creation_statements",
					Description: `(Required) The database statements to execute when creating a user.`,
				},
				resource.Attribute{
					Name:        "revocation_statements",
					Description: `(Optional) The database statements to execute when revoking a user.`,
				},
				resource.Attribute{
					Name:        "rollback_statements",
					Description: `(Optional) The database statements to execute when rolling back creation due to an error.`,
				},
				resource.Attribute{
					Name:        "renew_statements",
					Description: `(Optional) The database statements to execute when renewing a user.`,
				},
				resource.Attribute{
					Name:        "default_ttl",
					Description: `(Optional) The default number of seconds for leases for this role.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum number of seconds for leases for this role. ## Attributes Reference No additional attributes are exported by this resource. ## Import Database secret backend roles can be imported using the ` + "`" + `backend` + "`" + `, ` + "`" + `/roles/` + "`" + `, and the ` + "`" + `name` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_database_secret_backend_role.example postgres/roles/my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_egp_policy",
			Category:         "Resources",
			ShortDescription: `Writes Sentinel endpoint governing policies for Vault`,
			Description:      ``,
			Keywords: []string{
				"egp",
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy`,
				},
				resource.Attribute{
					Name:        "paths",
					Description: `(Required) List of paths to which the policy will be applied to`,
				},
				resource.Attribute{
					Name:        "enforcement_level",
					Description: `(Required) Enforcement level of Sentinel policy. Can be either ` + "`" + `advisory` + "`" + ` or ` + "`" + `soft-mandatory` + "`" + ` or ` + "`" + `hard-mandatory` + "`" + ``,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) String containing a Sentinel policy ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_gcp_auth_backend",
			Category:         "Resources",
			ShortDescription: `Managing roles in an GCP auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"gcp",
				"auth",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) A JSON string containing the contents of a GCP credentials file. For more details on the usage of each argument consult the [Vault GCP API documentation](https://www.vaultproject.io/api/auth/gcp/index.html#configure). ## Attribute Reference In addition to the fields above, the following attributes are also exposed:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the credentials`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `The ID of the private key from the credentials`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The GCP Project ID`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `The clients email assosiated with the credentials`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the credentials`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `The ID of the private key from the credentials`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The GCP Project ID`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `The clients email assosiated with the credentials`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_gcp_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Managing roles in an GCP auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"gcp",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Name of the GCP role`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of GCP authentication role (either ` + "`" + `gce` + "`" + ` or ` + "`" + `iam` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional; Deprecated, use ` + "`" + `bound_projects` + "`" + ` instead) GCP Project that the role exists within`,
				},
				resource.Attribute{
					Name:        "bound_projects",
					Description: `(Optional) An array of GCP project IDs. Only entities belonging to this project can authenticate under the role.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Default TTL of tokens issued by the backend`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) Maximum TTL of tokens issued by the backend`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Duration in seconds for token. If set, the issued token is a periodic token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Policies to grant on the issued token`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Path to the mounted GCP auth backend`,
				},
				resource.Attribute{
					Name:        "bound_service_accounts",
					Description: `(Optional) GCP Service Accounts allowed to issue tokens under this role. (Note:`,
				},
				resource.Attribute{
					Name:        "max_jwt_exp",
					Description: `(Optional) The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter ` + "`" + `exp` + "`" + `. The GCE metadata tokens currently do not allow the ` + "`" + `exp` + "`" + ` claim to be customized.`,
				},
				resource.Attribute{
					Name:        "allow_gce_inference",
					Description: `(Optional) A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token. ### ` + "`" + `gce` + "`" + `-only Parameters The following parameters are only valid when the role is of type ` + "`" + `"gce"` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "bound_zones",
					Description: `(Optional) The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.`,
				},
				resource.Attribute{
					Name:        "bound_regions",
					Description: `(Optional) The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.`,
				},
				resource.Attribute{
					Name:        "bound_instance_groups",
					Description: `(Optional) The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either ` + "`" + `bound_zones` + "`" + ` or ` + "`" + `bound_regions` + "`" + ` must be set too.`,
				},
				resource.Attribute{
					Name:        "bound_labels",
					Description: `(Optional) A comma-separated list of GCP labels formatted as ` + "`" + `"key:value"` + "`" + ` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.`,
				},
				resource.Attribute{
					Name:        "bound_projects",
					Description: `(Optional) GCP Projects that the role exists within For more details on the usage of each argument consult the [Vault GCP API documentation](https://www.vaultproject.io/api/auth/gcp/index.html). ## Attribute Reference No additional attributes are exposed by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_gcp_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an GCP secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"gcp",
				"secret",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "credentials",
					Description: `(Optional) The GCP service account credentails in JSON format. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `gcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default TTL for credentials issued by this backend. Defaults to '0'.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend. Defaults to '0'. ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_gcp_secret_roleset",
			Category:         "Resources",
			ShortDescription: `Creates a Roleset for the GCP Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"gcp",
				"secret",
				"roleset",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required, Forces new resource) Path where the GCP Secrets Engine is mounted`,
				},
				resource.Attribute{
					Name:        "roleset",
					Description: `(Required, Forces new resource) Name of the Roleset to create`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required, Forces new resource) Name of the GCP project that this roleset's service account will belong to.`,
				},
				resource.Attribute{
					Name:        "secret_type",
					Description: `(Optional, Forces new resource) Type of secret generated for this role set. Accepted values: ` + "`" + `access_token` + "`" + `, ` + "`" + `service_account_key` + "`" + `. Defaults to ` + "`" + `access_token` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token_scopes",
					Description: `(Optional, Required for ` + "`" + `secret_type = "access_token"` + "`" + `) List of OAuth scopes to assign to ` + "`" + `access_token` + "`" + ` secrets generated under this role set (` + "`" + `access_token` + "`" + ` role sets only).`,
				},
				resource.Attribute{
					Name:        "binding",
					Description: `(Required) Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below. The ` + "`" + `binding` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource. ## Attributes Reference In addition to the fields above, the following attributes are also exposed:`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_generic_endpoint",
			Category:         "Resources",
			ShortDescription: `Writes arbitrary data to a given path in Vault`,
			Description:      ``,
			Keywords: []string{
				"generic",
				"endpoint",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The full logical path at which to write the given data. Consult each backend's documentation to see which endpoints support the ` + "`" + `PUT` + "`" + ` methods and to determine whether they also support ` + "`" + `DELETE` + "`" + ` and ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `(Required) String containing a JSON-encoded object that will be written to the given path as the secret data.`,
				},
				resource.Attribute{
					Name:        "disable_read",
					Description: `(Optional) True/false. Set this to true if your vault authentication is not able to read the data or if the endpoint does not support the ` + "`" + `GET` + "`" + ` method. Setting this to ` + "`" + `true` + "`" + ` will break drift detection. You should set this to ` + "`" + `true` + "`" + ` for endpoints that are write-only. Defaults to false.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_generic_secret",
			Category:         "Resources",
			ShortDescription: `Writes arbitrary data to a given path in Vault`,
			Description:      ``,
			Keywords: []string{
				"generic",
				"secret",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The full logical path at which to write the given data. To write data into the "generic" secret backend mounted in Vault by default, this should be prefixed with ` + "`" + `secret/` + "`" + `. Writing to other backends with this resource is possible; consult each backend's documentation to see which endpoints support the ` + "`" + `PUT` + "`" + ` and ` + "`" + `DELETE` + "`" + ` methods.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `(Required) String containing a JSON-encoded object that will be written as the secret data at the given path.`,
				},
				resource.Attribute{
					Name:        "allow_read",
					Description: `(Optional, Deprecated) True/false. Set this to true if your vault authentication is able to read the data, this allows the resource to be compared and updated. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "disable_read",
					Description: `(Optional) True/false. Set this to true if your vault authentication is not able to read the data. Setting this to ` + "`" + `true` + "`" + ` will break drift detection. Defaults to false. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `create` + "`" + ` or ` + "`" + `update` + "`" + ` capability (depending on whether the resource already exists) on the given path, along with the ` + "`" + `delete` + "`" + ` capbility if the resource is removed from configuration. This resource does not`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON. ## Import Generic secrets can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_generic_secret.example secret/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON. ## Import Generic secrets can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_generic_secret.example secret/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_github_auth_backend",
			Category:         "Resources",
			ShortDescription: `Manages Github Auth mounts in Vault.`,
			Description:      ``,
			Keywords: []string{
				"github",
				"auth",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path where the auth backend is mounted. Defaults to ` + "`" + `auth/github` + "`" + ` if not specified.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) The organization configured users must be part of.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Optional) The API endpoint to use. Useful if you are running GitHub Enterprise or an API-compatible authentication server.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the description of the mount. This overrides the current stored value, if any.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Duration after which authentication will be expired. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) Maximum duration after which authentication will be expired. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration). The ` + "`" + `tune` + "`" + ` block is used to tune the auth backend:`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl",
					Description: `(Optional) Specifies the default time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl",
					Description: `(Optional) Specifies the maximum time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_response_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_request_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `(Optional) Specifies whether to show this mount in the UI-specific listing endpoint. Valid values are "unauth" or "hidden".`,
				},
				resource.Attribute{
					Name:        "passthrough_request_headers",
					Description: `(Optional) List of headers to whitelist and pass from the request to the backend. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html). ## Import Github authentication mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "accessor",
					Description: `The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html). ## Import Github authentication mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_github_team",
			Category:         "Resources",
			ShortDescription: `Manages Team mappings for Github Auth backend mounts in Vault.`,
			Description:      ``,
			Keywords: []string{
				"github",
				"team",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) Path where the github auth backend is mounted. Defaults to ` + "`" + `github` + "`" + ` if not specified.`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Required) GitHub team name in "slugified" format, for example: Terraform Developers -> ` + "`" + `terraform-developers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) A list of policies to be assigned to this team. ## Attributes Reference No additional attributes are exported by this resource. ## Import Github team mappings can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_team.tf_devs auth/github/map/teams/terraform-developers ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_github_user",
			Category:         "Resources",
			ShortDescription: `Manages User mappings for Github Auth backend mounts in Vault.`,
			Description:      ``,
			Keywords: []string{
				"github",
				"user",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) Path where the github auth backend is mounted. Defaults to ` + "`" + `github` + "`" + ` if not specified.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) GitHub user name.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) A list of policies to be assigned to this user. ## Attributes Reference No additional attributes are exported by this resource. ## Import Github user mappings can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_user.tf_user auth/github/map/users/john.doe ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_entity",
			Category:         "Resources",
			ShortDescription: `Creates an Identity Entity for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"entity",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the identity entity to create.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) A list of policies to apply to the entity.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A Map of additional metadata to associate with the user.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) True/false Is this entity currently disabled. Defaults to ` + "`" + `false` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created entity.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created entity.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_entity_alias",
			Category:         "Resources",
			ShortDescription: `Creates an Identity Entity Alias for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"entity",
				"alias",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `(Required) Accessor of the mount to which the alias should belong to.`,
				},
				resource.Attribute{
					Name:        "canonical_id",
					Description: `(Required) Entity ID to which this alias belongs to. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the entity alias.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the entity alias.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_group",
			Category:         "Resources",
			ShortDescription: `Creates an Identity Group for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the identity group to create.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, Forces new resource) Type of the group, internal or external. Defaults to ` + "`" + `internal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) A list of policies to apply to the group.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A Map of additional metadata to associate with the group.`,
				},
				resource.Attribute{
					Name:        "member_group_ids",
					Description: `(Optional) A list of Group IDs to be assigned as group members.`,
				},
				resource.Attribute{
					Name:        "member_entity_ids",
					Description: `(Optional) A list of Entity IDs to be assigned as group members. Not allowed on ` + "`" + `external` + "`" + ` groups. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created group.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_group_alias",
			Category:         "Resources",
			ShortDescription: `Creates an Identity Group Alias for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
				"alias",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the group alias to create.`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `(Required) Mount accessor of the authentication backend to which this alias belongs to.`,
				},
				resource.Attribute{
					Name:        "canonical_id",
					Description: `(Required) ID of the group to which this is an alias. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created group alias.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created group alias.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_jwt_auth_backend",
			Category:         "Resources",
			ShortDescription: `Managing JWT/OIDC auth backends in Vault`,
			Description:      ``,
			Keywords: []string{
				"jwt",
				"auth",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path to mount the JWT/OIDC auth backend`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of auth backend. Should be one of ` + "`" + `jwt` + "`" + ` or ` + "`" + `oidc` + "`" + `. Default - ` + "`" + `jwt` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the auth backend`,
				},
				resource.Attribute{
					Name:        "oidc_discovery_url",
					Description: `(Optional) The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with ` + "`" + `jwt_validation_pubkeys` + "`" + ``,
				},
				resource.Attribute{
					Name:        "oidc_client_id",
					Description: `(Optional) Client ID used for OIDC backends`,
				},
				resource.Attribute{
					Name:        "oidc_client_secret",
					Description: `(Optional) Client Secret used for OIDC backends`,
				},
				resource.Attribute{
					Name:        "bound_issuer",
					Description: `(Optional) The value against which to match the iss claim in a JWT`,
				},
				resource.Attribute{
					Name:        "oidc_discovery_ca_pem",
					Description: `(Optional) The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used`,
				},
				resource.Attribute{
					Name:        "jwt_validation_pubkeys",
					Description: `(Optional) A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with ` + "`" + `oidc_discovery_url` + "`" + ``,
				},
				resource.Attribute{
					Name:        "jwt_supported_algs",
					Description: `(Optional) A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ The ` + "`" + `tune` + "`" + ` block is used to tune the auth backend:`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl",
					Description: `(Optional) Specifies the default time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl",
					Description: `(Optional) Specifies the maximum time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_response_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_request_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `(Optional) Specifies whether to show this mount in the UI-specific listing endpoint. Valid values are "unauth" or "hidden".`,
				},
				resource.Attribute{
					Name:        "passthrough_request_headers",
					Description: `(Optional) List of headers to whitelist and pass from the request to the backend. ## Attributes Reference No additional attributes are exposed by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_jwt_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages JWT/OIDC auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"jwt",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `(Optional) Type of role, either "oidc" (default) or "jwt".`,
				},
				resource.Attribute{
					Name:        "bound_audiences",
					Description: `(Required) List of ` + "`" + `aud` + "`" + ` claims to match against. Any match is sufficient.`,
				},
				resource.Attribute{
					Name:        "user_claim",
					Description: `(Required) The claim to use to uniquely identify the user; this will be used as the name for the Identity entity alias created due to a successful login.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The initial/renewal TTL of tokens issued using this role, in seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum allowed lifetime of tokens issued using this role, in seconds.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire, but instead always use the value set here as the TTL for every renewal.`,
				},
				resource.Attribute{
					Name:        "num_uses",
					Description: `(Optional) If set, puts a use-count limitation on the issued token.`,
				},
				resource.Attribute{
					Name:        "bound_subject",
					Description: `(Optional) If set, requires that the ` + "`" + `sub` + "`" + ` claim matches this value.`,
				},
				resource.Attribute{
					Name:        "bound_cidrs",
					Description: `(Optional) If set, a list of CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.`,
				},
				resource.Attribute{
					Name:        "bound_claims",
					Description: `(Optional) If set, a map of claims/values to match against. The expected value may be a single string or a list of strings.`,
				},
				resource.Attribute{
					Name:        "claim_mappings",
					Description: `(Optional) If set, a map of claims (keys) to be copied to specified metadata fields (values).`,
				},
				resource.Attribute{
					Name:        "oidc_scopes",
					Description: `(Optional) If set, a list of OIDC scopes to be used with an OIDC role. The standard scope "openid" is automatically included and need not be specified.`,
				},
				resource.Attribute{
					Name:        "groups_claim",
					Description: `(Optional) The claim to use to uniquely identify the set of groups to which the user belongs; this will be used as the names for the Identity group aliases created due to a successful login. The claim value must be a list of strings.`,
				},
				resource.Attribute{
					Name:        "groups_claim_delimiter_pattern",
					Description: `(Optional; Deprecated. This field has been removed since Vault 1.1. If the groups claim is not at the top level, it can now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).) A pattern of delimiters used to allow the groups_claim to live outside of the top-level JWT structure. For instance, a groups_claim of meta/user.name/groups with this field set to // will expect nested structures named meta, user.name, and groups. If this field was set to /./ the groups information would expect to be via nested structures of meta, user, name, and groups.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name of the auth backend to configure. Defaults to ` + "`" + `jwt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allowed_redirect_uris",
					Description: `(Optional) The list of allowed values for redirect_uri during OIDC logins. Required for OIDC roles ## Attributes Reference No additional attributes are exported by this resource. ## Import JWT authentication backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_jwt_auth_backend_role.example auth/jwt/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_auth_backend_config",
			Category:         "Resources",
			ShortDescription: `Manages Kubernetes auth backend configs in Vault.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"auth",
				"backend",
				"config",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "kubernetes_host",
					Description: `(Required) Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubernetes_ca_cert",
					Description: `(Optional) PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.`,
				},
				resource.Attribute{
					Name:        "token_reviewer_jwt",
					Description: `(Optional) A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.`,
				},
				resource.Attribute{
					Name:        "pem_keys",
					Description: `(Optional) List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys. ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages Kubernetes auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) Name of the role.`,
				},
				resource.Attribute{
					Name:        "bound_service_account_names",
					Description: `(Required) List of service account names able to access this role. If set to ` + "`" + `["`,
				},
				resource.Attribute{
					Name:        "bound_service_account_namespaces",
					Description: `(Required) List of namespaces allowed to access this role. If set to ` + "`" + `["`,
				},
				resource.Attribute{
					Name:        "bound_cirs",
					Description: `(Optional) List of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL period of tokens issued using this role in seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum allowed lifetime of tokens issued in seconds using this role.`,
				},
				resource.Attribute{
					Name:        "num_uses",
					Description: `(Optional) Number of times issued tokens can be used. Setting this to 0 or leaving it unset means unlimited uses.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this parameter.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Unique name of the kubernetes backend to configure. ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ldap_auth_backend",
			Category:         "Resources",
			ShortDescription: `Managing LDAP auth backends in Vault`,
			Description:      ``,
			Keywords: []string{
				"ldap",
				"auth",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL of the LDAP server`,
				},
				resource.Attribute{
					Name:        "starttls",
					Description: `(Optional) Control use of TLS when conecting to LDAP`,
				},
				resource.Attribute{
					Name:        "tls_min_version",
					Description: `(Optional) Minimum acceptable version of TLS`,
				},
				resource.Attribute{
					Name:        "tls_max_version",
					Description: `(Optional) Maximum acceptable version of TLS`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Control whether or TLS certificates must be validated`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional) Trusted CA to validate TLS certificate`,
				},
				resource.Attribute{
					Name:        "binddn",
					Description: `(Optional) DN of object to bind when performing user search`,
				},
				resource.Attribute{
					Name:        "bindpass",
					Description: `(Optional) Password to use with ` + "`" + `binddn` + "`" + ` when performing user search`,
				},
				resource.Attribute{
					Name:        "userdn",
					Description: `(Optional) Base DN under which to perform user search`,
				},
				resource.Attribute{
					Name:        "userattr",
					Description: `(Optional) Attribute on user object matching username passed in`,
				},
				resource.Attribute{
					Name:        "upndomain",
					Description: `(Optional) The userPrincipalDomain used to construct UPN string`,
				},
				resource.Attribute{
					Name:        "groupfilter",
					Description: `(Optional) Go template used to construct group membership query`,
				},
				resource.Attribute{
					Name:        "groupdn",
					Description: `(Optional) Base DN under which to perform group search`,
				},
				resource.Attribute{
					Name:        "groupattr",
					Description: `(Optional) LDAP attribute to follow on objects returned by groupfilter`,
				},
				resource.Attribute{
					Name:        "use_token_groups",
					Description: `(Optional) Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to mount the LDAP auth backend under`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the LDAP auth backend mount For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api/auth/ldap/index.html). ~>`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth mount. ## Import LDAP authentication backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ldap_auth_backend.ldap ldap ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth mount. ## Import LDAP authentication backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ldap_auth_backend.ldap ldap ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ldap_auth_backend_group",
			Category:         "Resources",
			ShortDescription: `Managing groups in an LDAP auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"ldap",
				"auth",
				"backend",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "groupname",
					Description: `(Required) The LDAP groupname`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Policies which should be granted to members of the group`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Path to the authentication backend For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api/auth/ldap/index.html). ## Attribute Reference No additional attributes are exposed by this resource. ## Import LDAP authentication backend groups can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ldap_auth_backend_group.foo auth/ldap/groups/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ldap_auth_backend_user",
			Category:         "Resources",
			ShortDescription: `Managing users in an LDAP auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"ldap",
				"auth",
				"backend",
				"user",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The LDAP username`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Policies which should be granted to user`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Override LDAP groups which should be granted to user`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Path to the authentication backend For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api/auth/ldap/index.html). ## Attribute Reference No additional attributes are exposed by this resource. ## Import LDAP authentication backend users can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ldap_auth_backend_user.foo auth/ldap/users/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_mount",
			Category:         "Resources",
			ShortDescription: `Managing the mounting of secret backends in Vault`,
			Description:      ``,
			Keywords: []string{
				"mount",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Where the secret backend will be mounted`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the backend, such as "aws"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description of the mount`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) Default lease duration for tokens and secrets in seconds`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) Maximum possible lease duration for tokens and secrets in seconds`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Specifies mount type specific options that are passed to the backend ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this mount. ## Import Mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_mount.example dummy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this mount. ## Import Mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_mount.example dummy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_namespace",
			Category:         "Resources",
			ShortDescription: `Writes namespaces for Vault`,
			Description:      ``,
			Keywords: []string{
				"namespace",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of the namespace. Must not have a trailing ` + "`" + `/` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_okta_auth_backend",
			Category:         "Resources",
			ShortDescription: `Managing Okta auth backends in Vault`,
			Description:      ``,
			Keywords: []string{
				"okta",
				"auth",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path to mount the Okta auth backend`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the auth backend`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) The Okta organization. This will be the first part of the url ` + "`" + `https://XXX.okta.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The Okta API token. This is required to query Okta for user group membership. If this is not supplied only locally configured groups will be enabled.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Optional) The Okta url. Examples: oktapreview.com, okta.com`,
				},
				resource.Attribute{
					Name:        "bypass_okta_mfa",
					Description: `(Optional) When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Duration after which authentication will be expired. [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) Maximum duration after which authentication will be expired [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Associate Okta groups with policies within Vault. [See below for more details](#okta-group).`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) Associate Okta users with groups or policies within Vault. [See below for more details](#okta-user). ### Okta Group`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) Name of the group within the Okta`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Vault policies to associate with this group ### Okta User`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required Optional) Name of the user within Okta`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of Okta groups to associate with this user`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) List of Vault policies to associate with this user ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "accessor",
					Description: `The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_okta_auth_backend_group",
			Category:         "Resources",
			ShortDescription: `Managing groups in an Okta auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"okta",
				"auth",
				"backend",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path where the Okta auth backend is mounted`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) Name of the group within the Okta`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Vault policies to associate with this group ## Attributes Reference No additional attributes are exposed by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_okta_auth_backend_user",
			Category:         "Resources",
			ShortDescription: `Managing users in an Okta auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"okta",
				"auth",
				"backend",
				"user",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path where the Okta auth backend is mounted`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required Optional) Name of the user within Okta`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of Okta groups to associate with this user`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) List of Vault policies to associate with this user ## Attributes Reference No additional attributes are exposed by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an PKI secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default TTL for credentials issued by this backend.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend. ## Attributes Reference No additional attributes are exported by this resource. ## Import PKI secret backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_pki_secret_backend.pki pki ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_cert",
			Category:         "Resources",
			ShortDescription: `Generate an PKI certificate.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"cert",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the role to create the certificate against`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of certificate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "private_key_format",
					Description: `(Optional) The private key format`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "min_seconds_remaining",
					Description: `(Optional) Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days) ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "private_key_type",
					Description: `The private key type`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The serial number`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "private_key_type",
					Description: `The private key type`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The serial number`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_config_ca",
			Category:         "Resources",
			ShortDescription: `Submit the CA information to PKI.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"config",
				"ca",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "pem_bundle",
					Description: `(Required) The key and certificate PEM bundle ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_config_urls",
			Category:         "Resources",
			ShortDescription: `Sets the config URL's on an PKI Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"config",
				"urls",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the PKI secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "issuing_certificates",
					Description: `(Optional) Specifies the URL values for the Issuing Certificate field. Comma-separated string if multiple.`,
				},
				resource.Attribute{
					Name:        "crl_distribution_points",
					Description: `(Optional) Specifies the URL values for the CRL Distribution Points field. Comma-separated string if multiple.`,
				},
				resource.Attribute{
					Name:        "ocsp_servers",
					Description: `(Optional) Specifies the URL values for the OCSP Servers field. Comma-separated string if multiple. ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_intermediate_cert_request",
			Category:         "Resources",
			ShortDescription: `Generate a new private key and a CSR for signing the PKI.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"intermediate",
				"cert",
				"request",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of intermediate to create. Must be either \"exported\" or \"internal\"`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of intermediate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "uri_sans",
					Description: `(Optional) List of alternative URIs`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "private_key_format",
					Description: `(Optional) The private key format`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Optional) The desired key type`,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Optional) The number of bits to use`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `(Optional) The organization unit`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) The organization`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country`,
				},
				resource.Attribute{
					Name:        "locality",
					Description: `(Optional) The locality`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `(Optional) The province`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `The CSR`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "private_key_type",
					Description: `The private key type`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The serial number`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "csr",
					Description: `The CSR`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "private_key_type",
					Description: `The private key type`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The serial number`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_intermediate_set_signed",
			Category:         "Resources",
			ShortDescription: `Submit the PKI CA certificate.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"intermediate",
				"set",
				"signed",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The certificate ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Create a role on an PKI Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the PKI secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to identify this role within the backend. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum TTL`,
				},
				resource.Attribute{
					Name:        "allow_localhost",
					Description: `(Optional) Flag to allow certificates for localhost`,
				},
				resource.Attribute{
					Name:        "allowed_domains",
					Description: `(Optional) List of allowed domains for certificates`,
				},
				resource.Attribute{
					Name:        "allow_bare_domains",
					Description: `(Optional) Flag to allow certificates matching the actual domain`,
				},
				resource.Attribute{
					Name:        "allow_subdomains",
					Description: `(Optional) Flag to allow certificates matching subdomains`,
				},
				resource.Attribute{
					Name:        "allow_glob_domains",
					Description: `(Optional) Flag to allow names containing glob patterns.`,
				},
				resource.Attribute{
					Name:        "allow_any_name",
					Description: `(Optional) Flag to allow any name`,
				},
				resource.Attribute{
					Name:        "enforce_hostnames",
					Description: `(Optional) Flag to allow only valid host names`,
				},
				resource.Attribute{
					Name:        "allow_ip_sans",
					Description: `(Optional) Flag to allow IP SANs`,
				},
				resource.Attribute{
					Name:        "allowed_uri_sans",
					Description: `(Optional) Defines allowed URI SANs`,
				},
				resource.Attribute{
					Name:        "allowed_other_sans",
					Description: `(Optional) Defines allowed custom SANs`,
				},
				resource.Attribute{
					Name:        "server_flag",
					Description: `(Optional) Flag to specify certificates for server use`,
				},
				resource.Attribute{
					Name:        "client_flag",
					Description: `(Optional) Flag to specify certificates for client use`,
				},
				resource.Attribute{
					Name:        "code_signing_flag",
					Description: `(Optional) Flag to specify certificates for code signing use`,
				},
				resource.Attribute{
					Name:        "email_protection_flag",
					Description: `(Optional) Flag to specify certificates for email protection use`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Optional) The type of generated keys`,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Optional) The number of bits of generated keys`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `(Optional) Specify the allowed key usage constraint on issued certificates`,
				},
				resource.Attribute{
					Name:        "ext_key_usage",
					Description: `(Optional) Specify the allowed extended key usage constraint on issued certificates`,
				},
				resource.Attribute{
					Name:        "use_csr_common_name",
					Description: `(Optional) Flag to use the CN in the CSR`,
				},
				resource.Attribute{
					Name:        "use_csr_sans",
					Description: `(Optional) Flag to use the SANs in the CSR`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `(Optional) The organization unit of generated certificates`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) The organization of generated certificates`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country of generated certificates`,
				},
				resource.Attribute{
					Name:        "locality",
					Description: `(Optional) The locality of generated certificates`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `(Optional) The province of generated certificates`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address of generated certificates`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code of generated certificates`,
				},
				resource.Attribute{
					Name:        "generate_lease",
					Description: `(Optional) Flag to generate leases with certificates`,
				},
				resource.Attribute{
					Name:        "no_store",
					Description: `(Optional) Flag to not store certificates in the storage backend`,
				},
				resource.Attribute{
					Name:        "require_cn",
					Description: `(Optional) Flag to force CN usage`,
				},
				resource.Attribute{
					Name:        "policy_identifiers",
					Description: `(Optional) Specify the list of allowed policies IODs`,
				},
				resource.Attribute{
					Name:        "basic_constraints_valid_for_non_ca",
					Description: `(Optional) Flag to mark basic constraints valid when issuing non-CA certificates ## Attributes Reference No additional attributes are exported by this resource. ## Import PKI secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_pki_secret_backend_role.role pki/roles/my_role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_root_cert",
			Category:         "Resources",
			ShortDescription: `Generate root.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"root",
				"cert",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of intermediate to create. Must be either \"exported\" or \"internal\"`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of intermediate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "uri_sans",
					Description: `(Optional) List of alternative URIs`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "private_key_format",
					Description: `(Optional) The private key format`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Optional) The desired key type`,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Optional) The number of bits to use`,
				},
				resource.Attribute{
					Name:        "max_path_length",
					Description: `(Optional) The maximum path length to encode in the generated certificate`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "permitted_dns_domains",
					Description: `(Optional) List of domains for which certificates are allowed to be issued`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `(Optional) The organization unit`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) The organization`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country`,
				},
				resource.Attribute{
					Name:        "locality",
					Description: `(Optional) The locality`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `(Optional) The province`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_root_sign_intermediate",
			Category:         "Resources",
			ShortDescription: `Signs intermediate certificate.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"root",
				"sign",
				"intermediate",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `(Required) The CSR`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of intermediate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "uri_sans",
					Description: `(Optional) List of alternative URIs`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "private_key_format",
					Description: `(Optional) The private key format`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Optional) The desired key type`,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Optional) The number of bits to use`,
				},
				resource.Attribute{
					Name:        "max_path_length",
					Description: `(Optional) The maximum path length to encode in the generated certificate`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "use_csr_values",
					Description: `(Optional) Preserve CSR values`,
				},
				resource.Attribute{
					Name:        "permitted_dns_domains",
					Description: `(Optional) List of domains for which certificates are allowed to be issued`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `(Optional) The organization unit`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) The organization`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country`,
				},
				resource.Attribute{
					Name:        "locality",
					Description: `(Optional) The locality`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `(Optional) The province`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_sign",
			Category:         "Resources",
			ShortDescription: `Sign a new certificate based on the CSR by the PKI.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"sign",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the role to create the certificate against`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `(Required) The CSR`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of certificate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "uri_sans",
					Description: `(Optional) List of alterative URIs`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "min_seconds_remaining",
					Description: `(Optional) Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days) ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_policy",
			Category:         "Resources",
			ShortDescription: `Writes arbitrary policies for Vault`,
			Description:      ``,
			Keywords: []string{
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) String containing a Vault policy ## Attributes Reference No additional attributes are exported by this resource. ## Import Policies can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_policy.example dev-team ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_rabbitmq_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an RabbitMQ secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"rabbitmq",
				"secret",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "connection_uri",
					Description: `(Required) Specifies the RabbitMQ connection URI.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Specifies the RabbitMQ management administrator username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Specifies the RabbitMQ management administrator password.`,
				},
				resource.Attribute{
					Name:        "verify_connection",
					Description: `(Optional) Specifies whether to verify connection URI, username, and password. Defaults to ` + "`" + `true` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `aws` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default TTL for credentials issued by this backend.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend. ## Attributes Reference No additional attributes are exported by this resource. ## Import RabbitMQ secret backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_rabbitmq_secret_backend.rabbitmq rabbitmq ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_rabbitmq_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Creates a role on an RabbitMQ Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"rabbitmq",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the RabbitMQ secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to identify this role within the backend. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Specifies a comma-separated RabbitMQ management tags.`,
				},
				resource.Attribute{
					Name:        "vhosts",
					Description: `(Optional) Specifies a map of virtual hosts to permissions. ## Attributes Reference No additional attributes are exported by this resource. ## Import RabbitMQ secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_rabbitmq_secret_backend_role.role rabbitmq/roles/deploy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_rgp_policy",
			Category:         "Resources",
			ShortDescription: `Writes Sentinel role governing policies for Vault`,
			Description:      ``,
			Keywords: []string{
				"rgp",
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy`,
				},
				resource.Attribute{
					Name:        "enforcement_level",
					Description: `(Required) Enforcement level of Sentinel policy. Can be either ` + "`" + `advisory` + "`" + ` or ` + "`" + `soft-mandatory` + "`" + ` or ` + "`" + `hard-mandatory` + "`" + ``,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) String containing a Sentinel policy ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ssh_secret_backend_ca",
			Category:         "Resources",
			ShortDescription: `Managing CA information in an SSH secret backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"secret",
				"backend",
				"ca",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path where the SSH secret backend is mounted. Defaults to 'ssh'`,
				},
				resource.Attribute{
					Name:        "generate_signing_key",
					Description: `(Optional) Whether Vault should generate the signing key pair internally. Defaults to true`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The public key part the SSH CA key pair; required if generate_signing_key is false.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The private key part the SSH CA key pair; required if generate_signing_key is false. ~>`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ssh_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Managing roles in an SSH secret backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the role to create.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path where the SSH secret backend is mounted.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Required) Specifies the type of credentials generated by this role. This can be either ` + "`" + `otp` + "`" + `, ` + "`" + `dynamic` + "`" + ` or ` + "`" + `ca` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_bare_domains",
					Description: `(Optional) Specifies if host certificates that are requested are allowed to use the base domains listed in ` + "`" + `allowed_domains` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_host_certificates",
					Description: `(Optional) Specifies if certificates are allowed to be signed for use as a 'host'.`,
				},
				resource.Attribute{
					Name:        "allow_subdomains",
					Description: `(Optional) Specifies if host certificates that are requested are allowed to be subdomains of those listed in ` + "`" + `allowed_domains` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_user_certificates",
					Description: `(Optional) Specifies if certificates are allowed to be signed for use as a 'user'.`,
				},
				resource.Attribute{
					Name:        "allow_user_key_ids",
					Description: `(Optional) Specifies if users can override the key ID for a signed certificate with the ` + "`" + `key_id` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "allowed_critical_options",
					Description: `(Optional) Specifies a comma-separated list of critical options that certificates can have when signed.`,
				},
				resource.Attribute{
					Name:        "allowed_domains",
					Description: `(Optional) The list of domains for which a client can request a host certificate.`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Optional) The comma-separated string of CIDR blocks for which this role is applicable.`,
				},
				resource.Attribute{
					Name:        "allowed_extensions",
					Description: `(Optional) Specifies a comma-separated list of extensions that certificates can have when signed.`,
				},
				resource.Attribute{
					Name:        "default_extensions",
					Description: `(Optional) Specifies a map of extensions that certificates have when signed.`,
				},
				resource.Attribute{
					Name:        "default_critical_options",
					Description: `(Optional) Specifies a map of critical options that certificates have when signed.`,
				},
				resource.Attribute{
					Name:        "allowed_users",
					Description: `(Optional) Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.`,
				},
				resource.Attribute{
					Name:        "default_user",
					Description: `(Optional) Specifies the default username for which a credential will be generated.`,
				},
				resource.Attribute{
					Name:        "key_id_format",
					Description: `(Optional) Specifies a custom format for the key id of a signed certificate.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) Specifies the Time To Live value.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Specifies the maximum Time To Live value. ## Attributes Reference No additional attributes are exposed by this resource. ## Import SSH secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ssh_secret_backend_role.foo ssh/roles/my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_token",
			Category:         "Resources",
			ShortDescription: `Writes token for Vault`,
			Description:      ``,
			Keywords: []string{
				"token",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Optional) The token role name`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) List of policies to attach to this token`,
				},
				resource.Attribute{
					Name:        "no_parent",
					Description: `(Optional) Flag to create a token without parent`,
				},
				resource.Attribute{
					Name:        "no_default_policy",
					Description: `(Optional) Flag to not attach the default policy to this token`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `(Optional) Flag to allow to renew this token`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL period of this token`,
				},
				resource.Attribute{
					Name:        "explicit_max_ttl",
					Description: `(Optional) The explicit max TTL of this token`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) String containing the token display name`,
				},
				resource.Attribute{
					Name:        "num_uses",
					Description: `(Optional) The number of allowed uses of this token`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The period of this token`,
				},
				resource.Attribute{
					Name:        "renew_min_lease",
					Description: `(Optional) The minimal lease to renew this token`,
				},
				resource.Attribute{
					Name:        "renew_increment",
					Description: `(Optional) The renew increment ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `String containing the token lease duration if present in state file`,
				},
				resource.Attribute{
					Name:        "lease_started",
					Description: `String containing the token lease started time if present in state file`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `String containing the client token if stored in present file ## Import Tokens can be imported using its ` + "`" + `id` + "`" + ` as accessor id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_token.example <accessor_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "lease_duration",
					Description: `String containing the token lease duration if present in state file`,
				},
				resource.Attribute{
					Name:        "lease_started",
					Description: `String containing the token lease started time if present in state file`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `String containing the client token if stored in present file ## Import Tokens can be imported using its ` + "`" + `id` + "`" + ` as accessor id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_token.example <accessor_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_token_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages Token auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"token",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role.`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"vault_approle_auth_backend_login":          0,
		"vault_approle_auth_backend_role":           1,
		"vault_approle_auth_backend_role_secret_id": 2,
		"vault_audit":                                        3,
		"vault_auth_backend":                                 4,
		"vault_aws_auth_backend_cert":                        5,
		"vault_aws_auth_backend_client":                      6,
		"vault_aws_auth_backend_identity_whitelist":          7,
		"vault_aws_auth_backend_login":                       8,
		"vault_aws_auth_backend_role":                        9,
		"vault_aws_auth_backend_role_tag":                    10,
		"vault_aws_auth_backend_roletag_blacklist":           11,
		"vault_aws_auth_backend_sts_role":                    12,
		"vault_aws_secret_backend":                           13,
		"vault_aws_secret_backend_role":                      14,
		"vault_azure_auth_backend_config":                    15,
		"vault_azure_auth_backend_role":                      16,
		"vault_cert_auth_backend_role":                       17,
		"vault_consul_secret_backend":                        18,
		"vault_database_secret_backend_connection":           19,
		"vault_database_secret_backend_role":                 20,
		"vault_egp_policy":                                   21,
		"vault_gcp_auth_backend":                             22,
		"vault_gcp_auth_backend_role":                        23,
		"vault_gcp_secret_backend":                           24,
		"vault_gcp_secret_roleset":                           25,
		"vault_generic_endpoint":                             26,
		"vault_generic_secret":                               27,
		"vault_github_auth_backend":                          28,
		"vault_github_team":                                  29,
		"vault_github_user":                                  30,
		"vault_identity_entity":                              31,
		"vault_identity_entity_alias":                        32,
		"vault_identity_group":                               33,
		"vault_identity_group_alias":                         34,
		"vault_jwt_auth_backend":                             35,
		"vault_jwt_auth_backend_role":                        36,
		"vault_kubernetes_auth_backend_config":               37,
		"vault_kubernetes_auth_backend_role":                 38,
		"vault_ldap_auth_backend":                            39,
		"vault_ldap_auth_backend_group":                      40,
		"vault_ldap_auth_backend_user":                       41,
		"vault_mount":                                        42,
		"vault_namespace":                                    43,
		"vault_okta_auth_backend":                            44,
		"vault_okta_auth_backend_group":                      45,
		"vault_okta_auth_backend_user":                       46,
		"vault_pki_secret_backend":                           47,
		"vault_pki_secret_backend_cert":                      48,
		"vault_pki_secret_backend_config_ca":                 49,
		"vault_pki_secret_backend_config_urls":               50,
		"vault_pki_secret_backend_intermediate_cert_request": 51,
		"vault_pki_secret_backend_intermediate_set_signed":   52,
		"vault_pki_secret_backend_role":                      53,
		"vault_pki_secret_backend_root_cert":                 54,
		"vault_pki_secret_backend_root_sign_intermediate":    55,
		"vault_pki_secret_backend_sign":                      56,
		"vault_policy":                                       57,
		"vault_rabbitmq_secret_backend":                      58,
		"vault_rabbitmq_secret_backend_role":                 59,
		"vault_rgp_policy":                                   60,
		"vault_ssh_secret_backend_ca":                        61,
		"vault_ssh_secret_backend_role":                      62,
		"vault_token":                                        63,
		"vault_token_auth_backend_role":                      64,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}