package constellix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "constellix_a_record",
			Category:         "Resources",
			ShortDescription: `Manages one or more A records within the specified domain.`,
			Description:      ``,
			Keywords: []string{
				"a",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the A record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Required) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Required) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "geo_location",
					Description: `(Optional) Details of IP filter / Geo proximity to be applied. Default is null.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_user_region",
					Description: `(Optional) For Geo proximity to be applied. geoipUserRegion should not be provided.`,
				},
				resource.Attribute{
					Name:        "geo_location.drop",
					Description: `(Optional) drop flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_proximity",
					Description: `(Optional) a valid geoipProximity id.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_user_region",
					Description: `(Optional) For Geo IP Filter to be applied. geoipUserRegion should be [1].`,
				},
				resource.Attribute{
					Name:        "geo_location.drop",
					Description: `(Optional) drop flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_proximity",
					Description: `(Optional) for Geo IP Filter, geoipProximity must not be provided. please create an A record with "World (Default)" IP Filter first before a more specific IP Filter is applied. The "World (Default)" record would only be used if no matching Filter or Proximity records are found.`,
				},
				resource.Attribute{
					Name:        "record_option",
					Description: `(Optional) Type of record. "roundRobin" for Standard record (Default). "failover" for Failover. "pools" for Pools. "roundRobinFailover" for Round Robin with Failover.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active).`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional)Record note.`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type A.`,
				},
				resource.Attribute{
					Name:        "contact_ids",
					Description: `(Optional) Applied contact list id. Only applicable to record with type roundRobin with failover and failover.`,
				},
				resource.Attribute{
					Name:        "pools",
					Description: `(Optional) Ids of Apool.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover",
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.value",
					Description: `(Required) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.disable_flag",
					Description: `(Required) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.sort_order",
					Description: `(Required) Integer value which decides in which order the rounrobinfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover",
					Description: `(Optional) To create a record failover object pass the following attributes.`,
				},
				resource.Attribute{
					Name:        "record_failover_values",
					Description: `(Required) Set.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Required) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Required) Integer value which decides in which order the recordfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Required) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Required) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Required) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false. ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the A resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_aaaa_record",
			Category:         "Resources",
			ShortDescription: `Manages one or more AAAA records within the specified domain.`,
			Description:      ``,
			Keywords: []string{
				"aaaa",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the AAAA record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Required) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Required) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "geo_location",
					Description: `(Optional) Details of IP filter / Geo proximity to be applied. Default is null.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_user_region",
					Description: `(Optional) For Geo proximity to be applied. geoipUserRegion should not be provided.`,
				},
				resource.Attribute{
					Name:        "geo_location.drop",
					Description: `(Optional) drop flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_proximity",
					Description: `(Optional) a valid geoipProximity id.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_user_region",
					Description: `(Optional) For Geo IP Filter to be applied. geoipUserRegion should be [1].`,
				},
				resource.Attribute{
					Name:        "geo_location.drop",
					Description: `(Optional) drop flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_proximity",
					Description: `(Optional) for Geo IP Filter, geoipProximity must not be provided. please create an A record with "World (Default)" IP Filter first before a more specific IP Filter is applied. The "World (Default)" record would only be used if no matching Filter or Proximity records are found.`,
				},
				resource.Attribute{
					Name:        "record_option",
					Description: `(Optional) Type of record. "roundRobin" for Standard record (Default). "failover" for Failover. "pools" for Pools. "roundRobinFailover" for Round Robin with Failover.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active).`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional)Record note.`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type AAAA.`,
				},
				resource.Attribute{
					Name:        "contact_ids",
					Description: `(Optional) Applied contact list id. Only applicable to record with type roundRobin with failover and failover.`,
				},
				resource.Attribute{
					Name:        "pools",
					Description: `(Optional) Ids of AAAApool.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover",
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.value",
					Description: `(Required) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.disable_flag",
					Description: `(Required) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.sort_order",
					Description: `(Required) Integer value which decides in which order the roundrobinfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover",
					Description: `(Optional) To create a record failover object pass the following attributes.`,
				},
				resource.Attribute{
					Name:        "record_failover_values",
					Description: `(Required) Set.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Required) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Required) Integer value which decides in which order the recordfailover should be sorted`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Required) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Required) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Required) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false. ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the AAAA resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_aaaa_record_pool",
			Category:         "Resources",
			ShortDescription: `Manages the pools of AAAA records.`,
			Description:      ``,
			Keywords: []string{
				"aaaa",
				"record",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Pool name should be unique.`,
				},
				resource.Attribute{
					Name:        "num_return",
					Description: `(Required) minimum number of value object to return. Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "min_available_failover",
					Description: `(Required)minimum number of Available Failover . Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "failed_flag",
					Description: `(Optional) failed flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values`,
				},
				resource.Attribute{
					Name:        "values.value",
					Description: `(Required) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "values.weight",
					Description: `(Required) weight number to sort the priorty. Weight must be in between 1 and 1000000.`,
				},
				resource.Attribute{
					Name:        "values.disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values.checkid",
					Description: `(Optional) Sonar check id is required when you want to apply the ITO feature on a pool.`,
				},
				resource.Attribute{
					Name:        "values.policy",
					Description: `(Required) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Description. ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the aaaa record pool resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_aname_record",
			Category:         "Resources",
			ShortDescription: `Manages one or more domain ANAME record.`,
			Description:      ``,
			Keywords: []string{
				"aname",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Required) Host name. If "Host" value does not end in a dot, your domain name will be appended to it.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Required) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "record_option",
					Description: `(Optional) Type of record. "roundRobin" for Standard record (Default). "failover" for Failover`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active)`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Record note`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type ANAME`,
				},
				resource.Attribute{
					Name:        "contact_ids",
					Description: `(Optional) Applied contact list id. Only applicable to record with type failover.`,
				},
				resource.Attribute{
					Name:        "record_failover",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "record_failover_values",
					Description: `(Required) Set`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Required) Host name`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Required) Enable or Disable the recordfailover values object. Default is false. Atleast one object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Required) Integer value which decides in which order recordfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Optional) 1 for Normal (always lowest level), 2 for Off on any Failover event, 3 for One Way (move to higher level)`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Optional) Enable or Disable the recordfailover object. Default is false. Atleast one recordfailover object should be false. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of aname resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_a_record_pool",
			Category:         "Resources",
			ShortDescription: `Manages the pools of A records.`,
			Description:      ``,
			Keywords: []string{
				"a",
				"record",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Pool name should be unique.`,
				},
				resource.Attribute{
					Name:        "num_return",
					Description: `(Required) minimum number of value object to return. Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "min_available_failover",
					Description: `(Required)minimum number of Available Failover . Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) System generated version number.`,
				},
				resource.Attribute{
					Name:        "failed_flag",
					Description: `(Optional) failed flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values`,
				},
				resource.Attribute{
					Name:        "values.value",
					Description: `(Required) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "values.weight",
					Description: `(Required) weight number to sort the priorty. Weight must be in between 1 and 1000000.`,
				},
				resource.Attribute{
					Name:        "values.disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values.check_id",
					Description: `(Optional) Sonar check id is required when you want to apply the ITO feature on a pool.`,
				},
				resource.Attribute{
					Name:        "values.policy",
					Description: `(Required) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Description. ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the A record pool resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_caa_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type CAA for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"caa",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.caa_provider_id",
					Description: `(Required) 1 for [ Custom ], 2 for [ No Provider ], 3 for Comodo, 4 for Digicert, 5 for Entrust, 6 for GeoTrust, 7 for Izenpe, 8 for Lets Encrypt, 9 for Symantec, 10 for Thawte`,
				},
				resource.Attribute{
					Name:        "roundrobin.tag",
					Description: `(Required) "issue" for Issue, "IssueWild" for IssueWild, "iodef" for iodef. Type allows you to choose how you want certificates to be issued by the CA. Each CAA record can contain only one tag-value pair. Options: issue: Explicitly authorizes a single certificate authority to issue a certificate (any type) for the hostname. issuewild: Authorization to issue certificates that specify a wildcard domain. Please note: issuewild properties take precedence over issue properties when specified. iodef: (Incident Description Exchange Format) Specifies a means of reporting certificate issue requests or cases of certificate issue for the corresponding domain that violate the security policy of the issuer or the domain name holder.`,
				},
				resource.Attribute{
					Name:        "roundrobin.data",
					Description: `(Required) "" for [ Custom ] if CAA provider Id is 1, ";" for [ No Provider ], "comodoca.com" for Comodo, "digicert.com" for Digicert, "entrust.net" for Entrust, "geotrust.com" for GeoTrust, "izenpe.com" for Izenpe, "letsencrypt.org" for Lets Encrypt, "symantec.com" for Symantec, "thawte.com" for Thawte`,
				},
				resource.Attribute{
					Name:        "roundrobin.flag",
					Description: `(Required) roundRobin.Issuer Critical There is currently only one flag defined, “issuer critical” at a value of 1. If a CA does not understand the flag value for an issuer critical record, then the CA will return with “no issue” for the certification. All records will have the default issuer critical value of 0, which means they are “not critical”. Issuer Critical Value should be between 0 to 255.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Required) Disable flag. Default is false`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active)`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Record note`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type CAA ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of caa resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_cert_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type CERT for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"cert",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the CERT record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active).`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional)Record note.`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.certificate_type",
					Description: `(Required) certificateType 0 - 65,535`,
				},
				resource.Attribute{
					Name:        "roundrobin.key_tag",
					Description: `(Required) 0 - 65,535`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) disable flag. Default is false`,
				},
				resource.Attribute{
					Name:        "roundrobin.certificate",
					Description: `(Required) certificate.`,
				},
				resource.Attribute{
					Name:        "roundrobin.algorithm",
					Description: `(Required) 0-255.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type CERT. ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the CERT resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_cname_record",
			Category:         "Resources",
			ShortDescription: `Manages one or more domain CNAME records.`,
			Description:      ``,
			Keywords: []string{
				"cname",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the CName record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "geo_location",
					Description: `(Optional) Details of IP filter / Geo proximity to be applied. Default is null.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_user_region",
					Description: `(Optional) For Geo proximity to be applied. geoipUserRegion should not be provided.`,
				},
				resource.Attribute{
					Name:        "geo_location.drop",
					Description: `(Optional) drop flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_proximity",
					Description: `(Optional) a valid geoipProximity id.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_user_region",
					Description: `(Optional) For Geo IP Filter to be applied. geoipUserRegion should be [1].`,
				},
				resource.Attribute{
					Name:        "geo_location.drop",
					Description: `(Optional) drop flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "geo_location.geo_ip_proximity",
					Description: `(Optional) for Geo IP Filter, geoipProximity must not be provided. please create an A record with "World (Default)" IP Filter first before a more specific IP Filter is applied. The "World (Default)" record would only be used if no matching Filter or Proximity records are found.`,
				},
				resource.Attribute{
					Name:        "record_option",
					Description: `(Optional) Type of record. "roundRobin" for Standard record (Default). "failover" for Failover. "pools" for Pools. "roundRobinFailover" for Round Robin with Failover.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active).`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional)Record note.`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type A.`,
				},
				resource.Attribute{
					Name:        "contact_ids",
					Description: `(Optional) Applied contact list id. Only applicable to record with type roundRobin with failover and failover.`,
				},
				resource.Attribute{
					Name:        "pools",
					Description: `(Optional) Ids of CNamepool.`,
				},
				resource.Attribute{
					Name:        "record_failover",
					Description: `(Optional) To create a record failover object pass the following attributes.`,
				},
				resource.Attribute{
					Name:        "record_failover_values",
					Description: `(Required) Set.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Required) Host name.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Required) Integer value which decides in which order the recordfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Required) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Required) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Required) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false. ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the CName resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_cname_record_pool",
			Category:         "Resources",
			ShortDescription: `Manages the pools of CNAME records.`,
			Description:      ``,
			Keywords: []string{
				"cname",
				"record",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Pool name should be unique.`,
				},
				resource.Attribute{
					Name:        "num_return",
					Description: `(Required) minimum number of value object to return. Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "min_available_failover",
					Description: `(Required) minimum number of Available Failover . Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) System generated version number.`,
				},
				resource.Attribute{
					Name:        "failed_flag",
					Description: `(Optional) failed flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values`,
				},
				resource.Attribute{
					Name:        "values.value",
					Description: `(Required) Host name. If "Host" value does not end in a dot, your domain name will be appended to it.`,
				},
				resource.Attribute{
					Name:        "values.weight",
					Description: `(Required)weight number to sort the priorty. Weight must be in between 1 and 1000000`,
				},
				resource.Attribute{
					Name:        "values.disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values.check_id",
					Description: `(Optional) Sonar check id is required when you want to apply the ITO feature on a pool.`,
				},
				resource.Attribute{
					Name:        "values.policy",
					Description: `(Required) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Description. ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the cname record pool resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_contact_lists",
			Category:         "Resources",
			ShortDescription: `Manages Constellix DNS contact lists.`,
			Description:      ``,
			Keywords: []string{
				"contact",
				"lists",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "email_addresses",
					Description: `(Required) List of email addresses ## Attribute Reference ## No attributes are exported`,
				},
			},
			Attributes: []resource.Attribute{},
		},
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
					Description: `(Optional) Ip Address where DNS provided in the FQDN should resolved to in ideal conditions. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of DNS check resource.`,
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
					Description: `(Optional) The Default value of SOA Primary Nameserver is "ns0.constellix.com.". However, it is possible to create a custom SOA record with differing values if required.`,
				},
				resource.Attribute{
					Name:        "soa.email",
					Description: `(Optional) An Email Address specifies the mailbox of the person responsible for this zone. The default value is "dns.constellix.com."`,
				},
				resource.Attribute{
					Name:        "soa.ttl",
					Description: `(Optional) The number of seconds that this SOA record will be cached in other resolving name servers. The Default value is "86400".TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "soa.refresh",
					Description: `(Optional) The time interval (in seconds) before the zone should be refreshed. The recommended value – 86400 (24 Hours). The default value is 43200 (12 hours)`,
				},
				resource.Attribute{
					Name:        "soa.serial",
					Description: `(Optional) The starting serial number for the version of the zone. If the SOA record is applied to a domain that is already created (and thus already has a starting serial number), the existing serial number will be incremented by one. e.g 2015010196`,
				},
				resource.Attribute{
					Name:        "soa.retry",
					Description: `(Optional) The time interval (in seconds) before a failed refresh should be retried. Recommended value – 7200 (2 Hours). The default value is 1 hour`,
				},
				resource.Attribute{
					Name:        "soa.expire",
					Description: `(Optional) The time internal (in seconds) that specifies the upper limit on the time internally that can elapse before the zone is no longer authoritative. This is when the secondary name servers will expire if they are unable to refresh. Recommended value – up to 1209600`,
				},
				resource.Attribute{
					Name:        "soa.negcache",
					Description: `(Optional) The amount of time a record not found is cached. Recommended values can vary, between 180 and 172800 (3 min – 2 days). The default value is 180 ## Attribute Reference ## No attributes are exported`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_geo_filter",
			Category:         "Resources",
			ShortDescription: `Manage Geofilters for A, AAAA, CNAME or ANAME records.`,
			Description:      ``,
			Keywords: []string{
				"geo",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Geo Filter name should be unique.`,
				},
				resource.Attribute{
					Name:        "geoip_continents",
					Description: `(Optional) Two digit Continents Code.`,
				},
				resource.Attribute{
					Name:        "geoip_countries",
					Description: `(Optional)Two digit Countries Code.`,
				},
				resource.Attribute{
					Name:        "geoip_regions",
					Description: `(Optional) Two digit country code followed by "/" followed by two digit region code.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Optional) Autonomous System Number (ASN). ASN code should be a number between 0 and 4294967295.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) IPV4 Address.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) IPV6 Address.`,
				},
				resource.Attribute{
					Name:        "filter_rules_limit",
					Description: `(Optional) Default is 100. For more than 100 rules, parameter should be set explicitly for ADD and Update API calls. Value should be in mulitple of 100 like 200, 300 ...upto the quota limit assigned to the account. Check quota details for IP Filter Rule Limit. ## Attributes Reference No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_geo_proximity",
			Category:         "Resources",
			ShortDescription: `Manage Geoproximity for A, AAAA, CNAME or ANAME records.`,
			Description:      ``,
			Keywords: []string{
				"geo",
				"proximity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Geo Proximity name should be unique.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) Country code. Default is null.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional)Region or state or province code. Default is null.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Optional) Latitude value.`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Optional) Longitude value.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `(Optional)City code. Default is null. ## Attributes Reference No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_hinfo_record",
			Category:         "Resources",
			ShortDescription: `Manages one or more HINFO records.`,
			Description:      ``,
			Keywords: []string{
				"hinfo",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.cpu",
					Description: `(Required) A description of basic system hardware`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin.os",
					Description: `(Required) A description of the operating system and version`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active)`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Record note`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type HINFO ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of hinfo resource.`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Optional) Expected HTTP status code for this check. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of HTTP check resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_http_redirection_record",
			Category:         "Resources",
			ShortDescription: `Manages HTTP redirection for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"http",
				"redirection",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) URL link to redirect`,
				},
				resource.Attribute{
					Name:        "redirect_type_id",
					Description: `(Required) 1 for Hidden Frame Masked, 2 for Standard - 301, 3 for Standard - 302`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active)`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Record note`,
				},
				resource.Attribute{
					Name:        "hardlink_flag",
					Description: `(Optional) Hardlink flag. Default is false`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Optional) Title of redirection`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type HTTP Redirection ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of httpredirection resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_mx_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type MX for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"mx",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Required) The mail server that will accept mail for the host that is specified in the name field. Your domain name is automatically appended to your value if it does not end it a dot.`,
				},
				resource.Attribute{
					Name:        "roundrobin.level",
					Description: `(Required) Level must be in between 0 and 65535. The MX level determines the order (by priority) that remote mail servers will attempt to deliver email. The mail server with the lowest MX level will be the first priority.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active)`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Record note`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type MX ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of mx resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_naptr_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type NAPTR for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"naptr",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the Naptr record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active).`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional)Record note.`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type Naptr.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Set.`,
				},
				resource.Attribute{
					Name:        "roundrobin.order",
					Description: `(Required) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.`,
				},
				resource.Attribute{
					Name:        "roundrobin.preference",
					Description: `(Required) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal). A 16-bit value ranging from 0 to 63535, the lowest number having the highest order.`,
				},
				resource.Attribute{
					Name:        "roundrobin.flags",
					Description: `(Required) A Flag is a single character from the set A-Z and 0-9, defined to be application specific, such that each application may define a specific use of the flag or which flags are valid. The flag is enclosed in quotes (“”). Currently defined values are: U – a terminal condition – the result of the regexp is a valid URI. S – a terminal condition – the replace field contains the FQDN of an SRV record. A – a terminal condition – the replace field contains the FQDN of an A or AAAA record. P – a non-terminal condition – the protocol/services part of the params field determines the application specific behavior and subsequent processing is external to the record. “” (empty string) – a non-terminal condition to indicate that regexp is empty and the replace field contains the FQDN of a further NAPTR record.`,
				},
				resource.Attribute{
					Name:        "roundrobin.service",
					Description: `(Required) Defines the application specific service parameters. The generic format is: protocol+rs. Where “protocol” defines the protocol used by the application and “rs” is the resolution service. There may be 0 or more resolution services each separated by +.`,
				},
				resource.Attribute{
					Name:        "roundrobin.regular_expression",
					Description: `(Required) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.`,
				},
				resource.Attribute{
					Name:        "roundrobin.replacement",
					Description: `(Required) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal).`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Required) disable flag. Default is false ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the NAPTR resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_ns_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type NS for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"ns",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the NS record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active).`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional)Record note.`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type NS.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Set.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Required) This will be the host name for the name server, for example ns0.nameserver.com. It is important to note, the domain name is automatically appended to the end of this field unless it ends with a dot (.).`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Required) disable flag. Default is false ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the NS resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_ptr_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type PTR for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"ptr",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the PTR record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active).`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional)Record note.`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type A.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Required) This will be the host name of the computer or server the IP resolves to, for example mail.example.com. It is important to note, the domain name is automatically appended to the end of this field unless it ends with a dot (.).`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false. ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the PTR resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_rp_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type RP for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"rp",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.mailbox",
					Description: `(Required) A mailbox for the responsible person of the domain`,
				},
				resource.Attribute{
					Name:        "roundrobin.txt",
					Description: `(Required) A hostname for the responsible person of the domain`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active)`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Record note`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type RP ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of rp resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_spf_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type SPF for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"spf",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the PTR record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active).`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional)Record note.`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type A.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Required) Value may contain multiple strings (each string enclosed in double quotes). Individual string length should not exceed 255 characters.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false. ## Attributes Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of the SPF resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_srv_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type SRV for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"srv",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Required) The system that will receive the service.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin.port",
					Description: `(Required) The port of the service offered. Value should be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "roundrobin.priority",
					Description: `(Required) The lower the number in the priority field, the higher the preference of the associated target. 0 is the highest priority (lowest number). Value should be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "roundrobin.weight",
					Description: `(Required) The weight of the record allows an administrator to distribute load to multiple targets (load balance). Value should be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active)`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Record note`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type SRV ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of srv resource.`,
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
					Description: `(Required) Name of record. Name should be unique. ## Attribute Reference ## No attributes are exported`,
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
					Description: `(Optional) String which should be received as a result of TCP check. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of TCP check resource.`,
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
					Description: `(Optional) System generated template history version. ## Attributes Reference No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_txt_record",
			Category:         "Resources",
			ShortDescription: `Manages records of type TXT for a specific domain.`,
			Description:      ``,
			Keywords: []string{
				"txt",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Required) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Required) Free form text data of any type which may be no longer than 255 characters unless divided into multiple strings with sets of quotation marks..`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Disable flag. Default is false`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "noanswer",
					Description: `(Optional) Shows if record is enabled or disabled. Default is false (Active)`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Record note`,
				},
				resource.Attribute{
					Name:        "gtd_region",
					Description: `(Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type TXT ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of txt resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_vanity_nameserver",
			Category:         "Resources",
			ShortDescription: `Manages Vanity Nameservers for domain records.`,
			Description:      ``,
			Keywords: []string{
				"vanity",
				"nameserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Vanity nameserver name should be unique.`,
				},
				resource.Attribute{
					Name:        "nameserver_group",
					Description: `(Required) Name server group id. 1 .. Available nameserver groups`,
				},
				resource.Attribute{
					Name:        "nameserver_list_string",
					Description: `(Required) Comma separedted name servers list`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Default flag. Default is false.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) isPublic flag. Default is false`,
				},
				resource.Attribute{
					Name:        "nameserver_group_name",
					Description: `(Optional) Name server group name ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the constellix calculated id of vanitynameserver resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"constellix_a_record":                0,
		"constellix_aaaa_record":             1,
		"constellix_aaaa_record_pool":        2,
		"constellix_aname_record":            3,
		"constellix_a_record_pool":           4,
		"constellix_caa_record":              5,
		"constellix_cert_record":             6,
		"constellix_cname_record":            7,
		"constellix_cname_record_pool":       8,
		"constellix_contact_lists":           9,
		"constellix_dns_check":               10,
		"constellix_domain":                  11,
		"constellix_geo_filter":              12,
		"constellix_geo_proximity":           13,
		"constellix_hinfo_record":            14,
		"constellix_http_check":              15,
		"constellix_http_redirection_record": 16,
		"constellix_mx_record":               17,
		"constellix_naptr_record":            18,
		"constellix_ns_record":               19,
		"constellix_ptr_record":              20,
		"constellix_rp_record":               21,
		"constellix_spf_record":              22,
		"constellix_srv_record":              23,
		"constellix_tags":                    24,
		"constellix_tcp_check":               25,
		"constellix_template":                26,
		"constellix_txt_record":              27,
		"constellix_vanity_nameserver":       28,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
