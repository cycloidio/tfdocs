package linode

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "linode_account",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode account.`,
			Description: `\_account

Provides information about a Linode account.

This data source should not be used in conjuction with the ` + "`" + `LINODE_DEBUG` + "`" + ` option.  See the [debugging notes](/providers/linode/linode/latest/docs#debugging) for more details.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The email address for this Account, for account management communications, and may be used for other communications as configured.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The first name of the person associated with this Account.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The last name of the person associated with this Account.`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `The company name associated with this Account.`,
				},
				resource.Attribute{
					Name:        "address_1",
					Description: `First line of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "address_2",
					Description: `Second line of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `The phone number associated with this Account.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city for this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `If billing address is in the United States, this is the State portion of the Account's billing address. If the address is outside the US, this is the Province associated with the Account's billing address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The two-letter country code of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "zip",
					Description: `The zip code of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "balance",
					Description: `This Account's balance, in US dollars.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The email address for this Account, for account management communications, and may be used for other communications as configured.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The first name of the person associated with this Account.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The last name of the person associated with this Account.`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `The company name associated with this Account.`,
				},
				resource.Attribute{
					Name:        "address_1",
					Description: `First line of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "address_2",
					Description: `Second line of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `The phone number associated with this Account.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city for this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `If billing address is in the United States, this is the State portion of the Account's billing address. If the address is outside the US, this is the Province associated with the Account's billing address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The two-letter country code of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "zip",
					Description: `The zip code of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "balance",
					Description: `This Account's balance, in US dollars.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_account_login",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode account login.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique ID of this login object. ## Attributes Reference The Linode Account Login resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this login object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The remote IP address that requested the login.`,
				},
				resource.Attribute{
					Name:        "datetime",
					Description: `When the login was initiated.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the User that was logged into.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `True if the User that was logged into was a restricted User, false otherwise.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this login object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The remote IP address that requested the login.`,
				},
				resource.Attribute{
					Name:        "datetime",
					Description: `When the login was initiated.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the User that was logged into.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `True if the User that was logged into was a restricted User, false otherwise.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_account_logins",
			Category:         "Data Sources",
			ShortDescription: `Provides information about Linode account logins that match a set of filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by. See the [Filterable Fields section](#filterable-fields) for a complete list of filterable fields.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each Linode account login will be stored in the ` + "`" + `logins` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this login object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The remote IP address that requested the login.`,
				},
				resource.Attribute{
					Name:        "datetime",
					Description: `When the login was initiated.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the User that was logged into.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `True if the User that was logged into was a restricted User, false otherwise. ## Filterable Fields`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this login object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The remote IP address that requested the login.`,
				},
				resource.Attribute{
					Name:        "datetime",
					Description: `When the login was initiated.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the User that was logged into.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `True if the User that was logged into was a restricted User, false otherwise. ## Filterable Fields`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_database_backups",
			Category:         "Data Sources",
			ShortDescription: `Provides information about Linode Database Backups that match a set of filters.`,
			Description: `\_database\_backups

Provides information about Linode Database Backups that match a set of filters.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "database_id",
					Description: `(Required) The ID of the database to retrieve backups for.`,
				},
				resource.Attribute{
					Name:        "database_type",
					Description: `(Required) The type of the database to retrieve backups for. (` + "`" + `mysql` + "`" + `, ` + "`" + `postgresql` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If true, only the latest backup will be returned.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) The attribute to order the results by. (` + "`" + `created` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order in which results should be returned. (` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `; default ` + "`" + `asc` + "`" + `) ### Filter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each backup will be stored in the ` + "`" + `backups` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `A time value given in a combined date and time format that represents when the database backup was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database backup object.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The database backup’s label, for display purposes only.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of database backup, determined by how the backup was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `A time value given in a combined date and time format that represents when the database backup was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database backup object.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The database backup’s label, for display purposes only.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of database backup, determined by how the backup was created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_database_engines",
			Category:         "Data Sources",
			ShortDescription: `Provides information about Linode Managed Database engines.`,
			Description: `\_database\_engines

Provides information about Linode Managed Database engines that match a set of filters.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If true, only the latest engine version will be returned.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) The attribute to order the results by. (` + "`" + `version` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order in which results should be returned. (` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `; default ` + "`" + `asc` + "`" + `) ### Filter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each engine will be stored in the ` + "`" + `engines` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The Managed Database engine type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Managed Database engine ID in engine/version format.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Managed Database engine version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "engine",
					Description: `The Managed Database engine type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Managed Database engine ID in engine/version format.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Managed Database engine version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_database_mysql",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Linode MySQL Database.`,
			Description: `\_database\_mysql

Provides information about a Linode MySQL Database.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "database_id",
					Description: `The ID of the MySQL database. ## Attributes Reference The ` + "`" + `linode_database_mysql` + "`" + ` data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "allow_list",
					Description: `A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `The base64-encoded SSL CA certificate for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "cluster_size",
					Description: `The number of Linode Instance nodes deployed to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Managed Database was created.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the Managed Databases is encrypted.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The Managed Database engine. (e.g. ` + "`" + `mysql` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `The Managed Database engine in engine/version format. (e.g. ` + "`" + `mysql/8.0.26` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "host_primary",
					Description: `The primary host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "host_secondary",
					Description: `The secondary/private network host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A unique, user-defined string referring to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region that hosts this Linode Managed Database.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `The randomly-generated root password for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "root_username",
					Description: `The root username for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `The replication method used for the Managed Database. (` + "`" + `none` + "`" + `, ` + "`" + `asynch` + "`" + `, ` + "`" + `semi_synch` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ssl_connection",
					Description: `Whether to require SSL credentials to establish a connection to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The operating status of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Linode Instance type used for the nodes of the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When this Managed Database was last updated.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Managed Database engine version. (e.g. ` + "`" + `v8.0.26` + "`" + `) ## updates The following arguments are exported by the ` + "`" + `updates` + "`" + ` specification block:`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `The day to perform maintenance. (` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ...)`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `The maximum maintenance window time in hours. (` + "`" + `1` + "`" + `..` + "`" + `3` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Whether maintenance occurs on a weekly or monthly basis. (` + "`" + `weekly` + "`" + `, ` + "`" + `monthly` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "hour_of_day",
					Description: `The hour to begin maintenance based in UTC time. (` + "`" + `0` + "`" + `..` + "`" + `23` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "week_of_month",
					Description: `The week of the month to perform monthly frequency updates. Required for ` + "`" + `monthly` + "`" + ` frequency updates. (` + "`" + `1` + "`" + `..` + "`" + `4` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_list",
					Description: `A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `The base64-encoded SSL CA certificate for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "cluster_size",
					Description: `The number of Linode Instance nodes deployed to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Managed Database was created.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the Managed Databases is encrypted.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The Managed Database engine. (e.g. ` + "`" + `mysql` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `The Managed Database engine in engine/version format. (e.g. ` + "`" + `mysql/8.0.26` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "host_primary",
					Description: `The primary host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "host_secondary",
					Description: `The secondary/private network host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A unique, user-defined string referring to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region that hosts this Linode Managed Database.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `The randomly-generated root password for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "root_username",
					Description: `The root username for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `The replication method used for the Managed Database. (` + "`" + `none` + "`" + `, ` + "`" + `asynch` + "`" + `, ` + "`" + `semi_synch` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ssl_connection",
					Description: `Whether to require SSL credentials to establish a connection to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The operating status of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Linode Instance type used for the nodes of the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When this Managed Database was last updated.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Managed Database engine version. (e.g. ` + "`" + `v8.0.26` + "`" + `) ## updates The following arguments are exported by the ` + "`" + `updates` + "`" + ` specification block:`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `The day to perform maintenance. (` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ...)`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `The maximum maintenance window time in hours. (` + "`" + `1` + "`" + `..` + "`" + `3` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Whether maintenance occurs on a weekly or monthly basis. (` + "`" + `weekly` + "`" + `, ` + "`" + `monthly` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "hour_of_day",
					Description: `The hour to begin maintenance based in UTC time. (` + "`" + `0` + "`" + `..` + "`" + `23` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "week_of_month",
					Description: `The week of the month to perform monthly frequency updates. Required for ` + "`" + `monthly` + "`" + ` frequency updates. (` + "`" + `1` + "`" + `..` + "`" + `4` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_database_mysql_backups",
			Category:         "Data Sources",
			ShortDescription: `Provides information about Linode MySQL Database Backups that match a set of filters.`,
			Description: `\_database\_mysql\_backups

~> **NOTICE:** This data source has been deprecated in favor of ` + "`" + `linode_database_backups` + "`" + `.

Provides information about Linode MySQL Database Backups that match a set of filters.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "database_id",
					Description: `(Required) The ID of the database to retrieve backups for.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If true, only the latest backup will be returned.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) The attribute to order the results by. (` + "`" + `created` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order in which results should be returned. (` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `; default ` + "`" + `asc` + "`" + `) ### Filter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each backup will be stored in the ` + "`" + `backups` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `A time value given in a combined date and time format that represents when the database backup was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database backup object.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The database backup’s label, for display purposes only.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of database backup, determined by how the backup was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `A time value given in a combined date and time format that represents when the database backup was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database backup object.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The database backup’s label, for display purposes only.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of database backup, determined by how the backup was created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_database_postgresql",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Linode PostgreSQL Database.`,
			Description: `\_database\_postgresql

Provides information about a Linode PostgreSQL Database.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "database_id",
					Description: `The ID of the PostgreSQL database. ## Attributes Reference The ` + "`" + `linode_database_postgresql` + "`" + ` data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "allow_list",
					Description: `A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `The base64-encoded SSL CA certificate for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "cluster_size",
					Description: `The number of Linode Instance nodes deployed to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Managed Database was created.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the Managed Databases is encrypted.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The Managed Database engine. (e.g. ` + "`" + `postgresql` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `The Managed Database engine in engine/version format. (e.g. ` + "`" + `postgresql/13.2` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "host_primary",
					Description: `The primary host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "host_secondary",
					Description: `The secondary/private network host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A unique, user-defined string referring to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region that hosts this Linode Managed Database.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `The randomly-generated root password for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "root_username",
					Description: `The root username for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `The replication method used for the Managed Database. (` + "`" + `none` + "`" + `, ` + "`" + `asynch` + "`" + `, ` + "`" + `semi_synch` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "replication_commit_type",
					Description: `(Optional) The synchronization level of the replicating server. (` + "`" + `on` + "`" + `, ` + "`" + `local` + "`" + `, ` + "`" + `remote_write` + "`" + `, ` + "`" + `remote_apply` + "`" + `, ` + "`" + `off` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ssl_connection",
					Description: `Whether to require SSL credentials to establish a connection to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The operating status of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Linode Instance type used for the nodes of the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When this Managed Database was last updated.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Managed Database engine version. (e.g. ` + "`" + `v8.0.26` + "`" + `) ## updates The following arguments are exported by the ` + "`" + `updates` + "`" + ` specification block:`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `The day to perform maintenance. (` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ...)`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `The maximum maintenance window time in hours. (` + "`" + `1` + "`" + `..` + "`" + `3` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Whether maintenance occurs on a weekly or monthly basis. (` + "`" + `weekly` + "`" + `, ` + "`" + `monthly` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "hour_of_day",
					Description: `The hour to begin maintenance based in UTC time. (` + "`" + `0` + "`" + `..` + "`" + `23` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "week_of_month",
					Description: `The week of the month to perform monthly frequency updates. Required for ` + "`" + `monthly` + "`" + ` frequency updates. (` + "`" + `1` + "`" + `..` + "`" + `4` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_list",
					Description: `A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `The base64-encoded SSL CA certificate for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "cluster_size",
					Description: `The number of Linode Instance nodes deployed to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Managed Database was created.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the Managed Databases is encrypted.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The Managed Database engine. (e.g. ` + "`" + `postgresql` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `The Managed Database engine in engine/version format. (e.g. ` + "`" + `postgresql/13.2` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "host_primary",
					Description: `The primary host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "host_secondary",
					Description: `The secondary/private network host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A unique, user-defined string referring to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region that hosts this Linode Managed Database.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `The randomly-generated root password for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "root_username",
					Description: `The root username for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `The replication method used for the Managed Database. (` + "`" + `none` + "`" + `, ` + "`" + `asynch` + "`" + `, ` + "`" + `semi_synch` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "replication_commit_type",
					Description: `(Optional) The synchronization level of the replicating server. (` + "`" + `on` + "`" + `, ` + "`" + `local` + "`" + `, ` + "`" + `remote_write` + "`" + `, ` + "`" + `remote_apply` + "`" + `, ` + "`" + `off` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ssl_connection",
					Description: `Whether to require SSL credentials to establish a connection to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The operating status of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Linode Instance type used for the nodes of the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When this Managed Database was last updated.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Managed Database engine version. (e.g. ` + "`" + `v8.0.26` + "`" + `) ## updates The following arguments are exported by the ` + "`" + `updates` + "`" + ` specification block:`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `The day to perform maintenance. (` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ...)`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `The maximum maintenance window time in hours. (` + "`" + `1` + "`" + `..` + "`" + `3` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Whether maintenance occurs on a weekly or monthly basis. (` + "`" + `weekly` + "`" + `, ` + "`" + `monthly` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "hour_of_day",
					Description: `The hour to begin maintenance based in UTC time. (` + "`" + `0` + "`" + `..` + "`" + `23` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "week_of_month",
					Description: `The week of the month to perform monthly frequency updates. Required for ` + "`" + `monthly` + "`" + ` frequency updates. (` + "`" + `1` + "`" + `..` + "`" + `4` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_databases",
			Category:         "Data Sources",
			ShortDescription: `Provides information about Linode Managed Databases.`,
			Description: `\_databases

Provides information about Linode Managed Databases that match a set of filters.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If true, only the latest create database will be returned.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) The attribute to order the results by. (` + "`" + `version` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order in which results should be returned. (` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `; default ` + "`" + `asc` + "`" + `) ### Filter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each engine will be stored in the ` + "`" + `databases` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "allow_list",
					Description: `A list of IP addresses that can access the Managed Database.`,
				},
				resource.Attribute{
					Name:        "cluster_size",
					Description: `The number of Linode Instance nodes deployed to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Managed Database was created.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the Managed Databases is encrypted.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The Managed Database engine.`,
				},
				resource.Attribute{
					Name:        "host_primary",
					Description: `The primary host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "host_secondary",
					Description: `The secondary/private network host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A unique, user-defined string referring to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region to use for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `The replication method used for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "ssl_connection",
					Description: `Whether to require SSL credentials to establish a connection to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The operating status of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Linode Instance type used for the nodes of the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When this Managed Database was last updated.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Managed Database engine version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_list",
					Description: `A list of IP addresses that can access the Managed Database.`,
				},
				resource.Attribute{
					Name:        "cluster_size",
					Description: `The number of Linode Instance nodes deployed to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Managed Database was created.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the Managed Databases is encrypted.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The Managed Database engine.`,
				},
				resource.Attribute{
					Name:        "host_primary",
					Description: `The primary host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "host_secondary",
					Description: `The secondary/private network host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A unique, user-defined string referring to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region to use for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `The replication method used for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "ssl_connection",
					Description: `Whether to require SSL credentials to establish a connection to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The operating status of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Linode Instance type used for the nodes of the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When this Managed Database was last updated.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Managed Database engine version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_domain",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode domain.`,
			Description: `\_domain

Provides information about a Linode domain.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique numeric ID of the Domain record to query.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The unique domain name of the Domain record to query. ## Attributes Reference The Linode Domain resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this Domain.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave) (` + "`" + `master` + "`" + `, ` + "`" + `slave` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The group this Domain belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Used to control whether this Domain is currently being rendered. (` + "`" + `disabled` + "`" + `, ` + "`" + `active` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for this Domain.`,
				},
				resource.Attribute{
					Name:        "master_ips",
					Description: `The IP addresses representing the master DNS for this Domain.`,
				},
				resource.Attribute{
					Name:        "axfr_ips",
					Description: `The list of IPs that may perform a zone transfer for this Domain.`,
				},
				resource.Attribute{
					Name:        "ttl_sec",
					Description: `'Time to Live'-the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers.`,
				},
				resource.Attribute{
					Name:        "retry_sec",
					Description: `The interval, in seconds, at which a failed refresh should be retried.`,
				},
				resource.Attribute{
					Name:        "expire_sec",
					Description: `The amount of time in seconds that may pass before this Domain is no longer authoritative.`,
				},
				resource.Attribute{
					Name:        "refresh_sec",
					Description: `The amount of time in seconds before this Domain should be refreshed.`,
				},
				resource.Attribute{
					Name:        "soa_email",
					Description: `Start of Authority email address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `An array of tags applied to this object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this Domain.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave) (` + "`" + `master` + "`" + `, ` + "`" + `slave` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The group this Domain belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Used to control whether this Domain is currently being rendered. (` + "`" + `disabled` + "`" + `, ` + "`" + `active` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for this Domain.`,
				},
				resource.Attribute{
					Name:        "master_ips",
					Description: `The IP addresses representing the master DNS for this Domain.`,
				},
				resource.Attribute{
					Name:        "axfr_ips",
					Description: `The list of IPs that may perform a zone transfer for this Domain.`,
				},
				resource.Attribute{
					Name:        "ttl_sec",
					Description: `'Time to Live'-the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers.`,
				},
				resource.Attribute{
					Name:        "retry_sec",
					Description: `The interval, in seconds, at which a failed refresh should be retried.`,
				},
				resource.Attribute{
					Name:        "expire_sec",
					Description: `The amount of time in seconds that may pass before this Domain is no longer authoritative.`,
				},
				resource.Attribute{
					Name:        "refresh_sec",
					Description: `The amount of time in seconds before this Domain should be refreshed.`,
				},
				resource.Attribute{
					Name:        "soa_email",
					Description: `Start of Authority email address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `An array of tags applied to this object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_domain_record",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode Domain Record.`,
			Description: `

Provides information about a Linode Domain Record.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_domain_zonefile",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode Domain Zonefile.`,
			Description: `

Provides information about a Linode Domain Zonefile.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_firewall",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Firewall.`,
			Description: `\_firewall

Provides details about a Linode Firewall.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The Firewall's ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label for the firewall.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags applied to the firewall.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `If true, the firewall is inactive.`,
				},
				resource.Attribute{
					Name:        "inbound_policy",
					Description: `The default behavior for inbound traffic. (` + "`" + `ACCEPT` + "`" + `, ` + "`" + `DROP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "outbound_policy",
					Description: `The default behavior for outbound traffic. (` + "`" + `ACCEPT` + "`" + `, ` + "`" + `DROP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `The IDs of Linodes to apply this firewall to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the firewall. (` + "`" + `enabled` + "`" + `, ` + "`" + `disabled` + "`" + `, ` + "`" + `deleted` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Used to identify this rule. For display purposes only.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Controls whether traffic is accepted or dropped by this rule. Overrides the Firewall’s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The network protocol this rule controls. (` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `A list of IPv4 addresses or networks. Must be in IP/mask format.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `A list of IPv6 addresses or networks. Must be in IP/mask format. ### devices The following attributes are available on devices:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Device.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID of the underlying entity this device references (i.e. the Linode's ID).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of Firewall Device.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the underlying entity this device references.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `The label for the firewall.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags applied to the firewall.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `If true, the firewall is inactive.`,
				},
				resource.Attribute{
					Name:        "inbound_policy",
					Description: `The default behavior for inbound traffic. (` + "`" + `ACCEPT` + "`" + `, ` + "`" + `DROP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "outbound_policy",
					Description: `The default behavior for outbound traffic. (` + "`" + `ACCEPT` + "`" + `, ` + "`" + `DROP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `The IDs of Linodes to apply this firewall to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the firewall. (` + "`" + `enabled` + "`" + `, ` + "`" + `disabled` + "`" + `, ` + "`" + `deleted` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Used to identify this rule. For display purposes only.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Controls whether traffic is accepted or dropped by this rule. Overrides the Firewall’s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The network protocol this rule controls. (` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `A list of IPv4 addresses or networks. Must be in IP/mask format.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `A list of IPv6 addresses or networks. Must be in IP/mask format. ### devices The following attributes are available on devices:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Device.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID of the underlying entity this device references (i.e. the Linode's ID).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of Firewall Device.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the underlying entity this device references.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_image",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode image.`,
			Description: `\_image

Provides information about a Linode image

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique ID of this Image. The ID of private images begin with ` + "`" + `private/` + "`" + ` followed by the numeric identifier of the private image, for example ` + "`" + `private/12345` + "`" + `. ## Attributes Reference The Linode Image resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A short description of the Image.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Image was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The name of the User who created this Image, or "linode" for official Images.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Whether or not this Image is deprecated. Will only be true for deprecated public Images.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A detailed description of this Image.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `True if the Image is public.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The minimum size this Image needs to deploy. Size is in MB. example: 2500`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of this image. (` + "`" + `creating` + "`" + `, ` + "`" + `pending_upload` + "`" + `, ` + "`" + `available` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `How the Image was created. Manual Images can be created at any time. "Automatic" Images are created automatically from a deleted Linode. (` + "`" + `manual` + "`" + `, ` + "`" + `automatic` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The upstream distribution vendor. ` + "`" + `None` + "`" + ` for private Images.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `A short description of the Image.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Image was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The name of the User who created this Image, or "linode" for official Images.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Whether or not this Image is deprecated. Will only be true for deprecated public Images.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A detailed description of this Image.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `True if the Image is public.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The minimum size this Image needs to deploy. Size is in MB. example: 2500`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of this image. (` + "`" + `creating` + "`" + `, ` + "`" + `pending_upload` + "`" + `, ` + "`" + `available` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `How the Image was created. Manual Images can be created at any time. "Automatic" Images are created automatically from a deleted Linode. (` + "`" + `manual` + "`" + `, ` + "`" + `automatic` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The upstream distribution vendor. ` + "`" + `None` + "`" + ` for private Images.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_images",
			Category:         "Data Sources",
			ShortDescription: `Provides information about Linode images that match a set of filters.`,
			Description: `\_images

Provides information about Linode images that match a set of filters.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If true, only the latest image will be returned. Images without a valid ` + "`" + `created` + "`" + ` field are not included in the result.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) The attribute to order the results by. See the [Filterable Fields section](#filterable-fields) for a list of valid fields.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order in which results should be returned. (` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `; default ` + "`" + `asc` + "`" + `) ### Filter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by. See the [Filterable Fields section](#filterable-fields) for a complete list of filterable fields.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each Linode image will be stored in the ` + "`" + `images` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this Image. The ID of private images begin with ` + "`" + `private/` + "`" + ` followed by the numeric identifier of the private image, for example ` + "`" + `private/12345` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A short description of the Image.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Image was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The name of the User who created this Image, or "linode" for official Images.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Whether or not this Image is deprecated. Will only be true for deprecated public Images.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A detailed description of this Image.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `True if the Image is public.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The minimum size this Image needs to deploy. Size is in MB. example: 2500`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of this image. (` + "`" + `creating` + "`" + `, ` + "`" + `pending_upload` + "`" + `, ` + "`" + `available` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `How the Image was created. Manual Images can be created at any time. "Automatic" Images are created automatically from a deleted Linode. (` + "`" + `manual` + "`" + `, ` + "`" + `automatic` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The upstream distribution vendor. ` + "`" + `None` + "`" + ` for private Images. ## Filterable Fields`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this Image. The ID of private images begin with ` + "`" + `private/` + "`" + ` followed by the numeric identifier of the private image, for example ` + "`" + `private/12345` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A short description of the Image.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Image was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The name of the User who created this Image, or "linode" for official Images.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Whether or not this Image is deprecated. Will only be true for deprecated public Images.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A detailed description of this Image.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `True if the Image is public.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The minimum size this Image needs to deploy. Size is in MB. example: 2500`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of this image. (` + "`" + `creating` + "`" + `, ` + "`" + `pending_upload` + "`" + `, ` + "`" + `available` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `How the Image was created. Manual Images can be created at any time. "Automatic" Images are created automatically from a deleted Linode. (` + "`" + `manual` + "`" + `, ` + "`" + `automatic` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The upstream distribution vendor. ` + "`" + `None` + "`" + ` for private Images. ## Filterable Fields`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_instance_backups",
			Category:         "Data Sources",
			ShortDescription: `Provides details about the backups of an Instance.`,
			Description: `\_instance_backups

Provides details about the backups of an Instance.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "linode_id",
					Description: `(Required) The Linode instance's ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this Backup.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A label for Backups that are of type snapshot.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current state of a specific Backup. (` + "`" + `paused` + "`" + `, ` + "`" + `pending` + "`" + `, ` + "`" + `running` + "`" + `, ` + "`" + `needsPostProcessing` + "`" + `, ` + "`" + `successful` + "`" + `, ` + "`" + `failed` + "`" + `, ` + "`" + `userAborted` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `This indicates whether the Backup is an automatic Backup or manual snapshot taken by the User at a specific point in time. (` + "`" + `auto` + "`" + `, ` + "`" + `snapshot` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date the Backup was taken.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date the Backup was most recently updated.`,
				},
				resource.Attribute{
					Name:        "finished",
					Description: `The date the Backup completed.`,
				},
				resource.Attribute{
					Name:        "configs",
					Description: `A list of the labels of the Configuration profiles that are part of the Backup.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of this disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of this disk.`,
				},
				resource.Attribute{
					Name:        "filesystem",
					Description: `The filesystem of this disk.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this Backup.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A label for Backups that are of type snapshot.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current state of a specific Backup. (` + "`" + `paused` + "`" + `, ` + "`" + `pending` + "`" + `, ` + "`" + `running` + "`" + `, ` + "`" + `needsPostProcessing` + "`" + `, ` + "`" + `successful` + "`" + `, ` + "`" + `failed` + "`" + `, ` + "`" + `userAborted` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `This indicates whether the Backup is an automatic Backup or manual snapshot taken by the User at a specific point in time. (` + "`" + `auto` + "`" + `, ` + "`" + `snapshot` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date the Backup was taken.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date the Backup was most recently updated.`,
				},
				resource.Attribute{
					Name:        "finished",
					Description: `The date the Backup completed.`,
				},
				resource.Attribute{
					Name:        "configs",
					Description: `A list of the labels of the Configuration profiles that are part of the Backup.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of this disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of this disk.`,
				},
				resource.Attribute{
					Name:        "filesystem",
					Description: `The filesystem of this disk.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_instance_type",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode instance type.`,
			Description: `\_instance\_type

Provides information about a Linode instance type

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Label used to identify instance type ## Attributes Reference The Linode Instance Type resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID representing the Linode Type`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Linode Type's label is for display purposes only`,
				},
				resource.Attribute{
					Name:        "class",
					Description: `The class of the Linode Type. See all classes [here](https://www.linode.com/docs/api/linode-types/#type-view__responses).`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The Disk size, in MB, of the Linode Type`,
				},
				resource.Attribute{
					Name:        "price.0.hourly",
					Description: `Cost (in US dollars) per hour.`,
				},
				resource.Attribute{
					Name:        "price.0.monthly",
					Description: `Cost (in US dollars) per month.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.hourly",
					Description: `The cost (in US dollars) per hour to add Backups service.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.monthly",
					Description: `The cost (in US dollars) per month to add Backups service.`,
				},
				resource.Attribute{
					Name:        "network_out",
					Description: `The Mbits outbound bandwidth allocation.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM included in this Linode Type.`,
				},
				resource.Attribute{
					Name:        "transfer",
					Description: `The monthly outbound transfer amount, in MB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of VCPU cores this Linode Type offers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID representing the Linode Type`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Linode Type's label is for display purposes only`,
				},
				resource.Attribute{
					Name:        "class",
					Description: `The class of the Linode Type. See all classes [here](https://www.linode.com/docs/api/linode-types/#type-view__responses).`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The Disk size, in MB, of the Linode Type`,
				},
				resource.Attribute{
					Name:        "price.0.hourly",
					Description: `Cost (in US dollars) per hour.`,
				},
				resource.Attribute{
					Name:        "price.0.monthly",
					Description: `Cost (in US dollars) per month.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.hourly",
					Description: `The cost (in US dollars) per hour to add Backups service.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.monthly",
					Description: `The cost (in US dollars) per month to add Backups service.`,
				},
				resource.Attribute{
					Name:        "network_out",
					Description: `The Mbits outbound bandwidth allocation.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM included in this Linode Type.`,
				},
				resource.Attribute{
					Name:        "transfer",
					Description: `The monthly outbound transfer amount, in MB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of VCPU cores this Linode Type offers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_instance_types",
			Category:         "Data Sources",
			ShortDescription: `Provides information about Linode Instance types that match a set of filters.`,
			Description: `\_instance_types

Provides information about Linode Instance types that match a set of filters.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) The attribute to order the results by. See the [Filterable Fields section](#filterable-fields) for a list of valid fields.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order in which results should be returned. (` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `; default ` + "`" + `asc` + "`" + `) ### Filter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by. See the [Filterable Fields section](#filterable-fields) for a complete list of filterable fields.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each Linode Instance type will be stored in the ` + "`" + `types` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID representing the Linode Type.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Linode Type's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "class",
					Description: `The class of the Linode Type. See all classes [here](https://www.linode.com/docs/api/linode-types/#type-view__responses).`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The Disk size, in MB, of the Linode Type.`,
				},
				resource.Attribute{
					Name:        "price.0.hourly",
					Description: `Cost (in US dollars) per hour.`,
				},
				resource.Attribute{
					Name:        "price.0.monthly",
					Description: `Cost (in US dollars) per month.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.hourly",
					Description: `The cost (in US dollars) per hour to add Backups service.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.monthly",
					Description: `The cost (in US dollars) per month to add Backups service.`,
				},
				resource.Attribute{
					Name:        "network_out",
					Description: `The Mbits outbound bandwidth allocation.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM included in this Linode Type.`,
				},
				resource.Attribute{
					Name:        "transfer",
					Description: `The monthly outbound transfer amount, in MB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of VCPU cores this Linode Type offers. ## Filterable Fields`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID representing the Linode Type.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Linode Type's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "class",
					Description: `The class of the Linode Type. See all classes [here](https://www.linode.com/docs/api/linode-types/#type-view__responses).`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The Disk size, in MB, of the Linode Type.`,
				},
				resource.Attribute{
					Name:        "price.0.hourly",
					Description: `Cost (in US dollars) per hour.`,
				},
				resource.Attribute{
					Name:        "price.0.monthly",
					Description: `Cost (in US dollars) per month.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.hourly",
					Description: `The cost (in US dollars) per hour to add Backups service.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.monthly",
					Description: `The cost (in US dollars) per month to add Backups service.`,
				},
				resource.Attribute{
					Name:        "network_out",
					Description: `The Mbits outbound bandwidth allocation.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM included in this Linode Type.`,
				},
				resource.Attribute{
					Name:        "transfer",
					Description: `The monthly outbound transfer amount, in MB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of VCPU cores this Linode Type offers. ## Filterable Fields`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides information about Linode instances that match a set of filters.`,
			Description: `\_instances

Provides information about Linode instances that match a set of filters.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) The attribute to order the results by. See the [Filterable Fields section](#filterable-fields) for a list of valid fields.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order in which results should be returned. (` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `; default ` + "`" + `asc` + "`" + `) ### Filter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by. See the [Filterable Fields section](#filterable-fields) for a list of filterable fields.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each Linode instance will be stored in the ` + "`" + `instances` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Linode instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `This is the location where the Linode is deployed. Examples are ` + "`" + `"us-east"` + "`" + `, ` + "`" + `"us-west"` + "`" + `, ` + "`" + `"ap-south"` + "`" + `, etc. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Linode type defines the pricing, CPU, disk, and RAM specs of the instance. Examples are ` + "`" + `"g6-nanode-1"` + "`" + `, ` + "`" + `"g6-standard-2"` + "`" + `, ` + "`" + `"g6-highmem-16"` + "`" + `, ` + "`" + `"g6-dedicated-16"` + "`" + `, etc. See all types [here](https://api.linode.com/v4/linode/types).`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Linode's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The display group of the Linode instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to this object. Tags are for organizational purposes only.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `If true, the Linode has private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region.`,
				},
				resource.Attribute{
					Name:        "alerts.0.cpu",
					Description: `The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we'll send you an alert. If this is set to 0, the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.network_in",
					Description: `The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.network_out",
					Description: `The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.transfer_quota",
					Description: `The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we'll alert you. If this is set to 0 (zero), the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.io",
					Description: `The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we'll send you an alert. If set to 0, this alert is disabled.`,
				},
				resource.Attribute{
					Name:        "watchdog_enabled",
					Description: `The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with ` + "`" + `private/` + "`" + `. See [images](https://api.linode.com/v4/images) for more information on the Images available for you to use. Examples are ` + "`" + `linode/debian9` + "`" + `, ` + "`" + `linode/fedora28` + "`" + `, ` + "`" + `linode/ubuntu16.04lts` + "`" + `, ` + "`" + `linode/arch` + "`" + `, and ` + "`" + `private/12345` + "`" + `. See all images [here](https://api.linode.com/v4/linode/images) (Requires a personal access token; docs [here](https://developers.linode.com/api/v4/images)).`,
				},
				resource.Attribute{
					Name:        "swap_size",
					Description: `When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance, indicating the current readiness state. (` + "`" + `running` + "`" + `, ` + "`" + `offline` + "`" + `, ...)`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `A string containing the Linode's public IP address.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `This Linode's Private IPv4 Address, if enabled. The regional private IP address range, 192.168.128.0/17, is shared by all Linode Instances in a region.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared. The prefix (` + "`" + `/64` + "`" + `) is included in this attribute.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "specs.0.disk",
					Description: `The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an image through POST /linode/instances.`,
				},
				resource.Attribute{
					Name:        "specs.0.memory",
					Description: `The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile.`,
				},
				resource.Attribute{
					Name:        "specs.0.vcpus",
					Description: `The number of vcpus this Linode has access to. Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile.`,
				},
				resource.Attribute{
					Name:        "specs.0.transfer",
					Description: `The amount of network transfer this Linode is allotted each month.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The disks label, which acts as an identifier in Terraform. This must be unique within each Linode Instance.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the Disk in MB.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the disk in the Linode API.`,
				},
				resource.Attribute{
					Name:        "filesystem",
					Description: `The Disk filesystem can be one of: ` + "`" + `"raw"` + "`" + `, ` + "`" + `"swap"` + "`" + `, ` + "`" + `"ext3"` + "`" + `, ` + "`" + `"ext4"` + "`" + `, or ` + "`" + `"initrd"` + "`" + ` which has a max size of 32mb and can be used in the config ` + "`" + `initrd` + "`" + ` (not currently supported in this Terraform Provider). ### Configs Configuration profiles define the VM settings and boot behavior of the Linode Instance. Multiple configurations profiles can be provided but their ` + "`" + `label` + "`" + ` values must be unique.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Config's label for display purposes. Also used by ` + "`" + `boot_config_label` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kernel",
					Description: `A Kernel ID to boot a Linode with. Default is based on image choice. Examples are ` + "`" + `linode/latest-64bit` + "`" + `, ` + "`" + `linode/grub2` + "`" + `, ` + "`" + `linode/direct-disk` + "`" + `, etc. See all kernels [here](https://api.linode.com/v4/linode/kernels). Note that this is a paginated API endpoint ([docs](https://developers.linode.com/api/v4/linode-kernels)).`,
				},
				resource.Attribute{
					Name:        "run_level",
					Description: `Defines the state of your Linode after booting.`,
				},
				resource.Attribute{
					Name:        "virt_mode",
					Description: `Controls the virtualization mode.`,
				},
				resource.Attribute{
					Name:        "root_device",
					Description: `The root device to boot.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Arbitrary user comments about this ` + "`" + `config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory_limit",
					Description: `Defaults to the total RAM of the Linode`,
				},
				resource.Attribute{
					Name:        "helpers",
					Description: `Helpers enabled when booting to this Linode Config.`,
				},
				resource.Attribute{
					Name:        "updatedb_disabled",
					Description: `Disables updatedb cron job to avoid disk thrashing.`,
				},
				resource.Attribute{
					Name:        "distro",
					Description: `Controls the behavior of the Linode Config's Distribution Helper setting.`,
				},
				resource.Attribute{
					Name:        "modules_dep",
					Description: `Creates a modules dependency file for the Kernel you run.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Controls the behavior of the Linode Config's Network Helper setting, used to automatically configure additional IP addresses assigned to this instance.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `A list of ` + "`" + `disk` + "`" + ` or ` + "`" + `volume` + "`" + ` attachments for this ` + "`" + `config` + "`" + `. If the ` + "`" + `boot_config_label` + "`" + ` omits a ` + "`" + `devices` + "`" + ` block, the Linode will not be booted.`,
				},
				resource.Attribute{
					Name:        "disk_label",
					Description: `The ` + "`" + `label` + "`" + ` of the ` + "`" + `disk` + "`" + ` to map to this ` + "`" + `device` + "`" + ` slot.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The Volume ID to map to this ` + "`" + `device` + "`" + ` slot.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `The Disk ID of the associated ` + "`" + `disk_label` + "`" + `, if used`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `(Required) The type of interface. (` + "`" + `public` + "`" + `, ` + "`" + `vlan` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The name of this interface. If the interface is a ` + "`" + `vlan` + "`" + `, a label is required. Must be undefined for ` + "`" + `public` + "`" + ` purpose interfaces.`,
				},
				resource.Attribute{
					Name:        "ipam_address",
					Description: `(Optional) This Network Interface’s private IP address in Classless Inter-Domain Routing (CIDR) notation. ### Backups`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If this Linode has the Backup service enabled.`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `The day of the week that your Linode's weekly Backup is taken. If not set manually, a day will be chosen for you. Backups are taken every day, but backups taken on this day are preferred when selecting backups to retain for a longer period. If not set manually, then when backups are initially enabled, this may come back as "Scheduling" until the day is automatically selected.`,
				},
				resource.Attribute{
					Name:        "window",
					Description: `The window ('W0'-'W22') in which your backups will be taken, in UTC. A backups window is a two-hour span of time in which the backup may occur. For example, 'W10' indicates that your backups should be taken between 10:00 and 12:00. If you do not choose a backup window, one will be selected for you automatically. If not set manually, when backups are initially enabled this may come back as Scheduling until the window is automatically selected. ## Filterable Fields`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Linode instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `This is the location where the Linode is deployed. Examples are ` + "`" + `"us-east"` + "`" + `, ` + "`" + `"us-west"` + "`" + `, ` + "`" + `"ap-south"` + "`" + `, etc. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Linode type defines the pricing, CPU, disk, and RAM specs of the instance. Examples are ` + "`" + `"g6-nanode-1"` + "`" + `, ` + "`" + `"g6-standard-2"` + "`" + `, ` + "`" + `"g6-highmem-16"` + "`" + `, ` + "`" + `"g6-dedicated-16"` + "`" + `, etc. See all types [here](https://api.linode.com/v4/linode/types).`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Linode's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The display group of the Linode instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to this object. Tags are for organizational purposes only.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `If true, the Linode has private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region.`,
				},
				resource.Attribute{
					Name:        "alerts.0.cpu",
					Description: `The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we'll send you an alert. If this is set to 0, the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.network_in",
					Description: `The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.network_out",
					Description: `The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.transfer_quota",
					Description: `The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we'll alert you. If this is set to 0 (zero), the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.io",
					Description: `The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we'll send you an alert. If set to 0, this alert is disabled.`,
				},
				resource.Attribute{
					Name:        "watchdog_enabled",
					Description: `The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with ` + "`" + `private/` + "`" + `. See [images](https://api.linode.com/v4/images) for more information on the Images available for you to use. Examples are ` + "`" + `linode/debian9` + "`" + `, ` + "`" + `linode/fedora28` + "`" + `, ` + "`" + `linode/ubuntu16.04lts` + "`" + `, ` + "`" + `linode/arch` + "`" + `, and ` + "`" + `private/12345` + "`" + `. See all images [here](https://api.linode.com/v4/linode/images) (Requires a personal access token; docs [here](https://developers.linode.com/api/v4/images)).`,
				},
				resource.Attribute{
					Name:        "swap_size",
					Description: `When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance, indicating the current readiness state. (` + "`" + `running` + "`" + `, ` + "`" + `offline` + "`" + `, ...)`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `A string containing the Linode's public IP address.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `This Linode's Private IPv4 Address, if enabled. The regional private IP address range, 192.168.128.0/17, is shared by all Linode Instances in a region.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared. The prefix (` + "`" + `/64` + "`" + `) is included in this attribute.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "specs.0.disk",
					Description: `The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an image through POST /linode/instances.`,
				},
				resource.Attribute{
					Name:        "specs.0.memory",
					Description: `The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile.`,
				},
				resource.Attribute{
					Name:        "specs.0.vcpus",
					Description: `The number of vcpus this Linode has access to. Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile.`,
				},
				resource.Attribute{
					Name:        "specs.0.transfer",
					Description: `The amount of network transfer this Linode is allotted each month.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The disks label, which acts as an identifier in Terraform. This must be unique within each Linode Instance.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the Disk in MB.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the disk in the Linode API.`,
				},
				resource.Attribute{
					Name:        "filesystem",
					Description: `The Disk filesystem can be one of: ` + "`" + `"raw"` + "`" + `, ` + "`" + `"swap"` + "`" + `, ` + "`" + `"ext3"` + "`" + `, ` + "`" + `"ext4"` + "`" + `, or ` + "`" + `"initrd"` + "`" + ` which has a max size of 32mb and can be used in the config ` + "`" + `initrd` + "`" + ` (not currently supported in this Terraform Provider). ### Configs Configuration profiles define the VM settings and boot behavior of the Linode Instance. Multiple configurations profiles can be provided but their ` + "`" + `label` + "`" + ` values must be unique.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Config's label for display purposes. Also used by ` + "`" + `boot_config_label` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kernel",
					Description: `A Kernel ID to boot a Linode with. Default is based on image choice. Examples are ` + "`" + `linode/latest-64bit` + "`" + `, ` + "`" + `linode/grub2` + "`" + `, ` + "`" + `linode/direct-disk` + "`" + `, etc. See all kernels [here](https://api.linode.com/v4/linode/kernels). Note that this is a paginated API endpoint ([docs](https://developers.linode.com/api/v4/linode-kernels)).`,
				},
				resource.Attribute{
					Name:        "run_level",
					Description: `Defines the state of your Linode after booting.`,
				},
				resource.Attribute{
					Name:        "virt_mode",
					Description: `Controls the virtualization mode.`,
				},
				resource.Attribute{
					Name:        "root_device",
					Description: `The root device to boot.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Arbitrary user comments about this ` + "`" + `config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory_limit",
					Description: `Defaults to the total RAM of the Linode`,
				},
				resource.Attribute{
					Name:        "helpers",
					Description: `Helpers enabled when booting to this Linode Config.`,
				},
				resource.Attribute{
					Name:        "updatedb_disabled",
					Description: `Disables updatedb cron job to avoid disk thrashing.`,
				},
				resource.Attribute{
					Name:        "distro",
					Description: `Controls the behavior of the Linode Config's Distribution Helper setting.`,
				},
				resource.Attribute{
					Name:        "modules_dep",
					Description: `Creates a modules dependency file for the Kernel you run.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Controls the behavior of the Linode Config's Network Helper setting, used to automatically configure additional IP addresses assigned to this instance.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `A list of ` + "`" + `disk` + "`" + ` or ` + "`" + `volume` + "`" + ` attachments for this ` + "`" + `config` + "`" + `. If the ` + "`" + `boot_config_label` + "`" + ` omits a ` + "`" + `devices` + "`" + ` block, the Linode will not be booted.`,
				},
				resource.Attribute{
					Name:        "disk_label",
					Description: `The ` + "`" + `label` + "`" + ` of the ` + "`" + `disk` + "`" + ` to map to this ` + "`" + `device` + "`" + ` slot.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The Volume ID to map to this ` + "`" + `device` + "`" + ` slot.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `The Disk ID of the associated ` + "`" + `disk_label` + "`" + `, if used`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `(Required) The type of interface. (` + "`" + `public` + "`" + `, ` + "`" + `vlan` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The name of this interface. If the interface is a ` + "`" + `vlan` + "`" + `, a label is required. Must be undefined for ` + "`" + `public` + "`" + ` purpose interfaces.`,
				},
				resource.Attribute{
					Name:        "ipam_address",
					Description: `(Optional) This Network Interface’s private IP address in Classless Inter-Domain Routing (CIDR) notation. ### Backups`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If this Linode has the Backup service enabled.`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `The day of the week that your Linode's weekly Backup is taken. If not set manually, a day will be chosen for you. Backups are taken every day, but backups taken on this day are preferred when selecting backups to retain for a longer period. If not set manually, then when backups are initially enabled, this may come back as "Scheduling" until the day is automatically selected.`,
				},
				resource.Attribute{
					Name:        "window",
					Description: `The window ('W0'-'W22') in which your backups will be taken, in UTC. A backups window is a two-hour span of time in which the backup may occur. For example, 'W10' indicates that your backups should be taken between 10:00 and 12:00. If you do not choose a backup window, one will be selected for you automatically. If not set manually, when backups are initially enabled this may come back as Scheduling until the window is automatically selected. ## Filterable Fields`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_ipv6_range",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode IPv6 Range.`,
			Description: `\_ipv6\_range

Provides information about a Linode IPv6 Range.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "range",
					Description: `(Required) The IPv6 range to retrieve information about. ## Attributes Reference The ` + "`" + `linode_ipv6_range` + "`" + ` data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip_bgp",
					Description: `Whether this IPv6 range is shared.`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `A set of Linodes targeted by this IPv6 range. Includes Linodes with IP sharing.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The prefix length of the address, denoting how many addresses can be assigned from this range.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region for this range of IPv6 addresses.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_bgp",
					Description: `Whether this IPv6 range is shared.`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `A set of Linodes targeted by this IPv6 range. Includes Linodes with IP sharing.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The prefix length of the address, denoting how many addresses can be assigned from this range.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region for this range of IPv6 addresses.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_lke_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides details about an LKE Cluster.`,
			Description: `\_lke_cluster

Provides details about an LKE Cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The LKE Cluster's ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `The Kubernetes version for this Kubernetes cluster in the format of ` + "`" + `major.minor` + "`" + ` (e.g. ` + "`" + `1.17` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `This Kubernetes cluster's location.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags applied to the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the cluster.`,
				},
				resource.Attribute{
					Name:        "api_endpoints",
					Description: `The endpoints for the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The base64 encoded kubeconfig for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "dashboard_url",
					Description: `The Kubernetes Dashboard access URL for this cluster.`,
				},
				resource.Attribute{
					Name:        "pools",
					Description: `Node pools associated with this cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Node Pool.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The linode type for all of the nodes in the Node Pool. See all node types [here](https://api.linode.com/v4/linode/types).`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `The nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the node.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the underlying Linode instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the node. (` + "`" + `ready` + "`" + `, ` + "`" + `not_ready` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "autoscaler",
					Description: `The configuration options for the autoscaler. This field only contains an autoscaler configuration if autoscaling is enabled on this cluster.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `The minimum number of nodes to autoscale to.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `The maximum number of nodes to autoscale to.`,
				},
				resource.Attribute{
					Name:        "control_plane.0.high_availability",
					Description: `Whether High Availability is enabled for the cluster Control Plane.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "k8s_version",
					Description: `The Kubernetes version for this Kubernetes cluster in the format of ` + "`" + `major.minor` + "`" + ` (e.g. ` + "`" + `1.17` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `This Kubernetes cluster's location.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags applied to the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the cluster.`,
				},
				resource.Attribute{
					Name:        "api_endpoints",
					Description: `The endpoints for the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The base64 encoded kubeconfig for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "dashboard_url",
					Description: `The Kubernetes Dashboard access URL for this cluster.`,
				},
				resource.Attribute{
					Name:        "pools",
					Description: `Node pools associated with this cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Node Pool.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The linode type for all of the nodes in the Node Pool. See all node types [here](https://api.linode.com/v4/linode/types).`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `The nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the node.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the underlying Linode instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the node. (` + "`" + `ready` + "`" + `, ` + "`" + `not_ready` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "autoscaler",
					Description: `The configuration options for the autoscaler. This field only contains an autoscaler configuration if autoscaling is enabled on this cluster.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `The minimum number of nodes to autoscale to.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `The maximum number of nodes to autoscale to.`,
				},
				resource.Attribute{
					Name:        "control_plane.0.high_availability",
					Description: `Whether High Availability is enabled for the cluster Control Plane.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_networking_ip",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode Networking IP Address.`,
			Description: `\_network\_ip

Provides information about a Linode Networking IP Address

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The IP Address to access. The address must be associated with the account and a resource that the user has access to view. ## Attributes Reference The Linode Network IP Address resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The default gateway for this address.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The mask that separates host bits from network bits for this address.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The number of bits set in the subnet mask.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of address this is (ipv4, ipv6, ipv6/pool, ipv6/range).`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether this is a public or private IP address.`,
				},
				resource.Attribute{
					Name:        "rdns",
					Description: `The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.`,
				},
				resource.Attribute{
					Name:        "linode_id",
					Description: `The ID of the Linode this address currently belongs to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Region this IP address resides in. See all regions [here](https://api.linode.com/v4/regions).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `The IP address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The default gateway for this address.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The mask that separates host bits from network bits for this address.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The number of bits set in the subnet mask.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of address this is (ipv4, ipv6, ipv6/pool, ipv6/range).`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether this is a public or private IP address.`,
				},
				resource.Attribute{
					Name:        "rdns",
					Description: `The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.`,
				},
				resource.Attribute{
					Name:        "linode_id",
					Description: `The ID of the Linode this address currently belongs to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Region this IP address resides in. See all regions [here](https://api.linode.com/v4/regions).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_nodebalancer",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a NodeBalancer.`,
			Description: `\_nodebalancer

Provides details about a Linode NodeBalancer.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The NodeBalancer's ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the Linode NodeBalancer`,
				},
				resource.Attribute{
					Name:        "client_conn_throttle",
					Description: `Throttle connections per second (0-20).`,
				},
				resource.Attribute{
					Name:        "linode_id",
					Description: `The ID of a Linode Instance where the NodeBalancer should be attached.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to this object. Tags are for organizational purposes only.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `This NodeBalancer's hostname, ending with .nodebalancer.linode.com`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The Public IPv4 Address of this NodeBalancer`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The Public IPv6 Address of this NodeBalancer`,
				},
				resource.Attribute{
					Name:        "in",
					Description: `The total transfer, in MB, used by this NodeBalancer for the current month`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `The total inbound transfer, in MB, used for this NodeBalancer for the current month`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `The total outbound transfer, in MB, used for this NodeBalancer for the current month`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `The label of the Linode NodeBalancer`,
				},
				resource.Attribute{
					Name:        "client_conn_throttle",
					Description: `Throttle connections per second (0-20).`,
				},
				resource.Attribute{
					Name:        "linode_id",
					Description: `The ID of a Linode Instance where the NodeBalancer should be attached.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to this object. Tags are for organizational purposes only.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `This NodeBalancer's hostname, ending with .nodebalancer.linode.com`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The Public IPv4 Address of this NodeBalancer`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The Public IPv6 Address of this NodeBalancer`,
				},
				resource.Attribute{
					Name:        "in",
					Description: `The total transfer, in MB, used by this NodeBalancer for the current month`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `The total inbound transfer, in MB, used for this NodeBalancer for the current month`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `The total outbound transfer, in MB, used for this NodeBalancer for the current month`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_nodebalancer_config",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a NodeBalancer config.`,
			Description: `\_nodebalancer_config

Provides details about a Linode NodeBalancer Config.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The config's ID.`,
				},
				resource.Attribute{
					Name:        "nodebalancer_id",
					Description: `(Required) The ID of the NodeBalancer that contains the config. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where this nodebalancer_config will be deployed. Examples are ` + "`" + `"us-east"` + "`" + `, ` + "`" + `"us-west"` + "`" + `, ` + "`" + `"ap-south"` + "`" + `, etc. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key. (` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + `) (Defaults to ` + "`" + `http` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `The version of ProxyProtocol to use for the underlying NodeBalancer. This requires protocol to be ` + "`" + `tcp` + "`" + `. (` + "`" + `none` + "`" + `, ` + "`" + `v1` + "`" + `, and ` + "`" + `v2` + "`" + `) (Defaults to ` + "`" + `none` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The TCP port this Config is for.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `What algorithm this NodeBalancer should use for routing traffic to backends (` + "`" + `roundrobin` + "`" + `, ` + "`" + `leastconn` + "`" + `, ` + "`" + `source` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "stickiness",
					Description: `Controls how session stickiness is handled on this port. (` + "`" + `none` + "`" + `, ` + "`" + `table` + "`" + `, ` + "`" + `http_cookie` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "check",
					Description: `The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected. (` + "`" + `none` + "`" + `, ` + "`" + `connection` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `http_body` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `How often, in seconds, to check that backends are up and serving requests.`,
				},
				resource.Attribute{
					Name:        "check_timeout",
					Description: `How long, in seconds, to wait for a check attempt before considering it failed. (1-30)`,
				},
				resource.Attribute{
					Name:        "check_attempts",
					Description: `How many times to attempt a check before considering a backend to be down. (1-30)`,
				},
				resource.Attribute{
					Name:        "check_path",
					Description: `The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.`,
				},
				resource.Attribute{
					Name:        "check_passive",
					Description: `If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.`,
				},
				resource.Attribute{
					Name:        "cipher_suite",
					Description: `What ciphers to use for SSL connections served by this NodeBalancer. ` + "`" + `legacy` + "`" + ` is considered insecure and should only be used if necessary. (` + "`" + `recommended` + "`" + `, ` + "`" + `legacy` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ssl_commonname",
					Description: `The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.`,
				},
				resource.Attribute{
					Name:        "ssl_fingerprint",
					Description: `The read-only fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.`,
				},
				resource.Attribute{
					Name:        "up",
					Description: `The number of backends considered to be 'UP' and healthy, and that are serving requests.`,
				},
				resource.Attribute{
					Name:        "down",
					Description: `The number of backends considered to be 'DOWN' and unhealthy. These are not in rotation, and not serving requests.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `The region where this nodebalancer_config will be deployed. Examples are ` + "`" + `"us-east"` + "`" + `, ` + "`" + `"us-west"` + "`" + `, ` + "`" + `"ap-south"` + "`" + `, etc. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key. (` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + `) (Defaults to ` + "`" + `http` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `The version of ProxyProtocol to use for the underlying NodeBalancer. This requires protocol to be ` + "`" + `tcp` + "`" + `. (` + "`" + `none` + "`" + `, ` + "`" + `v1` + "`" + `, and ` + "`" + `v2` + "`" + `) (Defaults to ` + "`" + `none` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The TCP port this Config is for.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `What algorithm this NodeBalancer should use for routing traffic to backends (` + "`" + `roundrobin` + "`" + `, ` + "`" + `leastconn` + "`" + `, ` + "`" + `source` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "stickiness",
					Description: `Controls how session stickiness is handled on this port. (` + "`" + `none` + "`" + `, ` + "`" + `table` + "`" + `, ` + "`" + `http_cookie` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "check",
					Description: `The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected. (` + "`" + `none` + "`" + `, ` + "`" + `connection` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `http_body` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `How often, in seconds, to check that backends are up and serving requests.`,
				},
				resource.Attribute{
					Name:        "check_timeout",
					Description: `How long, in seconds, to wait for a check attempt before considering it failed. (1-30)`,
				},
				resource.Attribute{
					Name:        "check_attempts",
					Description: `How many times to attempt a check before considering a backend to be down. (1-30)`,
				},
				resource.Attribute{
					Name:        "check_path",
					Description: `The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.`,
				},
				resource.Attribute{
					Name:        "check_passive",
					Description: `If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.`,
				},
				resource.Attribute{
					Name:        "cipher_suite",
					Description: `What ciphers to use for SSL connections served by this NodeBalancer. ` + "`" + `legacy` + "`" + ` is considered insecure and should only be used if necessary. (` + "`" + `recommended` + "`" + `, ` + "`" + `legacy` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ssl_commonname",
					Description: `The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.`,
				},
				resource.Attribute{
					Name:        "ssl_fingerprint",
					Description: `The read-only fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.`,
				},
				resource.Attribute{
					Name:        "up",
					Description: `The number of backends considered to be 'UP' and healthy, and that are serving requests.`,
				},
				resource.Attribute{
					Name:        "down",
					Description: `The number of backends considered to be 'DOWN' and unhealthy. These are not in rotation, and not serving requests.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_nodebalancer_node",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a NodeBalancer node.`,
			Description: `\_nodebalancer_node

Provides details about a Linode NodeBalancer node.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The node's ID.`,
				},
				resource.Attribute{
					Name:        "nodebalancer_id",
					Description: `(Required) The ID of the NodeBalancer that contains the node.`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the config that contains the Node. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the Linode NodeBalancer Node. This is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The private IP Address where this backend can be reached.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The mode this NodeBalancer should use when sending traffic to this backend. If set to ` + "`" + `accept` + "`" + ` this backend is accepting traffic. If set to ` + "`" + `reject` + "`" + ` this backend will not receive traffic. If set to ` + "`" + `drain` + "`" + ` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. (` + "`" + `accept` + "`" + `, ` + "`" + `reject` + "`" + `, ` + "`" + `drain` + "`" + `, ` + "`" + `backup` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of this node, based on the configured checks of its NodeBalancer Config. (` + "`" + `unknown` + "`" + `, ` + "`" + `UP` + "`" + `, ` + "`" + `DOWN` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `The label of the Linode NodeBalancer Node. This is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The private IP Address where this backend can be reached.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The mode this NodeBalancer should use when sending traffic to this backend. If set to ` + "`" + `accept` + "`" + ` this backend is accepting traffic. If set to ` + "`" + `reject` + "`" + ` this backend will not receive traffic. If set to ` + "`" + `drain` + "`" + ` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. (` + "`" + `accept` + "`" + `, ` + "`" + `reject` + "`" + `, ` + "`" + `drain` + "`" + `, ` + "`" + `backup` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of this node, based on the configured checks of its NodeBalancer Config. (` + "`" + `unknown` + "`" + `, ` + "`" + `UP` + "`" + `, ` + "`" + `DOWN` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_object_storage_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode Object Storage Cluster.`,
			Description: `\_object\_storage\_cluster

Provides information about a Linode Object Storage Cluster

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique ID of this cluster. ## Attributes Reference The Linode Object Storage Cluster resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The base URL for this cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `This cluster's status. (` + "`" + `available` + "`" + `, ` + "`" + `unavailable` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this cluster is located in. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "static_site_domain",
					Description: `The base URL for this cluster used when hosting static sites.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `The base URL for this cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `This cluster's status. (` + "`" + `available` + "`" + `, ` + "`" + `unavailable` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this cluster is located in. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "static_site_domain",
					Description: `The base URL for this cluster used when hosting static sites.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_profile",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode profile.`,
			Description: `\_profile

Provides information about a Linode profile.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The profile email address. This address will be used for communication with Linode as necessary.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The profile's preferred timezone. This is not used by the API, and is for the benefit of clients only. All times the API returns are in UTC.`,
				},
				resource.Attribute{
					Name:        "email_notifications",
					Description: `If true, email notifications will be sent about account activity. If false, when false business-critical communications may still be sent through email.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username for logging in to Linode services.`,
				},
				resource.Attribute{
					Name:        "ip_whitelist_enabled",
					Description: `If true, logins for the user will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled.`,
				},
				resource.Attribute{
					Name:        "lish_auth_method",
					Description: `The methods of authentication allowed when connecting via Lish. 'keys_only' is the most secure with the intent to use Lish, and 'disabled' is recommended for users that will not use Lish at all.`,
				},
				resource.Attribute{
					Name:        "authorized_keys",
					Description: `The list of SSH Keys authorized to use Lish for this user. This value is ignored if lish_auth_method is 'disabled'.`,
				},
				resource.Attribute{
					Name:        "two_factor_auth",
					Description: `If true, logins from untrusted computers will require Two Factor Authentication.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `If true, the user has restrictions on what can be accessed on the Account.`,
				},
				resource.Attribute{
					Name:        "referrals",
					Description: `Credit Card information associated with this Account.`,
				},
				resource.Attribute{
					Name:        "referrals.0.total",
					Description: `The number of users who have signed up with the referral code.`,
				},
				resource.Attribute{
					Name:        "referrals.0.credit",
					Description: `The amount of account credit in US Dollars issued to the account through the referral program.`,
				},
				resource.Attribute{
					Name:        "referrals.0.completed",
					Description: `The number of completed signups with the referral code.`,
				},
				resource.Attribute{
					Name:        "referrals.0.pending",
					Description: `The number of pending signups for the referral code. To receive credit the signups must be completed.`,
				},
				resource.Attribute{
					Name:        "referrals.0.code",
					Description: `The Profile referral code. If new accounts use this when signing up for Linode, referring account will receive credit.`,
				},
				resource.Attribute{
					Name:        "referrals.0.url",
					Description: `The referral URL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The profile email address. This address will be used for communication with Linode as necessary.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The profile's preferred timezone. This is not used by the API, and is for the benefit of clients only. All times the API returns are in UTC.`,
				},
				resource.Attribute{
					Name:        "email_notifications",
					Description: `If true, email notifications will be sent about account activity. If false, when false business-critical communications may still be sent through email.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username for logging in to Linode services.`,
				},
				resource.Attribute{
					Name:        "ip_whitelist_enabled",
					Description: `If true, logins for the user will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled.`,
				},
				resource.Attribute{
					Name:        "lish_auth_method",
					Description: `The methods of authentication allowed when connecting via Lish. 'keys_only' is the most secure with the intent to use Lish, and 'disabled' is recommended for users that will not use Lish at all.`,
				},
				resource.Attribute{
					Name:        "authorized_keys",
					Description: `The list of SSH Keys authorized to use Lish for this user. This value is ignored if lish_auth_method is 'disabled'.`,
				},
				resource.Attribute{
					Name:        "two_factor_auth",
					Description: `If true, logins from untrusted computers will require Two Factor Authentication.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `If true, the user has restrictions on what can be accessed on the Account.`,
				},
				resource.Attribute{
					Name:        "referrals",
					Description: `Credit Card information associated with this Account.`,
				},
				resource.Attribute{
					Name:        "referrals.0.total",
					Description: `The number of users who have signed up with the referral code.`,
				},
				resource.Attribute{
					Name:        "referrals.0.credit",
					Description: `The amount of account credit in US Dollars issued to the account through the referral program.`,
				},
				resource.Attribute{
					Name:        "referrals.0.completed",
					Description: `The number of completed signups with the referral code.`,
				},
				resource.Attribute{
					Name:        "referrals.0.pending",
					Description: `The number of pending signups for the referral code. To receive credit the signups must be completed.`,
				},
				resource.Attribute{
					Name:        "referrals.0.code",
					Description: `The Profile referral code. If new accounts use this when signing up for Linode, referring account will receive credit.`,
				},
				resource.Attribute{
					Name:        "referrals.0.url",
					Description: `The referral URL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_region",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific service region`,
			Description: `\_region

` + "`" + `linode_region` + "`" + ` provides details about a specific Linode region. See all regions [here](https://api.linode.com/v4/regions).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The code name of the region to select. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country the region resides in.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Detailed location information for this Region, including city, state or region, and country.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `A list of capabilities of this region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `This region’s current operational status (ok or outage).`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The IPv4 addresses for this region’s DNS resolvers, separated by commas.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The IPv6 addresses for this region’s DNS resolvers, separated by commas.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "country",
					Description: `The country the region resides in.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Detailed location information for this Region, including city, state or region, and country.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `A list of capabilities of this region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `This region’s current operational status (ok or outage).`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The IPv4 addresses for this region’s DNS resolvers, separated by commas.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The IPv6 addresses for this region’s DNS resolvers, separated by commas.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_sshkey",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a profile SSH Key`,
			Description: `\_sshkey

` + "`" + `linode_sshkey` + "`" + ` provides access to a specifically labeled SSH Key in the Profile of the User identified by the access token.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_stackscript",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode StackScript.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique numeric ID of the StackScript to query. ## Attributes Reference This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The StackScript's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The script to execute when provisioning a new Linode with this StackScript.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the StackScript.`,
				},
				resource.Attribute{
					Name:        "rev_note",
					Description: `This field allows you to add notes for the set of revisions made to this StackScript.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A set of Image IDs representing the Images that this StackScript is compatible for deploying with. ` + "`" + `any/all` + "`" + ` indicates that all available image distributions, including private images, are accepted.`,
				},
				resource.Attribute{
					Name:        "deployments_active",
					Description: `Count of currently active, deployed Linodes created from this StackScript.`,
				},
				resource.Attribute{
					Name:        "user_gravatar_id",
					Description: `The Gravatar ID for the User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "deployments_total",
					Description: `The total number of times this StackScript has been deployed.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this StackScript was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date this StackScript was updated.`,
				},
				resource.Attribute{
					Name:        "user_defined_fields",
					Description: `This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A human-readable label for the field that will serve as the input prompt for entering the value during deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the field.`,
				},
				resource.Attribute{
					Name:        "example",
					Description: `An example value for the field.`,
				},
				resource.Attribute{
					Name:        "one_of",
					Description: `A list of acceptable single values for the field.`,
				},
				resource.Attribute{
					Name:        "many_of",
					Description: `A list of acceptable values for the field in any quantity, combination or order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `The default value. If not specified, this value will be used. ## Import Linodes StackScripts can be imported using the Linode StackScript ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_stackscript.mystackscript 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `The StackScript's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The script to execute when provisioning a new Linode with this StackScript.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the StackScript.`,
				},
				resource.Attribute{
					Name:        "rev_note",
					Description: `This field allows you to add notes for the set of revisions made to this StackScript.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A set of Image IDs representing the Images that this StackScript is compatible for deploying with. ` + "`" + `any/all` + "`" + ` indicates that all available image distributions, including private images, are accepted.`,
				},
				resource.Attribute{
					Name:        "deployments_active",
					Description: `Count of currently active, deployed Linodes created from this StackScript.`,
				},
				resource.Attribute{
					Name:        "user_gravatar_id",
					Description: `The Gravatar ID for the User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "deployments_total",
					Description: `The total number of times this StackScript has been deployed.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this StackScript was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date this StackScript was updated.`,
				},
				resource.Attribute{
					Name:        "user_defined_fields",
					Description: `This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A human-readable label for the field that will serve as the input prompt for entering the value during deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the field.`,
				},
				resource.Attribute{
					Name:        "example",
					Description: `An example value for the field.`,
				},
				resource.Attribute{
					Name:        "one_of",
					Description: `A list of acceptable single values for the field.`,
				},
				resource.Attribute{
					Name:        "many_of",
					Description: `A list of acceptable values for the field in any quantity, combination or order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `The default value. If not specified, this value will be used. ## Import Linodes StackScripts can be imported using the Linode StackScript ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_stackscript.mystackscript 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_stackscripts",
			Category:         "Data Sources",
			ShortDescription: `Provides information about Linode StackScripts that match a set of filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If true, only the latest StackScript will be returned. StackScripts without a valid ` + "`" + `created` + "`" + ` field are not included in the result.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) The attribute to order the results by. See the [Filterable Fields section](#filterable-fields) for a list of valid fields.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order in which results should be returned. (` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `; default ` + "`" + `asc` + "`" + `) ### Filter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by. See the [Filterable Fields section](#filterable-fields) for a complete list of filterable fields.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each Linode StackScript will be stored in the ` + "`" + `stackscripts` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the StackScript.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The StackScript's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The script to execute when provisioning a new Linode with this StackScript.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the StackScript.`,
				},
				resource.Attribute{
					Name:        "rev_note",
					Description: `This field allows you to add notes for the set of revisions made to this StackScript.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `An array of Image IDs representing the Images that this StackScript is compatible for deploying with.`,
				},
				resource.Attribute{
					Name:        "deployments_active",
					Description: `Count of currently active, deployed Linodes created from this StackScript.`,
				},
				resource.Attribute{
					Name:        "user_gravatar_id",
					Description: `The Gravatar ID for the User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "deployments_total",
					Description: `The total number of times this StackScript has been deployed.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this StackScript was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date this StackScript was updated.`,
				},
				resource.Attribute{
					Name:        "user_defined_fields",
					Description: `This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A human-readable label for the field that will serve as the input prompt for entering the value during deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the field.`,
				},
				resource.Attribute{
					Name:        "example",
					Description: `An example value for the field.`,
				},
				resource.Attribute{
					Name:        "one_of",
					Description: `A list of acceptable single values for the field.`,
				},
				resource.Attribute{
					Name:        "many_of",
					Description: `A list of acceptable values for the field in any quantity, combination or order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `The default value. If not specified, this value will be used. ## Filterable Fields`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the StackScript.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The StackScript's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The script to execute when provisioning a new Linode with this StackScript.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the StackScript.`,
				},
				resource.Attribute{
					Name:        "rev_note",
					Description: `This field allows you to add notes for the set of revisions made to this StackScript.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `An array of Image IDs representing the Images that this StackScript is compatible for deploying with.`,
				},
				resource.Attribute{
					Name:        "deployments_active",
					Description: `Count of currently active, deployed Linodes created from this StackScript.`,
				},
				resource.Attribute{
					Name:        "user_gravatar_id",
					Description: `The Gravatar ID for the User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "deployments_total",
					Description: `The total number of times this StackScript has been deployed.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this StackScript was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date this StackScript was updated.`,
				},
				resource.Attribute{
					Name:        "user_defined_fields",
					Description: `This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A human-readable label for the field that will serve as the input prompt for entering the value during deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the field.`,
				},
				resource.Attribute{
					Name:        "example",
					Description: `An example value for the field.`,
				},
				resource.Attribute{
					Name:        "one_of",
					Description: `A list of acceptable single values for the field.`,
				},
				resource.Attribute{
					Name:        "many_of",
					Description: `A list of acceptable values for the field in any quantity, combination or order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `The default value. If not specified, this value will be used. ## Filterable Fields`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_user",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode user.`,
			Description: `\_user

Provides information about a Linode user

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The unique username of this User. ## Attributes Reference The Linode User resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the authorized_users field of a create Linode, rebuild Linode, or create Disk request.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address for this User, for account management communications, and may be used for other communications as configured.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `If true, this User must be granted access to perform actions or access entities on this Account.`,
				},
				resource.Attribute{
					Name:        "global_grants",
					Description: `The Account-level grants a User has.`,
				},
				resource.Attribute{
					Name:        "domain_grant",
					Description: `The grants this User has pertaining to Domains on this Account.`,
				},
				resource.Attribute{
					Name:        "firewall_grant",
					Description: `The grants this User has pertaining to Firewalls on this Account.`,
				},
				resource.Attribute{
					Name:        "image_grant",
					Description: `The grants this User has pertaining to Images on this Account.`,
				},
				resource.Attribute{
					Name:        "linode_grant",
					Description: `The grants this User has pertaining to Linodes on this Account.`,
				},
				resource.Attribute{
					Name:        "longview_grant",
					Description: `The grants this User has pertaining to Longview Clients on this Account.`,
				},
				resource.Attribute{
					Name:        "nodebalancer_grant",
					Description: `The grants this User has pertaining to NodeBalancers on this Account.`,
				},
				resource.Attribute{
					Name:        "stackscript_grant",
					Description: `The grants this User has pertaining to StackScripts on this Account.`,
				},
				resource.Attribute{
					Name:        "volume_grant",
					Description: `The grants this User has pertaining to Volumes on this Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the authorized_users field of a create Linode, rebuild Linode, or create Disk request.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address for this User, for account management communications, and may be used for other communications as configured.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `If true, this User must be granted access to perform actions or access entities on this Account.`,
				},
				resource.Attribute{
					Name:        "global_grants",
					Description: `The Account-level grants a User has.`,
				},
				resource.Attribute{
					Name:        "domain_grant",
					Description: `The grants this User has pertaining to Domains on this Account.`,
				},
				resource.Attribute{
					Name:        "firewall_grant",
					Description: `The grants this User has pertaining to Firewalls on this Account.`,
				},
				resource.Attribute{
					Name:        "image_grant",
					Description: `The grants this User has pertaining to Images on this Account.`,
				},
				resource.Attribute{
					Name:        "linode_grant",
					Description: `The grants this User has pertaining to Linodes on this Account.`,
				},
				resource.Attribute{
					Name:        "longview_grant",
					Description: `The grants this User has pertaining to Longview Clients on this Account.`,
				},
				resource.Attribute{
					Name:        "nodebalancer_grant",
					Description: `The grants this User has pertaining to NodeBalancers on this Account.`,
				},
				resource.Attribute{
					Name:        "stackscript_grant",
					Description: `The grants this User has pertaining to StackScripts on this Account.`,
				},
				resource.Attribute{
					Name:        "volume_grant",
					Description: `The grants this User has pertaining to Volumes on this Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_vlans",
			Category:         "Data Sources",
			ShortDescription: `Provides details about Linode VLANs.`,
			Description: `\_vlans

Provides details about Linode VLANs.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) The attribute to order the results by. See the [Filterable Fields section](#filterable-fields) for a list of valid fields.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order in which results should be returned. (` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `; default ` + "`" + `asc` + "`" + `) ### Filter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by. See the [Filterable Fields section](#filterable-fields) for a complete list of filterable fields.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values for the filter to allow. These values should all be in string form.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) The method to match the field by. (` + "`" + `exact` + "`" + `, ` + "`" + `regex` + "`" + `, ` + "`" + `substring` + "`" + `; default ` + "`" + `exact` + "`" + `) ## Attributes Reference Each Linode VLAN will be stored in the ` + "`" + `vlans` + "`" + ` attribute and will export the following attributes:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The unique label of the VLAN.`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `The running Linodes currently attached to the VLAN.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the VLAN is located in. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When the VLAN was created. ## Filterable Fields`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `The unique label of the VLAN.`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `The running Linodes currently attached to the VLAN.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the VLAN is located in. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When the VLAN was created. ## Filterable Fields`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_volume",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode Volume.`,
			Description: `

Provides information about a Linode Volume.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"linode_account":                0,
		"linode_account_login":          1,
		"linode_account_logins":         2,
		"linode_database_backups":       3,
		"linode_database_engines":       4,
		"linode_database_mysql":         5,
		"linode_database_mysql_backups": 6,
		"linode_database_postgresql":    7,
		"linode_databases":              8,
		"linode_domain":                 9,
		"linode_domain_record":          10,
		"linode_domain_zonefile":        11,
		"linode_firewall":               12,
		"linode_image":                  13,
		"linode_images":                 14,
		"linode_instance_backups":       15,
		"linode_instance_type":          16,
		"linode_instance_types":         17,
		"linode_instances":              18,
		"linode_ipv6_range":             19,
		"linode_lke_cluster":            20,
		"linode_networking_ip":          21,
		"linode_nodebalancer":           22,
		"linode_nodebalancer_config":    23,
		"linode_nodebalancer_node":      24,
		"linode_object_storage_cluster": 25,
		"linode_profile":                26,
		"linode_region":                 27,
		"linode_sshkey":                 28,
		"linode_stackscript":            29,
		"linode_stackscripts":           30,
		"linode_user":                   31,
		"linode_vlans":                  32,
		"linode_volume":                 33,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
