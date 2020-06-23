package constellix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "constellix_a_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for A record.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the A record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain id of the A record. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
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
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.sort_order",
					Description: `(Optional) Integer value which decides in which order the rounrobinfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover",
					Description: `(Optional) To create a record failover object pass the following attributes.`,
				},
				resource.Attribute{
					Name:        "record_failover_values",
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Optional) Integer value which decides in which order the recordfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Optional) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Optional) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
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
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.sort_order",
					Description: `(Optional) Integer value which decides in which order the rounrobinfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover",
					Description: `(Optional) To create a record failover object pass the following attributes.`,
				},
				resource.Attribute{
					Name:        "record_failover_values",
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Optional) Integer value which decides in which order the recordfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Optional) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Optional) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_aaaa_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for AAAA record.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the AAAA record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain id of the AAAA record. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
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
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.sort_order",
					Description: `(Optional) Integer value which decides in which order the roundrobinfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover",
					Description: `(Optional) To create a record failover object pass the following attributes.`,
				},
				resource.Attribute{
					Name:        "record_failover_values",
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Optional) Integer value which decides in which order the recordfailover should be sorted`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Optional) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Optional) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
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
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin_failover.sort_order",
					Description: `(Optional) Integer value which decides in which order the roundrobinfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover",
					Description: `(Optional) To create a record failover object pass the following attributes.`,
				},
				resource.Attribute{
					Name:        "record_failover_values",
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Optional) Integer value which decides in which order the recordfailover should be sorted`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Optional) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Optional) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_aaaa_record_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for the pools of AAAA records.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Pool name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "num_return",
					Description: `(Optional) minimum number of value object to return. Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "min_available_failover",
					Description: `(Optional)minimum number of Available Failover . Value must be in between 0 and 64.`,
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
					Description: `(Optional) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values`,
				},
				resource.Attribute{
					Name:        "values.value",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "values.weight",
					Description: `(Optional) weight number to sort the priorty. Weight must be in between 1 and 1000000.`,
				},
				resource.Attribute{
					Name:        "values.disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values.check_id",
					Description: `(Optional) Sonar check id is Optional when you want to apply the ITO feature on a pool.`,
				},
				resource.Attribute{
					Name:        "values.policy",
					Description: `(Optional) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Description.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_aname_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for Aname record`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Record id of Aname record ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) Host name. If "Host" value does not end in a dot, your domain name will be appended to it.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
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
					Description: `(Optional) Host name`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Optional) Enable or Disable the recordfailover values object. Default is false. Atleast one object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Optional) Integer value which decides in which order recordfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Optional) 1 for Normal (always lowest level), 2 for Off on any Failover event, 3 for One Way (move to higher level)`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Optional) Enable or Disable the recordfailover object. Default is false. Atleast one recordfailover object should be false.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) Host name. If "Host" value does not end in a dot, your domain name will be appended to it.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
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
					Description: `(Optional) Host name`,
				},
				resource.Attribute{
					Name:        "record_failover_values.check_id",
					Description: `(Optional) Sonar check id`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Optional) Enable or Disable the recordfailover values object. Default is false. Atleast one object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Optional) Integer value which decides in which order recordfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Optional) 1 for Normal (always lowest level), 2 for Off on any Failover event, 3 for One Way (move to higher level)`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Optional) Enable or Disable the recordfailover object. Default is false. Atleast one recordfailover object should be false.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_a_record_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for the pools of A records.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Pool name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "num_return",
					Description: `(Optional) minimum number of value object to return. Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "min_available_failover",
					Description: `(Optional)minimum number of Available Failover . Value must be in between 0 and 64.`,
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
					Description: `(Optional) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values`,
				},
				resource.Attribute{
					Name:        "values.value",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "values.weight",
					Description: `(Optional) weight number to sort the priorty. Weight must be in between 1 and 1000000.`,
				},
				resource.Attribute{
					Name:        "values.disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values.check_id",
					Description: `(Optional) Sonar check id is Optional when you want to apply the ITO feature on a pool.`,
				},
				resource.Attribute{
					Name:        "values.policy",
					Description: `(Optional) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Description.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "num_return",
					Description: `(Optional) minimum number of value object to return. Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "min_available_failover",
					Description: `(Optional)minimum number of Available Failover . Value must be in between 0 and 64.`,
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
					Description: `(Optional) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values`,
				},
				resource.Attribute{
					Name:        "values.value",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "values.weight",
					Description: `(Optional) weight number to sort the priorty. Weight must be in between 1 and 1000000.`,
				},
				resource.Attribute{
					Name:        "values.disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values.check_id",
					Description: `(Optional) Sonar check id is Optional when you want to apply the ITO feature on a pool.`,
				},
				resource.Attribute{
					Name:        "values.policy",
					Description: `(Optional) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Description.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_caa_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for CAA record`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Record id of CAA record ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.caa_provider_id",
					Description: `(Optional) 1 for [ Custom ], 2 for [ No Provider ], 3 for Comodo, 4 for Digicert, 5 for Entrust, 6 for GeoTrust, 7 for Izenpe, 8 for Lets Encrypt, 9 for Symantec, 10 for Thawte`,
				},
				resource.Attribute{
					Name:        "roundrobin.tag",
					Description: `(Optional) "issue" for Issue, "IssueWild" for IssueWild, "iodef" for iodef. Type allows you to choose how you want certificates to be issued by the CA. Each CAA record can contain only one tag-value pair. Options: issue: Explicitly authorizes a single certificate authority to issue a certificate (any type) for the hostname. issuewild: Authorization to issue certificates that specify a wildcard domain. Please note: issuewild properties take precedence over issue properties when specified. iodef: (Incident Description Exchange Format) Specifies a means of reporting certificate issue requests or cases of certificate issue for the corresponding domain that violate the security policy of the issuer or the domain name holder.`,
				},
				resource.Attribute{
					Name:        "roundrobin.data",
					Description: `(Optional) "" for [ Custom ] if CAA provider Id is 1, ";" for [ No Provider ], "comodoca.com" for Comodo, "digicert.com" for Digicert, "entrust.net" for Entrust, "geotrust.com" for GeoTrust, "izenpe.com" for Izenpe, "letsencrypt.org" for Lets Encrypt, "symantec.com" for Symantec, "thawte.com" for Thawte`,
				},
				resource.Attribute{
					Name:        "roundrobin.flag",
					Description: `(Optional) roundRobin.Issuer Critical There is currently only one flag defined, “issuer critical” at a value of 1. If a CA does not understand the flag value for an issuer critical record, then the CA will return with “no issue” for the certification. All records will have the default issuer critical value of 0, which means they are “not critical”. Issuer Critical Value should be between 0 to 255.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Disable flag. Default is false`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type CAA`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.caa_provider_id",
					Description: `(Optional) 1 for [ Custom ], 2 for [ No Provider ], 3 for Comodo, 4 for Digicert, 5 for Entrust, 6 for GeoTrust, 7 for Izenpe, 8 for Lets Encrypt, 9 for Symantec, 10 for Thawte`,
				},
				resource.Attribute{
					Name:        "roundrobin.tag",
					Description: `(Optional) "issue" for Issue, "IssueWild" for IssueWild, "iodef" for iodef. Type allows you to choose how you want certificates to be issued by the CA. Each CAA record can contain only one tag-value pair. Options: issue: Explicitly authorizes a single certificate authority to issue a certificate (any type) for the hostname. issuewild: Authorization to issue certificates that specify a wildcard domain. Please note: issuewild properties take precedence over issue properties when specified. iodef: (Incident Description Exchange Format) Specifies a means of reporting certificate issue requests or cases of certificate issue for the corresponding domain that violate the security policy of the issuer or the domain name holder.`,
				},
				resource.Attribute{
					Name:        "roundrobin.data",
					Description: `(Optional) "" for [ Custom ] if CAA provider Id is 1, ";" for [ No Provider ], "comodoca.com" for Comodo, "digicert.com" for Digicert, "entrust.net" for Entrust, "geotrust.com" for GeoTrust, "izenpe.com" for Izenpe, "letsencrypt.org" for Lets Encrypt, "symantec.com" for Symantec, "thawte.com" for Thawte`,
				},
				resource.Attribute{
					Name:        "roundrobin.flag",
					Description: `(Optional) roundRobin.Issuer Critical There is currently only one flag defined, “issuer critical” at a value of 1. If a CA does not understand the flag value for an issuer critical record, then the CA will return with “no issue” for the certification. All records will have the default issuer critical value of 0, which means they are “not critical”. Issuer Critical Value should be between 0 to 255.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Disable flag. Default is false`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type CAA`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_cert_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for records of type CERT for a specific domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the CERT record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain id of the CERT record. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.certificate_type",
					Description: `(Optional) certificateType 0 - 65,535`,
				},
				resource.Attribute{
					Name:        "roundrobin.key_tag",
					Description: `(Optional) 0 - 65,535`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) disable flag. Default is false`,
				},
				resource.Attribute{
					Name:        "roundrobin.certificate",
					Description: `(Optional) certificate.`,
				},
				resource.Attribute{
					Name:        "roundrobin.algorithm",
					Description: `(Optional) 0-255.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type CERT.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.certificate_type",
					Description: `(Optional) certificateType 0 - 65,535`,
				},
				resource.Attribute{
					Name:        "roundrobin.key_tag",
					Description: `(Optional) 0 - 65,535`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) disable flag. Default is false`,
				},
				resource.Attribute{
					Name:        "roundrobin.certificate",
					Description: `(Optional) certificate.`,
				},
				resource.Attribute{
					Name:        "roundrobin.algorithm",
					Description: `(Optional) 0-255.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type CERT.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_cname_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for one or more domain CNAME records.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the CName record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain id of the CName record. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Optional) Host name.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.checkid",
					Description: `(Optional) Sonar check id.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Optional) Integer value which decides in which order the recordfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Optional) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Optional) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.value",
					Description: `(Optional) Host name.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.checkid",
					Description: `(Optional) Sonar check id.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.sort_order",
					Description: `(Optional) Integer value which decides in which order the recordfailover should be sorted.`,
				},
				resource.Attribute{
					Name:        "record_failover_values.disable_flag",
					Description: `(Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.`,
				},
				resource.Attribute{
					Name:        "record_failover_failover_type",
					Description: `(Optional) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).`,
				},
				resource.Attribute{
					Name:        "record_failover_disable_flag",
					Description: `(Optional) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_cname_record_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for the pools of CNAME records.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Pool name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "num_return",
					Description: `(Optional) minimum number of value object to return. Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "min_available_failover",
					Description: `(Optional) minimum number of Available Failover . Value must be in between 0 and 64.`,
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
					Name:        "disable_flag1",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values`,
				},
				resource.Attribute{
					Name:        "values.value",
					Description: `(Optional) Host name. If "Host" value does not end in a dot, your domain name will be appended to it.`,
				},
				resource.Attribute{
					Name:        "values.weight",
					Description: `(Optional)weight number to sort the priorty. Weight must be in between 1 and 1000000`,
				},
				resource.Attribute{
					Name:        "values.disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values.check_id",
					Description: `(Optional) Sonar check id is Optional when you want to apply the ITO feature on a pool.`,
				},
				resource.Attribute{
					Name:        "values.policy",
					Description: `(Optional) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Description.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "num_return",
					Description: `(Optional) minimum number of value object to return. Value must be in between 0 and 64.`,
				},
				resource.Attribute{
					Name:        "min_available_failover",
					Description: `(Optional) minimum number of Available Failover . Value must be in between 0 and 64.`,
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
					Name:        "disable_flag1",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values`,
				},
				resource.Attribute{
					Name:        "values.value",
					Description: `(Optional) Host name. If "Host" value does not end in a dot, your domain name will be appended to it.`,
				},
				resource.Attribute{
					Name:        "values.weight",
					Description: `(Optional)weight number to sort the priorty. Weight must be in between 1 and 1000000`,
				},
				resource.Attribute{
					Name:        "values.disable_flag",
					Description: `(Optional) Enable or disable pool values. Default is false.`,
				},
				resource.Attribute{
					Name:        "values.check_id",
					Description: `(Optional) Sonar check id is Optional when you want to apply the ITO feature on a pool.`,
				},
				resource.Attribute{
					Name:        "values.policy",
					Description: `(Optional) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Description.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_contact_lists",
			Category:         "Data Sources",
			ShortDescription: `Data source for Contact List`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "email_addresses",
					Description: `(Optional) List of email addresses`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "email_addresses",
					Description: `(Optional) List of email addresses`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_dns_check",
			Category:         "Data Sources",
			ShortDescription: `Data source for DNS check resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of resource. Name should be unique. ## Attribute Reference ##`,
				},
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
			},
			Attributes: []resource.Attribute{
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for Constellix domain`,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional) The amount of time a record not found is cached. Recommended values can vary, between 180 and 172800 (3 min – 2 days). The default value is 180`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Optional) The amount of time a record not found is cached. Recommended values can vary, between 180 and 172800 (3 min – 2 days). The default value is 180`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_geo_filter",
			Category:         "Data Sources",
			ShortDescription: `Data source for Geofilters for A, AAAA, CNAME or ANAME records.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Geo Filter name should be unique. ## Attribute Reference ##`,
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
					Description: `(Optional) Default is 100. For more than 100 rules, parameter should be set explicitly for ADD and Update API calls. Value should be in mulitple of 100 like 200, 300 ...upto the quota limit assigned to the account. Check quota details for IP Filter Rule Limit.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Optional) Default is 100. For more than 100 rules, parameter should be set explicitly for ADD and Update API calls. Value should be in mulitple of 100 like 200, 300 ...upto the quota limit assigned to the account. Check quota details for IP Filter Rule Limit.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_geo_proximity",
			Category:         "Data Sources",
			ShortDescription: `Data source for Geoproximity for A, AAAA, CNAME or ANAME records.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Geo Proximity name should be unique. ## Attribute Reference ##`,
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
					Description: `(Optional)City code. Default is null.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Optional)City code. Default is null.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_hinfo_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for HINFO record`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Record id of HINFO record ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.cpu",
					Description: `(Optional) A description of basic system hardware`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin.os",
					Description: `(Optional) A description of the operating system and version`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type HINFO`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.cpu",
					Description: `(Optional) A description of basic system hardware`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin.os",
					Description: `(Optional) A description of the operating system and version`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type HINFO`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_http_check",
			Category:         "Data Sources",
			ShortDescription: `Data source for HTTP check resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of resource. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Optional) Specifies upper layer protocol like HTTP, HTTPs, etc.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Optional) Site ids to check.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Optional) Specifies upper layer protocol like HTTP, HTTPs, etc.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Optional) Site ids to check.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_http_redirection_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for HTTP Redirection record`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Record id of HTTP Redirection record ## Attribute Reference ##`,
				},
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
					Description: `(Optional) Record type HTTP Redirection`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Optional) Record type HTTP Redirection`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_mx_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for MX record`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Record id of MX record ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) The mail server that will accept mail for the host that is specified in the name field. Your domain name is automatically appended to your value if it does not end it a dot.`,
				},
				resource.Attribute{
					Name:        "roundrobin.level",
					Description: `(Optional) Level must be in between 0 and 65535. The MX level determines the order (by priority) that remote mail servers will attempt to deliver email. The mail server with the lowest MX level will be the first priority.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type MX`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) The mail server that will accept mail for the host that is specified in the name field. Your domain name is automatically appended to your value if it does not end it a dot.`,
				},
				resource.Attribute{
					Name:        "roundrobin.level",
					Description: `(Optional) Level must be in between 0 and 65535. The MX level determines the order (by priority) that remote mail servers will attempt to deliver email. The mail server with the lowest MX level will be the first priority.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type MX`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_naptr_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for records of type NAPTR for a specific domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the NAPTR record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain id of the NAPTR record. ## Attributes Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "roundrobin.order",
					Description: `(Optional) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.`,
				},
				resource.Attribute{
					Name:        "roundrobin.preference",
					Description: `(Optional) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal). A 16-bit value ranging from 0 to 63535, the lowest number having the highest order.`,
				},
				resource.Attribute{
					Name:        "roundrobin.flags",
					Description: `(Optional) A Flag is a single character from the set A-Z and 0-9, defined to be application specific, such that each application may define a specific use of the flag or which flags are valid. The flag is enclosed in quotes (“”). Currently defined values are: U – a terminal condition – the result of the regexp is a valid URI. S – a terminal condition – the replace field contains the FQDN of an SRV record. A – a terminal condition – the replace field contains the FQDN of an A or AAAA record. P – a non-terminal condition – the protocol/services part of the params field determines the application specific behavior and subsequent processing is external to the record. “” (empty string) – a non-terminal condition to indicate that regexp is empty and the replace field contains the FQDN of a further NAPTR record.`,
				},
				resource.Attribute{
					Name:        "roundrobin.service",
					Description: `(Optional) Defines the application specific service parameters. The generic format is: protocol+rs. Where “protocol” defines the protocol used by the application and “rs” is the resolution service. There may be 0 or more resolution services each separated by +.`,
				},
				resource.Attribute{
					Name:        "roundrobin.regular_expression",
					Description: `(Optional) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.`,
				},
				resource.Attribute{
					Name:        "roundrobin.replacement",
					Description: `(Optional) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal).`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) disable flag. Default is false`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "roundrobin.order",
					Description: `(Optional) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.`,
				},
				resource.Attribute{
					Name:        "roundrobin.preference",
					Description: `(Optional) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal). A 16-bit value ranging from 0 to 63535, the lowest number having the highest order.`,
				},
				resource.Attribute{
					Name:        "roundrobin.flags",
					Description: `(Optional) A Flag is a single character from the set A-Z and 0-9, defined to be application specific, such that each application may define a specific use of the flag or which flags are valid. The flag is enclosed in quotes (“”). Currently defined values are: U – a terminal condition – the result of the regexp is a valid URI. S – a terminal condition – the replace field contains the FQDN of an SRV record. A – a terminal condition – the replace field contains the FQDN of an A or AAAA record. P – a non-terminal condition – the protocol/services part of the params field determines the application specific behavior and subsequent processing is external to the record. “” (empty string) – a non-terminal condition to indicate that regexp is empty and the replace field contains the FQDN of a further NAPTR record.`,
				},
				resource.Attribute{
					Name:        "roundrobin.service",
					Description: `(Optional) Defines the application specific service parameters. The generic format is: protocol+rs. Where “protocol” defines the protocol used by the application and “rs” is the resolution service. There may be 0 or more resolution services each separated by +.`,
				},
				resource.Attribute{
					Name:        "roundrobin.regular_expression",
					Description: `(Optional) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.`,
				},
				resource.Attribute{
					Name:        "roundrobin.replacement",
					Description: `(Optional) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal).`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) disable flag. Default is false`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_ns_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for records of type NS for a specific domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the NS record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain id of the NS record. ## Attributes Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) This will be the host name for the name server, for example ns0.nameserver.com. It is important to note, the domain name is automatically appended to the end of this field unless it ends with a dot (.).`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) disable flag. Default is false`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Set.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) This will be the host name for the name server, for example ns0.nameserver.com. It is important to note, the domain name is automatically appended to the end of this field unless it ends with a dot (.).`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) disable flag. Default is false`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_ptr_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for records of type PTR for a specific domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the PTR record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain id of the PTR record. ## Attributes Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) This will be the host name of the computer or server the IP resolves to, for example mail.example.com. It is important to note, the domain name is automatically appended to the end of this field unless it ends with a dot (.).`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) This will be the host name of the computer or server the IP resolves to, for example mail.example.com. It is important to note, the domain name is automatically appended to the end of this field unless it ends with a dot (.).`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_rp_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for RP record`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Record id of RP record ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.mailbox",
					Description: `(Optional) A mailbox for the responsible person of the domain`,
				},
				resource.Attribute{
					Name:        "roundrobin.txt",
					Description: `(Optional) A hostname for the responsible person of the domain`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type RP`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.mailbox",
					Description: `(Optional) A mailbox for the responsible person of the domain`,
				},
				resource.Attribute{
					Name:        "roundrobin.txt",
					Description: `(Optional) A hostname for the responsible person of the domain`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type RP`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_spf_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for records of type SPF for a specific domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Type of the SPF record. The values which can be applied are "domains" or "templates".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain id of the SPF record. ## Attributes Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) Value may contain multiple strings (each string enclosed in double quotes). Individual string length should not exceed 255 characters.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647.`,
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
					Description: `(Optional) Object.`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) Value may contain multiple strings (each string enclosed in double quotes). Individual string length should not exceed 255 characters.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_srv_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for SRV record`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Record id of SRV record ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) The system that will receive the service.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin.port",
					Description: `(Optional) The port of the service offered. Value should be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "roundrobin.priority",
					Description: `(Optional) The lower the number in the priority field, the higher the preference of the associated target. 0 is the highest priority (lowest number). Value should be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "roundrobin.weight",
					Description: `(Optional) The weight of the record allows an administrator to distribute load to multiple targets (load balance). Value should be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type SRV`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) The system that will receive the service.`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.`,
				},
				resource.Attribute{
					Name:        "roundrobin.port",
					Description: `(Optional) The port of the service offered. Value should be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "roundrobin.priority",
					Description: `(Optional) The lower the number in the priority field, the higher the preference of the associated target. 0 is the highest priority (lowest number). Value should be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "roundrobin.weight",
					Description: `(Optional) The weight of the record allows an administrator to distribute load to multiple targets (load balance). Value should be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type SRV`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_tags",
			Category:         "Data Sources",
			ShortDescription: `Data source for Tags`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_tcp_check",
			Category:         "Data Sources",
			ShortDescription: `Data source for TCP check resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of resource. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Optional) Site ids to check.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Optional) Site ids to check.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for the DNS record templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Template names. e.g "sampletemplate". ## Attributes Reference ##`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Id of domain to be applied.`,
				},
				resource.Attribute{
					Name:        "has_gtd_region",
					Description: `(Optional) Enable/Disable GTD Region of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "has_geoip",
					Description: `(Optional) Enable/Disable GEO IP. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) System generated template history version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Id of domain to be applied.`,
				},
				resource.Attribute{
					Name:        "has_gtd_region",
					Description: `(Optional) Enable/Disable GTD Region of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "has_geoip",
					Description: `(Optional) Enable/Disable GEO IP. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) System generated template history version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_txt_record",
			Category:         "Data Sources",
			ShortDescription: `Data source for TXT record`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Record id of TXT record ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) Free form text data of any type which may be no longer than 255 characters unless divided into multiple strings with sets of quotation marks..`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Disable flag. Default is false`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type TXT`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) "domains" for Domain records and "template" for Template records`,
				},
				resource.Attribute{
					Name:        "roundrobin",
					Description: `(Optional) Set`,
				},
				resource.Attribute{
					Name:        "roundrobin.value",
					Description: `(Optional) Free form text data of any type which may be no longer than 255 characters unless divided into multiple strings with sets of quotation marks..`,
				},
				resource.Attribute{
					Name:        "roundrobin.disable_flag",
					Description: `(Optional) Disable flag. Default is false`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
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
					Description: `(Optional) Record type TXT`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_vanity_nameserver",
			Category:         "Data Sources",
			ShortDescription: `Data source for Vanitynameserver record`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Vanity nameserver name should be unique.`,
				},
				resource.Attribute{
					Name:        "nameserver_group",
					Description: `(Optional) Name server group id. 1 .. Available nameserver groups`,
				},
				resource.Attribute{
					Name:        "nameserver_list_string",
					Description: `(Optional) Comma separedted name servers list`,
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
					Description: `(Optional) Name server group name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Vanity nameserver name should be unique.`,
				},
				resource.Attribute{
					Name:        "nameserver_group",
					Description: `(Optional) Name server group id. 1 .. Available nameserver groups`,
				},
				resource.Attribute{
					Name:        "nameserver_list_string",
					Description: `(Optional) Comma separedted name servers list`,
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
					Description: `(Optional) Name server group name`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

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

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
