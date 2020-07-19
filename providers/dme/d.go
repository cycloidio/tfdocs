package dme

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dme_transfer_acl",
			Category:         "Data Sources",
			ShortDescription: `Manages ACL (Access Control List) within the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) ACL Identifiable name. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(Optional) The list of IP addresses defined in the ACL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ips",
					Description: `(Optional) The list of IP addresses defined in the ACL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_contact_list",
			Category:         "Data Sources",
			ShortDescription: `Data source for Contact List`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Contact List action. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of contact list action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "emails",
					Description: `List of emails assigned in the contact list.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the dme calculated id of Contact list action.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of contact list action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "emails",
					Description: `List of emails assigned in the contact list.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the dme calculated id of Contact list action.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_dns_record",
			Category:         "Data Sources",
			ShortDescription: `Manages records in a domain within the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The record type. Values: A, AAAA, ANAME, CNAME, HTTPRED, MX, NS, PTR, SRV, TXT, CAA or SPF.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain id of the record. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of record. For A record Ipv4 address is required. For example: value: "1.2.3.4" For CNAME record alias name is required. For example: value: "www" For ANAME record FQDN is required. For example: value: "www.google.com." For MX record server name is required. For example: value: "document." For HTTPRED record URL is required. For example: value: "http://www.google.com" For TXT record text data is required. For example: value: "practice" For SPF record string value is required. For example: value: "1.2.3.4" For PTR record host name is required. For example: value: "mail.domainDocument." For NS record host name is required. For example: value: "mail.domainDocument." For AAAA record IPv6 address is required. For example: value: "0::0:0:0:0:0:6" For SRV record host name is required. For example: value: "mail.domainDocument." For CAA record text data is required. For example: value: "comodoca.com"`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live or TTL of the record.`,
				},
				resource.Attribute{
					Name:        "gtd_location",
					Description: `(Optional) Global Traffic Director location. Values:DEFAULT, US_EAST, US_WEST, EUROPE, ASIA_PAC, OCREANIA.`,
				},
				resource.Attribute{
					Name:        "dynamic_dns",
					Description: `(Optional) Indicates if the record has dynamic DNS enabled or not.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The per record password for a dynamic DNS update.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) For HTTP Redirection Records, A description of the HTTP Redirection Record`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `(Optional) For HTTPRED records. Keywords associated with the HTTPRED record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Optional) For HTTPRED records. The title of the HTTPRED record.`,
				},
				resource.Attribute{
					Name:        "redirect_type",
					Description: `(Optional) For HTTPRED records. Type of redirection performed. Values: Hidden Frame Masked, Standard - 301, Standard - 302.`,
				},
				resource.Attribute{
					Name:        "hardlink",
					Description: `(Optional) For HTTP Redirection Records.`,
				},
				resource.Attribute{
					Name:        "mx_level",
					Description: `(Optional) The priority for an MX record.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The weight for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "caa_type",
					Description: `(Optional) The type for an CAA Record. Caa type can be "issue", "issuewild", "iodef"`,
				},
				resource.Attribute{
					Name:        "issuer_critical",
					Description: `(Optional) The issuer critical value for an CAA Record. Value will be integer less than or equals to 255.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of record. For A record Ipv4 address is required. For example: value: "1.2.3.4" For CNAME record alias name is required. For example: value: "www" For ANAME record FQDN is required. For example: value: "www.google.com." For MX record server name is required. For example: value: "document." For HTTPRED record URL is required. For example: value: "http://www.google.com" For TXT record text data is required. For example: value: "practice" For SPF record string value is required. For example: value: "1.2.3.4" For PTR record host name is required. For example: value: "mail.domainDocument." For NS record host name is required. For example: value: "mail.domainDocument." For AAAA record IPv6 address is required. For example: value: "0::0:0:0:0:0:6" For SRV record host name is required. For example: value: "mail.domainDocument." For CAA record text data is required. For example: value: "comodoca.com"`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live or TTL of the record.`,
				},
				resource.Attribute{
					Name:        "gtd_location",
					Description: `(Optional) Global Traffic Director location. Values:DEFAULT, US_EAST, US_WEST, EUROPE, ASIA_PAC, OCREANIA.`,
				},
				resource.Attribute{
					Name:        "dynamic_dns",
					Description: `(Optional) Indicates if the record has dynamic DNS enabled or not.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The per record password for a dynamic DNS update.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) For HTTP Redirection Records, A description of the HTTP Redirection Record`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `(Optional) For HTTPRED records. Keywords associated with the HTTPRED record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Optional) For HTTPRED records. The title of the HTTPRED record.`,
				},
				resource.Attribute{
					Name:        "redirect_type",
					Description: `(Optional) For HTTPRED records. Type of redirection performed. Values: Hidden Frame Masked, Standard - 301, Standard - 302.`,
				},
				resource.Attribute{
					Name:        "hardlink",
					Description: `(Optional) For HTTP Redirection Records.`,
				},
				resource.Attribute{
					Name:        "mx_level",
					Description: `(Optional) The priority for an MX record.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The weight for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "caa_type",
					Description: `(Optional) The type for an CAA Record. Caa type can be "issue", "issuewild", "iodef"`,
				},
				resource.Attribute{
					Name:        "issuer_critical",
					Description: `(Optional) The issuer critical value for an CAA Record. Value will be integer less than or equals to 255.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for domain action`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of domain action. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of domain action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "gtd_enabled",
					Description: `(Optional) Indicator of whether or not this domain uses the Global Traffic Director service.`,
				},
				resource.Attribute{
					Name:        "soa_id",
					Description: `(Optional) The ID of a custom SOA record.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional) The ID of a template applied to the domain.`,
				},
				resource.Attribute{
					Name:        "vanity_id",
					Description: `(Optional) The ID of a vanity DNS configuration.`,
				},
				resource.Attribute{
					Name:        "transfer_acl_id",
					Description: `(Optional) The ID of an applied transfer ACL.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of a domain folder.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `(Optional) The number of seconds since the domain was last updated in Epoch time.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `(Optional) The number of seconds since the domain was last created in Epoch time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of domain action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "gtd_enabled",
					Description: `(Optional) Indicator of whether or not this domain uses the Global Traffic Director service.`,
				},
				resource.Attribute{
					Name:        "soa_id",
					Description: `(Optional) The ID of a custom SOA record.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional) The ID of a template applied to the domain.`,
				},
				resource.Attribute{
					Name:        "vanity_id",
					Description: `(Optional) The ID of a vanity DNS configuration.`,
				},
				resource.Attribute{
					Name:        "transfer_acl_id",
					Description: `(Optional) The ID of an applied transfer ACL.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of a domain folder.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `(Optional) The number of seconds since the domain was last updated in Epoch time.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `(Optional) The number of seconds since the domain was last created in Epoch time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_failover",
			Category:         "Data Sources",
			ShortDescription: `Manages failover in A record within the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "record_id",
					Description: `(Required) Record id of the record. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional) True indicates System Monitoring is Enabled. If monitor is enabled, system_description and max_emails value is required.`,
				},
				resource.Attribute{
					Name:        "system_description",
					Description: `(Optional) The system description of the failover configuration.`,
				},
				resource.Attribute{
					Name:        "max_emails",
					Description: `(Optional) The maximum number of emails to send per failover event.Value can be less than or equal to 150.`,
				},
				resource.Attribute{
					Name:        "sensitivity",
					Description: `(Optional) The number of checks placed against the primary IP before a Failover event occurs. List of Sensitivity ID’s:. Low (slower failover) = 8. Medium = 5. High = 3.`,
				},
				resource.Attribute{
					Name:        "protocol_id",
					Description: `(Optional)The protocol for DNS Failover to monitor on. List of Protocol IDs:. TCP = 1, UDP = 2, HTTP = 3, DNS = 4, SMTP = 5, HTTPS = 6.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port for the DNS Failover service to monitor on the specified protocol.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional) True indicates DNS Failover is enabled. If failover minimum 2 Ip address values are required.`,
				},
				resource.Attribute{
					Name:        "auto_failover",
					Description: `(Optional) True indicates the failback to the primary IP address is a manual process. False indicates the failback to the primary IP is an automatic process.`,
				},
				resource.Attribute{
					Name:        "ip1",
					Description: `(Optional) The primary IP address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `(Optional) The secondary IP address.`,
				},
				resource.Attribute{
					Name:        "ip3",
					Description: `(Optional) The teriary IP address.`,
				},
				resource.Attribute{
					Name:        "ip4",
					Description: `(Optional) The quaternary IP address.`,
				},
				resource.Attribute{
					Name:        "ip5",
					Description: `(Optional) The quinary IP address.`,
				},
				resource.Attribute{
					Name:        "contact_list",
					Description: `(Optional) The ID of the contact list for system monitoring notifications.`,
				},
				resource.Attribute{
					Name:        "http_fqdn",
					Description: `(Optional) The FQDN to monitor for HTTP or HTTPS checks.`,
				},
				resource.Attribute{
					Name:        "http_file",
					Description: `(Optional) The file to query for HTTP or HTTPS checks.`,
				},
				resource.Attribute{
					Name:        "http_query_string",
					Description: `(Optional) The string to query for HTTP or HTTPS checks.`,
				},
				resource.Attribute{
					Name:        "send_string",
					Description: `(Optional) The send string value for protocol_id = 2. The value will be string.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout value for the protocol_id = 2. The timeout value can be less than or equal to 7.`,
				},
				resource.Attribute{
					Name:        "dns_fqdn",
					Description: `(Optional) The string value for the protocol_id = 4.`,
				},
				resource.Attribute{
					Name:        "dns_timeout",
					Description: `(Optional) The timeout value for the protocol_id = 4. The timeout value can be less than 100.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional) True indicates System Monitoring is Enabled. If monitor is enabled, system_description and max_emails value is required.`,
				},
				resource.Attribute{
					Name:        "system_description",
					Description: `(Optional) The system description of the failover configuration.`,
				},
				resource.Attribute{
					Name:        "max_emails",
					Description: `(Optional) The maximum number of emails to send per failover event.Value can be less than or equal to 150.`,
				},
				resource.Attribute{
					Name:        "sensitivity",
					Description: `(Optional) The number of checks placed against the primary IP before a Failover event occurs. List of Sensitivity ID’s:. Low (slower failover) = 8. Medium = 5. High = 3.`,
				},
				resource.Attribute{
					Name:        "protocol_id",
					Description: `(Optional)The protocol for DNS Failover to monitor on. List of Protocol IDs:. TCP = 1, UDP = 2, HTTP = 3, DNS = 4, SMTP = 5, HTTPS = 6.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port for the DNS Failover service to monitor on the specified protocol.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional) True indicates DNS Failover is enabled. If failover minimum 2 Ip address values are required.`,
				},
				resource.Attribute{
					Name:        "auto_failover",
					Description: `(Optional) True indicates the failback to the primary IP address is a manual process. False indicates the failback to the primary IP is an automatic process.`,
				},
				resource.Attribute{
					Name:        "ip1",
					Description: `(Optional) The primary IP address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `(Optional) The secondary IP address.`,
				},
				resource.Attribute{
					Name:        "ip3",
					Description: `(Optional) The teriary IP address.`,
				},
				resource.Attribute{
					Name:        "ip4",
					Description: `(Optional) The quaternary IP address.`,
				},
				resource.Attribute{
					Name:        "ip5",
					Description: `(Optional) The quinary IP address.`,
				},
				resource.Attribute{
					Name:        "contact_list",
					Description: `(Optional) The ID of the contact list for system monitoring notifications.`,
				},
				resource.Attribute{
					Name:        "http_fqdn",
					Description: `(Optional) The FQDN to monitor for HTTP or HTTPS checks.`,
				},
				resource.Attribute{
					Name:        "http_file",
					Description: `(Optional) The file to query for HTTP or HTTPS checks.`,
				},
				resource.Attribute{
					Name:        "http_query_string",
					Description: `(Optional) The string to query for HTTP or HTTPS checks.`,
				},
				resource.Attribute{
					Name:        "send_string",
					Description: `(Optional) The send string value for protocol_id = 2. The value will be string.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout value for the protocol_id = 2. The timeout value can be less than or equal to 7.`,
				},
				resource.Attribute{
					Name:        "dns_fqdn",
					Description: `(Optional) The string value for the protocol_id = 4.`,
				},
				resource.Attribute{
					Name:        "dns_timeout",
					Description: `(Optional) The timeout value for the protocol_id = 4. The timeout value can be less than 100.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "folder_record",
			Category:         "Data Sources",
			ShortDescription: `Manages Custom Folder Records for the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Folder Record name ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Folder Record name`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Optional) A list of the primary domain IDs assigned to the folder`,
				},
				resource.Attribute{
					Name:        "secondaries",
					Description: `(Optional) A list of the secondary domain ID's assigned to the folder.`,
				},
				resource.Attribute{
					Name:        "default_folder",
					Description: `(Optional) Indicator of the folder being marked as the Default folder. Default value is false`,
				},
				resource.Attribute{
					Name:        "folder_permissions",
					Description: `(Optional) Permissions for the folder.`,
				},
				resource.Attribute{
					Name:        "folder_permissions.group_id",
					Description: `(Optional) This is static value for assigning group to the folder.`,
				},
				resource.Attribute{
					Name:        "folder_permissions.group_name",
					Description: `(Optional) This is static value for assigning group to the folder.`,
				},
				resource.Attribute{
					Name:        "folder_permissions.permission",
					Description: `(Optional) Assigning permissions for the folder. 4 for Read Only, 5 for Read and Create/Delete, 6 for Read and Edit and 7 for Read-Edit-Create/Delete all permissions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Folder Record name`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Optional) A list of the primary domain IDs assigned to the folder`,
				},
				resource.Attribute{
					Name:        "secondaries",
					Description: `(Optional) A list of the secondary domain ID's assigned to the folder.`,
				},
				resource.Attribute{
					Name:        "default_folder",
					Description: `(Optional) Indicator of the folder being marked as the Default folder. Default value is false`,
				},
				resource.Attribute{
					Name:        "folder_permissions",
					Description: `(Optional) Permissions for the folder.`,
				},
				resource.Attribute{
					Name:        "folder_permissions.group_id",
					Description: `(Optional) This is static value for assigning group to the folder.`,
				},
				resource.Attribute{
					Name:        "folder_permissions.group_name",
					Description: `(Optional) This is static value for assigning group to the folder.`,
				},
				resource.Attribute{
					Name:        "folder_permissions.permission",
					Description: `(Optional) Assigning permissions for the folder. 4 for Read Only, 5 for Read and Create/Delete, 6 for Read and Edit and 7 for Read-Edit-Create/Delete all permissions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_secondary_dns",
			Category:         "Data Sources",
			ShortDescription: `Data source for Secondary DNS action`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Secondary DNS action. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of domain action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "ipset_id",
					Description: `id of the Secondary Ip set which is associated with the secondary DNS.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `id of the Folder record which is associated with the secondary DNS.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the dme calculated id of secondary DNS action.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of domain action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "ipset_id",
					Description: `id of the Secondary Ip set which is associated with the secondary DNS.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `id of the Folder record which is associated with the secondary DNS.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the dme calculated id of secondary DNS action.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_secondary_ip_set",
			Category:         "Data Sources",
			ShortDescription: `Data source for Secondary IP Set action`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of secondary ip set action. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of secondary ip set action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `List of ip addresses assigned in the secondary ip set.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the dme calculated id of secondary Ip set action.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of secondary ip set action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `List of ip addresses assigned in the secondary ip set.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the dme calculated id of secondary Ip set action.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_custom_soa_record",
			Category:         "Data Sources",
			ShortDescription: `Manages Custom SOA Records for the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) SOA Record name ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) SOA Record name`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Contact email address.`,
				},
				resource.Attribute{
					Name:        "comp",
					Description: `(Optional) Primary name server.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL of the SOA Record (in seconds). TTl value should be greater than or equal to 21600`,
				},
				resource.Attribute{
					Name:        "refresh",
					Description: `(Optional) Zone refresh time (in seconds). Refresh value should be greater than or equal to 14400`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Optional) Starting zone serial number`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `(Optional) Failed Refresh retry time (in seconds). Retry value should be greater than or equal to 300`,
				},
				resource.Attribute{
					Name:        "expire",
					Description: `(Optional) Expire time of zone the (in seconds). Expire value should be greater than or equal to 86400`,
				},
				resource.Attribute{
					Name:        "negative_cache",
					Description: `(Optional) Record not found cache (in seconds). Negative Cache value should be greater than or equal to 180`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) SOA Record name`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Contact email address.`,
				},
				resource.Attribute{
					Name:        "comp",
					Description: `(Optional) Primary name server.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL of the SOA Record (in seconds). TTl value should be greater than or equal to 21600`,
				},
				resource.Attribute{
					Name:        "refresh",
					Description: `(Optional) Zone refresh time (in seconds). Refresh value should be greater than or equal to 14400`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Optional) Starting zone serial number`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `(Optional) Failed Refresh retry time (in seconds). Retry value should be greater than or equal to 300`,
				},
				resource.Attribute{
					Name:        "expire",
					Description: `(Optional) Expire time of zone the (in seconds). Expire value should be greater than or equal to 86400`,
				},
				resource.Attribute{
					Name:        "negative_cache",
					Description: `(Optional) Record not found cache (in seconds). Negative Cache value should be greater than or equal to 180`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for template action`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of domain action. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of domain action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "domain_ids",
					Description: `ids of the domain to which this template is associated.`,
				},
				resource.Attribute{
					Name:        "public_template",
					Description: `True represents a system defined public template rather than a user defined account specific template.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the dme calculated id of domain action.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of domain action. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "domain_ids",
					Description: `ids of the domain to which this template is associated.`,
				},
				resource.Attribute{
					Name:        "public_template",
					Description: `True represents a system defined public template rather than a user defined account specific template.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the dme calculated id of domain action.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_template_record",
			Category:         "Data Sources",
			ShortDescription: `Manages records in a template within the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The record type. Values: A, AAAA, ANAME, CNAME, HTTPRED, MX, NS, PTR, SRV, TXT, CAA or SPF.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Required) Template id of the record. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of record. For A record Ipv4 address is required. For example: value: "1.2.3.4" For CNAME record alias name is required. For example: value: "www" For ANAME record FQDN is required. For example: value: "www.google.com." For MX record server name is required. For example: value: "document." For HTTPRED record URL is required. For example: value: "http://www.google.com" For TXT record text data is required. For example: value: "practice" For SPF record string value is required. For example: value: "1.2.3.4" For PTR record host name is required. For example: value: "mail.domainDocument." For NS record host name is required. For example: value: "mail.domainDocument." For AAAA record IPv6 address is required. For example: value: "0::0:0:0:0:0:6" For SRV record host name is required. For example: value: "mail.domainDocument." For CAA record text data is required. For example: value: "comodoca.com"`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live or TTL of the record.`,
				},
				resource.Attribute{
					Name:        "gtd_location",
					Description: `(Optional) Global Traffic Director location. Values:DEFAULT, US_EAST, US_WEST, EUROPE, ASIA_PAC, OCREANIA.`,
				},
				resource.Attribute{
					Name:        "dynamic_dns",
					Description: `(Optional) Indicates if the record has dynamic DNS enabled or not.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The per record password for a dynamic DNS update.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) For HTTP Redirection Records, A description of the HTTP Redirection Record`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `(Optional) For HTTPRED records. Keywords associated with the HTTPRED record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Optional) For HTTPRED records. The title of the HTTPRED record.`,
				},
				resource.Attribute{
					Name:        "redirect_type",
					Description: `(Optional) For HTTPRED records. Type of redirection performed. Values: Hidden Frame Masked, Standard - 301, Standard - 302.`,
				},
				resource.Attribute{
					Name:        "hardlink",
					Description: `(Optional) For HTTP Redirection Records.`,
				},
				resource.Attribute{
					Name:        "mx_level",
					Description: `(Optional) The priority for an MX record.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The weight for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "caa_type",
					Description: `(Optional) The type for an CAA Record. Caa type can be "issue", "issuewild", "iodef"`,
				},
				resource.Attribute{
					Name:        "issuer_critical",
					Description: `(Optional) The issuer critical value for an CAA Record. Value will be integer less than or equals to 255.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of record. For A record Ipv4 address is required. For example: value: "1.2.3.4" For CNAME record alias name is required. For example: value: "www" For ANAME record FQDN is required. For example: value: "www.google.com." For MX record server name is required. For example: value: "document." For HTTPRED record URL is required. For example: value: "http://www.google.com" For TXT record text data is required. For example: value: "practice" For SPF record string value is required. For example: value: "1.2.3.4" For PTR record host name is required. For example: value: "mail.domainDocument." For NS record host name is required. For example: value: "mail.domainDocument." For AAAA record IPv6 address is required. For example: value: "0::0:0:0:0:0:6" For SRV record host name is required. For example: value: "mail.domainDocument." For CAA record text data is required. For example: value: "comodoca.com"`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live or TTL of the record.`,
				},
				resource.Attribute{
					Name:        "gtd_location",
					Description: `(Optional) Global Traffic Director location. Values:DEFAULT, US_EAST, US_WEST, EUROPE, ASIA_PAC, OCREANIA.`,
				},
				resource.Attribute{
					Name:        "dynamic_dns",
					Description: `(Optional) Indicates if the record has dynamic DNS enabled or not.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The per record password for a dynamic DNS update.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) For HTTP Redirection Records, A description of the HTTP Redirection Record`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `(Optional) For HTTPRED records. Keywords associated with the HTTPRED record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Optional) For HTTPRED records. The title of the HTTPRED record.`,
				},
				resource.Attribute{
					Name:        "redirect_type",
					Description: `(Optional) For HTTPRED records. Type of redirection performed. Values: Hidden Frame Masked, Standard - 301, Standard - 302.`,
				},
				resource.Attribute{
					Name:        "hardlink",
					Description: `(Optional) For HTTP Redirection Records.`,
				},
				resource.Attribute{
					Name:        "mx_level",
					Description: `(Optional) The priority for an MX record.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The weight for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port for an SRV Record.`,
				},
				resource.Attribute{
					Name:        "caa_type",
					Description: `(Optional) The type for an CAA Record. Caa type can be "issue", "issuewild", "iodef"`,
				},
				resource.Attribute{
					Name:        "issuer_critical",
					Description: `(Optional) The issuer critical value for an CAA Record. Value will be integer less than or equals to 255.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dme_vanity_nameserver_record",
			Category:         "Data Sources",
			ShortDescription: `Manages Custom Vanity Name Server Records for the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) SOA Record name ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) SOA Record name`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Optional) The vanity host names defined in the configuration.`,
				},
				resource.Attribute{
					Name:        "public_config",
					Description: `(Optional) True represents a system defined rather than user defined vanity configuration. Default is false.`,
				},
				resource.Attribute{
					Name:        "default_config",
					Description: `(Optional) Indicates if the vanity configuration is the system default. Default is false.`,
				},
				resource.Attribute{
					Name:        "name_server_group_id",
					Description: `(Optional) The name server group the configuration is assigned`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) SOA Record name`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Optional) The vanity host names defined in the configuration.`,
				},
				resource.Attribute{
					Name:        "public_config",
					Description: `(Optional) True represents a system defined rather than user defined vanity configuration. Default is false.`,
				},
				resource.Attribute{
					Name:        "default_config",
					Description: `(Optional) Indicates if the vanity configuration is the system default. Default is false.`,
				},
				resource.Attribute{
					Name:        "name_server_group_id",
					Description: `(Optional) The name server group the configuration is assigned`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"dme_transfer_acl":             0,
		"dme_contact_list":             1,
		"dme_dns_record":               2,
		"dme_domain":                   3,
		"dme_failover":                 4,
		"folder_record":                5,
		"dme_secondary_dns":            6,
		"dme_secondary_ip_set":         7,
		"dme_custom_soa_record":        8,
		"dme_template":                 9,
		"dme_template_record":          10,
		"dme_vanity_nameserver_record": 11,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
