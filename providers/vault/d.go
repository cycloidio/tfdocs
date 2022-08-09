package vault

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vault_ad_access_credentials",
			Category:         "Data Sources",
			ShortDescription: `Reads Active Directory credentials from an AD secret backend in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path to the AD secret backend to read credentials from, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the AD secret backend role to read credentials from, with no leading or trailing ` + "`" + `/` + "`" + `s. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "current_password",
					Description: `The current set password on the Active Directory service account.`,
				},
				resource.Attribute{
					Name:        "last_password",
					Description: `The current set password on the Active Directory service account, provided because AD is eventually consistent.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The Active Directory service account username.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "current_password",
					Description: `The current set password on the Active Directory service account.`,
				},
				resource.Attribute{
					Name:        "last_password",
					Description: `The current set password on the Active Directory service account, provided because AD is eventually consistent.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The Active Directory service account username.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_approle_auth_backend_role_id",
			Category:         "Data Sources",
			ShortDescription: `Manages AppRole auth backend roles in Vault.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role to retrieve the Role ID for.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name for the AppRole backend the role to retrieve a RoleID for resides in. Defaults to "approle". ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The RoleID of the role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `The RoleID of the role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_auth_backend",
			Category:         "Data Sources",
			ShortDescription: `Lookup an Auth Backend from Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The auth backend mount point. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The name of the auth method type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the auth method.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `The default lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `The maximum lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `Specifies whether to show this mount in the UI-specific listing endpoint.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `Specifies if the auth method is local only.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth method.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The name of the auth method type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the auth method.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `The default lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `The maximum lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `Specifies whether to show this mount in the UI-specific listing endpoint.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `Specifies if the auth method is local only.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth method.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_access_credentials",
			Category:         "Data Sources",
			ShortDescription: `Reads AWS credentials from an AWS secret backend in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path to the AWS secret backend to read credentials from, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the AWS secret backend role to read credentials from, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of credentials to read. Defaults to ` + "`" + `"creds"` + "`" + `, which just returns an AWS Access Key ID and Secret Key. Can also be set to ` + "`" + `"sts"` + "`" + `, which will return a security token in addition to the keys.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required if role has multiple ARNs) The specific AWS ARN to use from the configured role. If the role does not have multiple ARNs, this does not need to be specified.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Specifies the TTL for the use of the STS token. This is specified as a string with a duration suffix. Valid only when ` + "`" + `credential_type` + "`" + ` is ` + "`" + `assumed_role` + "`" + ` or ` + "`" + `federation_token` + "`" + ` ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `The AWS Access Key ID returned by Vault.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The AWS Secret Key returned by Vault.`,
				},
				resource.Attribute{
					Name:        "security_token",
					Description: `The STS token returned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `As a convenience, this records the current time on the computer where Terraform is running when the data is requested. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to the Vault server.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key",
					Description: `The AWS Access Key ID returned by Vault.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The AWS Secret Key returned by Vault.`,
				},
				resource.Attribute{
					Name:        "security_token",
					Description: `The STS token returned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `As a convenience, this records the current time on the computer where Terraform is running when the data is requested. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to the Vault server.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_azure_access_credentials",
			Category:         "Data Sources",
			ShortDescription: `Reads Azure credentials from an Azure secret backend in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path to the Azure secret backend to read credentials from, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the Azure secret backend role to read credentials from, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "validate_creds",
					Description: `(Optional) Whether generated credentials should be validated before being returned. Defaults to ` + "`" + `false` + "`" + `, which returns credentials without checking whether they have fully propagated throughout Azure Active Directory. Designating ` + "`" + `true` + "`" + ` activates testing.`,
				},
				resource.Attribute{
					Name:        "num_sequential_successes",
					Description: `(Optional) If 'validate_creds' is true, the number of sequential successes required to validate generated credentials. Defaults to 8.`,
				},
				resource.Attribute{
					Name:        "num_seconds_between_tests",
					Description: `(Optional) If 'validate_creds' is true, the number of seconds to wait between each test of generated credentials. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "max_cred_validation_seconds",
					Description: `(Optional) If 'validate_creds' is true, the number of seconds after which to give up validating credentials. Defaults to 300.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Optional) The subscription ID to use during credential validation. Defaults to the subscription ID configured in the Vault ` + "`" + `backend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The tenant ID to use during credential validation. Defaults to the tenant ID configured in the Vault ` + "`" + `backend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) The Azure environment to use during credential validation. Defaults to the environment configured in the Vault backend. Some possible values: ` + "`" + `AzurePublicCloud` + "`" + `, ` + "`" + `AzureGovernmentCloud` + "`" + ``,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The client id for credentials to query the Azure APIs.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `The client secret for credentials to query the Azure APIs.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `As a convenience, this records the current time on the computer where Terraform is running when the data is requested. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to the Vault server.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `The client id for credentials to query the Azure APIs.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `The client secret for credentials to query the Azure APIs.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `As a convenience, this records the current time on the computer where Terraform is running when the data is requested. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to the Vault server.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_generic_secret",
			Category:         "Data Sources",
			ShortDescription: `Reads arbitrary data from a given path in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The full logical path from which to request data. To read data from the "generic" secret backend mounted in Vault by default, this should be prefixed with ` + "`" + `secret/` + "`" + `. Reading from other backends with this data source is possible; consult each backend's documentation to see which endpoints support the ` + "`" + `GET` + "`" + ` method.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the secret to read. This is used by the Vault KV secrets engine - version 2 to indicate which version of the secret to read.`,
				},
				resource.Attribute{
					Name:        "with_lease_start_time",
					Description: `If set to true, stores ` + "`" + `lease_start_time` + "`" + ` in the TF state. Note that storing the ` + "`" + `lease_start_time` + "`" + ` in the TF state will cause a persistent drift on every ` + "`" + `terraform plan` + "`" + ` and will require a ` + "`" + `terraform apply` + "`" + `. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `read` + "`" + ` capability on the given path. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `A string containing the full data payload retrieved from Vault, serialized in JSON format.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `The date and time of Terraform execution. It is derived from the local machine's clock, and is recorded in RFC3339 format UTC. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to the Vault server. _Provided only as a convenience_.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data_json",
					Description: `A string containing the full data payload retrieved from Vault, serialized in JSON format.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `The date and time of Terraform execution. It is derived from the local machine's clock, and is recorded in RFC3339 format UTC. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to the Vault server. _Provided only as a convenience_.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_entity",
			Category:         "Data Sources",
			ShortDescription: `Lookup an Identity Entity from Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "entity_name",
					Description: `(Optional) Name of the entity.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `(Optional) ID of the entity.`,
				},
				resource.Attribute{
					Name:        "alias_id",
					Description: `(Optional) ID of the alias.`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `(Optional) Name of the alias. This should be supplied in conjunction with ` + "`" + `alias_mount_accessor` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alias_mount_accessor",
					Description: `(Optional) Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with ` + "`" + `alias_name` + "`" + `. The lookup criteria can be ` + "`" + `entity_name` + "`" + `, ` + "`" + `entity_id` + "`" + `, ` + "`" + `alias_id` + "`" + `, or a combination of ` + "`" + `alias_name` + "`" + ` and ` + "`" + `alias_mount_accessor` + "`" + `. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `create` + "`" + ` capability on ` + "`" + `/identity/lookup/entity` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `A string containing the full data payload retrieved from Vault, serialized in JSON format.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation timestamp of the entity`,
				},
				resource.Attribute{
					Name:        "direct_group_ids",
					Description: `List of Group IDs of which the entity is directly a member of`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Whether the entity is disabled`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of all Group IDs of which the entity is a member of`,
				},
				resource.Attribute{
					Name:        "inherited_group_ids",
					Description: `List of all Group IDs of which the entity is a member of transitively`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last updated time of the entity`,
				},
				resource.Attribute{
					Name:        "merged_entity_ids",
					Description: `Other entity IDs which is merged with this entity`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Arbitrary metadata`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Namespace of which the entity is part of`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `List of policies attached to the entity`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `A list of entity alias. Structure is documented below. ### Aliases`,
				},
				resource.Attribute{
					Name:        "canonical_id",
					Description: `Canonical ID of the Alias`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the Alias`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the alias`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last update time of the alias`,
				},
				resource.Attribute{
					Name:        "merged_from_canonical_ids",
					Description: `List of canonical IDs merged with this alias`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Arbitrary metadata`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `Authentication mount acccessor which this alias belongs to`,
				},
				resource.Attribute{
					Name:        "mount_path",
					Description: `Authentication mount path which this alias belongs to`,
				},
				resource.Attribute{
					Name:        "mount_type",
					Description: `Authentication mount type which this alias belongs to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the alias`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data_json",
					Description: `A string containing the full data payload retrieved from Vault, serialized in JSON format.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation timestamp of the entity`,
				},
				resource.Attribute{
					Name:        "direct_group_ids",
					Description: `List of Group IDs of which the entity is directly a member of`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Whether the entity is disabled`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of all Group IDs of which the entity is a member of`,
				},
				resource.Attribute{
					Name:        "inherited_group_ids",
					Description: `List of all Group IDs of which the entity is a member of transitively`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last updated time of the entity`,
				},
				resource.Attribute{
					Name:        "merged_entity_ids",
					Description: `Other entity IDs which is merged with this entity`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Arbitrary metadata`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Namespace of which the entity is part of`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `List of policies attached to the entity`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `A list of entity alias. Structure is documented below. ### Aliases`,
				},
				resource.Attribute{
					Name:        "canonical_id",
					Description: `Canonical ID of the Alias`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the Alias`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the alias`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last update time of the alias`,
				},
				resource.Attribute{
					Name:        "merged_from_canonical_ids",
					Description: `List of canonical IDs merged with this alias`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Arbitrary metadata`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `Authentication mount acccessor which this alias belongs to`,
				},
				resource.Attribute{
					Name:        "mount_path",
					Description: `Authentication mount path which this alias belongs to`,
				},
				resource.Attribute{
					Name:        "mount_type",
					Description: `Authentication mount type which this alias belongs to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the alias`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_group",
			Category:         "Data Sources",
			ShortDescription: `Lookup an Identity Group from Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Optional) Name of the group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) ID of the group.`,
				},
				resource.Attribute{
					Name:        "alias_id",
					Description: `(Optional) ID of the alias.`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `(Optional) Name of the alias. This should be supplied in conjunction with ` + "`" + `alias_mount_accessor` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alias_mount_accessor",
					Description: `(Optional) Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with ` + "`" + `alias_name` + "`" + `. The lookup criteria can be ` + "`" + `group_name` + "`" + `, ` + "`" + `group_id` + "`" + `, ` + "`" + `alias_id` + "`" + `, or a combination of ` + "`" + `alias_name` + "`" + ` and ` + "`" + `alias_mount_accessor` + "`" + `. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `create` + "`" + ` capability on ` + "`" + `/identity/lookup/group` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `A string containing the full data payload retrieved from Vault, serialized in JSON format.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation timestamp of the group`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last updated time of the group`,
				},
				resource.Attribute{
					Name:        "member_entity_ids",
					Description: `List of Entity IDs which are members of this group`,
				},
				resource.Attribute{
					Name:        "member_group_ids",
					Description: `List of Group IDs which are members of this group`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Arbitrary metadata`,
				},
				resource.Attribute{
					Name:        "modify_index",
					Description: `Modify index of the group`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Namespace of which the group is part of`,
				},
				resource.Attribute{
					Name:        "parent_group_ids",
					Description: `List of Group IDs which are parents of this group.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `List of policies attached to the group`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of group`,
				},
				resource.Attribute{
					Name:        "alias_canonical_id",
					Description: `Canonical ID of the Alias`,
				},
				resource.Attribute{
					Name:        "alias_creation_time",
					Description: `Creation time of the Alias`,
				},
				resource.Attribute{
					Name:        "alias_last_update_time",
					Description: `Last update time of the alias`,
				},
				resource.Attribute{
					Name:        "alias_merged_from_canonical_ids",
					Description: `List of canonical IDs merged with this alias`,
				},
				resource.Attribute{
					Name:        "alias_metadata",
					Description: `Arbitrary metadata`,
				},
				resource.Attribute{
					Name:        "alias_mount_path",
					Description: `Authentication mount path which this alias belongs to`,
				},
				resource.Attribute{
					Name:        "alias_mount_type",
					Description: `Authentication mount type which this alias belongs to`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data_json",
					Description: `A string containing the full data payload retrieved from Vault, serialized in JSON format.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation timestamp of the group`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last updated time of the group`,
				},
				resource.Attribute{
					Name:        "member_entity_ids",
					Description: `List of Entity IDs which are members of this group`,
				},
				resource.Attribute{
					Name:        "member_group_ids",
					Description: `List of Group IDs which are members of this group`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Arbitrary metadata`,
				},
				resource.Attribute{
					Name:        "modify_index",
					Description: `Modify index of the group`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Namespace of which the group is part of`,
				},
				resource.Attribute{
					Name:        "parent_group_ids",
					Description: `List of Group IDs which are parents of this group.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `List of policies attached to the group`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of group`,
				},
				resource.Attribute{
					Name:        "alias_canonical_id",
					Description: `Canonical ID of the Alias`,
				},
				resource.Attribute{
					Name:        "alias_creation_time",
					Description: `Creation time of the Alias`,
				},
				resource.Attribute{
					Name:        "alias_last_update_time",
					Description: `Last update time of the alias`,
				},
				resource.Attribute{
					Name:        "alias_merged_from_canonical_ids",
					Description: `List of canonical IDs merged with this alias`,
				},
				resource.Attribute{
					Name:        "alias_metadata",
					Description: `Arbitrary metadata`,
				},
				resource.Attribute{
					Name:        "alias_mount_path",
					Description: `Authentication mount path which this alias belongs to`,
				},
				resource.Attribute{
					Name:        "alias_mount_type",
					Description: `Authentication mount type which this alias belongs to`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_client_creds",
			Category:         "Data Sources",
			ShortDescription: `Reads client credentials from an OIDC Client provisioned in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the OIDC Client in Vault. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID returned by Vault.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `The Client Secret Key returned by Vault.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID returned by Vault.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `The Client Secret Key returned by Vault.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_openid_config",
			Category:         "Data Sources",
			ShortDescription: `Reads OpenID Configuration from an OIDC Provider provisioned in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the OIDC Provider in Vault. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The URL of the issuer for the provider.`,
				},
				resource.Attribute{
					Name:        "jwks_uri",
					Description: `The well known keys URI for the provider.`,
				},
				resource.Attribute{
					Name:        "authorization_endpoint",
					Description: `The Authorization Endpoint for the provider.`,
				},
				resource.Attribute{
					Name:        "token_endpoint",
					Description: `The Token Endpoint for the provider.`,
				},
				resource.Attribute{
					Name:        "userinfo_endpoint",
					Description: `The User Info Endpoint for the provider`,
				},
				resource.Attribute{
					Name:        "request_uri_parameter_supported",
					Description: `Specifies whether Request URI Parameter is supported by the provider.`,
				},
				resource.Attribute{
					Name:        "id_token_signing_alg_values_supported",
					Description: `The signing algorithms supported by the provider.`,
				},
				resource.Attribute{
					Name:        "response_types_supported",
					Description: `The response types supported by the provider.`,
				},
				resource.Attribute{
					Name:        "scopes_supported",
					Description: `The scopes supported by the provider.`,
				},
				resource.Attribute{
					Name:        "grant_types_supported",
					Description: `The grant types supported by the provider.`,
				},
				resource.Attribute{
					Name:        "subject_types_supported",
					Description: `The subject types supported by the provider.`,
				},
				resource.Attribute{
					Name:        "token_endpoint_auth_methods_supported",
					Description: `The token endpoint auth methods supported by the provider.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "issuer",
					Description: `The URL of the issuer for the provider.`,
				},
				resource.Attribute{
					Name:        "jwks_uri",
					Description: `The well known keys URI for the provider.`,
				},
				resource.Attribute{
					Name:        "authorization_endpoint",
					Description: `The Authorization Endpoint for the provider.`,
				},
				resource.Attribute{
					Name:        "token_endpoint",
					Description: `The Token Endpoint for the provider.`,
				},
				resource.Attribute{
					Name:        "userinfo_endpoint",
					Description: `The User Info Endpoint for the provider`,
				},
				resource.Attribute{
					Name:        "request_uri_parameter_supported",
					Description: `Specifies whether Request URI Parameter is supported by the provider.`,
				},
				resource.Attribute{
					Name:        "id_token_signing_alg_values_supported",
					Description: `The signing algorithms supported by the provider.`,
				},
				resource.Attribute{
					Name:        "response_types_supported",
					Description: `The response types supported by the provider.`,
				},
				resource.Attribute{
					Name:        "scopes_supported",
					Description: `The scopes supported by the provider.`,
				},
				resource.Attribute{
					Name:        "grant_types_supported",
					Description: `The grant types supported by the provider.`,
				},
				resource.Attribute{
					Name:        "subject_types_supported",
					Description: `The subject types supported by the provider.`,
				},
				resource.Attribute{
					Name:        "token_endpoint_auth_methods_supported",
					Description: `The token endpoint auth methods supported by the provider.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_public_keys",
			Category:         "Data Sources",
			ShortDescription: `Reads well known public keys from an OIDC Provider provisioned in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the OIDC Provider in Vault. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `The public portion of keys for an OIDC provider. Clients can use them to validate the authenticity of an identity token.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keys",
					Description: `The public portion of keys for an OIDC provider. Clients can use them to validate the authenticity of an identity token.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_auth_backend_config",
			Category:         "Data Sources",
			ShortDescription: `Manages Kubernetes auth backend configs in Vault.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name for the Kubernetes backend the config to retrieve Role attributes for resides in. Defaults to "kubernetes". ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "kubernetes_host",
					Description: `Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubernetes_ca_cert",
					Description: `PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.`,
				},
				resource.Attribute{
					Name:        "pem_keys",
					Description: `Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `Optional JWT issuer. If no issuer is specified, ` + "`" + `kubernetes.io/serviceaccount` + "`" + ` will be used as the default issuer.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "kubernetes_host",
					Description: `Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubernetes_ca_cert",
					Description: `PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.`,
				},
				resource.Attribute{
					Name:        "pem_keys",
					Description: `Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `Optional JWT issuer. If no issuer is specified, ` + "`" + `kubernetes.io/serviceaccount` + "`" + ` will be used as the default issuer.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_auth_backend_role",
			Category:         "Data Sources",
			ShortDescription: `Reads Kubernetes auth backend roles in Vault.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role to retrieve the Role attributes for.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name for the Kubernetes backend the role to retrieve Role attributes for resides in. Defaults to "kubernetes". ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bound_service_account_names",
					Description: `List of service account names able to access this role. If set to "`,
				},
				resource.Attribute{
					Name:        "bound_service_account_namespaces",
					Description: `List of namespaces allowed to access this role. If set to "`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `Audience claim to verify in the JWT.`,
				},
				resource.Attribute{
					Name:        "alias_name_source",
					Description: `Method used for generating identity aliases. (vault-1.9+) ### Common Token Attributes These attributes are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bound_service_account_names",
					Description: `List of service account names able to access this role. If set to "`,
				},
				resource.Attribute{
					Name:        "bound_service_account_namespaces",
					Description: `List of namespaces allowed to access this role. If set to "`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `Audience claim to verify in the JWT.`,
				},
				resource.Attribute{
					Name:        "alias_name_source",
					Description: `Method used for generating identity aliases. (vault-1.9+) ### Common Token Attributes These attributes are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kv_secret",
			Category:         "Data Sources",
			ShortDescription: `Reads a KV-V1 secret from a given path in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Full path of the KV-V1 secret. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `read` + "`" + ` capability on the given path. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `True if the duration of this lease can be extended through renewal.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data_json",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `True if the duration of this lease can be extended through renewal.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kv_secret_v2",
			Category:         "Data Sources",
			ShortDescription: `Reads a KV-V2 secret from a given path in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "mount",
					Description: `(Required) Path where KV-V2 engine is mounted.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Full name of the secret. For a nested secret the name is the nested path excluding the mount and data prefix. For example, for a secret at ` + "`" + `kvv2/data/foo/bar/baz` + "`" + ` the name is ` + "`" + `foo/bar/baz` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the secret to retrieve. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `read` + "`" + ` capability on the given path. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Full path where the KVV2 secret is written.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Time at which secret was created.`,
				},
				resource.Attribute{
					Name:        "custom_metadata",
					Description: `Custom metadata for the secret.`,
				},
				resource.Attribute{
					Name:        "deletion_time",
					Description: `Deletion time for the secret.`,
				},
				resource.Attribute{
					Name:        "destroyed",
					Description: `Indicates whether the secret has been destroyed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `Full path where the KVV2 secret is written.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Time at which secret was created.`,
				},
				resource.Attribute{
					Name:        "custom_metadata",
					Description: `Custom metadata for the secret.`,
				},
				resource.Attribute{
					Name:        "deletion_time",
					Description: `Deletion time for the secret.`,
				},
				resource.Attribute{
					Name:        "destroyed",
					Description: `Indicates whether the secret has been destroyed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kv_secrets_list",
			Category:         "Data Sources",
			ShortDescription: `Lists KV-V1 secrets at a given path in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Full KV-V1 path where secrets will be listed. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `read` + "`" + ` capability on the given path. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `List of all secret names listed under the given path.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `List of all secret names listed under the given path.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kv_secrets_list_v2",
			Category:         "Data Sources",
			ShortDescription: `Lists KV-V2 secrets at a given path in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "mount",
					Description: `(Required) Path where KV-V2 engine is mounted.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Full name of the secret. For a nested secret the name is the nested path excluding the mount and data prefix. For example, for a secret at ` + "`" + `kvv2/data/foo/bar/baz` + "`" + ` the name is ` + "`" + `foo/bar/baz` + "`" + `. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `read` + "`" + ` capability on the given path. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Full path where the KV-V2 secrets are listed.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `List of all secret names listed under the given path.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `Full path where the KV-V2 secrets are listed.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `List of all secret names listed under the given path.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kv_subkeys_v2",
			Category:         "Data Sources",
			ShortDescription: `Reads the subkeys for a KV-V2 secret written to Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the target resource. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "mount",
					Description: `(Required) Path where KV-V2 engine is mounted.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Full name of the secret. For a nested secret the name is the nested path excluding the mount and data prefix. For example, for a secret at ` + "`" + `kvv2/data/foo/bar/baz` + "`" + ` the name is ` + "`" + `foo/bar/baz` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Specifies the version to return. If not set the latest version is returned.`,
				},
				resource.Attribute{
					Name:        "depth",
					Description: `(Optional) Specifies the deepest nesting level to provide in the output. If non-zero, keys that reside at the specified depth value will be artificially treated as leaves and will thus be ` + "`" + `null` + "`" + ` even if further underlying sub-keys exist. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `read` + "`" + ` capability on the given path. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Full path where the KV-V2 secrets are listed.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `Subkeys for the KV-V2 secret read from Vault..`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `Full path where the KV-V2 secrets are listed.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `Subkeys for the KV-V2 secret read from Vault..`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_policy_document",
			Category:         "Data Sources",
			ShortDescription: `Generates an Vault policy document in HCL format.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) A path in Vault that this rule applies to.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `(Required) A list of capabilities that this rule apply to ` + "`" + `path` + "`" + `. For example, ["read", "write"].`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the rule. Will be added as a comment to rendered rule.`,
				},
				resource.Attribute{
					Name:        "required_parameters",
					Description: `(Optional) A list of parameters that must be specified.`,
				},
				resource.Attribute{
					Name:        "allowed_parameter",
					Description: `(Optional) Whitelists a list of keys and values that are permitted on the given path. See [Parameters](#Parameters) below.`,
				},
				resource.Attribute{
					Name:        "denied_parameter",
					Description: `(Optional) Blacklists a list of parameter and values. Any values specified here take precedence over ` + "`" + `allowed_parameter` + "`" + `. See [Parameters](#Parameters) below.`,
				},
				resource.Attribute{
					Name:        "min_wrapping_ttl",
					Description: `(Optional) The minimum allowed TTL that clients can specify for a wrapped response.`,
				},
				resource.Attribute{
					Name:        "max_wrapping_ttl",
					Description: `(Optional) The maximum allowed TTL that clients can specify for a wrapped response. ### Parameters Each of ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) name of permitted or denied parameter.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) list of values what are permitted or denied by policy rule. ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "hcl",
					Description: `The above arguments serialized as a standard Vault HCL policy document.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hcl",
					Description: `The above arguments serialized as a standard Vault HCL policy document.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"vault_ad_access_credentials":          0,
		"vault_approle_auth_backend_role_id":   1,
		"vault_auth_backend":                   2,
		"vault_aws_access_credentials":         3,
		"vault_azure_access_credentials":       4,
		"vault_generic_secret":                 5,
		"vault_identity_entity":                6,
		"vault_identity_group":                 7,
		"vault_identity_oidc_client_creds":     8,
		"vault_identity_oidc_openid_config":    9,
		"vault_identity_oidc_public_keys":      10,
		"vault_kubernetes_auth_backend_config": 11,
		"vault_kubernetes_auth_backend_role":   12,
		"vault_kv_secret":                      13,
		"vault_kv_secret_v2":                   14,
		"vault_kv_secrets_list":                15,
		"vault_kv_secrets_list_v2":             16,
		"vault_kv_subkeys_v2":                  17,
		"vault_policy_document":                18,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
