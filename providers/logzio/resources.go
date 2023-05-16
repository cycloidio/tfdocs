package logzio

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "logzio_alert_v2",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `(String) Alert title.`,
				},
				resource.Attribute{
					Name:        "search_timeframe_minutes",
					Description: `(Integer) The time frame for evaluating the log data is a sliding window, with 1 minute granularity.`,
				},
				resource.Attribute{
					Name:        "sub_components",
					Description: `(Block list) Sets the search criteria using a search query, filters, group by aggregations, accounts to search, and trigger conditions. See below for`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) A description of the event, its significance, and suggested next steps or instructions for the team.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(String list) Tags for filtering alerts and triggered alerts. Can be used in Kibana Discover, dashboards, and more.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Boolean) True by default. If ` + "`" + `true` + "`" + `, the alert is currently active.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `(String list) Array of email addresses to be notified when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "alert_notification_endpoints",
					Description: `(Integer list) Array of IDs of pre-configured endpoint channels to notify when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications_minutes",
					Description: `(Integer) Defaults to 5. Add a waiting period in minutes to space out notifications. (The alert will still trigger but will not send out notifications during the waiting period.)`,
				},
				resource.Attribute{
					Name:        "output_type",
					Description: `(String) Selects the output format for the alert notification. Can be: ` + "`" + `"JSON"` + "`" + ` or ` + "`" + `"TABLE""` + "`" + ` If the alert has no aggregations/group by fields, JSON offers the option to send full sample logs without selecting specific fields.`,
				},
				resource.Attribute{
					Name:        "correlation_operator",
					Description: `(String) Comma separated string of supported operators. Only applicable when multiple sub-components are in use. Selects a logic for correlating the alert’s sub-components. ` + "`" + `AND` + "`" + ` is currently the only supported operator. When AND is the correlation_operator, both sub-components must meet their triggering criteria for the alert to trigger.`,
				},
				resource.Attribute{
					Name:        "joins",
					Description: `(Map list) Specifies which group by fields must have the same values to trigger the alert. Joins the group by fields from the first and second sub-components. The key represents the index of the sub component in the array. The fields must be ordered pairs of the group by fields already in use in the ` + "`" + `sub_components.query_string` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "schedule_cron_expression",
					Description: `(String) Cron job for the intervals schedule.`,
				},
				resource.Attribute{
					Name:        "schedule_timezone",
					Description: `(String) Time zone for the cron job. If no time zone is selected, UTC will be used by default. For a list of all possible values, see ` + "`" + `Timezones` + "`" + ` section. #### Nested schema for ` + "`" + `sub_components` + "`" + `: ##### Required:`,
				},
				resource.Attribute{
					Name:        "query_string",
					Description: `(String) Provide a Kibana search query written in Lucene syntax. The search query together with the filters select for the relevant logs. Cannot be null - send an asterisk wildcard ` + "`" + `"`,
				},
				resource.Attribute{
					Name:        "value_aggregation_type",
					Description: `(String) Specifies the aggregation operator. Can be: ` + "`" + `"SUM"` + "`" + `, ` + "`" + `"MIN"` + "`" + `, ` + "`" + `"MAX"` + "`" + `, ` + "`" + `"AVG"` + "`" + `, ` + "`" + `"COUNT"` + "`" + `, ` + "`" + `"UNIQUE_COUNT"` + "`" + `, ` + "`" + `"NONE"` + "`" + `. If ` + "`" + `"COUNT"` + "`" + ` or ` + "`" + `"NONE"` + "`" + `, ` + "`" + `value_aggregation_field` + "`" + ` must be null, and ` + "`" + `group_by_aggregation_fields` + "`" + ` fields must not be empty. If any other operator type (other than ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"COUNT"` + "`" + `), ` + "`" + `value_aggregation_field` + "`" + ` must not be null.`,
				},
				resource.Attribute{
					Name:        "severity_threshold_tiers",
					Description: `(Block) Sets a severity label per trigger threshold. If using more than one sub-component, only 1 severityThresholdTiers is allowed. Otherwise, 1 per enum are allowed (for a total of 5 thresholds of increasing severities). Increasing severity must adhere to the logic of the operator. See below for`,
				},
				resource.Attribute{
					Name:        "filter_must_not",
					Description: `(String) Runs Elasticsearch Bool Query filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "group_by_aggregation_fields",
					Description: `(String list) Specify 1-3 fields by which to group the results and count them. If you apply a group by operation, the alert returns a count of the results aggregated by unique values.`,
				},
				resource.Attribute{
					Name:        "value_aggregation_field",
					Description: `(String) Selects the field on which to run the aggregation for the trigger condition. Cannot be a field already in use for ` + "`" + `group_by_aggregation_fields` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "should_query_on_all_accounts",
					Description: `(Boolean) Defaults to true. Only applicable when the alert is run from the main account. If true, the alert runs on the main account and all associated searchable sub accounts. If false, specify relevant account IDs for the alert to monitor using the ` + "`" + `account_ids_to_query_on` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "account_ids_to_query_on",
					Description: `(Integer list) Specify Account IDs to select which accounts the alert should monitor. The alert will be checked only on these accounts.`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(String) Specifies the operator for evaluating the results. Can be: ` + "`" + `"LESS_THAN"` + "`" + `, ` + "`" + `"GREATER_THAN"` + "`" + `, ` + "`" + `"LESS_THAN_OR_EQUALS"` + "`" + `, ` + "`" + `"GREATER_THAN_OR_EQUALS"` + "`" + `, ` + "`" + `"EQUALS"` + "`" + `, ` + "`" + `"NOT_EQUALS"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "columns",
					Description: `(Block list) See below for`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(String) Labels the event with a severity tag. Available severity tags are: ` + "`" + `"INFO"` + "`" + `, ` + "`" + `"LOW"` + "`" + `, ` + "`" + `"MEDIUM"` + "`" + `, ` + "`" + `"HIGH"` + "`" + `, ` + "`" + `"SEVERE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Integer) Number of logs per search timeframe. #### Nested schema for ` + "`" + `sub_components.columns` + "`" + `: ##### Optional:`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `(String) Specify the fields to be included in the notification.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(String) Trims the data using regex filters. [Learn more](https://docs.logz.io/user-guide/alerts/regex-filters.html).`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(String) Specify a single field to sort by. The field cannot be an analyzed field (a field that supports free text search or searching by part of a message, such as the 'message' field). Should be: ` + "`" + `"DESC"` + "`" + `, ` + "`" + `"ASC"` + "`" + `. #### Timezones: This is a list of possible timezone values for field ` + "`" + `schedule_timezone` + "`" + `: ` + "`" + `Africa/Abidjan` + "`" + `, ` + "`" + `Africa/Accra` + "`" + `, ` + "`" + `Africa/Addis_Ababa` + "`" + `, ` + "`" + `Africa/Algiers` + "`" + `, ` + "`" + `Africa/Asmara` + "`" + `, ` + "`" + `Africa/Asmera` + "`" + `, ` + "`" + `Africa/Bamako` + "`" + `, ` + "`" + `Africa/Bangui` + "`" + `, ` + "`" + `Africa/Banjul` + "`" + `, ` + "`" + `Africa/Bissau` + "`" + `, ` + "`" + `Africa/Blantyre` + "`" + `, ` + "`" + `Africa/Brazzaville` + "`" + `, ` + "`" + `Africa/Bujumbura` + "`" + `, ` + "`" + `Africa/Cairo` + "`" + `, ` + "`" + `Africa/Casablanca` + "`" + `, ` + "`" + `Africa/Ceuta` + "`" + `, ` + "`" + `Africa/Conakry` + "`" + `, ` + "`" + `Africa/Dakar` + "`" + `, ` + "`" + `Africa/Dar_es_Salaam` + "`" + `, ` + "`" + `Africa/Djibouti` + "`" + `, ` + "`" + `Africa/Douala` + "`" + `, ` + "`" + `Africa/El_Aaiun` + "`" + `, ` + "`" + `Africa/Freetown` + "`" + `, ` + "`" + `Africa/Gaborone` + "`" + `, ` + "`" + `Africa/Harare` + "`" + `, ` + "`" + `Africa/Johannesburg` + "`" + `, ` + "`" + `Africa/Juba` + "`" + `, ` + "`" + `Africa/Kampala` + "`" + `, ` + "`" + `Africa/Khartoum` + "`" + `, ` + "`" + `Africa/Kigali` + "`" + `, ` + "`" + `Africa/Kinshasa` + "`" + `, ` + "`" + `Africa/Lagos` + "`" + `, ` + "`" + `Africa/Libreville` + "`" + `, ` + "`" + `Africa/Lome` + "`" + `, ` + "`" + `Africa/Luanda` + "`" + `, ` + "`" + `Africa/Lubumbashi` + "`" + `, ` + "`" + `Africa/Lusaka` + "`" + `, ` + "`" + `Africa/Malabo` + "`" + `, ` + "`" + `Africa/Maputo` + "`" + `, ` + "`" + `Africa/Maseru` + "`" + `, ` + "`" + `Africa/Mbabane` + "`" + `, ` + "`" + `Africa/Mogadishu` + "`" + `, ` + "`" + `Africa/Monrovia` + "`" + `, ` + "`" + `Africa/Nairobi` + "`" + `, ` + "`" + `Africa/Ndjamena` + "`" + `, ` + "`" + `Africa/Niamey` + "`" + `, ` + "`" + `Africa/Nouakchott` + "`" + `, ` + "`" + `Africa/Ouagadougou` + "`" + `, ` + "`" + `Africa/Porto-Novo` + "`" + `, ` + "`" + `Africa/Sao_Tome` + "`" + `, ` + "`" + `Africa/Timbuktu` + "`" + `, ` + "`" + `Africa/Tripoli` + "`" + `, ` + "`" + `Africa/Tunis` + "`" + `, ` + "`" + `Africa/Windhoek` + "`" + `, ` + "`" + `America/Adak` + "`" + `, ` + "`" + `America/Anchorage` + "`" + `, ` + "`" + `America/Anguilla` + "`" + `, ` + "`" + `America/Antigua` + "`" + `, ` + "`" + `America/Araguaina` + "`" + `, ` + "`" + `America/Argentina/Buenos_Aires` + "`" + `, ` + "`" + `America/Argentina/Catamarca` + "`" + `, ` + "`" + `America/Argentina/ComodRivadavia` + "`" + `, ` + "`" + `America/Argentina/Cordoba` + "`" + `, ` + "`" + `America/Argentina/Jujuy` + "`" + `, ` + "`" + `America/Argentina/La_Rioja` + "`" + `, ` + "`" + `America/Argentina/Mendoza` + "`" + `, ` + "`" + `America/Argentina/Rio_Gallegos` + "`" + `, ` + "`" + `America/Argentina/Salta` + "`" + `, ` + "`" + `America/Argentina/San_Juan` + "`" + `, ` + "`" + `America/Argentina/San_Luis` + "`" + `, ` + "`" + `America/Argentina/Tucuman` + "`" + `, ` + "`" + `America/Argentina/Ushuaia` + "`" + `, ` + "`" + `America/Aruba` + "`" + `, ` + "`" + `America/Asuncion` + "`" + `, ` + "`" + `America/Atikokan` + "`" + `, ` + "`" + `America/Atka` + "`" + `, ` + "`" + `America/Bahia` + "`" + `, ` + "`" + `America/Bahia_Banderas` + "`" + `, ` + "`" + `America/Barbados` + "`" + `, ` + "`" + `America/Belem` + "`" + `, ` + "`" + `America/Belize` + "`" + `, ` + "`" + `America/Blanc-Sablon` + "`" + `, ` + "`" + `America/Boa_Vista` + "`" + `, ` + "`" + `America/Bogota` + "`" + `, ` + "`" + `America/Boise` + "`" + `, ` + "`" + `America/Buenos_Aires` + "`" + `, ` + "`" + `America/Cambridge_Bay` + "`" + `, ` + "`" + `America/Campo_Grande` + "`" + `, ` + "`" + `America/Cancun` + "`" + `, ` + "`" + `America/Caracas` + "`" + `, ` + "`" + `America/Catamarca` + "`" + `, ` + "`" + `America/Cayenne` + "`" + `, ` + "`" + `America/Cayman` + "`" + `, ` + "`" + `America/Chicago` + "`" + `, ` + "`" + `America/Chihuahua` + "`" + `, ` + "`" + `America/Coral_Harbour` + "`" + `, ` + "`" + `America/Cordoba` + "`" + `, ` + "`" + `America/Costa_Rica` + "`" + `, ` + "`" + `America/Creston` + "`" + `, ` + "`" + `America/Cuiaba` + "`" + `, ` + "`" + `America/Curacao` + "`" + `, ` + "`" + `America/Danmarkshavn` + "`" + `, ` + "`" + `America/Dawson` + "`" + `, ` + "`" + `America/Dawson_Creek` + "`" + `, ` + "`" + `America/Denver` + "`" + `, ` + "`" + `America/Detroit` + "`" + `, ` + "`" + `America/Dominica` + "`" + `, ` + "`" + `America/Edmonton` + "`" + `, ` + "`" + `America/Eirunepe` + "`" + `, ` + "`" + `America/El_Salvador` + "`" + `, ` + "`" + `America/Ensenada` + "`" + `, ` + "`" + `America/Fort_Nelson` + "`" + `, ` + "`" + `America/Fort_Wayne` + "`" + `, ` + "`" + `America/Fortaleza` + "`" + `, ` + "`" + `America/Glace_Bay` + "`" + `, ` + "`" + `America/Godthab` + "`" + `, ` + "`" + `America/Goose_Bay` + "`" + `, ` + "`" + `America/Grand_Turk` + "`" + `, ` + "`" + `America/Grenada` + "`" + `, ` + "`" + `America/Guadeloupe` + "`" + `, ` + "`" + `America/Guatemala` + "`" + `, ` + "`" + `America/Guayaquil` + "`" + `, ` + "`" + `America/Guyana` + "`" + `, ` + "`" + `America/Halifax` + "`" + `, ` + "`" + `America/Havana` + "`" + `, ` + "`" + `America/Hermosillo` + "`" + `, ` + "`" + `America/Indiana/Indianapolis` + "`" + `, ` + "`" + `America/Indiana/Knox` + "`" + `, ` + "`" + `America/Indiana/Marengo` + "`" + `, ` + "`" + `America/Indiana/Petersburg` + "`" + `, ` + "`" + `America/Indiana/Tell_City` + "`" + `, ` + "`" + `America/Indiana/Vevay` + "`" + `, ` + "`" + `America/Indiana/Vincennes` + "`" + `, ` + "`" + `America/Indiana/Winamac` + "`" + `, ` + "`" + `America/Indianapolis` + "`" + `, ` + "`" + `America/Inuvik` + "`" + `, ` + "`" + `America/Iqaluit` + "`" + `, ` + "`" + `America/Jamaica` + "`" + `, ` + "`" + `America/Jujuy` + "`" + `, ` + "`" + `America/Juneau` + "`" + `, ` + "`" + `America/Kentucky/Louisville` + "`" + `, ` + "`" + `America/Kentucky/Monticello` + "`" + `, ` + "`" + `America/Knox_IN` + "`" + `, ` + "`" + `America/Kralendijk` + "`" + `, ` + "`" + `America/La_Paz` + "`" + `, ` + "`" + `America/Lima` + "`" + `, ` + "`" + `America/Los_Angeles` + "`" + `, ` + "`" + `America/Louisville` + "`" + `, ` + "`" + `America/Lower_Princes` + "`" + `, ` + "`" + `America/Maceio` + "`" + `, ` + "`" + `America/Managua` + "`" + `, ` + "`" + `America/Manaus` + "`" + `, ` + "`" + `America/Marigot` + "`" + `, ` + "`" + `America/Martinique` + "`" + `, ` + "`" + `America/Matamoros` + "`" + `, ` + "`" + `America/Mazatlan` + "`" + `, ` + "`" + `America/Mendoza` + "`" + `, ` + "`" + `America/Menominee` + "`" + `, ` + "`" + `America/Merida` + "`" + `, ` + "`" + `America/Metlakatla` + "`" + `, ` + "`" + `America/Mexico_City` + "`" + `, ` + "`" + `America/Miquelon` + "`" + `, ` + "`" + `America/Moncton` + "`" + `, ` + "`" + `America/Monterrey` + "`" + `, ` + "`" + `America/Montevideo` + "`" + `, ` + "`" + `America/Montreal` + "`" + `, ` + "`" + `America/Montserrat` + "`" + `, ` + "`" + `America/Nassau` + "`" + `, ` + "`" + `America/New_York` + "`" + `, ` + "`" + `America/Nipigon` + "`" + `, ` + "`" + `America/Nome` + "`" + `, ` + "`" + `America/Noronha` + "`" + `, ` + "`" + `America/North_Dakota/Beulah` + "`" + `, ` + "`" + `America/North_Dakota/Center` + "`" + `, ` + "`" + `America/North_Dakota/New_Salem` + "`" + `, ` + "`" + `America/Nuuk` + "`" + `, ` + "`" + `America/Ojinaga` + "`" + `, ` + "`" + `America/Panama` + "`" + `, ` + "`" + `America/Pangnirtung` + "`" + `, ` + "`" + `America/Paramaribo` + "`" + `, ` + "`" + `America/Phoenix` + "`" + `, ` + "`" + `America/Port-au-Prince` + "`" + `, ` + "`" + `America/Port_of_Spain` + "`" + `, ` + "`" + `America/Porto_Acre` + "`" + `, ` + "`" + `America/Porto_Velho` + "`" + `, ` + "`" + `America/Puerto_Rico` + "`" + `, ` + "`" + `America/Punta_Arenas` + "`" + `, ` + "`" + `America/Rainy_River` + "`" + `, ` + "`" + `America/Rankin_Inlet` + "`" + `, ` + "`" + `America/Recife` + "`" + `, ` + "`" + `America/Regina` + "`" + `, ` + "`" + `America/Resolute` + "`" + `, ` + "`" + `America/Rio_Branco` + "`" + `, ` + "`" + `America/Rosario` + "`" + `, ` + "`" + `America/Santa_Isabel` + "`" + `, ` + "`" + `America/Santarem` + "`" + `, ` + "`" + `America/Santiago` + "`" + `, ` + "`" + `America/Santo_Domingo` + "`" + `, ` + "`" + `America/Sao_Paulo` + "`" + `, ` + "`" + `America/Scoresbysund` + "`" + `, ` + "`" + `America/Shiprock` + "`" + `, ` + "`" + `America/Sitka` + "`" + `, ` + "`" + `America/St_Barthelemy` + "`" + `, ` + "`" + `America/St_Johns` + "`" + `, ` + "`" + `America/St_Kitts` + "`" + `, ` + "`" + `America/St_Lucia` + "`" + `, ` + "`" + `America/St_Thomas` + "`" + `, ` + "`" + `America/St_Vincent` + "`" + `, ` + "`" + `America/Swift_Current` + "`" + `, ` + "`" + `America/Tegucigalpa` + "`" + `, ` + "`" + `America/Thule` + "`" + `, ` + "`" + `America/Thunder_Bay` + "`" + `, ` + "`" + `America/Tijuana` + "`" + `, ` + "`" + `America/Toronto` + "`" + `, ` + "`" + `America/Tortola` + "`" + `, ` + "`" + `America/Vancouver` + "`" + `, ` + "`" + `America/Virgin` + "`" + `, ` + "`" + `America/Whitehorse` + "`" + `, ` + "`" + `America/Winnipeg` + "`" + `, ` + "`" + `America/Yakutat` + "`" + `, ` + "`" + `America/Yellowknife` + "`" + `, ` + "`" + `Antarctica/Casey` + "`" + `, ` + "`" + `Antarctica/Davis` + "`" + `, ` + "`" + `Antarctica/DumontDUrville` + "`" + `, ` + "`" + `Antarctica/Macquarie` + "`" + `, ` + "`" + `Antarctica/Mawson` + "`" + `, ` + "`" + `Antarctica/McMurdo` + "`" + `, ` + "`" + `Antarctica/Palmer` + "`" + `, ` + "`" + `Antarctica/Rothera` + "`" + `, ` + "`" + `Antarctica/South_Pole` + "`" + `, ` + "`" + `Antarctica/Syowa` + "`" + `, ` + "`" + `Antarctica/Troll` + "`" + `, ` + "`" + `Antarctica/Vostok` + "`" + `, ` + "`" + `Arctic/Longyearbyen` + "`" + `, ` + "`" + `Asia/Aden` + "`" + `, ` + "`" + `Asia/Almaty` + "`" + `, ` + "`" + `Asia/Amman` + "`" + `, ` + "`" + `Asia/Anadyr` + "`" + `, ` + "`" + `Asia/Aqtau` + "`" + `, ` + "`" + `Asia/Aqtobe` + "`" + `, ` + "`" + `Asia/Ashgabat` + "`" + `, ` + "`" + `Asia/Ashkhabad` + "`" + `, ` + "`" + `Asia/Atyrau` + "`" + `, ` + "`" + `Asia/Baghdad` + "`" + `, ` + "`" + `Asia/Bahrain` + "`" + `, ` + "`" + `Asia/Baku` + "`" + `, ` + "`" + `Asia/Bangkok` + "`" + `, ` + "`" + `Asia/Barnaul` + "`" + `, ` + "`" + `Asia/Beirut` + "`" + `, ` + "`" + `Asia/Bishkek` + "`" + `, ` + "`" + `Asia/Brunei` + "`" + `, ` + "`" + `Asia/Calcutta` + "`" + `, ` + "`" + `Asia/Chita` + "`" + `, ` + "`" + `Asia/Choibalsan` + "`" + `, ` + "`" + `Asia/Chongqing` + "`" + `, ` + "`" + `Asia/Chungking` + "`" + `, ` + "`" + `Asia/Colombo` + "`" + `, ` + "`" + `Asia/Dacca` + "`" + `, ` + "`" + `Asia/Damascus` + "`" + `, ` + "`" + `Asia/Dhaka` + "`" + `, ` + "`" + `Asia/Dili` + "`" + `, ` + "`" + `Asia/Dubai` + "`" + `, ` + "`" + `Asia/Dushanbe` + "`" + `, ` + "`" + `Asia/Famagusta` + "`" + `, ` + "`" + `Asia/Gaza` + "`" + `, ` + "`" + `Asia/Harbin` + "`" + `, ` + "`" + `Asia/Hebron` + "`" + `, ` + "`" + `Asia/Ho_Chi_Minh` + "`" + `, ` + "`" + `Asia/Hong_Kong` + "`" + `, ` + "`" + `Asia/Hovd` + "`" + `, ` + "`" + `Asia/Irkutsk` + "`" + `, ` + "`" + `Asia/Istanbul` + "`" + `, ` + "`" + `Asia/Jakarta` + "`" + `, ` + "`" + `Asia/Jayapura` + "`" + `, ` + "`" + `Asia/Jerusalem` + "`" + `, ` + "`" + `Asia/Kabul` + "`" + `, ` + "`" + `Asia/Kamchatka` + "`" + `, ` + "`" + `Asia/Karachi` + "`" + `, ` + "`" + `Asia/Kashgar` + "`" + `, ` + "`" + `Asia/Kathmandu` + "`" + `, ` + "`" + `Asia/Katmandu` + "`" + `, ` + "`" + `Asia/Khandyga` + "`" + `, ` + "`" + `Asia/Kolkata` + "`" + `, ` + "`" + `Asia/Krasnoyarsk` + "`" + `, ` + "`" + `Asia/Kuala_Lumpur` + "`" + `, ` + "`" + `Asia/Kuching` + "`" + `, ` + "`" + `Asia/Kuwait` + "`" + `, ` + "`" + `Asia/Macao` + "`" + `, ` + "`" + `Asia/Macau` + "`" + `, ` + "`" + `Asia/Magadan` + "`" + `, ` + "`" + `Asia/Makassar` + "`" + `, ` + "`" + `Asia/Manila` + "`" + `, ` + "`" + `Asia/Muscat` + "`" + `, ` + "`" + `Asia/Nicosia` + "`" + `, ` + "`" + `Asia/Novokuznetsk` + "`" + `, ` + "`" + `Asia/Novosibirsk` + "`" + `, ` + "`" + `Asia/Omsk` + "`" + `, ` + "`" + `Asia/Oral` + "`" + `, ` + "`" + `Asia/Phnom_Penh` + "`" + `, ` + "`" + `Asia/Pontianak` + "`" + `, ` + "`" + `Asia/Pyongyang` + "`" + `, ` + "`" + `Asia/Qatar` + "`" + `, ` + "`" + `Asia/Qostanay` + "`" + `, ` + "`" + `Asia/Qyzylorda` + "`" + `, ` + "`" + `Asia/Rangoon` + "`" + `, ` + "`" + `Asia/Riyadh` + "`" + `, ` + "`" + `Asia/Saigon` + "`" + `, ` + "`" + `Asia/Sakhalin` + "`" + `, ` + "`" + `Asia/Samarkand` + "`" + `, ` + "`" + `Asia/Seoul` + "`" + `, ` + "`" + `Asia/Shanghai` + "`" + `, ` + "`" + `Asia/Singapore` + "`" + `, ` + "`" + `Asia/Srednekolymsk` + "`" + `, ` + "`" + `Asia/Taipei` + "`" + `, ` + "`" + `Asia/Tashkent` + "`" + `, ` + "`" + `Asia/Tbilisi` + "`" + `, ` + "`" + `Asia/Tehran` + "`" + `, ` + "`" + `Asia/Tel_Aviv` + "`" + `, ` + "`" + `Asia/Thimbu` + "`" + `, ` + "`" + `Asia/Thimphu` + "`" + `, ` + "`" + `Asia/Tokyo` + "`" + `, ` + "`" + `Asia/Tomsk` + "`" + `, ` + "`" + `Asia/Ujung_Pandang` + "`" + `, ` + "`" + `Asia/Ulaanbaatar` + "`" + `, ` + "`" + `Asia/Ulan_Bator` + "`" + `, ` + "`" + `Asia/Urumqi` + "`" + `, ` + "`" + `Asia/Ust-Nera` + "`" + `, ` + "`" + `Asia/Vientiane` + "`" + `, ` + "`" + `Asia/Vladivostok` + "`" + `, ` + "`" + `Asia/Yakutsk` + "`" + `, ` + "`" + `Asia/Yangon` + "`" + `, ` + "`" + `Asia/Yekaterinburg` + "`" + `, ` + "`" + `Asia/Yerevan` + "`" + `, ` + "`" + `Atlantic/Azores` + "`" + `, ` + "`" + `Atlantic/Bermuda` + "`" + `, ` + "`" + `Atlantic/Canary` + "`" + `, ` + "`" + `Atlantic/Cape_Verde` + "`" + `, ` + "`" + `Atlantic/Faeroe` + "`" + `, ` + "`" + `Atlantic/Faroe` + "`" + `, ` + "`" + `Atlantic/Jan_Mayen` + "`" + `, ` + "`" + `Atlantic/Madeira` + "`" + `, ` + "`" + `Atlantic/Reykjavik` + "`" + `, ` + "`" + `Atlantic/South_Georgia` + "`" + `, ` + "`" + `Atlantic/St_Helena` + "`" + `, ` + "`" + `Atlantic/Stanley` + "`" + `, ` + "`" + `Australia/ACT` + "`" + `, ` + "`" + `Australia/Adelaide` + "`" + `, ` + "`" + `Australia/Brisbane` + "`" + `, ` + "`" + `Australia/Broken_Hill` + "`" + `, ` + "`" + `Australia/Canberra` + "`" + `, ` + "`" + `Australia/Currie` + "`" + `, ` + "`" + `Australia/Darwin` + "`" + `, ` + "`" + `Australia/Eucla` + "`" + `, ` + "`" + `Australia/Hobart` + "`" + `, ` + "`" + `Australia/LHI` + "`" + `, ` + "`" + `Australia/Lindeman` + "`" + `, ` + "`" + `Australia/Lord_Howe` + "`" + `, ` + "`" + `Australia/Melbourne` + "`" + `, ` + "`" + `Australia/NSW` + "`" + `, ` + "`" + `Australia/North` + "`" + `, ` + "`" + `Australia/Perth` + "`" + `, ` + "`" + `Australia/Queensland` + "`" + `, ` + "`" + `Australia/South` + "`" + `, ` + "`" + `Australia/Sydney` + "`" + `, ` + "`" + `Australia/Tasmania` + "`" + `, ` + "`" + `Australia/Victoria` + "`" + `, ` + "`" + `Australia/West` + "`" + `, ` + "`" + `Australia/Yancowinna` + "`" + `, ` + "`" + `Brazil/Acre` + "`" + `, ` + "`" + `Brazil/DeNoronha` + "`" + `, ` + "`" + `Brazil/East` + "`" + `, ` + "`" + `Brazil/West` + "`" + `, ` + "`" + `CET` + "`" + `, ` + "`" + `CST6CDT` + "`" + `, ` + "`" + `Canada/Atlantic` + "`" + `, ` + "`" + `Canada/Central` + "`" + `, ` + "`" + `Canada/Eastern` + "`" + `, ` + "`" + `Canada/Mountain` + "`" + `, ` + "`" + `Canada/Newfoundland` + "`" + `, ` + "`" + `Canada/Pacific` + "`" + `, ` + "`" + `Canada/Saskatchewan` + "`" + `, ` + "`" + `Canada/Yukon` + "`" + `, ` + "`" + `Chile/Continental` + "`" + `, ` + "`" + `Chile/EasterIsland` + "`" + `, ` + "`" + `Cuba` + "`" + `, ` + "`" + `EET` + "`" + `, ` + "`" + `EST5EDT` + "`" + `, ` + "`" + `Egypt` + "`" + `, ` + "`" + `Eire` + "`" + `, ` + "`" + `Etc/GMT` + "`" + `, ` + "`" + `Etc/GMT+0` + "`" + `, ` + "`" + `Etc/GMT+1` + "`" + `, ` + "`" + `Etc/GMT+10` + "`" + `, ` + "`" + `Etc/GMT+11` + "`" + `, ` + "`" + `Etc/GMT+12` + "`" + `, ` + "`" + `Etc/GMT+2` + "`" + `, ` + "`" + `Etc/GMT+3` + "`" + `, ` + "`" + `Etc/GMT+4` + "`" + `, ` + "`" + `Etc/GMT+5` + "`" + `, ` + "`" + `Etc/GMT+6` + "`" + `, ` + "`" + `Etc/GMT+7` + "`" + `, ` + "`" + `Etc/GMT+8` + "`" + `, ` + "`" + `Etc/GMT+9` + "`" + `, ` + "`" + `Etc/GMT-0` + "`" + `, ` + "`" + `Etc/GMT-1` + "`" + `, ` + "`" + `Etc/GMT-10` + "`" + `, ` + "`" + `Etc/GMT-11` + "`" + `, ` + "`" + `Etc/GMT-12` + "`" + `, ` + "`" + `Etc/GMT-13` + "`" + `, ` + "`" + `Etc/GMT-14` + "`" + `, ` + "`" + `Etc/GMT-2` + "`" + `, ` + "`" + `Etc/GMT-3` + "`" + `, ` + "`" + `Etc/GMT-4` + "`" + `, ` + "`" + `Etc/GMT-5` + "`" + `, ` + "`" + `Etc/GMT-6` + "`" + `, ` + "`" + `Etc/GMT-7` + "`" + `, ` + "`" + `Etc/GMT-8` + "`" + `, ` + "`" + `Etc/GMT-9` + "`" + `, ` + "`" + `Etc/GMT0` + "`" + `, ` + "`" + `Etc/Greenwich` + "`" + `, ` + "`" + `Etc/UCT` + "`" + `, ` + "`" + `Etc/UTC` + "`" + `, ` + "`" + `Etc/Universal` + "`" + `, ` + "`" + `Etc/Zulu` + "`" + `, ` + "`" + `Europe/Amsterdam` + "`" + `, ` + "`" + `Europe/Andorra` + "`" + `, ` + "`" + `Europe/Astrakhan` + "`" + `, ` + "`" + `Europe/Athens` + "`" + `, ` + "`" + `Europe/Belfast` + "`" + `, ` + "`" + `Europe/Belgrade` + "`" + `, ` + "`" + `Europe/Berlin` + "`" + `, ` + "`" + `Europe/Bratislava` + "`" + `, ` + "`" + `Europe/Brussels` + "`" + `, ` + "`" + `Europe/Bucharest` + "`" + `, ` + "`" + `Europe/Budapest` + "`" + `, ` + "`" + `Europe/Busingen` + "`" + `, ` + "`" + `Europe/Chisinau` + "`" + `, ` + "`" + `Europe/Copenhagen` + "`" + `, ` + "`" + `Europe/Dublin` + "`" + `, ` + "`" + `Europe/Gibraltar` + "`" + `, ` + "`" + `Europe/Guernsey` + "`" + `, ` + "`" + `Europe/Helsinki` + "`" + `, ` + "`" + `Europe/Isle_of_Man` + "`" + `, ` + "`" + `Europe/Istanbul` + "`" + `, ` + "`" + `Europe/Jersey` + "`" + `, ` + "`" + `Europe/Kaliningrad` + "`" + `, ` + "`" + `Europe/Kiev` + "`" + `, ` + "`" + `Europe/Kirov` + "`" + `, ` + "`" + `Europe/Lisbon` + "`" + `, ` + "`" + `Europe/Ljubljana` + "`" + `, ` + "`" + `Europe/London` + "`" + `, ` + "`" + `Europe/Luxembourg` + "`" + `, ` + "`" + `Europe/Madrid` + "`" + `, ` + "`" + `Europe/Malta` + "`" + `, ` + "`" + `Europe/Mariehamn` + "`" + `, ` + "`" + `Europe/Minsk` + "`" + `, ` + "`" + `Europe/Monaco` + "`" + `, ` + "`" + `Europe/Moscow` + "`" + `, ` + "`" + `Europe/Nicosia` + "`" + `, ` + "`" + `Europe/Oslo` + "`" + `, ` + "`" + `Europe/Paris` + "`" + `, ` + "`" + `Europe/Podgorica` + "`" + `, ` + "`" + `Europe/Prague` + "`" + `, ` + "`" + `Europe/Riga` + "`" + `, ` + "`" + `Europe/Rome` + "`" + `, ` + "`" + `Europe/Samara` + "`" + `, ` + "`" + `Europe/San_Marino` + "`" + `, ` + "`" + `Europe/Sarajevo` + "`" + `, ` + "`" + `Europe/Saratov` + "`" + `, ` + "`" + `Europe/Simferopol` + "`" + `, ` + "`" + `Europe/Skopje` + "`" + `, ` + "`" + `Europe/Sofia` + "`" + `, ` + "`" + `Europe/Stockholm` + "`" + `, ` + "`" + `Europe/Tallinn` + "`" + `, ` + "`" + `Europe/Tirane` + "`" + `, ` + "`" + `Europe/Tiraspol` + "`" + `, ` + "`" + `Europe/Ulyanovsk` + "`" + `, ` + "`" + `Europe/Uzhgorod` + "`" + `, ` + "`" + `Europe/Vaduz` + "`" + `, ` + "`" + `Europe/Vatican` + "`" + `, ` + "`" + `Europe/Vienna` + "`" + `, ` + "`" + `Europe/Vilnius` + "`" + `, ` + "`" + `Europe/Volgograd` + "`" + `, ` + "`" + `Europe/Warsaw` + "`" + `, ` + "`" + `Europe/Zagreb` + "`" + `, ` + "`" + `Europe/Zaporozhye` + "`" + `, ` + "`" + `Europe/Zurich` + "`" + `, ` + "`" + `GB` + "`" + `, ` + "`" + `GB-Eire` + "`" + `, ` + "`" + `GMT` + "`" + `, ` + "`" + `GMT0` + "`" + `, ` + "`" + `Greenwich` + "`" + `, ` + "`" + `Hongkong` + "`" + `, ` + "`" + `Iceland` + "`" + `, ` + "`" + `Indian/Antananarivo` + "`" + `, ` + "`" + `Indian/Chagos` + "`" + `, ` + "`" + `Indian/Christmas` + "`" + `, ` + "`" + `Indian/Cocos` + "`" + `, ` + "`" + `Indian/Comoro` + "`" + `, ` + "`" + `Indian/Kerguelen` + "`" + `, ` + "`" + `Indian/Mahe` + "`" + `, ` + "`" + `Indian/Maldives` + "`" + `, ` + "`" + `Indian/Mauritius` + "`" + `, ` + "`" + `Indian/Mayotte` + "`" + `, ` + "`" + `Indian/Reunion` + "`" + `, ` + "`" + `Iran` + "`" + `, ` + "`" + `Israel` + "`" + `, ` + "`" + `Jamaica` + "`" + `, ` + "`" + `Japan` + "`" + `, ` + "`" + `Kwajalein` + "`" + `, ` + "`" + `Libya` + "`" + `, ` + "`" + `MET` + "`" + `, ` + "`" + `MST7MDT` + "`" + `, ` + "`" + `Mexico/BajaNorte` + "`" + `, ` + "`" + `Mexico/BajaSur` + "`" + `, ` + "`" + `Mexico/General` + "`" + `, ` + "`" + `NZ` + "`" + `, ` + "`" + `NZ-CHAT` + "`" + `, ` + "`" + `Navajo` + "`" + `, ` + "`" + `PRC` + "`" + `, ` + "`" + `PST8PDT` + "`" + `, ` + "`" + `Pacific/Apia` + "`" + `, ` + "`" + `Pacific/Auckland` + "`" + `, ` + "`" + `Pacific/Bougainville` + "`" + `, ` + "`" + `Pacific/Chatham` + "`" + `, ` + "`" + `Pacific/Chuuk` + "`" + `, ` + "`" + `Pacific/Easter` + "`" + `, ` + "`" + `Pacific/Efate` + "`" + `, ` + "`" + `Pacific/Enderbury` + "`" + `, ` + "`" + `Pacific/Fakaofo` + "`" + `, ` + "`" + `Pacific/Fiji` + "`" + `, ` + "`" + `Pacific/Funafuti` + "`" + `, ` + "`" + `Pacific/Galapagos` + "`" + `, ` + "`" + `Pacific/Gambier` + "`" + `, ` + "`" + `Pacific/Guadalcanal` + "`" + `, ` + "`" + `Pacific/Guam` + "`" + `, ` + "`" + `Pacific/Honolulu` + "`" + `, ` + "`" + `Pacific/Johnston` + "`" + `, ` + "`" + `Pacific/Kanton` + "`" + `, ` + "`" + `Pacific/Kiritimati` + "`" + `, ` + "`" + `Pacific/Kosrae` + "`" + `, ` + "`" + `Pacific/Kwajalein` + "`" + `, ` + "`" + `Pacific/Majuro` + "`" + `, ` + "`" + `Pacific/Marquesas` + "`" + `, ` + "`" + `Pacific/Midway` + "`" + `, ` + "`" + `Pacific/Nauru` + "`" + `, ` + "`" + `Pacific/Niue` + "`" + `, ` + "`" + `Pacific/Norfolk` + "`" + `, ` + "`" + `Pacific/Noumea` + "`" + `, ` + "`" + `Pacific/Pago_Pago` + "`" + `, ` + "`" + `Pacific/Palau` + "`" + `, ` + "`" + `Pacific/Pitcairn` + "`" + `, ` + "`" + `Pacific/Pohnpei` + "`" + `, ` + "`" + `Pacific/Ponape` + "`" + `, ` + "`" + `Pacific/Port_Moresby` + "`" + `, ` + "`" + `Pacific/Rarotonga` + "`" + `, ` + "`" + `Pacific/Saipan` + "`" + `, ` + "`" + `Pacific/Samoa` + "`" + `, ` + "`" + `Pacific/Tahiti` + "`" + `, ` + "`" + `Pacific/Tarawa` + "`" + `, ` + "`" + `Pacific/Tongatapu` + "`" + `, ` + "`" + `Pacific/Truk` + "`" + `, ` + "`" + `Pacific/Wake` + "`" + `, ` + "`" + `Pacific/Wallis` + "`" + `, ` + "`" + `Pacific/Yap` + "`" + `, ` + "`" + `Poland` + "`" + `, ` + "`" + `Portugal` + "`" + `, ` + "`" + `ROK` + "`" + `, ` + "`" + `Singapore` + "`" + `, ` + "`" + `SystemV/AST4` + "`" + `, ` + "`" + `SystemV/AST4ADT` + "`" + `, ` + "`" + `SystemV/CST6` + "`" + `, ` + "`" + `SystemV/CST6CDT` + "`" + `, ` + "`" + `SystemV/EST5` + "`" + `, ` + "`" + `SystemV/EST5EDT` + "`" + `, ` + "`" + `SystemV/HST10` + "`" + `, ` + "`" + `SystemV/MST7` + "`" + `, ` + "`" + `SystemV/MST7MDT` + "`" + `, ` + "`" + `SystemV/PST8` + "`" + `, ` + "`" + `SystemV/PST8PDT` + "`" + `, ` + "`" + `SystemV/YST9` + "`" + `, ` + "`" + `SystemV/YST9YDT` + "`" + `, ` + "`" + `Turkey` + "`" + `, ` + "`" + `UCT` + "`" + `, ` + "`" + `US/Alaska` + "`" + `, ` + "`" + `US/Aleutian` + "`" + `, ` + "`" + `US/Arizona` + "`" + `, ` + "`" + `US/Central` + "`" + `, ` + "`" + `US/East-Indiana` + "`" + `, ` + "`" + `US/Eastern` + "`" + `, ` + "`" + `US/Hawaii` + "`" + `, ` + "`" + `US/Indiana-Starke` + "`" + `, ` + "`" + `US/Michigan` + "`" + `, ` + "`" + `US/Mountain` + "`" + `, ` + "`" + `US/Pacific` + "`" + `, ` + "`" + `US/Samoa` + "`" + `, ` + "`" + `UTC` + "`" + `, ` + "`" + `Universal` + "`" + `, ` + "`" + `W-SU` + "`" + `, ` + "`" + `WET` + "`" + `, ` + "`" + `Zulu` + "`" + `, ` + "`" + `EST` + "`" + `, ` + "`" + `HST` + "`" + `, ` + "`" + `MST` + "`" + `, ` + "`" + `ACT` + "`" + `, ` + "`" + `AET` + "`" + `, ` + "`" + `AGT` + "`" + `, ` + "`" + `ART` + "`" + `, ` + "`" + `AST` + "`" + `, ` + "`" + `BET` + "`" + `, ` + "`" + `BST` + "`" + `, ` + "`" + `CAT` + "`" + `, ` + "`" + `CNT` + "`" + `, ` + "`" + `CST` + "`" + `, ` + "`" + `CTT` + "`" + `, ` + "`" + `EAT` + "`" + `, ` + "`" + `ECT` + "`" + `, ` + "`" + `IET` + "`" + `, ` + "`" + `IST` + "`" + `, ` + "`" + `JST` + "`" + `, ` + "`" + `MIT` + "`" + `, ` + "`" + `NET` + "`" + `, ` + "`" + `NST` + "`" + `, ` + "`" + `PLT` + "`" + `, ` + "`" + `PNT` + "`" + `, ` + "`" + `PRT` + "`" + `, ` + "`" + `PST` + "`" + `, ` + "`" + `SST` + "`" + `, ` + "`" + `VST` + "`" + ` ## Importing resource: To import an alert you'll need to specify your Logz.io alert's id. For example, if you have in your Terraform configuration the following: ` + "`" + `` + "`" + `` + "`" + `hcl resource "logzio_alert_v2" "imported" { } ` + "`" + `` + "`" + `` + "`" + ` And your alert's id is 123456, then your import command should be: ` + "`" + `` + "`" + `` + "`" + `bash terraform import logzio_alert_v2.imported 123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_archive_logs",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_type",
					Description: `(String) Specifies the storage provider. Applicable values: ` + "`" + `S3` + "`" + `, ` + "`" + `BLOB` + "`" + `. ### Optional:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) Defaults to ` + "`" + `true` + "`" + `. If ` + "`" + `true` + "`" + `, archiving is currently enabled.`,
				},
				resource.Attribute{
					Name:        "compressed",
					Description: `(Boolean) Defaults to ` + "`" + `true` + "`" + `. If ` + "`" + `true` + "`" + `, logs are compressed before they are archived. #### Required if ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "aws_credentials_type",
					Description: `(String) Specifies which credentials will be used for authentication. The options are either ` + "`" + `KEYS` + "`" + ` or ` + "`" + `IAM` + "`" + `. Authentication with S3 Secret Credentials is supported for backward compatibility. IAM roles are strongly recommended.`,
				},
				resource.Attribute{
					Name:        "aws_s3_path",
					Description: `(String) Specify a path to the`,
				},
				resource.Attribute{
					Name:        "aws_s3_iam_credentials_arn",
					Description: `(String) Applicable when ` + "`" + `aws_credentials_type` + "`" + ` is ` + "`" + `IAM` + "`" + `. Amazon Resource Name (ARN) to uniquely identify the S3 bucket.`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(String) Applicable when ` + "`" + `aws_credentials_type` + "`" + ` is ` + "`" + `KEYS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aws_secret_key",
					Description: `(String) Applicable when ` + "`" + `aws_credentials_type` + "`" + ` is ` + "`" + `KEYS` + "`" + `. ##### Required if ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "azure_tenant_id",
					Description: `(String) Azure Directory (tenant) ID. The Tenant ID of the AD app. Go to`,
				},
				resource.Attribute{
					Name:        "azure_client_id",
					Description: `(String) Azure application (client) ID. The Client ID of the AD app, found under the App Overview page. Go to`,
				},
				resource.Attribute{
					Name:        "azure_client_secret",
					Description: `(String) Azure client secret. Password of the Client secret, found in the app's`,
				},
				resource.Attribute{
					Name:        "azure_account_name",
					Description: `(String) Azure Storage account name. Name of the storage account that holds the container where the logs will be archived.`,
				},
				resource.Attribute{
					Name:        "azure_container_name",
					Description: `(String) Name of the container in the Storage account. This is where the logs will be archived. ##### Optional if ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "azure_blob_path",
					Description: `(String) Optional virtual sub-folder specifiying a path within the container. Logs will be archived under the path “{container-name}/{virtual sub-folder}”. Avoid leading and trailing slashes (/). For example, the prefix “region1” is good, but “/region1/” is not. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "archive_id",
					Description: `(String) Archive ID in the Logz.io database. ## Importing resource: To import an archive you'll need to specify your archive's id. For example, if you have in your Terraform configuration the following: ` + "`" + `` + "`" + `` + "`" + `hcl resource "logzio_archive_logs" "imported" { ... } ` + "`" + `` + "`" + `` + "`" + ` And your archives's id is 123456, then your import command should be: ` + "`" + `` + "`" + `` + "`" + `bash terraform import logzio_archive_logs.imported 123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_authentication_groups",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "authentication_group",
					Description: `(Block List) Sets details for the authentication groups. Must be at least one ` + "`" + `authentication_group` + "`" + ` in a resource. #### Nested schema for ` + "`" + `authentication_group` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(String) Name of authentication group.`,
				},
				resource.Attribute{
					Name:        "user_role",
					Description: `(String) User role for that group. Can be one of the following: ` + "`" + `USER_ROLE_ACCOUNT_ADMIN` + "`" + `, ` + "`" + `USER_ROLE_REGULAR` + "`" + `, ` + "`" + `USER_ROLE_READONLY` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "manage_groups_id",
					Description: `(String) Id for the resource, generated by the provider. It has no real use outside of Terraform. ## Importing resource: The import command expects an id for the import command, but the authentication groups API does not work with ids. Because of that, if you wish to import authentication groups, you can use whichever numeric id you'd like. That id won't affect anything, it will only be the id of the resource in Terraform. ` + "`" + `` + "`" + `` + "`" + `bash terraform import logzio_authentication_groups.imported 1234 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_drop_filter",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "field_conditions",
					Description: `(Block list) Filters for an exact match of a field:value pair.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `(String) Filters for the [log type](https://docs.logz.io/user-guide/log-shipping/built-in-log-types.html). Omit or leave empty if you want this filter to apply to all types.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Boolean) If true, the drop filter is active and logs that match the filter are dropped before indexing. If false, the drop filter is disabled.`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `(String) Exact field name in your Kibana mapping for the selected ` + "`" + `log_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Object) Exact field value. The filter looks for an exact value match of the entire object.`,
				},
				resource.Attribute{
					Name:        "drop_filter_id",
					Description: `(String) Drop filter ID in the Logz.io database. ### Import drop filter as resource You can import drop filters as follows: ` + "`" + `` + "`" + `` + "`" + ` terraform import logzio_drop_filter.my_filter <DROP-FILTER-ID> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_endpoint",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `(Required) Specifies the endpoint resource type: ` + "`" + `custom` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `bigpanda` + "`" + `, ` + "`" + `datadog` + "`" + `, ` + "`" + `victorops` + "`" + `, ` + "`" + `opsgenie` + "`" + `, ` + "`" + `servicenow` + "`" + `, ` + "`" + `microsoftteams` + "`" + `. Use the appropriate parameters for your selected endpoint type.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) Name of the endpoint.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Detailed description of the endpoint.`,
				},
				resource.Attribute{
					Name:        "slack",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `slack` + "`" + `. Manages a webhook to a specific Slack channel.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Slack webhook URL to a specific Slack channel.`,
				},
				resource.Attribute{
					Name:        "pagerduty",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `pagerduty` + "`" + `. Manages a webhook to PagerDuty.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `API key generated from PagerDuty for the purpose of the integration.`,
				},
				resource.Attribute{
					Name:        "bigpanda",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `bigpanda` + "`" + `. Manages a webhook to BigPanda.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `API authentication token from BigPanda.`,
				},
				resource.Attribute{
					Name:        "app_key",
					Description: `Application key from BigPanda.`,
				},
				resource.Attribute{
					Name:        "datadog",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `datadog` + "`" + `. Manages a webhook to Datadog.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `API key from Datadog.`,
				},
				resource.Attribute{
					Name:        "victorops",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `victorops` + "`" + `. Manages a webhook to VictorOps.`,
				},
				resource.Attribute{
					Name:        "routing_key",
					Description: `Alert routing key from VictorOps.`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `VictorOps REST API ` + "`" + `message_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_api_key",
					Description: `API key from VictorOps.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `custom` + "`" + `. Manages a custom webhook for your integration of choice.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Specifies the URL destination.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `Selects the HTTP request method.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `Header parameters for the request. String, sent as comma-separated key-value pairs.`,
				},
				resource.Attribute{
					Name:        "body_template",
					Description: `string of JSON object that serves as the template for the message body.`,
				},
				resource.Attribute{
					Name:        "opsgenie",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `opsgenie` + "`" + `. Manages a webhook to OpsGenie.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `API key from OpsGenie, see https://docs.opsgenie.com/docs/logz-io-integration.`,
				},
				resource.Attribute{
					Name:        "servicenow",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `servicenow` + "`" + `. Manages a webhook to ServiceNow.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `ServiceNow user name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `ServiceNow password.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Provide your instance URL to connect to your existing ServiceNow instance, i.e. https://xxxxxxxxx.service-now.com/.`,
				},
				resource.Attribute{
					Name:        "microsoftteams",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `microsoftteams` + "`" + `. Manages a webhook to Microsoft Teams.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Your Microsoft Teams webhook URL, see https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the notification endpoint. ### Import endpoint as resource You can import endpoint as follows: ` + "`" + `` + "`" + `` + "`" + ` terraform import logzio_endpoint.my_endpoint <ENDPOINT-ID> ` + "`" + `` + "`" + `` + "`" + ` ## Endpoints used Logz.io integrates with:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the notification endpoint. ### Import endpoint as resource You can import endpoint as follows: ` + "`" + `` + "`" + `` + "`" + ` terraform import logzio_endpoint.my_endpoint <ENDPOINT-ID> ` + "`" + `` + "`" + `` + "`" + ` ## Endpoints used Logz.io integrates with:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_grafana_dashboard",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_uid",
					Description: `(String) The unique identifier (uid) of a folder to store your dashboard. You cannot use ` + "`" + `General` + "`" + ` folder or the folder generated by logz.io - ` + "`" + `Logz.io Dashboards` + "`" + ` - to place your alerts.`,
				},
				resource.Attribute{
					Name:        "dashboard_json",
					Description: `(String) The complete dashboard model, to create a new dashboard, in a JSON format.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(String) A commit message for the version history.`,
				},
				resource.Attribute{
					Name:        "overwrite",
					Description: `(Boolean) Set to true if you want to overwrite existing dashboard with newer version. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "dashboard_id",
					Description: `(Int) The identifier (id) of a dashboard.`,
				},
				resource.Attribute{
					Name:        "dashboard_uid",
					Description: `(String) The unique identifier (uid) of a dashboard.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(String) Dashboard url.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Int) Dashboard version. ### Import Logz.io Grafana Dashboard as Terraform resource You can import existing dashboard as follows: ` + "`" + `` + "`" + `` + "`" + ` terraform import logzio_grafana_dashboard.my_dashboard <FETCHER-ID> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dashboard_id",
					Description: `(Int) The identifier (id) of a dashboard.`,
				},
				resource.Attribute{
					Name:        "dashboard_uid",
					Description: `(String) The unique identifier (uid) of a dashboard.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(String) Dashboard url.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Int) Dashboard version. ### Import Logz.io Grafana Dashboard as Terraform resource You can import existing dashboard as follows: ` + "`" + `` + "`" + `` + "`" + ` terraform import logzio_grafana_dashboard.my_dashboard <FETCHER-ID> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_kibana_object",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "kibana_version",
					Description: `(String) The version of Kibana used at the time of export.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(String) Exported Kibana objects. Should be a valid JSON that was retrieved from an export operation of the API.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_log_shipping_token",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(String) Descriptive name for this log shipping token. ### Optional:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) To enable this log shipping token, true. To disable, false.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `(Integer) The log shipping token's ID.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(String) The log shipping token itself.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `(Integer) Unix timestamp of when this log shipping token was last updated.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `(String) Email address of the last user to update this log shipping token.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Integer) Unix timestamp of when this log shipping token was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `(String) Email address of the user who created this log shipping token.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_restore_logs",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(String) Name of the restored account.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(String) Owner of the restored account. Effectively, the user's email address.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Integer) UNIX timestamp in milliseconds specifying the earliest logs to be restored.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Integer) UNIX timestamp in milliseconds specifying the latest logs to be restored.`,
				},
				resource.Attribute{
					Name:        "restore_operation_id",
					Description: `(Integer) ID of the restore operation in Logz.io.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Integer) ID of the restored account in Logz.io.`,
				},
				resource.Attribute{
					Name:        "restored_volume_gb",
					Description: `(Float) Volume of data restored so far. If the restore operation is still in progress, this will be continuously updated.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) Returns the current status of the restored account. See [documentation](https://docs.logz.io/api/#operation/getRestoreRequestByIdApi) for more info about the possible statuses and their meaning.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Integer) Timestamp when the restore process was created and entered the queue. Since only one account can be restored at a time, the process may not initiate immediately.`,
				},
				resource.Attribute{
					Name:        "started_at",
					Description: `(Integer) UNIX timestamp in milliseconds when the restore process initiated.`,
				},
				resource.Attribute{
					Name:        "finished_at",
					Description: `(Integer) UNIX timestamp in milliseconds when the restore process completed.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `(Integer) UNIX timestamp in milliseconds specifying when the account is due to expire. Restored accounts expire automatically after a number of days, as specified in the account's terms. ## Importing resource: To import a restore operation you'll need to specify the restore's id. For example, if you have in your Terraform the following configuration: ` + "`" + `` + "`" + `` + "`" + `hcl resource "logzio_restore_logs" "imported" { ... } ` + "`" + `` + "`" + `` + "`" + ` And your restore operation id is 123456, then your import command should be: ` + "`" + `` + "`" + `` + "`" + `bash terraform import logzio_restore_logs.imported 123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_s3_fetcher",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(String) AWS S3 bucket access key. Not applicable if you choose to authenticate with ` + "`" + `aws_arn` + "`" + `. If you choose to authenticate with AWS keys, both ` + "`" + `aws_access_key` + "`" + ` and ` + "`" + `aws_secret key` + "`" + ` must be set.`,
				},
				resource.Attribute{
					Name:        "aws_secret_key",
					Description: `(String) AWS S3 bucket secret key. Not applicable if you choose to authenticate with ` + "`" + `aws_arn` + "`" + `. If you choose to authenticate with AWS keys, both ` + "`" + `aws_access_key` + "`" + ` and ` + "`" + `aws_secret key` + "`" + ` must be set.`,
				},
				resource.Attribute{
					Name:        "aws_arn",
					Description: `(String) Amazon Resource Name (ARN) to uniquely identify the S3 bucket. To generate a new ARN, create a new IAM Role in your AWS admin console. Not applicable if you choose to authenticate with AWS keys (access key & secret key).`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(String) AWS S3 bucket name.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Boolean) If true, the S3 bucket connector is active and logs are being fetched to Logz.io. If false, the connector is disabled.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `(String) Bucket's region. Specify one supported AWS region: ` + "`" + `US_EAST_1` + "`" + `, ` + "`" + `US_EAST_2` + "`" + `, ` + "`" + `US_WEST_1` + "`" + `, ` + "`" + `US_WEST_2` + "`" + `, ` + "`" + `EU_WEST_1` + "`" + `, ` + "`" + `EU_WEST_2` + "`" + `, ` + "`" + `EU_WEST_3` + "`" + `, ` + "`" + `EU_CENTRAL_1` + "`" + `, ` + "`" + `AP_NORTHEAST_1` + "`" + `, ` + "`" + `AP_NORTHEAST_2` + "`" + `, ` + "`" + `AP_SOUTHEAST_1` + "`" + `, ` + "`" + `AP_SOUTHEAST_2` + "`" + `, ` + "`" + `SA_EAST_1` + "`" + `, ` + "`" + `AP_SOUTH_1` + "`" + `, ` + "`" + `CA_CENTRAL_1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logs_type",
					Description: `(String) Specifies the log type being sent to Logz.io. Determines the parsing pipeline used to parse and map the logs. [Learn more about parsing options supported by Logz.io](https://docs.logz.io/user-guide/log-shipping/built-in-log-types.html). Allowed values: ` + "`" + `elb` + "`" + `, ` + "`" + `vpcflow` + "`" + `, ` + "`" + `S3Access` + "`" + `, ` + "`" + `cloudfront` + "`" + `. ### Optional:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(String) Prefix of the AWS S3 bucket.`,
				},
				resource.Attribute{
					Name:        "add_s3_object_key_as_log_field",
					Description: `(Boolean) Defaults to ` + "`" + `false` + "`" + `. If ` + "`" + `true` + "`" + `, enriches logs with a new field detailing the S3 object key. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "fetcher_id",
					Description: `(Int) Id of the created fetcher. ### Import S3 Fetcher as Terraform resource You can import existing S3 Fetcher as follows: ` + "`" + `` + "`" + `` + "`" + ` terraform import logzio_s3_fetcher.my_fetcher <FETCHER-ID> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fetcher_id",
					Description: `(Int) Id of the created fetcher. ### Import S3 Fetcher as Terraform resource You can import existing S3 Fetcher as follows: ` + "`" + `` + "`" + `` + "`" + ` terraform import logzio_s3_fetcher.my_fetcher <FETCHER-ID> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_subaccount",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(String) Email address of an existing admin user on the main account which will also become the admin of the subaccount being created.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(String) Name of the subaccount.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `(Integer) Number of days that log data is retained.`,
				},
				resource.Attribute{
					Name:        "sharing_objects_accounts",
					Description: `(List) IDs of accounts that can access the account's Kibana objects. Can be an empty array. ### Optional`,
				},
				resource.Attribute{
					Name:        "max_daily_gb",
					Description: `(Float) Maximum daily log volume that the subaccount can index, in GB.`,
				},
				resource.Attribute{
					Name:        "searchable",
					Description: `(Boolean) False by default. Determines if other accounts can search logs indexed by the subaccount.`,
				},
				resource.Attribute{
					Name:        "accessible",
					Description: `(Boolean) False by default. Determines if users of main account can access the subaccount.`,
				},
				resource.Attribute{
					Name:        "doc_size_setting",
					Description: `(Boolean) False by default. If enabled, Logz.io adds a ` + "`" + `LogSize` + "`" + ` field to record the size of the log line in bytes, for the purpose of managing account utilization. [Learn more about managing account usage](https://docs.logz.io/user-guide/accounts/manage-account-usage.html#enabling-account-utilization-metrics-and-log-size)`,
				},
				resource.Attribute{
					Name:        "utilization_enabled",
					Description: `(Boolean) If enabled, account utilization metrics and expected utilization at the current indexing rate are measured at a pre-defined sampling rate. Useful for managing account utilization and avoiding running out of capacity. [Learn more about managing account capacity](https://docs.logz.io/user-guide/accounts/manage-account-usage.html).`,
				},
				resource.Attribute{
					Name:        "frequency_minutes",
					Description: `(Int) Determines the sampling rate in minutes of the utilization.`,
				},
				resource.Attribute{
					Name:        "flexible",
					Description: `(Boolean) Defaults to false. Whether the sub account that created is flexible or not. Can be set to flexible only if the main account is flexible.`,
				},
				resource.Attribute{
					Name:        "reserved_daily_gb",
					Description: `(Float) The maximum volume of data that an account can index per calendar day. Depends on ` + "`" + `flexible` + "`" + `. For further info see [the docs](https://docs.logz.io/api/#operation/createTimeBasedAccount). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `ID of the subaccount.`,
				},
				resource.Attribute{
					Name:        "account_token",
					Description: `Log shipping token for the subaccount. [Learn more](https://docs.logz.io/user-guide/tokens/log-shipping-tokens/)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fullname",
					Description: `(Required) First and last name of the user.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username credential.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) String. User role. Can be ` + "`" + `USER_ROLE_READONLY` + "`" + `, ` + "`" + `USER_ROLE_REGULAR` + "`" + ` or ` + "`" + `USER_ROLE_ACCOUNT_ADMIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Required) If ` + "`" + `true` + "`" + `, the user is active, if ` + "`" + `false` + "`" + `, the user is suspended.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) Logz.io account ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user in the Logz.io platform. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"logzio_alert_v2":              0,
		"logzio_archive_logs":          1,
		"logzio_authentication_groups": 2,
		"logzio_drop_filter":           3,
		"logzio_endpoint":              4,
		"logzio_grafana_dashboard":     5,
		"logzio_kibana_object":         6,
		"logzio_log_shipping_token":    7,
		"logzio_restore_logs":          8,
		"logzio_s3_fetcher":            9,
		"logzio_subaccount":            10,
		"logzio_user":                  11,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
