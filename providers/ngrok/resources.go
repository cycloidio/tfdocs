package ngrok

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ngrok_api_key",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_certificate_authority",
			Category:         "Resources",
			ShortDescription: `Certificate Authorities are x509 certificates that are used to sign other x509 certificates. Attach a Certificate Authority to the Mutual TLS module to verify that the TLS certificate presented by a client has been signed by this CA. Certificate Authorities are used only for mTLS validation only and thus a private key is not included in the resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_credential",
			Category:         "Resources",
			ShortDescription: `Tunnel Credentials are ngrok agent authtokens. They authorize the ngrok agent to connect the ngrok service as your account. They are installed with the ngrok authtoken command or by specifying it in the ngrok.yml configuration file with the authtoken property.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_endpoint_configuration",
			Category:         "Resources",
			ShortDescription: `Endpoint Configurations are a reusable group of modules that encapsulate how traffic to a domain or address is handled. Endpoint configurations are only applied to Domains and TCP Addresses they have been attached to.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_event_destination",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_event_subscription",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_ip_policy",
			Category:         "Resources",
			ShortDescription: `IP Policies are reusable groups of CIDR ranges with an allow or deny action. They can be attached to endpoints via the Endpoint Configuration IP Policy module. They can also be used with IP Restrictions to control source IP ranges that can start tunnel sessions and connect to the API and dashboard.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_ip_policy_rule",
			Category:         "Resources",
			ShortDescription: `IP Policy Rules are the IPv4 or IPv6 CIDRs entries that make up an IP Policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_ip_restriction",
			Category:         "Resources",
			ShortDescription: `An IP restriction is a restriction placed on the CIDRs that are allowed to initate traffic to a specific aspect of your ngrok account. An IP restriction has a type which defines the ingress it applies to. IP restrictions can be used to enforce the source IPs that can make API requests, log in to the dashboard, start ngrok agents, and connect to your public-facing endpoints.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_ip_whitelist_entry",
			Category:         "Resources",
			ShortDescription: `The IP Whitelist is deprecated and will be removed. Use an IP Restriction with an endpoints type instead.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_reserved_addr",
			Category:         "Resources",
			ShortDescription: `Reserved Addresses are TCP addresses that can be used to listen for traffic. TCP address hostnames and ports are assigned by ngrok, they cannot be chosen.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_reserved_domain",
			Category:         "Resources",
			ShortDescription: `Reserved Domains are hostnames that you can listen for traffic on. Domains can be used to listen for http, https or tls traffic. You may use a domain that you own by creating a CNAME record specified in the returned resource. This CNAME record points traffic for that domain to ngrok's edge servers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_ssh_certificate_authority",
			Category:         "Resources",
			ShortDescription: `An SSH Certificate Authority is a pair of an SSH Certificate and its private key that can be used to sign other SSH host and user certificates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_ssh_credential",
			Category:         "Resources",
			ShortDescription: `SSH Credentials are SSH public keys that can be used to start SSH tunnels via the ngrok SSH tunnel gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_ssh_host_certificate",
			Category:         "Resources",
			ShortDescription: `SSH Host Certificates along with the corresponding private key allows an SSH server to assert its authenticity to connecting SSH clients who trust the SSH Certificate Authority that was used to sign the certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_ssh_user_certificate",
			Category:         "Resources",
			ShortDescription: `SSH User Certificates are presented by SSH clients when connecting to an SSH server to authenticate their connection. The SSH server must trust the SSH Certificate Authority used to sign the certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ngrok_tls_certificate",
			Category:         "Resources",
			ShortDescription: `TLS Certificates are pairs of x509 certificates and their matching private key that can be used to terminate TLS traffic. TLS certificates are unused until they are attached to a Domain. TLS Certificates may also be provisioned by ngrok automatically for domains on which you have enabled automated certificate provisioning.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"ngrok_api_key":                   0,
		"ngrok_certificate_authority":     1,
		"ngrok_credential":                2,
		"ngrok_endpoint_configuration":    3,
		"ngrok_event_destination":         4,
		"ngrok_event_subscription":        5,
		"ngrok_ip_policy":                 6,
		"ngrok_ip_policy_rule":            7,
		"ngrok_ip_restriction":            8,
		"ngrok_ip_whitelist_entry":        9,
		"ngrok_reserved_addr":             10,
		"ngrok_reserved_domain":           11,
		"ngrok_ssh_certificate_authority": 12,
		"ngrok_ssh_credential":            13,
		"ngrok_ssh_host_certificate":      14,
		"ngrok_ssh_user_certificate":      15,
		"ngrok_tls_certificate":           16,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
