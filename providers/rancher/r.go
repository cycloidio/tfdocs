package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rancher_certificate",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher Certificate resource. This can be used to create certificates for rancher environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the certificate.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A certificate description.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the environment to create the certificate for.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `(Required) The certificate content.`,
				},
				resource.Attribute{
					Name:        "cert_chain",
					Description: `(Optional) The certificate chain.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The certificate key. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cn",
					Description: `The certificate CN.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The certificate algorithm.`,
				},
				resource.Attribute{
					Name:        "cert_fingerprint",
					Description: `The certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `The certificate expiration date.`,
				},
				resource.Attribute{
					Name:        "issued_at",
					Description: `The certificate creation date.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The certificate issuer.`,
				},
				resource.Attribute{
					Name:        "key_size",
					Description: `The certificate key size.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate serial number.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `The list of certificate Subject Alternative Names.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The certificate version. ## Import Certificates can be imported using the Certificate ID in the format ` + "`" + `<environment_id>/<certificate_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_certificate.mycert 1a5/1c605 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_certificate.mycert 1c605 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cn",
					Description: `The certificate CN.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The certificate algorithm.`,
				},
				resource.Attribute{
					Name:        "cert_fingerprint",
					Description: `The certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `The certificate expiration date.`,
				},
				resource.Attribute{
					Name:        "issued_at",
					Description: `The certificate creation date.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The certificate issuer.`,
				},
				resource.Attribute{
					Name:        "key_size",
					Description: `The certificate key size.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate serial number.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `The list of certificate Subject Alternative Names.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The certificate version. ## Import Certificates can be imported using the Certificate ID in the format ` + "`" + `<environment_id>/<certificate_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_certificate.mycert 1a5/1c605 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_certificate.mycert 1c605 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_environment",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher Environment resource. This can be used to create and manage environments on rancher.`,
			Description:      ``,
			Keywords: []string{
				"environment",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the environment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An environment description.`,
				},
				resource.Attribute{
					Name:        "orchestration",
					Description: `(Optional) Must be one of`,
				},
				resource.Attribute{
					Name:        "project_template_id",
					Description: `(Optional) This can be any valid project template ID. If this is set, then orchestration can not be. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Optional) Members to add to the environment. ### Member Parameters Reference A ` + "`" + `member` + "`" + ` takes three parameters:`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) The external ID of the member.`,
				},
				resource.Attribute{
					Name:        "external_id_type",
					Description: `(Required) The external ID type of the member.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role of the member in the environment. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the environment (ie ` + "`" + `1a11` + "`" + `) that can be used in other Terraform resources such as Rancher Stack definitions. ## Import Environments can be imported using their Rancher API ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_environment.dev 1a15 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the environment (ie ` + "`" + `1a11` + "`" + `) that can be used in other Terraform resources such as Rancher Stack definitions. ## Import Environments can be imported using their Rancher API ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_environment.dev 1a15 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_host",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher Host resource. This can be used to manage and delete hosts on Rancher.`,
			Description:      ``,
			Keywords: []string{
				"host",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the host.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A host description.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the environment the host is associated to.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The host name. Used as the primary key to detect the host ID.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A dictionary of labels to apply to the host. Computed internal labels are excluded from that list.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_registration_token",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher Registration Token resource. This can be used to create registration tokens for rancher environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"registration",
				"token",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the registration token.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A registration token description.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the environment to create the token for.`,
				},
				resource.Attribute{
					Name:        "host_labels",
					Description: `(Optional) A map of host labels to add to the registration command.`,
				},
				resource.Attribute{
					Name:        "agent_ip",
					Description: `(Optional) A string containing the CATTLE_AGENT_IP to add to the registration command. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `The command used to start a rancher agent for this environment.`,
				},
				resource.Attribute{
					Name:        "registration_url",
					Description: `The URL to use to register new nodes to the environment.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The token to use to register new nodes to the environment. ## Import Registration tokens can be imported using the Environment and Registration token IDs in the form ` + "`" + `<environment_id>/<registration_token_id>` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registration_token.dev_token 1a5/1c11 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registration_token.dev_token 1c11 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `The command used to start a rancher agent for this environment.`,
				},
				resource.Attribute{
					Name:        "registration_url",
					Description: `The URL to use to register new nodes to the environment.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The token to use to register new nodes to the environment. ## Import Registration tokens can be imported using the Environment and Registration token IDs in the form ` + "`" + `<environment_id>/<registration_token_id>` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registration_token.dev_token 1a5/1c11 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registration_token.dev_token 1c11 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_registry",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher Registy resource. This can be used to create registries for rancher environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"registry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the registry.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A registry description.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the environment to create the registry for.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Required) The server address for the registry. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource. ## Import Registries can be imported using the Environment and Registry IDs in the form ` + "`" + `<environment_id>/<registry_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registry.private_registry 1a5/1sp31 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registry.private_registry 1sp31 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource. ## Import Registries can be imported using the Environment and Registry IDs in the form ` + "`" + `<environment_id>/<registry_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registry.private_registry 1a5/1sp31 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registry.private_registry 1sp31 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_registry_credential",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher Registy Credential resource. This can be used to create registry credentials for rancher environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"registry",
				"credential",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the registry credential.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A registry credential description.`,
				},
				resource.Attribute{
					Name:        "registry_id",
					Description: `(Required) The ID of the registry to create the credential for.`,
				},
				resource.Attribute{
					Name:        "public_value",
					Description: `(Required) The public value (user name) of the account.`,
				},
				resource.Attribute{
					Name:        "secret_value",
					Description: `(Required) The secret value (password) of the account. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource. ## Import Registry credentials can be imported using the Registry and credentials IDs in the format ` + "`" + `<registry_id>/<credential_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registry_credential.private_registry 1sp31/1c605 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then then ` + "`" + `registry_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registry_credential.private_registry 1c605 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource. ## Import Registry credentials can be imported using the Registry and credentials IDs in the format ` + "`" + `<registry_id>/<credential_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registry_credential.private_registry 1sp31/1c605 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then then ` + "`" + `registry_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_registry_credential.private_registry 1c605 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_secret",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher Secret resource. This can be used to create secrets for rancher environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"secret",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the secret.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the secret.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the environment to create the secret for.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The secret value. ## Import Secrets can be imported using the Secret ID in the format ` + "`" + `<environment_id>/<secret_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_secret.mysec 1a5/1se10 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_secret.mysec 1se10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_stack",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher Stack resource. This can be used to create and manage stacks on rancher.`,
			Description:      ``,
			Keywords: []string{
				"stack",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the stack.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A stack description.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the environment to create the stack for.`,
				},
				resource.Attribute{
					Name:        "docker_compose",
					Description: `(Optional) The ` + "`" + `docker-compose.yml` + "`" + ` content to apply for the stack.`,
				},
				resource.Attribute{
					Name:        "rancher_compose",
					Description: `(Optional) The ` + "`" + `rancher-compose.yml` + "`" + ` content to apply for the stack.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) The environment to apply to interpret the docker-compose and rancher-compose files.`,
				},
				resource.Attribute{
					Name:        "catalog_id",
					Description: `(Optional) The catalog ID to link this stack to. When provided, ` + "`" + `docker_compose` + "`" + ` and ` + "`" + `rancher_compose` + "`" + ` will be retrieved from the catalog unless they are overridden.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope to attach the stack to. Must be one of`,
				},
				resource.Attribute{
					Name:        "start_on_create",
					Description: `(Optional) Whether to start the stack automatically.`,
				},
				resource.Attribute{
					Name:        "finish_upgrade",
					Description: `(Optional) Whether to automatically finish upgrades to this stack. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "rendered_docker_compose",
					Description: `The interpolated ` + "`" + `docker_compose` + "`" + ` applied to the stack.`,
				},
				resource.Attribute{
					Name:        "rendered_rancher_compose",
					Description: `The interpolated ` + "`" + `rancher_compose` + "`" + ` applied to the stack. ## Import Stacks can be imported using the Environment and Stack ID in the form ` + "`" + `<environment_id>/<stack_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_stack.foo 1a5/1e149 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_stack.foo 1e149 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "rendered_docker_compose",
					Description: `The interpolated ` + "`" + `docker_compose` + "`" + ` applied to the stack.`,
				},
				resource.Attribute{
					Name:        "rendered_rancher_compose",
					Description: `The interpolated ` + "`" + `rancher_compose` + "`" + ` applied to the stack. ## Import Stacks can be imported using the Environment and Stack ID in the form ` + "`" + `<environment_id>/<stack_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_stack.foo 1a5/1e149 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_stack.foo 1e149 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_volume",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher Volume resource. This can be used to create volumes for rancher environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the volume.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the volume.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the environment to create the volume for.`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Required) The volume driver. ## Import Volumes can be imported using the Volume ID in the format ` + "`" + `<environment_id>/<volume_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_volume.mysec 1a5/1v123456 ` + "`" + `` + "`" + `` + "`" + ` If the credentials for the Rancher provider have access to the global API, then ` + "`" + `environment_id` + "`" + ` can be omitted e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher_volume.mysec 1se10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"rancher_certificate":         0,
		"rancher_environment":         1,
		"rancher_host":                2,
		"rancher_registration_token":  3,
		"rancher_registry":            4,
		"rancher_registry_credential": 5,
		"rancher_secret":              6,
		"rancher_stack":               7,
		"rancher_volume":              8,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
