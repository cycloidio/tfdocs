package constellix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "constellix_dns_check",
			Category:         "Resources",
			ShortDescription: `Manages one or more DNS check resource`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Required) A website address. It can be set only once`,
				},
				resource.Attribute{
					Name:        "resolver",
					Description: `(Required) A website address. It can be set only once`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Required) Site ids to check.`,
				},
				resource.Attribute{
					Name:        "notification_groups",
					Description: `(Optional) List of group IDs for the notification group of DNS Check.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check Interval. Allowed values are ` + "`" + `THIRTYSECONDS` + "`" + `, ` + "`" + `ONEMINUTE` + "`" + `, ` + "`" + `TWOMINUTES` + "`" + `, ` + "`" + `THREEMINUTES` + "`" + `, ` + "`" + `FOURMINUTES` + "`" + `, ` + "`" + `FIVEMINUTES` + "`" + `, ` + "`" + `TENMINUTES` + "`" + `, ` + "`" + `THIRTYMINUTES` + "`" + `, ` + "`" + `HALFDAY` + "`" + ` and ` + "`" + `DAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval_policy",
					Description: `(Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are ` + "`" + `PARALLEL` + "`" + `, ` + "`" + `ONCEPERSITE` + "`" + ` and ` + "`" + `ONCEPERREGION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_policy",
					Description: `(Optional) Specifies how the check should be validated. Allowed values are ` + "`" + `SIMPLE` + "`" + ` and ` + "`" + `MAJORITY` + "`" + `. This parameter will only work with the ` + "`" + `interval_policy` + "`" + ` set to ` + "`" + `PARALLEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expected_response",
					Description: `(Optional) Ip Address where DNS provided in the FQDN should resolved to in ideal conditions. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of DNS check resource. ## Importing ## An existing Check can be [imported][docs-import] into this resource using its Id, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import constellix_dns_check.example <check-id> ` + "`" + `` + "`" + `` + "`" + ` Where check-id is the Id of check calculated via Constellix API.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_domain",
			Category:         "Resources",
			ShortDescription: `Manages one or more domains within the account.`,
			Description:      ``,
			Keywords: []string{
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the domain.`,
				},
				resource.Attribute{
					Name:        "has_gtd_regions",
					Description: `(Optional) GTD Region status of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "has_geoip",
					Description: `(Optional) GTD Region status of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "vanity_nameserver",
					Description: `(Optional) vanity nameserver of domain.`,
				},
				resource.Attribute{
					Name:        "nameserver_group",
					Description: `(Optional) Shows the nameserver group of domain. The Default nameserverGroup is 1.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Notes while creating the domain. The maximum length will be 1000 characters.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Id of tags applied on domain. The default value is empty.`,
				},
				resource.Attribute{
					Name:        "soa",
					Description: `(Optional) Object`,
				},
				resource.Attribute{
					Name:        "soa.primary_nameserver",
					Description: `(Optional) The Primary Nameserver is of SOA.`,
				},
				resource.Attribute{
					Name:        "soa.email",
					Description: `(Optional) An Email Address specifies the mailbox of the person responsible for this zone.`,
				},
				resource.Attribute{
					Name:        "soa.ttl",
					Description: `(Optional) The number of seconds that this SOA record will be cached in other resolving name servers.`,
				},
				resource.Attribute{
					Name:        "soa.refresh",
					Description: `(Optional) The time interval (in seconds) before the zone should be refreshed. The recommended value – 86400 (24 Hours).`,
				},
				resource.Attribute{
					Name:        "soa.retry",
					Description: `(Optional) The time interval (in seconds) before a failed refresh should be retried. Recommended value – 7200 (2 Hours).`,
				},
				resource.Attribute{
					Name:        "soa.expire",
					Description: `(Optional) The time internal (in seconds) that specifies the upper limit on the time internally that can elapse before the zone is no longer authoritative. This is when the secondary name servers will expire if they are unable to refresh. Recommended value – up to 1209600`,
				},
				resource.Attribute{
					Name:        "soa.negcache",
					Description: `(Optional) The amount of time a record not found is cached. Recommended values can vary, between 180 and 172800 (3 min – 2 days). ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "soa.serial",
					Description: `(Optional) The starting serial number for the version of the zone. If the SOA record is applied to a domain that is already created (and thus already has a starting serial number), the existing serial number will be incremented by one. e.g 2015010196 ## Importing ## An existing Record can be [imported][docs-import] into this resource using its Id, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import constellix_domain.example <record-id> ` + "`" + `` + "`" + `` + "`" + ` Where record-id is the Id of record calculated via Constellix API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "soa.serial",
					Description: `(Optional) The starting serial number for the version of the zone. If the SOA record is applied to a domain that is already created (and thus already has a starting serial number), the existing serial number will be incremented by one. e.g 2015010196 ## Importing ## An existing Record can be [imported][docs-import] into this resource using its Id, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import constellix_domain.example <record-id> ` + "`" + `` + "`" + `` + "`" + ` Where record-id is the Id of record calculated via Constellix API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_http_check",
			Category:         "Resources",
			ShortDescription: `Manages one or more HTTP check resource`,
			Description:      ``,
			Keywords: []string{
				"http",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Required) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Required) Specifies upper layer protocol like HTTP, HTTPs, etc.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Required) Site ids to check.`,
				},
				resource.Attribute{
					Name:        "notification_groups",
					Description: `(Optional) List of group IDs for the notification group of HTTP Check.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check Interval. Allowed values are ` + "`" + `THIRTYSECONDS` + "`" + `, ` + "`" + `ONEMINUTE` + "`" + `, ` + "`" + `TWOMINUTES` + "`" + `, ` + "`" + `THREEMINUTES` + "`" + `, ` + "`" + `FOURMINUTES` + "`" + `, ` + "`" + `FIVEMINUTES` + "`" + `, ` + "`" + `TENMINUTES` + "`" + `, ` + "`" + `THIRTYMINUTES` + "`" + `, ` + "`" + `HALFDAY` + "`" + ` and ` + "`" + `DAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval_policy",
					Description: `(Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are ` + "`" + `PARALLEL` + "`" + `, ` + "`" + `ONCEPERSITE` + "`" + ` and ` + "`" + `ONCEPERREGION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_policy",
					Description: `(Optional) Specifies how the check should be validated. Allowed values are ` + "`" + `SIMPLE` + "`" + ` and ` + "`" + `MAJORITY` + "`" + `. This parameter will only work with the ` + "`" + `interval_policy` + "`" + ` set to ` + "`" + `PARALLEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Optional) Fully qualified domain name of the URL should be checked.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) In case of multi-page site, which path should be checked.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Optional) String to search in the first 2KB of resonse received.`,
				},
				resource.Attribute{
					Name:        "expected_status_code",
					Description: `(Optional) Expected HTTP status code for this check. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of HTTP check resource. ## Importing ## An existing check can be [imported][docs-import] into this resource using its Id, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import constellix_http_check.example <check-id> ` + "`" + `` + "`" + `` + "`" + ` Where check-id is the Id of check calculated via Constellix API.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_tags",
			Category:         "Resources",
			ShortDescription: `Manages tags for Constellix DNS.`,
			Description:      ``,
			Keywords: []string{
				"tags",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique. ## Attribute Reference ## No attributes are exported ## Importing ## An existing Tag can be [imported][docs-import] into this resource using its Id, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import constellix_tags.example <tag-id> ` + "`" + `` + "`" + `` + "`" + ` Where tag-id is the Id of tag calculated via Constellix API.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_tcp_check",
			Category:         "Resources",
			ShortDescription: `Manages one or more TCP check resource`,
			Description:      ``,
			Keywords: []string{
				"tcp",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Required) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Required) Site ids to check.`,
				},
				resource.Attribute{
					Name:        "notification_groups",
					Description: `(Optional) List of group IDs for the notification group of TCP Check.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check Interval. Allowed values are ` + "`" + `THIRTYSECONDS` + "`" + `, ` + "`" + `ONEMINUTE` + "`" + `, ` + "`" + `TWOMINUTES` + "`" + `, ` + "`" + `THREEMINUTES` + "`" + `, ` + "`" + `FOURMINUTES` + "`" + `, ` + "`" + `FIVEMINUTES` + "`" + `, ` + "`" + `TENMINUTES` + "`" + `, ` + "`" + `THIRTYMINUTES` + "`" + `, ` + "`" + `HALFDAY` + "`" + ` and ` + "`" + `DAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval_policy",
					Description: `(Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are ` + "`" + `PARALLEL` + "`" + `, ` + "`" + `ONCEPERSITE` + "`" + ` and ` + "`" + `ONCEPERREGION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_policy",
					Description: `(Optional) Specifies how the check should be validated. Allowed values are ` + "`" + `SIMPLE` + "`" + ` and ` + "`" + `MAJORITY` + "`" + `. This parameter will only work with the ` + "`" + `interval_policy` + "`" + ` set to ` + "`" + `PARALLEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "string_to_send",
					Description: `(Optional) String to send along with the check. It can be any parameter to the endpoint.`,
				},
				resource.Attribute{
					Name:        "string_to_receive",
					Description: `(Optional) String which should be received as a result of TCP check. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of TCP check resource. ## Importing ## An existing Check can be [imported][docs-import] into this resource using its Id, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import constellix_tcp_check.example <check-id> ` + "`" + `` + "`" + `` + "`" + ` Where check-id is the Id of check calculated via Constellix API.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_template",
			Category:         "Resources",
			ShortDescription: `Manages the DNS record templates.`,
			Description:      ``,
			Keywords: []string{
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Template names. e.g "sampletemplate".`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Id of domain to be applied.`,
				},
				resource.Attribute{
					Name:        "has_gtd_regions",
					Description: `(Optional) Enable/Disable GTD Region of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "has_geoip",
					Description: `(Optional) Enable/Disable GEO IP. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) System generated template history version. ## Attributes Reference No attributes are exported. ## Importing ## An existing Template can be [imported][docs-import] into this resource using its Id, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import constellix_template.example <template-id> ` + "`" + `` + "`" + `` + "`" + ` Where template-id is the Id of template calculated via Constellix API.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"constellix_dns_check":  0,
		"constellix_domain":     1,
		"constellix_http_check": 2,
		"constellix_tags":       3,
		"constellix_tcp_check":  4,
		"constellix_template":   5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
