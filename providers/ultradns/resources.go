package ultradns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ultradns_dirpool",
			Category:         "DIR-Pool",
			ShortDescription: `Manages the Directional (DIR) pool records in UltraDNS.`,
			Description: `

Use this resource to manage Directional (DIR) pool records in UltraDNS.

`,
			Keywords: []string{
				"dir",
				"pool",
				"dirpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A, AAAA, TXT, etc.), or the corresponding number for the type; between 1 and 33.<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ` + "`" + `PTR (12)` + "`" + ` ` + "`" + `MX (15)` + "`" + ` ` + "`" + `TXT (16)` + "`" + ` ` + "`" + `AAAA (28)` + "`" + ` ` + "`" + `SRV (33)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Optional) (String) Allows for an additional description of the Directional (DIR) pool.`,
				},
				resource.Attribute{
					Name:        "conflict_resolve",
					Description: `(Optional) (String) When there is a conflict between a matching GeoIP group and a matching SourceIP group, this will determine which should take precedence. This only applies to a mixed pool (contains both GeoIP and SourceIP data). Valid values are ` + "`" + `GEO` + "`" + ` and ` + "`" + `IP` + "`" + `. Default value set to ` + "`" + `GEO` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignore_ecs",
					Description: `(Optional) (Boolean) Whether to ignore the EDNSO (which is an extended label type allowing for greater DNS messaging size) Client Subnet data when available in the DNS request.</br> ` + "`" + `false` + "`" + `= EDNSO data will be used for IP directional routing.</br> ` + "`" + `true` + "`" + ` = EDNSO data will not be used and IP directional routing decisions will always use the IP address of the recursive server.</br> Default value set to false.`,
				},
				resource.Attribute{
					Name:        "no_response",
					Description: `(Optional) (Block Set, Max: 1) Nested block describing certain geographical territories and IP addresses that will not get a response if they try to access the directional pool. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Required) (Block Set, Min: 1) List of nested blocks describing the pool records. The structure of this block is described below. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) (String) The IPv4/IPv6 address, CNAME, MX, TXT, or SRV format data.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) (String) The type for the specified pool record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) (Integer) The time to live (in seconds) for the corresponding record in rdata. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "all_non_configured",
					Description: `(Optional) (Boolean) Indicates whether or not the associated rdata is used for all of the non-configured geographical territories and SourceIP ranges. At most, one entry in rdataInfo can have this set to true. If this is set to true, then geoInfo (` + "`" + `geo_group_name` + "`" + ` and ` + "`" + `geo_codes` + "`" + ` ) and ipInfo (` + "`" + `ip_group_name` + "`" + ` and ` + "`" + `ip` + "`" + `) are ignored. Default value set to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "geo_group_name",
					Description: `(Optional) (String) The name of the GeoIP group.`,
				},
				resource.Attribute{
					Name:        "geo_codes",
					Description: `(Optional) (String List) The codes for the geographical territories that make up this group. [Valid GEO codes](#valid-geo-codes).`,
				},
				resource.Attribute{
					Name:        "ip_group_name",
					Description: `(Optional) (String) The name of the SourceIP group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) (Block Set) List of nested blocks describing the IP addresses and IP ranges this SourceIP group contains. The structure of this block is described below. ### Nested ` + "`" + `no_response` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "all_non_configured",
					Description: `(Optional) (Boolean) Indicates whether or not “no response” is returned for all of the non-configured geographical territories and IP ranges. This can only be set to ` + "`" + `true` + "`" + ` if there is no entry for rdataInfo, with allNonConfigured is set to ` + "`" + `true` + "`" + `. If this is set to true, then geoInfo (` + "`" + `geo_group_name` + "`" + ` and ` + "`" + `geo_codes` + "`" + ` ) and ipInfo (` + "`" + `ip_group_name` + "`" + ` and ` + "`" + `ip` + "`" + `) are ignored. Default value set to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "geo_group_name",
					Description: `(Optional) (String) The name for the “no response” GeoIP group.`,
				},
				resource.Attribute{
					Name:        "geo_codes",
					Description: `(Optional) (String List) The codes for the geographical territories that make up the “no response” group. [Valid GEO codes](#valid-geo-codes).`,
				},
				resource.Attribute{
					Name:        "ip_group_name",
					Description: `(Optional) (String) The name of the “no response” SourceIP group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) (Block Set) List of nested blocks describing the IP addresses and IP range for the “no response” SourceIP group. The structure of this block is described below. ### Nested ` + "`" + `ip` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Optional) (String) The starting IP address (IPv4 or IPv6). If the start value is present, the end value must be present as well. ` + "`" + `Cidr` + "`" + ` and ` + "`" + `address` + "`" + ` cannot be present.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Optional) (String) The ending IP address (IPv4 or IPv6). If the end value is present, the start value must be present as well. ` + "`" + `cidr` + "`" + ` and ` + "`" + `address` + "`" + ` cannot be present.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) (String) The CIDR format (IPv4 or IPv6) for an IP address range. If ` + "`" + `cidr` + "`" + ` is present, the ` + "`" + `start` + "`" + `, ` + "`" + `end` + "`" + `, and ` + "`" + `address` + "`" + ` cannot be present.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) (String) A single IPv4 or IPv6 address. If ` + "`" + `address` + "`" + ` is present, the ` + "`" + `start` + "`" + `, ` + "`" + `end` + "`" + `, and ` + "`" + `CIDR` + "`" + ` cannot be present. ## Import Directional (DIR) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_dirpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ` ## Valid GEO Codes: | Code | Meaning | Equivalent ISO codes | | :--- | :----: | :--- | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `A1` + "`" + ` | Anonymous Proxy | None | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `A2` + "`" + ` | Satellite Provider | None | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `A3` + "`" + ` | Unknown / Uncategorized IPs | None | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `NAM` + "`" + ` | North America (including Central America and the Caribbean) | ` + "`" + `AG` + "`" + `,` + "`" + `AI` + "`" + `,` + "`" + `AN` + "`" + `,` + "`" + `AW` + "`" + `,` + "`" + `BB` + "`" + `,` + "`" + `BL` + "`" + `,` + "`" + `BM` + "`" + `,</br>` + "`" + `BQ` + "`" + `,` + "`" + `BS` + "`" + `,` + "`" + `BZ` + "`" + `,` + "`" + `CA` + "`" + `,` + "`" + `CR` + "`" + `,` + "`" + `CU` + "`" + `,` + "`" + `CW` + "`" + `,</br>` + "`" + `DM` + "`" + `,` + "`" + `DO` + "`" + `,` + "`" + `GD` + "`" + `,` + "`" + `GL` + "`" + `,` + "`" + `GP` + "`" + `,` + "`" + `GT` + "`" + `,` + "`" + `HN` + "`" + `,</br>` + "`" + `HT` + "`" + `,` + "`" + `JM` + "`" + `,` + "`" + `KN` + "`" + `,` + "`" + `KY` + "`" + `,` + "`" + `LC` + "`" + `,` + "`" + `MF` + "`" + `,` + "`" + `MQ` + "`" + `,</br>` + "`" + `MS` + "`" + `,` + "`" + `MX` + "`" + `,` + "`" + `NI` + "`" + `,` + "`" + `PA` + "`" + `,` + "`" + `PM` + "`" + `,` + "`" + `PR` + "`" + `,` + "`" + `SV` + "`" + `,</br>` + "`" + `SX` + "`" + `,` + "`" + `TC` + "`" + `,` + "`" + `TT` + "`" + `,` + "`" + `U3` + "`" + `,` + "`" + `US` + "`" + `,` + "`" + `VC` + "`" + `,` + "`" + `VG` + "`" + `,</br>` + "`" + `VI` + "`" + ` | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `SAM` + "`" + ` | South America | ` + "`" + `AR` + "`" + `,` + "`" + `BO` + "`" + `,` + "`" + `BR` + "`" + `,` + "`" + `CL` + "`" + `,` + "`" + `CO` + "`" + `,` + "`" + `EC` + "`" + `,` + "`" + `FK` + "`" + `,</br>` + "`" + `GF` + "`" + `,` + "`" + `GS` + "`" + `,` + "`" + `GY` + "`" + `,` + "`" + `PE` + "`" + `,` + "`" + `PY` + "`" + `,` + "`" + `SR` + "`" + `,` + "`" + `U4` + "`" + `,</br>` + "`" + `UY` + "`" + `,` + "`" + `VE` + "`" + ` | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `EUR` + "`" + ` | Europe | ` + "`" + `AD` + "`" + `,` + "`" + `AL` + "`" + `,` + "`" + `AM` + "`" + `,` + "`" + `AT` + "`" + `,` + "`" + `AX` + "`" + `,` + "`" + `AZ` + "`" + `,` + "`" + `BA` + "`" + `,</br>` + "`" + `BE` + "`" + `,` + "`" + `BG` + "`" + `,` + "`" + `BY` + "`" + `,` + "`" + `CH` + "`" + `,` + "`" + `CZ` + "`" + `,` + "`" + `DE` + "`" + `,` + "`" + `DK` + "`" + `,</br>` + "`" + `EE` + "`" + `,` + "`" + `ES` + "`" + `,` + "`" + `FI` + "`" + `,` + "`" + `FO` + "`" + `,` + "`" + `FR` + "`" + `,` + "`" + `GB` + "`" + `,` + "`" + `GE` + "`" + `,</br>` + "`" + `GG` + "`" + `,` + "`" + `GI` + "`" + `,` + "`" + `GR` + "`" + `,` + "`" + `HR` + "`" + `,` + "`" + `HU` + "`" + `,` + "`" + `IE` + "`" + `,` + "`" + `IM` + "`" + `,</br>` + "`" + `IS` + "`" + `,` + "`" + `IT` + "`" + `,` + "`" + `JE` + "`" + `,` + "`" + `LI` + "`" + `,` + "`" + `LT` + "`" + `,` + "`" + `LU` + "`" + `,` + "`" + `LV` + "`" + `,</br>` + "`" + `MC` + "`" + `,` + "`" + `MD` + "`" + `,` + "`" + `ME` + "`" + `,` + "`" + `MK` + "`" + `,` + "`" + `MT` + "`" + `,` + "`" + `NL` + "`" + `,` + "`" + `NO` + "`" + `,</br>` + "`" + `PL` + "`" + `,` + "`" + `PT` + "`" + `,` + "`" + `RO` + "`" + `,` + "`" + `RS` + "`" + `,` + "`" + `SE` + "`" + `,` + "`" + `SI` + "`" + `,` + "`" + `SJ` + "`" + `,</br>` + "`" + `SK` + "`" + `,` + "`" + `SM` + "`" + `,` + "`" + `U5` + "`" + `,` + "`" + `UA` + "`" + `,` + "`" + `VA` + "`" + ` | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `AFR` + "`" + ` | Africa | ` + "`" + `AO` + "`" + `,` + "`" + `BF` + "`" + `,` + "`" + `BI` + "`" + `,` + "`" + `BJ` + "`" + `,` + "`" + `BW` + "`" + `,` + "`" + `CD` + "`" + `,` + "`" + `CF` + "`" + `,</br>` + "`" + `CG` + "`" + `,` + "`" + `CI` + "`" + `,` + "`" + `CM` + "`" + `,` + "`" + `CV` + "`" + `,` + "`" + `DJ` + "`" + `,` + "`" + `DZ` + "`" + `,` + "`" + `EG` + "`" + `,</br>` + "`" + `EH` + "`" + `,` + "`" + `ER` + "`" + `,` + "`" + `ET` + "`" + `,` + "`" + `GA` + "`" + `,` + "`" + `GH` + "`" + `,` + "`" + `GM` + "`" + `,` + "`" + `GN` + "`" + `,</br>` + "`" + `GQ` + "`" + `,` + "`" + `GW` + "`" + `,` + "`" + `KE` + "`" + `,` + "`" + `KM` + "`" + `,` + "`" + `LR` + "`" + `,` + "`" + `LS` + "`" + `,` + "`" + `LY` + "`" + `,</br>` + "`" + `MA` + "`" + `,` + "`" + `MG` + "`" + `,` + "`" + `ML` + "`" + `,` + "`" + `MR` + "`" + `,` + "`" + `MU` + "`" + `,` + "`" + `MW` + "`" + `,` + "`" + `MZ` + "`" + `,</br>` + "`" + `NA` + "`" + `,` + "`" + `NE` + "`" + `,` + "`" + `NG` + "`" + `,` + "`" + `RE` + "`" + `,` + "`" + `RW` + "`" + `,` + "`" + `SC` + "`" + `,` + "`" + `SD` + "`" + `,</br>` + "`" + `SH` + "`" + `,` + "`" + `SL` + "`" + `,` + "`" + `SN` + "`" + `,` + "`" + `SO` + "`" + `,` + "`" + `SS` + "`" + `,` + "`" + `ST` + "`" + `,` + "`" + `SZ` + "`" + `,</br>` + "`" + `TD` + "`" + `,` + "`" + `TG` + "`" + `,` + "`" + `TN` + "`" + `,` + "`" + `TZ` + "`" + `,` + "`" + `U7` + "`" + `,` + "`" + `UG` + "`" + `,` + "`" + `YT` + "`" + `,</br>` + "`" + `ZA` + "`" + `,` + "`" + `ZM` + "`" + `,` + "`" + `ZW` + "`" + ` | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `ASI` + "`" + ` | Asia (including Middle East and the Russian Federation) | ` + "`" + `AE` + "`" + `,` + "`" + `AF` + "`" + `,` + "`" + `BD` + "`" + `,` + "`" + `BH` + "`" + `,` + "`" + `BN` + "`" + `,` + "`" + `BT` + "`" + `,` + "`" + `CN` + "`" + `,</br>` + "`" + `CY` + "`" + `,` + "`" + `HK` + "`" + `,` + "`" + `ID` + "`" + `,` + "`" + `IL` + "`" + `,` + "`" + `IN` + "`" + `,` + "`" + `IO` + "`" + `,` + "`" + `IQ` + "`" + `,</br>` + "`" + `IR` + "`" + `,` + "`" + `JO` + "`" + `,` + "`" + `JP` + "`" + `,` + "`" + `KG` + "`" + `,` + "`" + `KH` + "`" + `,` + "`" + `KP` + "`" + `,` + "`" + `KR` + "`" + `,</br>` + "`" + `KW` + "`" + `,` + "`" + `KZ` + "`" + `,` + "`" + `LA` + "`" + `,` + "`" + `LB` + "`" + `,` + "`" + `LK` + "`" + `,` + "`" + `MM` + "`" + `,` + "`" + `MN` + "`" + `,</br>` + "`" + `MO` + "`" + `,` + "`" + `MV` + "`" + `,` + "`" + `MY` + "`" + `,` + "`" + `NP` + "`" + `,` + "`" + `OM` + "`" + `,` + "`" + `PH` + "`" + `,` + "`" + `PK` + "`" + `,</br>` + "`" + `PS` + "`" + `,` + "`" + `QA` + "`" + `,` + "`" + `RU` + "`" + `,` + "`" + `SA` + "`" + `,` + "`" + `SG` + "`" + `,` + "`" + `SY` + "`" + `,` + "`" + `TH` + "`" + `,</br>` + "`" + `TJ` + "`" + `,` + "`" + `TL` + "`" + `,` + "`" + `TM` + "`" + `,` + "`" + `TR` + "`" + `,` + "`" + `TW` + "`" + `,` + "`" + `U6` + "`" + `,` + "`" + `U8` + "`" + `,</br>` + "`" + `UZ` + "`" + `,` + "`" + `VN` + "`" + `,` + "`" + `YE` + "`" + ` | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `OCN` + "`" + ` | Australia / Oceania | ` + "`" + `AS` + "`" + `,` + "`" + `AU` + "`" + `,` + "`" + `CC` + "`" + `,` + "`" + `CK` + "`" + `,` + "`" + `CX` + "`" + `,` + "`" + `FJ` + "`" + `,` + "`" + `FM` + "`" + `,</br>` + "`" + `GU` + "`" + `,` + "`" + `HM` + "`" + `,` + "`" + `KI` + "`" + `,` + "`" + `MH` + "`" + `,` + "`" + `MP` + "`" + `,` + "`" + `NC` + "`" + `,` + "`" + `NF` + "`" + `,</br>` + "`" + `NR` + "`" + `,` + "`" + `NU` + "`" + `,` + "`" + `NZ` + "`" + `,` + "`" + `PF` + "`" + `,` + "`" + `PG` + "`" + `,` + "`" + `PN` + "`" + `,` + "`" + `PW` + "`" + `,</br>` + "`" + `SB` + "`" + `,` + "`" + `TK` + "`" + `,` + "`" + `TO` + "`" + `,` + "`" + `TV` + "`" + `,` + "`" + `U9` + "`" + `,` + "`" + `UM` + "`" + `,` + "`" + `VU` + "`" + `,</br>` + "`" + `WF` + "`" + `,` + "`" + `WS` + "`" + ` | |_________________________|__________________________________________________|__________________________________________________| | ` + "`" + `ANT` + "`" + ` | Antarctica | ` + "`" + `AQ` + "`" + `, ` + "`" + `TF` + "`" + `, ` + "`" + `BV` + "`" + ` | |_________________________|__________________________________________________|__________________________________________________|`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_probedns",
			Category:         "DNS-Probe",
			ShortDescription: `Manages the DNS probe records in UltraDNS.`,
			Description: `

Use this resource to manage the DNS probe records in UltraDNS.

`,
			Keywords: []string{
				"dns",
				"probe",
				"probedns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) (String) Length of time between probes in minutes. Valid values are ` + "`" + `HALF_MINUTE` + "`" + `, ` + "`" + `ONE_MINUTE` + "`" + `, ` + "`" + `TWO_MINUTES` + "`" + `, ` + "`" + `FIVE_MINUTES` + "`" + `, ` + "`" + `TEN_MINUTES` + "`" + `, and ` + "`" + `FIFTEEN_MINUTES` + "`" + `.</br>Default value set to ` + "`" + `FIVE_MINUTES` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "agents",
					Description: `(Required) (String List) Locations that will be used for probing. Multiple values can be comma separated. Valid values are: ` + "`" + `ASIA` + "`" + `, ` + "`" + `CHINA` + "`" + `, ` + "`" + `EUROPE_EAST` + "`" + `, ` + "`" + `EUROPE_WEST` + "`" + `, ` + "`" + `NORTH_AMERICA_CENTRAL` + "`" + `, ` + "`" + `NORTH_AMERICA_EAST` + "`" + `, ` + "`" + `NORTH_AMERICA_WEST` + "`" + `, ` + "`" + `SOUTH_AMERICA` + "`" + `, ` + "`" + `NEW_YORK` + "`" + `, ` + "`" + `PALO_ALTO` + "`" + `, ` + "`" + `DALLAS` + "`" + `, and ` + "`" + `AMSTERDAM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) (Integer) Number of agents that must agree for a probe state to be changed.`,
				},
				resource.Attribute{
					Name:        "pool_record",
					Description: `(Optional) (String) The pool record associated with this probe. Specified when creating a record-level probe.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) (String) The Port that should be used for DNS lookup. Default value set to 53.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) (String) Select the record type that the probe will check for. Valid values are ` + "`" + `NULL` + "`" + `, ` + "`" + `AXFR` + "`" + `, or any Resource Record Type. Default value set to ` + "`" + `NULL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query_name",
					Description: `(Optional) (String) The name that should be queried.`,
				},
				resource.Attribute{
					Name:        "tcp_only",
					Description: `(Optional) (Boolean) Indicates whether or not the probe should use TCP only, or first UDP then TCP. Default value set to false.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Optional) (Block Set, Max:1) Nested block describing the strings to match the response that will generate a warning or failure. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "run_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_run_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `limit` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Optional) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before a warning is generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Optional) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before a critical warning is generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Optional) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before causing the probe to fail. ### Nested ` + "`" + `response` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Optional) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a warning.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Optional) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a critical warning.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Optional) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a failure. -> ` + "`" + `warning` + "`" + ` and ` + "`" + `critical` + "`" + ` are only used for Traffic Controller pools. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "guid",
					Description: `(Computed) (String) The internal id for this probe. ## Import DNS probe records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, ` + "`" + `record_type` + "`" + `, and ` + "`" + `guid` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1):06084A729D56C85C` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_probe_dns.example "www.example.com.:example.com.:A (1):06084A729D56C85C" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "guid",
					Description: `(Computed) (String) The internal id for this probe. ## Import DNS probe records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, ` + "`" + `record_type` + "`" + `, and ` + "`" + `guid` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1):06084A729D56C85C` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_probe_dns.example "www.example.com.:example.com.:A (1):06084A729D56C85C" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_probehttp",
			Category:         "HTTP-Probe",
			ShortDescription: `Manages the HTTP probe records in UltraDNS.`,
			Description: `

Use this resource to manage the HTTP probe records in UltraDNS.

`,
			Keywords: []string{
				"http",
				"probe",
				"probehttp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) (String) Length of time between probes in minutes. Valid values are ` + "`" + `HALF_MINUTE` + "`" + `, ` + "`" + `ONE_MINUTE` + "`" + `, ` + "`" + `TWO_MINUTES` + "`" + `, ` + "`" + `FIVE_MINUTES` + "`" + `, ` + "`" + `TEN_MINUTES` + "`" + `, and ` + "`" + `FIFTEEN_MINUTES` + "`" + `.</br>Default value set to ` + "`" + `FIVE_MINUTES` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "agents",
					Description: `(Required) (String List) Locations that will be used for probing. One or more values must be specified. Valid values are ` + "`" + `ASIA` + "`" + `, ` + "`" + `CHINA` + "`" + `, ` + "`" + `EUROPE_EAST` + "`" + `, ` + "`" + `EUROPE_WEST` + "`" + `, ` + "`" + `NORTH_AMERICA_CENTRAL` + "`" + `, ` + "`" + `NORTH_AMERICA_EAST` + "`" + `, ` + "`" + `NORTH_AMERICA_WEST` + "`" + `, ` + "`" + `SOUTH_AMERICA` + "`" + `, ` + "`" + `NEW_YORK` + "`" + `, ` + "`" + `PALO_ALTO` + "`" + `, ` + "`" + `DALLAS` + "`" + `, and ` + "`" + `AMSTERDAM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) (Integer) Number of agents that must agree for a probe state to be changed.`,
				},
				resource.Attribute{
					Name:        "pool_record",
					Description: `(Optional) (String) The pool record associated with this probe. Specified when creating a record-level probe.`,
				},
				resource.Attribute{
					Name:        "transaction",
					Description: `(Required) (Block Set, Min:1) List of nested blocks describing the http requests sent for a single probe. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "total_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing the total amount of time spent on all http transactions. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `transaction` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) (String) HTTP method. Valid values are ` + "`" + `GET` + "`" + ` or ` + "`" + `POST` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol_version",
					Description: `(Required) (String) HTTP protocol version. Valid values are ` + "`" + `HTTP/1.0` + "`" + `, ` + "`" + `HTTP/1.1` + "`" + `, ` + "`" + `HTTP/2` + "`" + `. -> HTTP probes will only correctly work if the indicated server supports the configured HTTP protocol version, otherwise the probe will fail.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) (String) The URL that will be probed..`,
				},
				resource.Attribute{
					Name:        "transmitted_data",
					Description: `(Optional) (String) The data to send to the URL.`,
				},
				resource.Attribute{
					Name:        "follow_redirects",
					Description: `(Optional) (Boolean) Indicates whether or not to follow redirects. Default value set to false.`,
				},
				resource.Attribute{
					Name:        "expected_response",
					Description: `(Optional) (String) The Expected Response code for probes to be returned as Successful. Valid values are</br> ` + "`" + `2XX` + "`" + `: Probe will pass for any code between 200-299.</br> ` + "`" + `3XX` + "`" + `: Probe will pass for any code between 300-399.</br> ` + "`" + `2XX|3XX` + "`" + `: Probe will pass for any code between 200-399.</br> Any combination of HTTP codes between 100-599 separated by "|" </br>For example:</br> ` + "`" + `201|302` + "`" + ``,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Optional) (Block Set, Max:1) Nested block describing the strings required to be searched for a probe’s successful response. This does not search the status line or headers. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "connect_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing how long the probe stays connected to the resource. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_connect_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing the mean (average) time to connect for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "run_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_run_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `limit` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Optional) (Integer) Indicates how long (in seconds) the HTTP Transactional Probe should wait, before a warning is generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Optional) (Integer) Indicates how long (in seconds) the HTTP Transactional Probe should wait, before a critical warning is generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Optional) (Integer) Indicates how long the (in seconds) HTTP Transactional Probe should wait, before causing the probe to fail. -> ` + "`" + `warning` + "`" + ` and ` + "`" + `critical` + "`" + ` are only used for Traffic Controller pools. ### Nested ` + "`" + `search_string` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Optional) (String) If the probe does not find the serach string within the response, or does not match it as a regular expression, a warning will be generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Optional) (String) If the probe does not find the serach string within the response, or does not match it as a regular expression, a critical warning will be generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Optional) (String) If the probe does not find the serach string within the response, or does not match it as a regular expression, the probe will fail. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "guid",
					Description: `(Computed) (String) The internal id for this probe. ## Import HTTP probe records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, ` + "`" + `record_type` + "`" + `, and ` + "`" + `guid` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1):06084A729D56C85C` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_probe_http.example "www.example.com.:example.com.:A (1):06084A729D56C85C" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "guid",
					Description: `(Computed) (String) The internal id for this probe. ## Import HTTP probe records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, ` + "`" + `record_type` + "`" + `, and ` + "`" + `guid` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1):06084A729D56C85C` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_probe_http.example "www.example.com.:example.com.:A (1):06084A729D56C85C" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_probeping",
			Category:         "PING-Probe",
			ShortDescription: `Manages the PING probe records in UltraDNS.`,
			Description: `

Use this resource to manage the PING probe records in UltraDNS.

`,
			Keywords: []string{
				"ping",
				"probe",
				"probeping",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) (String) Length of time between probes in minutes. Valid values are ` + "`" + `HALF_MINUTE` + "`" + `, ` + "`" + `ONE_MINUTE` + "`" + `, ` + "`" + `TWO_MINUTES` + "`" + `, ` + "`" + `FIVE_MINUTES` + "`" + `, ` + "`" + `TEN_MINUTES` + "`" + `, and ` + "`" + `FIFTEEN_MINUTES` + "`" + `.</br>Default value set to ` + "`" + `FIVE_MINUTES` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "agents",
					Description: `(Required) (String List) Locations that will be used for probing. Multiple values can be comma separated. Valid values are: ` + "`" + `ASIA` + "`" + `, ` + "`" + `CHINA` + "`" + `, ` + "`" + `EUROPE_EAST` + "`" + `, ` + "`" + `EUROPE_WEST` + "`" + `, ` + "`" + `NORTH_AMERICA_CENTRAL` + "`" + `, ` + "`" + `NORTH_AMERICA_EAST` + "`" + `, ` + "`" + `NORTH_AMERICA_WEST` + "`" + `, ` + "`" + `SOUTH_AMERICA` + "`" + `, ` + "`" + `NEW_YORK` + "`" + `, ` + "`" + `PALO_ALTO` + "`" + `, ` + "`" + `DALLAS` + "`" + `, and ` + "`" + `AMSTERDAM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) (Integer) Number of agents that must agree for a probe state to be changed.`,
				},
				resource.Attribute{
					Name:        "pool_record",
					Description: `(Optional) (String) The pool record associated with this probe. Specified when creating a record-level probe.`,
				},
				resource.Attribute{
					Name:        "packets",
					Description: `(Optional) (String) Number of ICMP packets to send. Default value set to 3.`,
				},
				resource.Attribute{
					Name:        "packet_size",
					Description: `(Optional) (String) Size of packets in bytes. Default value set to 56.`,
				},
				resource.Attribute{
					Name:        "loss_percent_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing the acceptable percentage of packets lost, which will in turn, generate either a warning or a failure. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "total_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing how long the probe should run in total for all pings. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "average_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing the mean (average) time to connect for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "run_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_run_limit",
					Description: `(Optional) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `limit` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Optional) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before a warning is generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Optional) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before a critical warning is generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Optional) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before causing the probe to fail. -> ` + "`" + `warning` + "`" + ` and ` + "`" + `critical` + "`" + ` are only used for Traffic Controller pools. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "guid",
					Description: `(Computed) (String) The internal id for this probe. ## Import PING probe records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, ` + "`" + `record_type` + "`" + `, and ` + "`" + `guid` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1):06084A729D56C85C` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_probe_ping.example "www.example.com.:example.com.:A (1):06084A729D56C85C" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "guid",
					Description: `(Computed) (String) The internal id for this probe. ## Import PING probe records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, ` + "`" + `record_type` + "`" + `, and ` + "`" + `guid` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1):06084A729D56C85C` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_probe_ping.example "www.example.com.:example.com.:A (1):06084A729D56C85C" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_rdpool",
			Category:         "RD-Pool",
			ShortDescription: `Manages the Resource Distribution (RD) pool records in UltraDNS.`,
			Description: `

Use this resource to manage resource distribution pool records in UltraDNS

`,
			Keywords: []string{
				"rd",
				"pool",
				"rdpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as the well-known resource record type (A or AAAA) or the corresponding number for the type (1 or 28).<br/> Below are the supported resource record types with the corresponding number::<br/> ` + "`" + `A (1)` + "`" + ` ` + "`" + `AAAA (28)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "record_data",
					Description: `(Required) (String List) The list of IPv4 or IPv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) (String) The order of the records will be returned in. Valid values are ` + "`" + `FIXED` + "`" + `, ` + "`" + `RANDOM` + "`" + `, ` + "`" + `ROUND_ROBIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) An optional description of the RD pool. ## Import Resource Distribution (RD) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type as well as the corresponding number as shown in the example above. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_rdpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_record",
			Category:         "Record",
			ShortDescription: `Manages the standard DNS records in UltraDNS.`,
			Description: `

Use this resource to manage standard DNS records in UltraDNS

`,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as the well-known resource record type (A, AAAA, TXT, etc.) or the corresponding number for the type; between 1 and 65535.<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ` + "`" + `NS (2)` + "`" + ` ` + "`" + `CNAME (5)` + "`" + ` ` + "`" + `SOA (6)` + "`" + ` ` + "`" + `PTR (12)` + "`" + ` ` + "`" + `MX (15)` + "`" + ` ` + "`" + `TXT (16)` + "`" + ` ` + "`" + `AAAA (28)` + "`" + ` ` + "`" + `SRV (33)` + "`" + ` ` + "`" + `SSHFP (44)` + "`" + ` ` + "`" + `APEXALIAS (65282)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "record_data",
					Description: `(Required) (String List) The data for the record displayed as the BIND presentation format for the specified RRTYPE.<br/> Example : For a SRV record, the format of data is ["priority weight port target"] (["2 2 523 example.com."])<br/> Additionally for MX, CNAME, and PTR record types, the data value must be a FQDN, as it cannot be relative to the zone name.<br/> ## Import Records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` must have the type as well as the corresponding number as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ultradns_record.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_sbpool",
			Category:         "SB-Pool",
			ShortDescription: `Manages the SiteBacker (SB) pool records in UltraDNS.`,
			Description: `

Use this resource to manage SiteBacker (SB) pool records in UltraDNS.

`,
			Keywords: []string{
				"sb",
				"pool",
				"sbpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A), or the corresponding number for the type (1).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Optional) (String) An optional description of the SiteBacker (SB) field.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Optional) (Boolean) Indicates whether or not the probes are run for this pool. Default value set to true.`,
				},
				resource.Attribute{
					Name:        "act_on_probes",
					Description: `(Optional) (Boolean) Indicates whether or not pool records will be enabled (true) or disabled (false) when probes are run. Default value set to true.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) (String) Indicates the order of the records returned by the resolver for the SiteBacker pool. Valid values are ` + "`" + `FIXED` + "`" + `, ` + "`" + `RANDOM` + "`" + `, and ` + "`" + `ROUND_ROBIN` + "`" + `. Default value set to ` + "`" + `ROUND_ROBIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) (Integer) The minimum number of records that must fail for a pool to be labeled 'FAILED'. If the number of failed records in the pool is greater than or equal to the 'Failure Threshold' value, the pool will be labeled FAILED.<br/> For example, a pool with six priority records, one all-fail record, and the Failure Threshold value is set to four (4). If four or more priority records are not available to serve, the pool will be labeled FAILED, and the all-fail record will be served.<br/> Valid value between 0 and the number of priority records in the pool.`,
				},
				resource.Attribute{
					Name:        "max_active",
					Description: `(Optional) (Integer) Specifies the maximum number of active servers in a pool and determines when SiteBacker takes backup servers offline.<br/> For example, consider a pool with six servers. Setting Max Active to 4 means SiteBacker takes two servers offline and sends the four active records in the answer. Default value set to 1.`,
				},
				resource.Attribute{
					Name:        "max_served",
					Description: `(Optional) (Integer) Determines the number of record answers for each query. This is typically All Active records, or a subset of Max Active. Default value set to the value of ` + "`" + `max_active` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_record",
					Description: `(Required) (Block Set) List of nested blocks describing the information of backup records for the SiteBacker pool. Specifies the records to be served if all other records fail. There can be one or more A records used as backup records, or a single CNAME record. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Required) (Block Set) List of nested blocks describing the pool records. The structure of this block is described below. ### Nested ` + "`" + `backup_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) (String) The IPv4 address or CNAME for the backup record.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Optional) (Integer) Specifies the time, between 0 – 30 minutes, that SiteBacker waits after detecting that the pool record has failed, prior to activating the primary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool's backup record is active and available to serve records. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) (String) The IPv4 address or CNAME.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) (Integer) Indicates the serving preference for this pool record.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) (Integer) Specifies how many probes must agree before the record state is changed. Default value set to 1.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Optional) (Integer) Specifies the time, between 0 – 30 minutes, that SiteBacker waits after detecting that the pool record has failed, prior to activating the secondary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) (String) The current state of the pool record. Valid values are ` + "`" + `NORMAL` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, and ` + "`" + `INACTIVE` + "`" + `. Default value set to ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Optional) (Boolean) Indicates whether or not probes are run for this pool record. Default value set to true.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool record is active and available to serve records. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br> ` + "`" + `WARNING` + "`" + ` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br> For example, if a pool has a Max Active of 1, and the Priority 1 record is not serving and the Priority 2 record is serving.</br> ` + "`" + `CRITICAL` + "`" + ` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br> For example, if a pool has a Max Active of 2, and only one record is serving.</br> ` + "`" + `FAILED` + "`" + ` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br> For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve. ## Import SiteBacker (SB) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_sbpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br> ` + "`" + `WARNING` + "`" + ` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br> For example, if a pool has a Max Active of 1, and the Priority 1 record is not serving and the Priority 2 record is serving.</br> ` + "`" + `CRITICAL` + "`" + ` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br> For example, if a pool has a Max Active of 2, and only one record is serving.</br> ` + "`" + `FAILED` + "`" + ` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br> For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve. ## Import SiteBacker (SB) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_sbpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_sfpool",
			Category:         "SF-Pool",
			ShortDescription: `Manages the Simple Monitor/Failover (SF) pool records in UltraDNS.`,
			Description: `

Use this resource to manage Simple Monitor/Failover (SF) pool records in UltraDNS.

`,
			Keywords: []string{
				"sf",
				"pool",
				"sfpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A or AAAA), or the corresponding number for the type (1 or 28).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ` + "`" + `AAAA (28)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "record_data",
					Description: `(Required) (String List) The list of IPv4 or IPv6 addresses.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "region_failure_sensitivity",
					Description: `(Required) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active. Valid values are ` + "`" + `LOW` + "`" + `, ` + "`" + `HIGH` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "live_record_state",
					Description: `(Required) (String) Whether or not the live record is currently active. Valid values are:</br> ` + "`" + `FORCED_INACTIVE` + "`" + ` – the backup record should always be returned by a DNS query.</br> ` + "`" + `NOT_FORCED` + "`" + ` – the monitor should determine if the live record or the backup record is returned by a DNS query.`,
				},
				resource.Attribute{
					Name:        "live_record_description",
					Description: `(Optional) (String) An optional description of the live record.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Optional) (String) An optional description of the Simple Failover field.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Required) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "backup_record",
					Description: `(Optional) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below. ### Nested ` + "`" + `monitor` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) (String) Monitored URL. A full URL including the protocol, host, and URI. Valid protocols are HTTP and HTTPS.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) (String) HTTP method used to connect to the monitored URL. Valid values are ` + "`" + `GET` + "`" + `, ` + "`" + `POST` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "transmitted_data",
					Description: `(Optional) (String) If a monitor is sending a POST, this is the data sent as the body of the request.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Optional) (String) A string that is checked against the returned data from the request. The monitor fails if the searchString is not present. ### Nested ` + "`" + `backup_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) (String) An IPv4 address or IPv6 address as a backup record.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) An optional description for the backup record. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + ` – Live record is being served.</br> ` + "`" + `CRITICAL` + "`" + ` – The backup record is being served due to the monitor detecting a failure.</br> ` + "`" + `MANUAL` + "`" + ` – The backup record is being served due to user forcing the live record to be inactive. ## Import Simple Monitor/Failover (SF) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + ` and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type as well as the corresponding number as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_sfpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + ` – Live record is being served.</br> ` + "`" + `CRITICAL` + "`" + ` – The backup record is being served due to the monitor detecting a failure.</br> ` + "`" + `MANUAL` + "`" + ` – The backup record is being served due to user forcing the live record to be inactive. ## Import Simple Monitor/Failover (SF) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + ` and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type as well as the corresponding number as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_sfpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_slbpool",
			Category:         "SLB-Pool",
			ShortDescription: `Manages the Simple Load Balancing (SLB) pool records in UltraDNS.`,
			Description: `

Use this resource to manage Simple Load Balancing (SLB) pool records in UltraDNS.

`,
			Keywords: []string{
				"slb",
				"pool",
				"slbpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A or AAAA), or the corresponding number for the type (1 or 28).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ` + "`" + `AAAA (28)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "region_failure_sensitivity",
					Description: `(Required) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active. Valid values are ` + "`" + `LOW` + "`" + `, and ` + "`" + `HIGH` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_method",
					Description: `(Required) (String) The method used to select which record is returned from the primary record pool. Valid values are:</br> ` + "`" + `PRIORITY_HUNT` + "`" + ` – The sequence of the records listed in the primary record pool determines the priority. The first record listed is the highest priority record. Once a record starts being served, it will always be served until the probe detects a failure on this record, or, the record is set to FORCED_INACTIVE.</br> ` + "`" + `RANDOM` + "`" + ` – A random record is returned from the following set of primary records.</br> ` + "`" + `ROUND_ROBIN` + "`" + ` -A record is returned in (a round robin fashion) rotation, based upon the priority of the following active set of records. -> The top priority record is always returned amongst the set of primary records, when the following conditions are met:</br></br> 1) Pool Probe is determined to be passing successfully for the record (based upon Threshold configuration), along with the record forced_state is`,
				},
				resource.Attribute{
					Name:        "serving_preference",
					Description: `(Required) (String) It determines if records will be selected from the Primary Records pool or from the All Fail Record. Valid values are:</br> ` + "`" + `AUTO_SELECT` + "`" + `: Serving method switches from serving Primary Records, to All Fail Record based upon probe results, and the Forced State of the Primary Records.</br> ` + "`" + `SERVE_PRIMARY` + "`" + `: Only the Primary Records are served based upon the probe results and the Forced State of the Primary Records.</br> ` + "`" + `SERVE_ALL_FAIL` + "`" + `: Only the All Fail Record will be served, ignoring the probe results and the Forced State of the Primary Records. -> Please be aware that it may take up to 15 seconds to see the updated value, after switching between Auto Select/Serve Primary and Serve All Fail.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Optional) (String) An optional description of the Simple Load Balancing (SLB) field.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Required) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "all_fail_record",
					Description: `(Required) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Required) (Block Set, Max: 5) Nested block describing the pool records. The structure of this block is described below. ### Nested ` + "`" + `monitor` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) (String) Monitored URL. A full URL including the protocol, host, and URI. Valid protocols are HTTP and HTTPS.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) (String) HTTP method used to connect to the monitored URL. Valid values are ` + "`" + `GET` + "`" + `, and ` + "`" + `POST` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "transmitted_data",
					Description: `(Optional) (String) If a monitor is sending a POST, this is the data sent as the body of the request.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Optional) (String) A string that is checked against the returned data from the request. The monitor fails if the search string is not present. ### Nested ` + "`" + `all_fail_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) (String) An IPv4 or IPv6 address as a backup record.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) An optional description for the backup record.`,
				},
				resource.Attribute{
					Name:        "serving",
					Description: `(Computed) (Boolean) Serving status of the AllFail Record. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) (String) An IPv4 address or IPv6 address.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) An optional description of the record in the live pool.`,
				},
				resource.Attribute{
					Name:        "forced_state",
					Description: `(Optional) (String) The Forced State of the record that indicates whether the record needs to be: force served, forced to be inactive, or the force status not being considered (monitoring result decides the record state). Valid values are ` + "`" + `FORCED_ACTIVE` + "`" + `, ` + "`" + `FORCED_INACTIVE` + "`" + `, or ` + "`" + `NOT_FORCED` + "`" + `. Default value set to ` + "`" + `NOT_FORCED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "probing_enabled",
					Description: `(Optional) (Boolean) Can be set at the record level to indicate whether probing is required (true) or not (false) for the given record. Default value set to true.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the record is available to be served (true) or not (false), based upon the probe results or the forced state of the record. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- Priority record(s) are being served.</br> ` + "`" + `WARNING` + "`" + ` – One of the priority records is not being served due to the monitor detecting a failure, but there is still a priority record to be served.</br> ` + "`" + `CRITICAL` + "`" + ` – The backup All Fail record is being served due to the monitor detecting a failure. ## Import Simple Load Balancing (SLB) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_slbpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- Priority record(s) are being served.</br> ` + "`" + `WARNING` + "`" + ` – One of the priority records is not being served due to the monitor detecting a failure, but there is still a priority record to be served.</br> ` + "`" + `CRITICAL` + "`" + ` – The backup All Fail record is being served due to the monitor detecting a failure. ## Import Simple Load Balancing (SLB) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_slbpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_tcpool",
			Category:         "TC-Pool",
			ShortDescription: `Manages the Traffic Controller (TC) pool records in UltraDNS.`,
			Description: `

Use this resource to manage Traffic Controller (TC) pool records in UltraDNS.

`,
			Keywords: []string{
				"tc",
				"pool",
				"tcpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A), or the corresponding number for the type (1).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Optional) (String) An optional description of the Traffic Controller (TC) field.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Optional) (Boolean) Indicates whether or not the probes are run for this pool. Default value set to true.`,
				},
				resource.Attribute{
					Name:        "act_on_probes",
					Description: `(Optional) (Boolean) Indicates whether or not pool records will be enabled (true) or disabled (false) when probes are run. Default value set to true.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) (Integer) The minimum number of records that must fail for a pool to be labeled 'FAILED'. If the number of failed records in the pool is greater than or equal to the 'Failure Threshold' value, the pool will be labeled FAILED.<br/> For example, a pool with six priority records, one all-fail record, and the Failure Threshold value is set to four (4). If four or more priority records are not available to serve, the pool will be labeled FAILED, and the all-fail record will be served.<br/> Valid value between 0 and the number of priority records in the pool.`,
				},
				resource.Attribute{
					Name:        "max_to_lb",
					Description: `(Optional) (Integer) Specifies the maximum number of active servers in a pool. The maximum value is the number of pool records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "backup_record",
					Description: `(Required) (Block Set, Max: 1) Nested block describing the information of the backup record for the Traffic Controller pool. The backup record is served if all other records fail. There can be one A record, or a single CNAME record. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Required) (Block Set) List of nested blocks describing the pool records. The structure of this block is described below. ### Nested ` + "`" + `backup_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) (String) The IPv4 address or CNAME for the backup record.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Optional) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the primary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool's backup record is active and available to serve records. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) (String) The IPv4 address or CNAME.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) (Integer) Indicates the serving preference for this pool record.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) (Integer) Specifies how many probes must agree before the record state is changed. Default value set to 1.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) (Integer) Determines the traffic load to send to each server in the Traffic Controller pool. Even integers from 2 to 100. Default value set to 2.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Optional) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the secondary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) (String) The current state of the pool record. Valid values are ` + "`" + `NORMAL` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, and ` + "`" + `INACTIVE` + "`" + `. Default value set to ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Optional) (Boolean) Indicates whether or not probes are run for this pool record. Default value set to true.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool record is active and available to serve records. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br> ` + "`" + `WARNING` + "`" + ` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br> For example, if a pool has a Max Active of 1, and the Priority 1 record is not serving and the Priority 2 record is serving.</br> ` + "`" + `CRITICAL` + "`" + ` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br> For example, if a pool has a Max Active of 2, and only one record is serving.</br> ` + "`" + `FAILED` + "`" + ` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br> For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve. ## Import Traffic Controller (TC) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_tcpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br> ` + "`" + `WARNING` + "`" + ` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br> For example, if a pool has a Max Active of 1, and the Priority 1 record is not serving and the Priority 2 record is serving.</br> ` + "`" + `CRITICAL` + "`" + ` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br> For example, if a pool has a Max Active of 2, and only one record is serving.</br> ` + "`" + `FAILED` + "`" + ` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br> For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve. ## Import Traffic Controller (TC) pool records can be imported by combining their ` + "`" + `owner_name` + "`" + `, ` + "`" + `zone_name` + "`" + `, and ` + "`" + `record_type` + "`" + `, separated by a colon.<br/> Example : ` + "`" + `www.example.com.:example.com.:A (1)` + "`" + `. -> For import, the ` + "`" + `owner_name` + "`" + ` and ` + "`" + `zone_name` + "`" + ` must be a FQDN, and ` + "`" + `record_type` + "`" + ` should have the type, as well as the corresponding number, as shown in the example below. Example: ` + "`" + `` + "`" + `` + "`" + `terraform $ terraform import ultradns_tcpool.example "www.example.com.:example.com.:A (1)" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_zone",
			Category:         "Zone",
			ShortDescription: `Manages the zones in UltraDNS.`,
			Description: `

Use this resource to manage zones in UltraDNS

`,
			Keywords: []string{
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) (String) Name of the account. It must be provided, but it can also be sourced from the ` + "`" + `ULTRADNS_ACCOUNT` + "`" + ` environment variable.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) (String) This is the type of zone. Valid values are ` + "`" + `PRIMARY` + "`" + `, ` + "`" + `SECONDARY` + "`" + ` or ` + "`" + `ALIAS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "change_comment",
					Description: `(Optional) (String) This is used to provide comments on updates.`,
				},
				resource.Attribute{
					Name:        "primary_create_info",
					Description: `(Optional) (Block Set, Max: 1) Nested block describing the info of primary zone. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "secondary_create_info",
					Description: `(Optional) (Block Set, Max: 1) Nested block describing the info of secondary zone. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "alias_create_info",
					Description: `(Optional) (Block Set, Max: 1) Nested block describing the info of alias zone. The structure of this block is described below. -> When ` + "`" + `type` + "`" + ` is "PRIMARY" [` + "`" + `primary_create_info` + "`" + `](#nested-primary_create_info-block-has-the-following-structure) is required.<br> When ` + "`" + `type` + "`" + ` is "SECONDARY" [` + "`" + `secondary_create_info` + "`" + `](#nested-secondary_create_info-block-has-the-following-structure) is required<br> When ` + "`" + `type` + "`" + ` is "ALIAS" [` + "`" + `alias_create_info` + "`" + `](##nested-alias_create_info-block-has-the-following-structure) is required. ### Nested ` + "`" + `primary_create_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "create_type",
					Description: `(Required) (String) Indicates the method for creating the primary zone. Valid values are ` + "`" + `NEW` + "`" + `, ` + "`" + `COPY` + "`" + `, ` + "`" + `TRANSFER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_import",
					Description: `(Optional) (Boolean) Indicates whether or not to move existing records from zones into this new zone. Default value set to false.`,
				},
				resource.Attribute{
					Name:        "original_zone_name",
					Description: `(Optional) (String) The name of the zone being copied. The existing zone must be owned by the same account as the new zone. It needs to be provided if [` + "`" + `create_type` + "`" + `](#create_type) is ` + "`" + `COPY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "inherit",
					Description: `(Optional) (String) Defines whether this zone should inherit the zone transfer values from the Account, and also specifies which values to inherit. Valid values are ` + "`" + `ALL` + "`" + `, ` + "`" + `NONE` + "`" + `, any combination of ` + "`" + `IP_RANGE` + "`" + `, ` + "`" + `NOTIFY_IP` + "`" + `, ` + "`" + `TSIG` + "`" + `. Separate multiple values with a comma.<br/> Example: ` + "`" + `IP_RANGE, NOTIFY_IP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name_server",
					Description: `(Optional) (Block Set, Max: 1) Nested block describing the Primary zone's name server. It needs to be provided if [` + "`" + `create_type` + "`" + `](#create_type) is ` + "`" + `TRANSFER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tsig",
					Description: `(Optional) (Block Set, Max: 1) Nested block describing the TSIG information for the primary zone. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "restrict_ip",
					Description: `(Optional) (Block Set) Nested block describing the list of IPv4 or IPv6 ranges that are allowed to transfer primary zones out using zone transfer protocol (AXFR/IXFR). The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "notify_addresses",
					Description: `(Optional) (Block Set) Nested block describing the addresses that are notified when updates are made to the primary zone. The structure of this block is described below. ### Nested ` + "`" + `name_server` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) (String) The IPv4 or IPv6 address of the primary name server for the source zone.`,
				},
				resource.Attribute{
					Name:        "tsig_key",
					Description: `(Optional) (String) If TSIG is enabled for this name server, the name of the TSIG key.`,
				},
				resource.Attribute{
					Name:        "tsig_key_value",
					Description: `(Optional) (String) If TSIG is enabled for this name server, the TSIG key’s value.`,
				},
				resource.Attribute{
					Name:        "tsig_algorithm",
					Description: `(Optional) (String) The hash algorithm used to generate the TSIG key. Valid values are ` + "`" + `hmac-md5` + "`" + `, ` + "`" + `hmac-sha1` + "`" + `, ` + "`" + `hmac-sha224` + "`" + `, ` + "`" + `hmac-sha256` + "`" + `, ` + "`" + `hmac-sha384` + "`" + `, ` + "`" + `hmac-sha512` + "`" + `. ### Nested ` + "`" + `tsig` + "`" + ` block has the following structure: The following tsig values are required if TSIG is enabled for the zone.`,
				},
				resource.Attribute{
					Name:        "tsig_key_name",
					Description: `(Required) (String) The name of the TSIG key for the zone.`,
				},
				resource.Attribute{
					Name:        "tsig_key_value",
					Description: `(Required) (String) The value of the TSIG key for the zone.`,
				},
				resource.Attribute{
					Name:        "tsig_algorithm",
					Description: `(Required) (String) The hash algorithm used to generate the TSIG key. Valid values are ` + "`" + `hmac-md5` + "`" + `, ` + "`" + `hmac-sha1` + "`" + `, ` + "`" + `hmac-sha224` + "`" + `, ` + "`" + `hmac-sha256` + "`" + `, ` + "`" + `hmac-sha384` + "`" + `, ` + "`" + `hmac-sha512` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) A description of this key. ### Nested ` + "`" + `restrict_ip` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `(Optional) (String) The start of the IPv4 or IPv6 range that is allowed to transfer this primary zone out using zone transfer protocol.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `(Optional) (String) The end of the IPv4 or IPv6 range that is allowed to transfer this primary zone out using zone transfer protocol.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) (String) The IP Address ranges specified in CIDR.`,
				},
				resource.Attribute{
					Name:        "single_ip",
					Description: `(Optional) (String) The IPv4 or IPv6 address that is allowed to transfer this primary zone out using zone transfer protocol.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) (String) A description of this range of IP addresses. ### Nested ` + "`" + `notify_addresses` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "notify_address",
					Description: `(Required) (String) The IPv4 Address that is notified when the primary zone is updated.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) A description of this IP Address. ### Nested ` + "`" + `secondary_create_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "primary_name_server_1",
					Description: `(Required) (Block Set) The structure of this block follows the same structure as the [` + "`" + `name_server` + "`" + `](#nested-name_server-block-has-the-following-structure) block described above. It is the info of primary name server.`,
				},
				resource.Attribute{
					Name:        "primary_name_server_2",
					Description: `(Optional) (Block Set) The structure of this block follows the same structure as the [` + "`" + `name_server` + "`" + `](#nested-name_server-block-has-the-following-structure) block described above. It is the info of first backup primary name server.`,
				},
				resource.Attribute{
					Name:        "primary_name_server_3",
					Description: `(Optional) (Block Set) The structure of this block follows the same structure as the [` + "`" + `name_server` + "`" + `](#nested-name_server-block-has-the-following-structure) block described above. It is the info of second backup primary name server.`,
				},
				resource.Attribute{
					Name:        "notification_email_address",
					Description: `(Optional) (String) The Notification Email for a secondary zone. ### Nested ` + "`" + `alias_create_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "original_zone_name",
					Description: `(Required) (String) The name of the zone being aliased. The existing zone must be owned by the same account as the new zone. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Display the status of the zone.`,
				},
				resource.Attribute{
					Name:        "dnssec_status",
					Description: `(Computed) (String) Whether or not the zone is signed with DNSSEC. Valid values are ` + "`" + `SIGNED` + "`" + ` or ` + "`" + `UNSIGNED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Computed) (String) Name of the user that created the zone.`,
				},
				resource.Attribute{
					Name:        "resource_record_count",
					Description: `(Computed) (Integer) Number of records in the zone.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Computed) (String) The last date and time the zone was modified, represented in ISO8601 format (` + "`" + `yyyy-MM-ddTHH:mmZ` + "`" + `).<br/> Example: ` + "`" + `2021-12-07T11:25Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "registrar_info",
					Description: `(Computed) (Block Set) Nested block describing information about the name server configuration for this zone. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "transfer_status_details",
					Description: `(Computed) (Block Set) Nested block describing the zone transfer details. The structure of this block is described below. #### When ` + "`" + `type` + "`" + ` is "PRIMARY" [` + "`" + `registrar_info` + "`" + `](#nested-registrar_info-block-has-the-following-structure) will be exported. #### When ` + "`" + `type` + "`" + ` is "SECONDARY" [` + "`" + `transfer_status_details` + "`" + `](#nested-transfer_status_details-block-has-the-following-structure) will be exported. ### Nested ` + "`" + `registrar_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "registrar",
					Description: `(Computed) (String) The name of the domain registrar.`,
				},
				resource.Attribute{
					Name:        "who_is_expiration",
					Description: `(Computed) (String) The date (` + "`" + `yyyy-MM-dd HH:mm:ss.S` + "`" + `) when the domain name registration expires.<br/> Example: ` + "`" + `2022-08-17 03:59:59.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Computed) (Block Set) Nested block describing the name servers configuration of the zone. The structure of this block is described below. ### Nested ` + "`" + `registrar_info.name_servers` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "ok",
					Description: `(Computed) (List String) List of UltraDNS name servers that are configured for this domain.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `(Computed) (List String) List of name servers that are configured for this domain, but are not UltraDNS managed name servers.`,
				},
				resource.Attribute{
					Name:        "missing",
					Description: `(Computed) (List String) List of UltraDNS name servers that should be configured for this domain, but are not.`,
				},
				resource.Attribute{
					Name:        "incorrect",
					Description: `(Computed) (List String) List of any obsolete UltraDNS name servers that are still configured for this zone. ### Nested ` + "`" + `transfer_status_details` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "last_refresh",
					Description: `(Computed) (String) Displays the date (` + "`" + `MM/dd/yy HH:mm:ss tt vvv` + "`" + `) when the last transfer attempt or refresh was.<br/> Example: ` + "`" + `03/18/21 10:20:34 AM GMT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_refresh",
					Description: `(Computed) (String) Displays the date (` + "`" + `MM/dd/yy HH:mm:ss tt vvv` + "`" + `) when the next transfer attempt or refresh is.<br/> Example: ` + "`" + `03/18/21 10:20:34 AM GMT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_refresh_status",
					Description: `(Computed) (String) Displays the status of the last transfer that was attempted. Valid values are ` + "`" + `IN_PROGRESS` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `SUCCESSFUL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_refresh_status_message",
					Description: `(Computed) (String) Displays the last transfer’s status message. This is currently shown as failure reason. ## Import Zones can be imported using their name (must be a FQDN). Example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ultradns_zone.example "example.com." ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Display the status of the zone.`,
				},
				resource.Attribute{
					Name:        "dnssec_status",
					Description: `(Computed) (String) Whether or not the zone is signed with DNSSEC. Valid values are ` + "`" + `SIGNED` + "`" + ` or ` + "`" + `UNSIGNED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Computed) (String) Name of the user that created the zone.`,
				},
				resource.Attribute{
					Name:        "resource_record_count",
					Description: `(Computed) (Integer) Number of records in the zone.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Computed) (String) The last date and time the zone was modified, represented in ISO8601 format (` + "`" + `yyyy-MM-ddTHH:mmZ` + "`" + `).<br/> Example: ` + "`" + `2021-12-07T11:25Z` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "registrar_info",
					Description: `(Computed) (Block Set) Nested block describing information about the name server configuration for this zone. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "transfer_status_details",
					Description: `(Computed) (Block Set) Nested block describing the zone transfer details. The structure of this block is described below. #### When ` + "`" + `type` + "`" + ` is "PRIMARY" [` + "`" + `registrar_info` + "`" + `](#nested-registrar_info-block-has-the-following-structure) will be exported. #### When ` + "`" + `type` + "`" + ` is "SECONDARY" [` + "`" + `transfer_status_details` + "`" + `](#nested-transfer_status_details-block-has-the-following-structure) will be exported. ### Nested ` + "`" + `registrar_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "registrar",
					Description: `(Computed) (String) The name of the domain registrar.`,
				},
				resource.Attribute{
					Name:        "who_is_expiration",
					Description: `(Computed) (String) The date (` + "`" + `yyyy-MM-dd HH:mm:ss.S` + "`" + `) when the domain name registration expires.<br/> Example: ` + "`" + `2022-08-17 03:59:59.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Computed) (Block Set) Nested block describing the name servers configuration of the zone. The structure of this block is described below. ### Nested ` + "`" + `registrar_info.name_servers` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "ok",
					Description: `(Computed) (List String) List of UltraDNS name servers that are configured for this domain.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `(Computed) (List String) List of name servers that are configured for this domain, but are not UltraDNS managed name servers.`,
				},
				resource.Attribute{
					Name:        "missing",
					Description: `(Computed) (List String) List of UltraDNS name servers that should be configured for this domain, but are not.`,
				},
				resource.Attribute{
					Name:        "incorrect",
					Description: `(Computed) (List String) List of any obsolete UltraDNS name servers that are still configured for this zone. ### Nested ` + "`" + `transfer_status_details` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "last_refresh",
					Description: `(Computed) (String) Displays the date (` + "`" + `MM/dd/yy HH:mm:ss tt vvv` + "`" + `) when the last transfer attempt or refresh was.<br/> Example: ` + "`" + `03/18/21 10:20:34 AM GMT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_refresh",
					Description: `(Computed) (String) Displays the date (` + "`" + `MM/dd/yy HH:mm:ss tt vvv` + "`" + `) when the next transfer attempt or refresh is.<br/> Example: ` + "`" + `03/18/21 10:20:34 AM GMT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_refresh_status",
					Description: `(Computed) (String) Displays the status of the last transfer that was attempted. Valid values are ` + "`" + `IN_PROGRESS` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `SUCCESSFUL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_refresh_status_message",
					Description: `(Computed) (String) Displays the last transfer’s status message. This is currently shown as failure reason. ## Import Zones can be imported using their name (must be a FQDN). Example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ultradns_zone.example "example.com." ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ultradns_dirpool":   0,
		"ultradns_probedns":  1,
		"ultradns_probehttp": 2,
		"ultradns_probeping": 3,
		"ultradns_rdpool":    4,
		"ultradns_record":    5,
		"ultradns_sbpool":    6,
		"ultradns_sfpool":    7,
		"ultradns_slbpool":   8,
		"ultradns_tcpool":    9,
		"ultradns_zone":      10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
