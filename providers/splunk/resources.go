package splunk

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "splunk_acl",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
The ACL resource is a dependent resource. It is optional. The ACL context applies to other Splunk resources.
Typically, knowledge objects, such as saved searches or event types, have an app/user context that is the namespace.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app",
					Description: `(Optional) The app context for the resource. Required for updating saved search ACL properties. Allowed values are: The name of an app and system.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.`,
				},
				resource.Attribute{
					Name:        "sharing",
					Description: `(Optional) Indicates how the resource is shared. Required for updating any knowledge object ACL properties. <br>app: Shared within a specific app<br>global: (Default) Shared globally to all apps<br>user: Private to a user`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Optional) Properties that indicate resource read permissions`,
				},
				resource.Attribute{
					Name:        "write",
					Description: `(Optional) Properties that indicate write permissions of the resource`,
				},
				resource.Attribute{
					Name:        "can_change_perms",
					Description: `(Optional) Indicates if the active user can change permissions for this object. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "can_share_app",
					Description: `(Optional) Indicates if the active user can change sharing to app level. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "can_share_global",
					Description: `(Optional) Indicates if the active user can change sharing to system level. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "can_share_user",
					Description: `(Optional) Indicates if the active user can change sharing to user level. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "can_write",
					Description: `(Optional) Indicates if the active user can edit this object. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "removable",
					Description: `(Optional) Indicates whether an admin or user with sufficient permissions can delete the entity.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_admin_saml_groups",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Manage external groups in an IdP response to internal Splunk roles.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the external group.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) List of internal roles assigned to the group. ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID (group_name) of the resource ## Import SAML groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import splunk_admin_saml_groups.saml-group mygroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID (group_name) of the resource ## Import SAML groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import splunk_admin_saml_groups.saml-group mygroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_apps_local",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create, install and manage apps on your Splunk instance

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Literal app name or path for the file to install, depending on the value of filename. <br>filename = false indicates that name is the literal app name and that the app is created from Splunkbase using a template. <br>filename = true indicates that name is the URL or path to the local .tar, .tgz or .spl file. If name is the Splunkbase URL, set auth or session to authenticate the request. The app folder name cannot include spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "auth",
					Description: `(Optional) Splunkbase session token for operations like install and update that require login. Use auth or session when installing or updating an app through Splunkbase.`,
				},
				resource.Attribute{
					Name:        "author",
					Description: `(Optional) For apps posted to Splunkbase, use your Splunk account username. For internal apps, include your name and contact information.`,
				},
				resource.Attribute{
					Name:        "configured",
					Description: `(Optional) Custom setup complete indication: <br>true = Custom app setup complete. <br>false = Custom app setup not complete.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Short app description also displayed below the app title in Splunk Web Launcher.`,
				},
				resource.Attribute{
					Name:        "explicit_appname",
					Description: `(Optional) Custom app name. Overrides name when installing an app from a file where filename is set to true. See also filename.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional) Indicates whether to use the name value as the app source location. <br>true indicates that name is a path to a file to install. <br>false indicates that name is the literal app name and that the app is created from Splunkbase using a template.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) App name displayed in Splunk Web, from five to eighty characters excluding the prefix "Splunk for".`,
				},
				resource.Attribute{
					Name:        "session",
					Description: `(Optional) Login session token for installing or updating an app on Splunkbase. Alternatively, use auth.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Optional) File-based update indication: <br>true specifies that filename should be used to update an existing app. If not specified, update defaults to <br>false, which indicates that filename should not be used to update an existing app.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) App version.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `(Optional) Indicates whether the app is visible and navigable from Splunk Web. <br>true = App is visible and navigable. <br>false = App is not visible or navigable.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_authentication_users",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create and update user information or delete the user.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique user login name.`,
				},
				resource.Attribute{
					Name:        "default_app",
					Description: `(Optional) User default app. Overrides the default app inherited from the user roles.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) User email address.`,
				},
				resource.Attribute{
					Name:        "force_change_pass",
					Description: `Force user to change password indication`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) User login password.`,
				},
				resource.Attribute{
					Name:        "restart_background_jobs",
					Description: `(Optional) Restart background search job that has not completed when Splunk restarts indication.`,
				},
				resource.Attribute{
					Name:        "realname",
					Description: `(Optional) Full user name.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) Role to assign to this user. At least one existing role is required.`,
				},
				resource.Attribute{
					Name:        "tz",
					Description: `(Optional) User timezone. ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_authorization_roles",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create and update role information.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user role to create.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `(Optional) List of capabilities assigned to role.`,
				},
				resource.Attribute{
					Name:        "cumulative_realtime_search_jobs_quota",
					Description: `(Optional) Maximum number of concurrently running real-time searches that all members of this role can have.`,
				},
				resource.Attribute{
					Name:        "cumulative_search_jobs_quota",
					Description: `(Optional) Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.`,
				},
				resource.Attribute{
					Name:        "default_app",
					Description: `Specify the folder name of the default app to use for this role. A user-specific default app overrides this.`,
				},
				resource.Attribute{
					Name:        "imported_roles",
					Description: `(Optional) List of imported roles for this role. <br>Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.`,
				},
				resource.Attribute{
					Name:        "realtime_search_jobs_quota",
					Description: `(Optional) Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.`,
				},
				resource.Attribute{
					Name:        "search_disk_quota",
					Description: `(Optional) Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.`,
				},
				resource.Attribute{
					Name:        "search_filter",
					Description: `(Optional) Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.`,
				},
				resource.Attribute{
					Name:        "search_indexes_allowed",
					Description: `(Optional) List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.`,
				},
				resource.Attribute{
					Name:        "search_indexes_default",
					Description: `(Optional) List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '`,
				},
				resource.Attribute{
					Name:        "search_jobs_quota",
					Description: `(Optional) The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.`,
				},
				resource.Attribute{
					Name:        "search_time_win",
					Description: `(Optional) Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does. ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_configs_conf",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create and manage configuration file stanzas.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A '/' separated string consisting of {conf_file_name}/{stanza_name} ex. props/custom_stanza`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `(Optional) A map of key value pairs for a stanza.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_data_ui_views",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create and manage splunk dashboards/views.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Dashboard name.`,
				},
				resource.Attribute{
					Name:        "eai:data",
					Description: `(Required) Dashboard XML definition. ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dashboard`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dashboard`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_generic_acl",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Manage the ACL of any Splunk object not already managed in Terraform. To define the ACL of an object that is itself
managed in Terraform, use the ` + "`" + `acl` + "`" + ` block on that configured resource instead of using a ` + "`" + `splunk_generic_acl` + "`" + ` resource.

Note: This resource doesn't actually create any remote resources, because ACLs can only exist (and always exist) for
knowledge objects. They can, however, be managed separately.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The ACL to apply to the object, including app/owner to properly identify the object. Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for app and owner for objects that don't fit in the normal namespace. ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource ## Import Generic ACL resources can be imported by specifying their owner, app, and path with a colon-delimited string as the ID: ` + "`" + `` + "`" + `` + "`" + ` terraform import splunk_generic_acl <owner>:<app>:<path> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource ## Import Generic ACL resources can be imported by specifying their owner, app, and path with a colon-delimited string as the ID: ` + "`" + `` + "`" + `` + "`" + ` terraform import splunk_generic_acl <owner>:<app>:<path> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_global_http_event_collector",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Update Global HTTP Event Collector input configuration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Input disabled indicator.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) HTTP data input IP port.`,
				},
				resource.Attribute{
					Name:        "enable_ssl",
					Description: `(Optional) Enable SSL protocol for HTTP data input. ` + "`" + `true` + "`" + ` = SSL enabled, ` + "`" + `false` + "`" + ` = SSL disabled.`,
				},
				resource.Attribute{
					Name:        "dedicated_io_threads",
					Description: `(Optional) Number of threads used by HTTP Input server.`,
				},
				resource.Attribute{
					Name:        "max_sockets",
					Description: `(Optional) Maximum number of simultaneous HTTP connections accepted. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.`,
				},
				resource.Attribute{
					Name:        "max_threads",
					Description: `(Optional) Maximum number of threads that can be used by active HTTP transactions. Adjusting this value may cause server performance issues and is not generally recommended. Possible values for this setting vary by OS.`,
				},
				resource.Attribute{
					Name:        "use_deployment_server",
					Description: `(Optional) Indicates whether the event collector input writes its configuration to a deployment server repository. When this setting is set to 1 (enabled), the input writes its configuration to the directory specified as repositoryLocation in serverclass.conf. Copy the full contents of the splunk_httpinput app directory to this directory for the configuration to work. When enabled, only the tokens defined in the splunk_httpinput app in this repository are viewable and editable on the API and the Data Inputs page in Splunk Web. When disabled, the input writes its configuration to $SPLUNK_HOME/etc/apps by default. Defaults to 0 (disabled).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_indexes",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create and manage data indexes.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the index to create.`,
				},
				resource.Attribute{
					Name:        "block_sign_size",
					Description: `(Optional) Controls how many events make up a block for block signatures. If this is set to 0, block signing is disabled for this index. <br>A recommended value is 100.`,
				},
				resource.Attribute{
					Name:        "bucket_rebuild_memory_hint",
					Description: `(Optional) Suggestion for the bucket rebuild process for the size of the time-series (tsidx) file to make. <be>Caution: This is an advanced parameter. Inappropriate use of this parameter causes splunkd to not start if rebuild is required. Do not set this parameter unless instructed by Splunk Support. Default value, auto, varies by the amount of physical RAM on the host<br> less than 2GB RAM = 67108864 (64MB) tsidx 2GB to 8GB RAM = 134217728 (128MB) tsidx more than 8GB RAM = 268435456 (256MB) tsidx<br> Values other than "auto" must be 16MB-1GB. Highest legal value (of the numerical part) is 4294967295 You can specify the value using a size suffix: "16777216" or "16MB" are equivalent.`,
				},
				resource.Attribute{
					Name:        "cold_path",
					Description: `(Optional) An absolute path that contains the colddbs for the index. The path must be readable and writable. Cold databases are opened as needed when searching.`,
				},
				resource.Attribute{
					Name:        "cold_to_frozen_dir",
					Description: `(Optional) Destination path for the frozen archive. Use as an alternative to a coldToFrozenScript. Splunk software automatically puts frozen buckets in this directory. <br> Bucket freezing policy is as follows:<br> New style buckets (4.2 and on): removes all files but the rawdata<br> To thaw, run splunk rebuild <bucket dir> on the bucket, then move to the thawed directory<br> Old style buckets (Pre-4.2): gzip all the .data and .tsidx files<br> To thaw, gunzip the zipped files and move the bucket into the thawed directory<br> If both coldToFrozenDir and coldToFrozenScript are specified, coldToFrozenDir takes precedence`,
				},
				resource.Attribute{
					Name:        "cold_to_frozen_script",
					Description: `(Optional) Path to the archiving script. <br>If your script requires a program to run it (for example, python), specify the program followed by the path. The script must be in $SPLUNK_HOME/bin or one of its subdirectories. <br>Splunk software ships with an example archiving script in $SPLUNK_HOME/bin called coldToFrozenExample.py. DO NOT use this example script directly. It uses a default path, and if modified in place any changes are overwritten on upgrade. <br>It is best to copy the example script to a new file in bin and modify it for your system. Most importantly, change the default archive path to an existing directory that fits your needs.`,
				},
				resource.Attribute{
					Name:        "compress_rawdata",
					Description: `(Optional) This parameter is ignored. The splunkd process always compresses raw data.`,
				},
				resource.Attribute{
					Name:        "datatype",
					Description: `(Optional) Valid values: (event | metric). Specifies the type of index.`,
				},
				resource.Attribute{
					Name:        "enable_online_bucket_repair",
					Description: `(Optional) Enables asynchronous "online fsck" bucket repair, which runs concurrently with Splunk software. When enabled, you do not have to wait until buckets are repaired to start the Splunk platform. However, you might observe a slight performance degratation.`,
				},
				resource.Attribute{
					Name:        "frozen_time_period_in_secs",
					Description: `(Optional) Number of seconds after which indexed data rolls to frozen. Defaults to 188697600 (6 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.`,
				},
				resource.Attribute{
					Name:        "home_path",
					Description: `(Optional) An absolute path that contains the hot and warm buckets for the index. Required. Splunk software does not start if an index lacks a valid homePath. <br>Caution: The path must be readable and writable.`,
				},
				resource.Attribute{
					Name:        "max_bloom_backfill_bucket_age",
					Description: `(Optional) Valid values are: Integer[m|s|h|d]. <br>If a warm or cold bucket is older than the specified age, do not create or rebuild its bloomfilter. Specify 0 to never rebuild bloomfilters.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_optimizes",
					Description: `(Optional) The number of concurrent optimize processes that can run against a hot bucket. This number should be increased if instructed by Splunk Support. Typically the default value should suffice.`,
				},
				resource.Attribute{
					Name:        "max_data_size",
					Description: `(Optional) The maximum size in MB for a hot DB to reach before a roll to warm is triggered. Specifying "auto" or "auto_high_volume" causes Splunk software to autotune this parameter (recommended). Use "auto_high_volume" for high volume indexes (such as the main index); otherwise, use "auto". A "high volume index" would typically be considered one that gets over 10GB of data per day.`,
				},
				resource.Attribute{
					Name:        "max_hot_buckets",
					Description: `(Optional) Maximum hot buckets that can exist per index. Defaults to 3. <br>When maxHotBuckets is exceeded, Splunk software rolls the least recently used (LRU) hot bucket to warm. Both normal hot buckets and quarantined hot buckets count towards this total. This setting operates independently of maxHotIdleSecs, which can also cause hot buckets to roll.`,
				},
				resource.Attribute{
					Name:        "max_hot_idle_secs",
					Description: `(Optional) Maximum life, in seconds, of a hot bucket. Defaults to 0. If a hot bucket exceeds maxHotIdleSecs, Splunk software rolls it to warm. This setting operates independently of maxHotBuckets, which can also cause hot buckets to roll. A value of 0 turns off the idle check (equivalent to INFINITE idle time).`,
				},
				resource.Attribute{
					Name:        "max_hot_span_secs",
					Description: `(Optional) Upper bound of target maximum timespan of hot/warm buckets in seconds. Defaults to 7776000 seconds (90 days).`,
				},
				resource.Attribute{
					Name:        "max_mem_mb",
					Description: `(Optional) The amount of memory, expressed in MB, to allocate for buffering a single tsidx file into memory before flushing to disk. Defaults to 5. The default is recommended for all environments.`,
				},
				resource.Attribute{
					Name:        "max_meta_entries",
					Description: `(Optional) Sets the maximum number of unique lines in .data files in a bucket, which may help to reduce memory consumption. If set to 0, this setting is ignored (it is treated as infinite). If exceeded, a hot bucket is rolled to prevent further increase. If your buckets are rolling due to Strings.data hitting this limit, the culprit may be the punct field in your data. If you do not use punct, it may be best to simply disable this (see props.conf.spec in $SPLUNK_HOME/etc/system/README).`,
				},
				resource.Attribute{
					Name:        "max_meta_entries",
					Description: `(Optional) Upper limit, in seconds, on how long an event can sit in raw slice. Applies only if replication is enabled for this index. Otherwise ignored. If there are any acknowledged events sharing this raw slice, this paramater does not apply. In this case, maxTimeUnreplicatedWithAcks applies. Highest legal value is 2147483647. To disable this parameter, set to 0.`,
				},
				resource.Attribute{
					Name:        "max_time_unreplicated_no_acks",
					Description: `(Optional) Upper limit, in seconds, on how long an event can sit in raw slice. Applies only if replication is enabled for this index. Otherwise ignored. If there are any acknowledged events sharing this raw slice, this paramater does not apply. In this case, maxTimeUnreplicatedWithAcks applies. Highest legal value is 2147483647. To disable this parameter, set to 0.`,
				},
				resource.Attribute{
					Name:        "max_time_unreplicated_with_acks",
					Description: `(Optional) Upper limit, in seconds, on how long events can sit unacknowledged in a raw slice. Applies only if you have enabled acks on forwarders and have replication enabled (with clustering). Note: This is an advanced parameter. Make sure you understand the settings on all forwarders before changing this. This number should not exceed ack timeout configured on any forwarder, and should actually be set to at most half of the minimum value of that timeout. You can find this setting in outputs.conf readTimeout setting under the tcpout stanza. To disable, set to 0, but this is NOT recommended. Highest legal value is 2147483647.`,
				},
				resource.Attribute{
					Name:        "max_total_data_size_mb",
					Description: `(Optional) The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen.`,
				},
				resource.Attribute{
					Name:        "max_warm_db_count",
					Description: `(Optional) The maximum number of warm buckets. If this number is exceeded, the warm bucket/s with the lowest value for their latest times is moved to cold.`,
				},
				resource.Attribute{
					Name:        "min_raw_file_sync_secs",
					Description: `(Optional) Specify an integer (or "disable") for this parameter. This parameter sets how frequently splunkd forces a filesystem sync while compressing journal slices. During this period, uncompressed slices are left on disk even after they are compressed. Then splunkd forces a filesystem sync of the compressed journal and removes the accumulated uncompressed files. If 0 is specified, splunkd forces a filesystem sync after every slice completes compressing. Specifying "disable" disables syncing entirely: uncompressed slices are removed as soon as compression is complete.`,
				},
				resource.Attribute{
					Name:        "min_stream_group_queue_size",
					Description: `(Optional) Minimum size of the queue that stores events in memory before committing them to a tsidx file.`,
				},
				resource.Attribute{
					Name:        "partial_service_meta_period",
					Description: `(Optional) Related to serviceMetaPeriod. If set, it enables metadata sync every <integer> seconds, but only for records where the sync can be done efficiently in-place, without requiring a full re-write of the metadata file. Records that require full re-write are be sync'ed at serviceMetaPeriod. partialServiceMetaPeriod specifies, in seconds, how frequently it should sync. Zero means that this feature is turned off and serviceMetaPeriod is the only time when metadata sync happens. If the value of partialServiceMetaPeriod is greater than serviceMetaPeriod, this setting has no effect. By default it is turned off (zero).`,
				},
				resource.Attribute{
					Name:        "process_tracker_service_interval",
					Description: `(Optional) Specifies, in seconds, how often the indexer checks the status of the child OS processes it launched to see if it can launch new processes for queued requests. Defaults to 15. If set to 0, the indexer checks child process status every second. Highest legal value is 4294967295.`,
				},
				resource.Attribute{
					Name:        "quarantine_future_secs",
					Description: `(Optional) Events with timestamp of quarantineFutureSecs newer than "now" are dropped into quarantine bucket. Defaults to 2592000 (30 days). This is a mechanism to prevent main hot buckets from being polluted with fringe events.`,
				},
				resource.Attribute{
					Name:        "quarantine_past_secs",
					Description: `(Optional) Events with timestamp of quarantinePastSecs older than "now" are dropped into quarantine bucket. Defaults to 77760000 (900 days). This is a mechanism to prevent the main hot buckets from being polluted with fringe events.`,
				},
				resource.Attribute{
					Name:        "raw_chunk_size_bytes",
					Description: `(Optional) Target uncompressed size in bytes for individual raw slice in the rawdata journal of the index. Defaults to 131072 (128KB). 0 is not a valid value. If 0 is specified, rawChunkSizeBytes is set to the default value.`,
				},
				resource.Attribute{
					Name:        "rep_factor",
					Description: `(Optional) Index replication control. This parameter applies to only clustering slaves. auto = Use the master index replication configuration value. 0 = Turn off replication for this index.`,
				},
				resource.Attribute{
					Name:        "rotate_period_in_secs",
					Description: `(Optional) How frequently (in seconds) to check if a new hot bucket needs to be created. Also, how frequently to check if there are any warm/cold buckets that should be rolled/frozen.`,
				},
				resource.Attribute{
					Name:        "service_meta_period",
					Description: `(Optional) Defines how frequently metadata is synced to disk, in seconds. Defaults to 25 (seconds). You may want to set this to a higher value if the sum of your metadata file sizes is larger than many tens of megabytes, to avoid the hit on I/O in the indexing fast path.`,
				},
				resource.Attribute{
					Name:        "sync_meta",
					Description: `(Optional) When true, a sync operation is called before file descriptor is closed on metadata file updates. This functionality improves integrity of metadata files, especially in regards to operating system crashes/machine failures.`,
				},
				resource.Attribute{
					Name:        "thawed_path",
					Description: `(Optional) An absolute path that contains the thawed (resurrected) databases for the index. Cannot be defined in terms of a volume definition. Required. Splunk software does not start if an index lacks a valid thawedPath.`,
				},
				resource.Attribute{
					Name:        "throttle_check_period",
					Description: `(Optional) Defines how frequently Splunk software checks for index throttling condition, in seconds. Defaults to 15 (seconds).`,
				},
				resource.Attribute{
					Name:        "tstats_home_path",
					Description: `(Optional) Location to store datamodel acceleration TSIDX data for this index. Restart splunkd after changing this parameter. If specified, it must be defined in terms of a volume definition.`,
				},
				resource.Attribute{
					Name:        "warm_to_cold_script",
					Description: `(Optional) Path to a script to run when moving data from warm to cold. This attribute is supported for backwards compatibility with Splunk software versions older than 4.0. Contact Splunk support if you need help configuring this setting.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_inputs_http_event_collector",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create or update HTTP Event Collector input configuration tokens.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Token name (inputs.conf key)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) Token value for sending data to collector/event endpoint`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Index to store generated events`,
				},
				resource.Attribute{
					Name:        "indexes",
					Description: `(Optional) Set of indexes allowed for events with this token`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Default host value for events with this token`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Default source for events with this token`,
				},
				resource.Attribute{
					Name:        "sourcetype",
					Description: `(Optional) Default source type for events with this token`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Input disabled indicator`,
				},
				resource.Attribute{
					Name:        "use_ack",
					Description: `(Optional) Indexer acknowledgement for this token`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_inputs_monitor",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create or update a new file or directory monitor input.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The file or directory path to monitor on the system.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Which index events from this input should be stored in. Defaults to default.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The value to populate in the host field for events from this data input.`,
				},
				resource.Attribute{
					Name:        "sourcetype",
					Description: `(Optional) The value to populate in the sourcetype field for incoming events.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Indicates if input monitoring is disabled.`,
				},
				resource.Attribute{
					Name:        "rename_source",
					Description: `(Optional) The value to populate in the source field for events from this data input. The same source should not be used for multiple data inputs.`,
				},
				resource.Attribute{
					Name:        "blacklist",
					Description: `(Optional) Specify a regular expression for a file path. The file path that matches this regular expression is not indexed.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `(Optional) Specify a regular expression for a file path. Only file paths that match this regular expression are indexed.`,
				},
				resource.Attribute{
					Name:        "crc_salt",
					Description: `(Optional) A string that modifies the file tracking identity for files in this input. The magic value <SOURCE> invokes special behavior.`,
				},
				resource.Attribute{
					Name:        "follow_tail",
					Description: `(Optional) If set to true, files that are seen for the first time is read from the end.`,
				},
				resource.Attribute{
					Name:        "recursive",
					Description: `(Optional) Setting this to false prevents monitoring of any subdirectories encountered within this data input.`,
				},
				resource.Attribute{
					Name:        "host_regex",
					Description: `(Optional) Specify a regular expression for a file path. If the path for a file matches this regular expression, the captured value is used to populate the host field for events from this data input. The regular expression must have one capture group.`,
				},
				resource.Attribute{
					Name:        "host_segment",
					Description: `(Optional) Use the specified slash-separate segment of the filepath as the host field value.`,
				},
				resource.Attribute{
					Name:        "ignore_older_than",
					Description: `(Optional) Specify a time value. If the modification time of a file being monitored falls outside of this rolling time window, the file is no longer being monitored.`,
				},
				resource.Attribute{
					Name:        "time_before_close",
					Description: `(Optional) When Splunk software reaches the end of a file that is being read, the file is kept open for a minimum of the number of seconds specified in this value. After this period has elapsed, the file is checked again for more data.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_inputs_script",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create or update scripted inputs.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specify the name of the scripted input.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Specifies whether the input script is disabled.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Sets the index for events from this input. Defaults to the main index.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Sets the host for events from this input. Defaults to whatever host sent the event.`,
				},
				resource.Attribute{
					Name:        "sourcetype",
					Description: `(Optional) Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default. Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Sets the source key/field for events from this input. Defaults to the input file path. Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.`,
				},
				resource.Attribute{
					Name:        "rename_source",
					Description: `(Optional) Specify a new name for the source field for the script.`,
				},
				resource.Attribute{
					Name:        "passauth",
					Description: `(Optional) User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_inputs_tcp_cooked",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create or update cooked TCP input information and create new containers for managing cooked data.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The port number of this input.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Indicates if input is disabled.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host from which the indexer gets data.`,
				},
				resource.Attribute{
					Name:        "restrict_to_host",
					Description: `(Optional) Restrict incoming connections on this port to the host specified here.`,
				},
				resource.Attribute{
					Name:        "connection_host",
					Description: `(Optional) Valid values: (ip | dns | none) Set the host for the remote server that is sending data. ip sets the host to the IP address of the remote server sending data. dns sets the host to the reverse DNS entry for the IP address of the remote server sending data. none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname. Default value is dns.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_inputs_tcp_raw",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create or update raw TCP input information for managing raw tcp inputs from forwarders.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The input port which receives raw data.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Indicates if input is disabled.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Index to store generated events. Defaults to default.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host from which the indexer gets data.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Sets the source key/field for events from this input. Defaults to the input file path. Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.`,
				},
				resource.Attribute{
					Name:        "sourcetype",
					Description: `(Optional) Set the source type for events from this input. "sourcetype=" is automatically prepended to <string>. Defaults to audittrail (if signedaudit=true) or fschange (if signedaudit=false).`,
				},
				resource.Attribute{
					Name:        "restrict_to_host",
					Description: `(Optional) Allows for restricting this input to only accept data from the host specified here.`,
				},
				resource.Attribute{
					Name:        "queue",
					Description: `(Optional) Valid values: (parsingQueue | indexQueue) Specifies where the input processor should deposit the events it reads. Defaults to parsingQueue. Set queue to parsingQueue to apply props.conf and other parsing rules to your data. For more information about props.conf and rules for timestamping and linebreaking, refer to props.conf and the online documentation at "Monitor files and directories with inputs.conf" Set queue to indexQueue to send your data directly into the index.`,
				},
				resource.Attribute{
					Name:        "connection_host",
					Description: `(Optional) Valid values: (ip | dns | none) Set the host for the remote server that is sending data. ip sets the host to the IP address of the remote server sending data. dns sets the host to the reverse DNS entry for the IP address of the remote server sending data. none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname. Default value is dns.`,
				},
				resource.Attribute{
					Name:        "raw_tcp_done_timeout",
					Description: `(Optional) Specifies in seconds the timeout value for adding a Done-key. Default value is 10 seconds. If a connection over the port specified by name remains idle after receiving data for specified number of seconds, it adds a Done-key. This implies the last event is completely received.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_inputs_tcp_splunktcptoken",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Manage receiver access using tokens.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Required. Name for the token to create.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) Optional. Token value to use. If unspecified, a token is generated automatically.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_inputs_tcp_ssl",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Access or update the SSL configuration for the host.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Indicates if input is disabled.`,
				},
				resource.Attribute{
					Name:        "root_ca",
					Description: `(Optional) Certificate authority list (root file)`,
				},
				resource.Attribute{
					Name:        "server_cert",
					Description: `(Optional) Full path to the server certificate.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Server certificate password, if any.`,
				},
				resource.Attribute{
					Name:        "require_client_cert",
					Description: `(Optional) Determines whether a client must authenticate. ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_inputs_udp",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create and manage UDP data inputs.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The UDP port that this input should listen on.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Indicates if input is disabled.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Which index events from this input should be stored in. Defaults to default.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The value to populate in the host field for incoming events. This is used during parsing/indexing, in particular to set the host field. It is also the host field used at search time.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) The value to populate in the source field for incoming events. The same source should not be used for multiple data inputs.`,
				},
				resource.Attribute{
					Name:        "sourcetype",
					Description: `(Optional) The value to populate in the sourcetype field for incoming events.`,
				},
				resource.Attribute{
					Name:        "restrict_to_host",
					Description: `(Optional) Restrict incoming connections on this port to the host specified here. If this is not set, the value specified in [udp://<remote server>:<port>] in inputs.conf is used.`,
				},
				resource.Attribute{
					Name:        "queue",
					Description: `(Optional) Which queue events from this input should be sent to. Generally this does not need to be changed.`,
				},
				resource.Attribute{
					Name:        "connection_host",
					Description: `(Optional) Valid values: (ip | dns | none) Set the host for the remote server that is sending data. ip sets the host to the IP address of the remote server sending data. dns sets the host to the reverse DNS entry for the IP address of the remote server sending data. none leaves the host as specified in inputs.conf, which is typically the Splunk system hostname. Default value is dns.`,
				},
				resource.Attribute{
					Name:        "no_appending_timestamp",
					Description: `(Optional) If set to true, prevents Splunk software from prepending a timestamp and hostname to incoming events.`,
				},
				resource.Attribute{
					Name:        "no_priority_stripping",
					Description: `(Optional) If set to true, Splunk software does not remove the priority field from incoming syslog events.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_outputs_tcp_default",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Manage to global tcpout properties.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Configuration to be edited. The only valid value is "tcpout".`,
				},
				resource.Attribute{
					Name:        "default_group",
					Description: `(Optional) Comma-separated list of one or more target group names, specified later in [tcpout:<target_group>] stanzas of outputs.conf.spec file. The forwarder sends all data to the specified groups. If you do not want to forward data automatically, do not set this attribute. Can be overridden by an inputs.conf _TCP_ROUTING setting, which in turn can be overridden by a props.conf/transforms.conf modifier.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Disables default tcpout settings`,
				},
				resource.Attribute{
					Name:        "drop_events_on_queue_full",
					Description: `(Optional) If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events). <br>CAUTION: Do not set this value to a positive integer if you are monitoring files. Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group. Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.`,
				},
				resource.Attribute{
					Name:        "heartbeat_frequency",
					Description: `(Optional) How often (in seconds) to send a heartbeat packet to the receiving server. Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.`,
				},
				resource.Attribute{
					Name:        "index_and_forward",
					Description: `(Optional) Specifies whether to index all data locally, in addition to forwarding it. Defaults to false. This is known as an "index-and-forward" configuration. This attribute is only available for heavy forwarders. It is available only at the top level [tcpout] stanza in outputs.conf. It cannot be overridden in a target group.`,
				},
				resource.Attribute{
					Name:        "max_queue_size",
					Description: `(Optional) Specify an integer or integer[KB|MB|GB]. <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true). Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer. For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting. If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks. If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).`,
				},
				resource.Attribute{
					Name:        "send_cooked_data",
					Description: `(Optional) If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true. Set to false if you are sending to a third-party system.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_outputs_tcp_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Access to the configuration of a group of one or more data forwarding destinations.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the group of receivers.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) If true, disables the group.`,
				},
				resource.Attribute{
					Name:        "compressed",
					Description: `(Optional) If true, forwarder sends compressed data. If set to true, the receiver port must also have compression turned on.`,
				},
				resource.Attribute{
					Name:        "drop_events_on_queue_full",
					Description: `(Optional) If set to a positive number, wait the specified number of seconds before throwing out all new events until the output queue has space. Defaults to -1 (do not drop events). <br>CAUTION: Do not set this value to a positive integer if you are monitoring files. Setting this to -1 or 0 causes the output queue to block when it gets full, which causes further blocking up the processing chain. If any target group queue is blocked, no more data reaches any other target group. Using auto load-balancing is the best way to minimize this condition, because, in that case, multiple receivers must be down (or jammed up) before queue blocking can occur.`,
				},
				resource.Attribute{
					Name:        "heartbeat_frequency",
					Description: `(Optional) How often (in seconds) to send a heartbeat packet to the receiving server. Heartbeats are only sent if sendCookedData=true. Defaults to 30 seconds.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) Valid values: (tcpout | syslog). Specifies the type of output processor.`,
				},
				resource.Attribute{
					Name:        "max_queue_size",
					Description: `(Optional) Specify an integer or integer[KB|MB|GB]. <br>Sets the maximum size of the forwarder output queue. It also sets the maximum size of the wait queue to 3x this value, if you have enabled indexer acknowledgment (useACK=true). Although the wait queue and the output queues are both configured by this attribute, they are separate queues. The setting determines the maximum size of the queue in-memory (RAM) buffer. For heavy forwarders sending parsed data, maxQueueSize is the maximum number of events. Since events are typically much shorter than data blocks, the memory consumed by the queue on a parsing forwarder is likely to be much smaller than on a non-parsing forwarder, if you use this version of the setting. If specified as a lone integer (for example, maxQueueSize=100), maxQueueSize indicates the maximum number of queued events (for parsed data) or blocks of data (for unparsed data). A block of data is approximately 64KB. For non-parsing forwarders, such as universal forwarders, that send unparsed data, maxQueueSize is the maximum number of data blocks. If specified as an integer followed by KB, MB, or GB (for example, maxQueueSize=100MB), maxQueueSize indicates the maximum RAM allocated to the queue buffer. Defaults to 500KB (which means a maximum size of 500KB for the output queue and 1500KB for the wait queue, if any).`,
				},
				resource.Attribute{
					Name:        "send_cooked_data",
					Description: `(Optional) If true, events are cooked (processed by Splunk software). If false, events are raw and untouched prior to sending. Defaults to true. Set to false if you are sending to a third-party system.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Optional) Comma-separated list of servers to include in the group.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) Token value generated by the indexer after configuration.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_outputs_tcp_server",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Access data forwarding configurations.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) <host>:<port> of the Splunk receiver. <host> can be either an ip address or server name. <port> is the that port that the Splunk receiver is listening on.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) If true, disables the group.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `Valid values: (clone | balance | autobalance) The data distribution method used when two or more servers exist in the same forwarder group.`,
				},
				resource.Attribute{
					Name:        "ssl_alt_name_to_check",
					Description: `(Optional) The alternate name to match in the remote server's SSL certificate.`,
				},
				resource.Attribute{
					Name:        "ssl_cert_path",
					Description: `(Optional) Path to the client certificate. If specified, connection uses SSL.`,
				},
				resource.Attribute{
					Name:        "ssl_cipher",
					Description: `(Optional) SSL Cipher in the form ALL:!aNULL:!eNULL:!LOW:!EXP:RC4+RSA:+HIGH:+MEDIUM`,
				},
				resource.Attribute{
					Name:        "ssl_common_name_to_check",
					Description: `(Optional) Check the common name of the server's certificate against this name. If there is no match, assume that Splunk Enterprise is not authenticated against this server. You must specify this setting if sslVerifyServerCert is true.`,
				},
				resource.Attribute{
					Name:        "ssl_root_ca_path",
					Description: `(Optional) The path to the root certificate authority file.`,
				},
				resource.Attribute{
					Name:        "ssl_password",
					Description: `(Optional) The password associated with the CAcert. The default Splunk Enterprise CAcert uses the password "password."`,
				},
				resource.Attribute{
					Name:        "ssl_verify_server_cert",
					Description: `(Optional) If true, make sure that the server you are connecting to is a valid one (authenticated). Both the common name and the alternate name of the server are then checked for a match.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_outputs_tcp_syslog",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Access the configuration of a forwarded server configured to provide data in standard syslog format.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the syslog output group. This is name used when creating syslog configuration in outputs.conf.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) If true, disables global syslog settings.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Sets syslog priority value. The priority value should specified as an integer. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) host:port of the server where syslog data should be sent`,
				},
				resource.Attribute{
					Name:        "syslog_sourcetype",
					Description: `(Optional) Specifies a rule for handling data in addition to that provided by the "syslog" sourcetype. By default, there is no value for syslogSourceType. <br>This string is used as a substring match against the sourcetype key. For example, if the string is set to 'syslog', then all source types containing the string "syslog" receives this special treatment. To match a source type explicitly, use the pattern "sourcetype::sourcetype_name." For example syslogSourcetype = sourcetype::apache_common Data that is "syslog" or matches this setting is assumed to already be in syslog format. Data that does not match the rules has a header, potentially a timestamp, and a hostname added to the front of the event. This is how Splunk software causes arbitrary log data to match syslog expectations.`,
				},
				resource.Attribute{
					Name:        "timestamp_format",
					Description: `(Optional) Format of timestamp to add at start of the events to be forwarded. The format is a strftime-style timestamp formatting string. See $SPLUNK_HOME/etc/system/README/outputs.conf.spec for details.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Protocol to use to send syslog data. Valid values: (tcp | udp ).`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the http event collector resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_saved_searches",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create and manage saved searches.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the search.`,
				},
				resource.Attribute{
					Name:        "search",
					Description: `(Required) Required when creating a new search.`,
				},
				resource.Attribute{
					Name:        "action_email",
					Description: `(Optional) The state of the email action. Read-only attribute. Value ignored on POST. Use actions to specify a list of enabled actions. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_email_auth_password",
					Description: `(Optional) The password to use when authenticating with the SMTP server. Normally this value is set when editing the email settings, however you can set a clear text password here and it is encrypted on the next platform restart.Defaults to empty string.`,
				},
				resource.Attribute{
					Name:        "action_email_auth_username",
					Description: `(Optional) The username to use when authenticating with the SMTP server. If this is empty string, no authentication is attempted. Defaults to empty stringNOTE: Your SMTP server might reject unauthenticated emails.`,
				},
				resource.Attribute{
					Name:        "action_email_bcc",
					Description: `(Optional) BCC email address to use if action.email is enabled.`,
				},
				resource.Attribute{
					Name:        "action_email_cc",
					Description: `(Optional) CC email address to use if action.email is enabled.`,
				},
				resource.Attribute{
					Name:        "action_email_command",
					Description: `(Optional) The search command (or pipeline) which is responsible for executing the action.Generally the command is a template search pipeline which is realized with values from the saved search. To reference saved search field values wrap them in $, for example to reference the savedsearch name use $name$, to reference the search use $search$.`,
				},
				resource.Attribute{
					Name:        "action_email_format",
					Description: `(Optional) Valid values: (table | plain | html | raw | csv)Specify the format of text in the email. This value also applies to any attachments.`,
				},
				resource.Attribute{
					Name:        "action_email_from",
					Description: `(Optional) Email address from which the email action originates.Defaults to splunk@$LOCALHOST or whatever value is set in alert_actions.conf.`,
				},
				resource.Attribute{
					Name:        "action_email_hostname",
					Description: `(Optional) Sets the hostname used in the web link (url) sent in email actions.This value accepts two forms:hostname (for example, splunkserver, splunkserver.example.com)`,
				},
				resource.Attribute{
					Name:        "action_email_include_results_link",
					Description: `(Optional) Specify whether to include a link to the results. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_email_include_search",
					Description: `(Optional) Specify whether to include the search that caused an email to be sent. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_email_include_trigger",
					Description: `(Optional) Specify whether to show the trigger condition that caused the alert to fire. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_email_include_trigger_time",
					Description: `(Optional) Specify whether to show the time that the alert was fired. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_email_include_view_link",
					Description: `(Optional) Specify whether to show the title and a link to enable the user to edit the saved search. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_email_inline",
					Description: `(Optional) Indicates whether the search results are contained in the body of the email.Results can be either inline or attached to an email.`,
				},
				resource.Attribute{
					Name:        "action_email_mailserver",
					Description: `(Optional) Set the address of the MTA server to be used to send the emails.Defaults to <LOCALHOST> or whatever is set in alert_actions.conf.`,
				},
				resource.Attribute{
					Name:        "action_email_max_results",
					Description: `(Optional) Sets the global maximum number of search results to send when email.action is enabled. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "action_email_max_time",
					Description: `(Optional) Valid values are Integer[m|s|h|d].Specifies the maximum amount of time the execution of an email action takes before the action is aborted. Defaults to 5m.`,
				},
				resource.Attribute{
					Name:        "action_email_message_alert",
					Description: `(Optional) Customize the message sent in the emailed alert. Defaults to: The alert condition for '$name$' was triggered.`,
				},
				resource.Attribute{
					Name:        "action_email_message_report",
					Description: `(Optional) Customize the message sent in the emailed report. Defaults to: The scheduled report '$name$' has run`,
				},
				resource.Attribute{
					Name:        "action_email_pdfview",
					Description: `(Optional) The name of the view to deliver if sendpdf is enabled`,
				},
				resource.Attribute{
					Name:        "action_email_preprocess_results",
					Description: `(Optional) Search string to preprocess results before emailing them. Defaults to empty string (no preprocessing).Usually the preprocessing consists of filtering out unwanted internal fields.`,
				},
				resource.Attribute{
					Name:        "action_email_report_cid_font_list",
					Description: `(Optional) Space-separated list. Specifies the set (and load order) of CID fonts for handling Simplified Chinese(gb), Traditional Chinese(cns), Japanese(jp), and Korean(kor) in Integrated PDF Rendering.If multiple fonts provide a glyph for a given character code, the glyph from the first font specified in the list is used.To skip loading any CID fonts, specify the empty string.Defaults to 'gb cns jp kor'`,
				},
				resource.Attribute{
					Name:        "action_email_report_include_splunk_logo",
					Description: `(Optional) Indicates whether to include the Splunk logo with the report.`,
				},
				resource.Attribute{
					Name:        "action_email_report_paper_orientation",
					Description: `(Optional) Valid values: (portrait | landscape)Specifies the paper orientation: portrait or landscape. Defaults to portrait.`,
				},
				resource.Attribute{
					Name:        "action_email_report_paper_size",
					Description: `(Optional) Valid values: (letter | legal | ledger | a2 | a3 | a4 | a5)Specifies the paper size for PDFs. Defaults to letter.`,
				},
				resource.Attribute{
					Name:        "action_email_report_server_enabled",
					Description: `(Optional) No Supported`,
				},
				resource.Attribute{
					Name:        "action_email_report_server_url",
					Description: `(Optional) Not supported.For a default locally installed report server, the URL is http://localhost:8091/`,
				},
				resource.Attribute{
					Name:        "action_email_send_csv",
					Description: `(Optional) Specify whether to send results as a CSV file. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_email_send_pdf",
					Description: `(Optional) Indicates whether to create and send the results as a PDF. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "action_email_send_results",
					Description: `(Optional) Indicates whether to attach the search results in the email.Results can be either attached or inline. See action.email.inline.`,
				},
				resource.Attribute{
					Name:        "action_email_subject",
					Description: `(Optional) Specifies an alternate email subject.Defaults to SplunkAlert-<savedsearchname>.`,
				},
				resource.Attribute{
					Name:        "action_email_to",
					Description: `(Optional) A comma or semicolon separated list of recipient email addresses. Required if this search is scheduled and the email alert action is enabled.`,
				},
				resource.Attribute{
					Name:        "action_email_track_alert",
					Description: `(Optional) Indicates whether the execution of this action signifies a trackable alert.`,
				},
				resource.Attribute{
					Name:        "action_email_ttl",
					Description: `(Optional) Valid values are Integer[p].Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered. If p follows <Integer>, int is the number of scheduled periods. Defaults to 86400 (24 hours).If no actions are triggered, the artifacts have their ttl determined by dispatch.ttl in savedsearches.conf.`,
				},
				resource.Attribute{
					Name:        "action_email_use_ssl",
					Description: `(Optional) Indicates whether to use SSL when communicating with the SMTP server. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "action_email_use_tls",
					Description: `(Optional) Indicates whether to use TLS (transport layer security) when communicating with the SMTP server (starttls).Defaults to false.`,
				},
				resource.Attribute{
					Name:        "action_email_width_sort_columns",
					Description: `(Optional) Indicates whether columns should be sorted from least wide to most wide, left to right.Only valid if format=text.`,
				},
				resource.Attribute{
					Name:        "action_populate_lookup",
					Description: `(Optional) The state of the populate lookup action. Read-only attribute. Value ignored on POST. Use actions to specify a list of enabled actions. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_populate_lookup_command",
					Description: `(Optional) The search command (or pipeline) which is responsible for executing the action.`,
				},
				resource.Attribute{
					Name:        "action_populate_lookup_dest",
					Description: `(Optional) Lookup name of path of the lookup to populate`,
				},
				resource.Attribute{
					Name:        "action_populate_lookup_hostname",
					Description: `(Optional) Sets the hostname used in the web link (url) sent in alert actions.This value accepts two forms: hostname (for example, splunkserver, splunkserver.example.com)\n\nprotocol://hostname:port (for example, http://splunkserver:8000, https://splunkserver.example.com:443)`,
				},
				resource.Attribute{
					Name:        "action_populate_lookup_max_results",
					Description: `(Optional) Sets the maximum number of search results sent using alerts. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "action_populate_lookup_max_time",
					Description: `(Optional) Valid values are: Integer[m|s|h|d]Sets the maximum amount of time the execution of an action takes before the action is aborted. Defaults to 5m.`,
				},
				resource.Attribute{
					Name:        "action_populate_lookup_track_alert",
					Description: `(Optional) Indicates whether the execution of this action signifies a trackable alert.`,
				},
				resource.Attribute{
					Name:        "action_populate_lookup_ttl",
					Description: `(Optional) Valid values are Integer[p]Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered. If p follows Integer, then this specifies the number of scheduled periods. Defaults to 10p.`,
				},
				resource.Attribute{
					Name:        "action_rss",
					Description: `(Optional) The state of the rss action. Read-only attribute. Value ignored on POST.Use actions to specify a list of enabled actions. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_rss_command",
					Description: `(Optional) The search command (or pipeline) which is responsible for executing the action.Generally the command is a template search pipeline which is realized with values from the saved search. To reference saved search field values wrap them in $, for example to reference the savedsearch name use $name$, to reference the search use $search$.`,
				},
				resource.Attribute{
					Name:        "action_rss_hostname",
					Description: `(Optional) Sets the hostname used in the web link (url) sent in alert actions.This value accepts two forms:hostname (for example, splunkserver, splunkserver.example.com)\n\nprotocol://hostname:port (for example, http://splunkserver:8000, https://splunkserver.example.com:443)`,
				},
				resource.Attribute{
					Name:        "action_rss_max_results",
					Description: `(Optional) Sets the maximum number of search results sent using alerts. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "action_rss_max_time",
					Description: `(Optional) Valid values are Integer[m|s|h|d].Sets the maximum amount of time the execution of an action takes before the action is aborted. Defaults to 1m.`,
				},
				resource.Attribute{
					Name:        "action_rss_track_alert",
					Description: `(Optional) Indicates whether the execution of this action signifies a trackable alert.`,
				},
				resource.Attribute{
					Name:        "action_rss_ttl",
					Description: `(Optional) Valid values are: Integer[p] Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered. If p follows Integer, specifies the number of scheduled periods. Defaults to 86400 (24 hours).`,
				},
				resource.Attribute{
					Name:        "action_script",
					Description: `(Optional) The state of the script action. Read-only attribute. Value ignored on POST. Use actions to specify a list of enabled actions. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_script_command",
					Description: `(Optional) The search command (or pipeline) which is responsible for executing the action.Generally the command is a template search pipeline which is realized with values from the saved search. To reference saved search field values wrap them in $, for example to reference the savedsearch name use $name$, to reference the search use $search$.`,
				},
				resource.Attribute{
					Name:        "action_script_filename",
					Description: `(Optional) File name of the script to call. Required if script action is enabled`,
				},
				resource.Attribute{
					Name:        "action_script_hostname",
					Description: `(Optional) Sets the hostname used in the web link (url) sent in alert actions.This value accepts two forms:hostname (for example, splunkserver, splunkserver.example.com)\n\nprotocol://hostname:port (for example, http://splunkserver:8000, https://splunkserver.example.com:443)`,
				},
				resource.Attribute{
					Name:        "action_script_max_results",
					Description: `(Optional) Sets the maximum number of search results sent using alerts. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "action_script_max_time",
					Description: `(Optional) Valid values are Integer[m|s|h|d].Sets the maximum amount of time the execution of an action takes before the action is aborted. Defaults to 1m.`,
				},
				resource.Attribute{
					Name:        "action_script_track_alert",
					Description: `(Optional) Indicates whether the execution of this action signifies a trackable alert.`,
				},
				resource.Attribute{
					Name:        "action_script_ttl",
					Description: `(Optional) Valid values are: Integer[p] Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered. If p follows Integer, specifies the number of scheduled periods. Defaults to 86400 (24 hours).`,
				},
				resource.Attribute{
					Name:        "action_summary_index",
					Description: `(Optional) The state of the summary index action. Read-only attribute. Value ignored on POST. Use actions to specify a list of enabled actions. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "action_summary_index_command",
					Description: `(Optional) The search command (or pipeline) which is responsible for executing the action.Generally the command is a template search pipeline which is realized with values from the saved search. To reference saved search field values wrap them in $, for example to reference the savedsearch name use $name$, to reference the search use $search$.`,
				},
				resource.Attribute{
					Name:        "action_summary_index_hostname",
					Description: `(Optional) Sets the hostname used in the web link (url) sent in summary-index alert actions.This value accepts two forms:hostname (for example, splunkserver, splunkserver.example.com)protocol://hostname:port (for example, http://splunkserver:8000, https://splunkserver.example.com:443)`,
				},
				resource.Attribute{
					Name:        "action_summary_index_inline",
					Description: `(Optional) Determines whether to execute the summary indexing action as part of the scheduled search.NOTE: This option is considered only if the summary index action is enabled and is always executed (in other words, if counttype = always).Defaults to true`,
				},
				resource.Attribute{
					Name:        "action_summary_index_max_results",
					Description: `(Optional) Sets the maximum number of search results sent using alerts. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "action_summary_index_max_time",
					Description: `(Optional) Valid values are Integer[m|s|h|d].Sets the maximum amount of time the execution of an action takes before the action is aborted. Defaults to 1m.`,
				},
				resource.Attribute{
					Name:        "action_summary_index_name",
					Description: `(Optional) Specifies the name of the summary index where the results of the scheduled search are saved.Defaults to summary.`,
				},
				resource.Attribute{
					Name:        "action_summary_index_track_alert",
					Description: `(Optional) Indicates whether the execution of this action signifies a trackable alert.`,
				},
				resource.Attribute{
					Name:        "action_summary_index_ttl",
					Description: `(Optional) Valid values are: Integer[p] Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered. If p follows Integer, specifies the number of scheduled periods. Defaults to 86400 (24 hours).`,
				},
				resource.Attribute{
					Name:        "action_create_xsoar_incident",
					Description: `(Optional) Enable XSOAR alerting (Should by 1 (Enabled) or 0 (Disabled))`,
				},
				resource.Attribute{
					Name:        "action_create_xsoar_incident_param_send_all_servers",
					Description: `(Optional) Enable XSOAR alerting sending to all servers (Should by 1 (Enabled) or 0 (Disabled)`,
				},
				resource.Attribute{
					Name:        "action_create_xsoar_incident_param_server_url",
					Description: `(Optional) XSOAR Server instance URL (Should start with https:// || http://)`,
				},
				resource.Attribute{
					Name:        "action_create_xsoar_incident_param_incident_name",
					Description: `(Optional) XSOAR incident name`,
				},
				resource.Attribute{
					Name:        "action_create_xsoar_incident_param_details",
					Description: `(Optional) XSOAR incident description`,
				},
				resource.Attribute{
					Name:        "action_create_xsoar_incident_param_custom_fields",
					Description: `(Optional) XSOAR custom incident fields (should be a comma separated list)`,
				},
				resource.Attribute{
					Name:        "action_create_xsoar_incident_param_severity",
					Description: `(Optional) XSOAR Severity (1 - Low, 2 - Medium, 3 - High, 4 - Critical)`,
				},
				resource.Attribute{
					Name:        "action_create_xsoar_incident_param_occurred",
					Description: `(Optional) XSOAR incident time`,
				},
				resource.Attribute{
					Name:        "action_create_xsoar_incident_param_type",
					Description: `(Optional) XSOAR incident type`,
				},
				resource.Attribute{
					Name:        "action_slack_param_channel",
					Description: `(Optional) Slack channel to send the message to (Should start with # or @)`,
				},
				resource.Attribute{
					Name:        "action_slack_param_fields",
					Description: `(Optional) Show one or more fields from the search results below your Slack message. Comma-separated list of field names. Allows wildcards. eg. index,source`,
				},
				resource.Attribute{
					Name:        "action_slack_param_attachment",
					Description: `(Optional) Include a message attachment. Valid values are message, none, or alert_link`,
				},
				resource.Attribute{
					Name:        "action_slack_param_message",
					Description: `(Optional) Enter the chat message to send to the Slack channel. The message can include tokens that insert text based on the results of the search.`,
				},
				resource.Attribute{
					Name:        "action_slack_param_webhook_url_override",
					Description: `(Optional) You can override the Slack webhook URL here if you need to send the alert message to a different Slack team`,
				},
				resource.Attribute{
					Name:        "action_jira_service_desk_param_account",
					Description: `(Optional) Jira Service Desk account name`,
				},
				resource.Attribute{
					Name:        "action_jira_service_desk_param_jira_project",
					Description: `(Optional) Jira Project name`,
				},
				resource.Attribute{
					Name:        "action_jira_service_desk_param_jira_issue_type",
					Description: `(Optional) Jira issue type name`,
				},
				resource.Attribute{
					Name:        "action_jira_service_desk_param_jira_summary",
					Description: `(Optional) Jira issue title/summary`,
				},
				resource.Attribute{
					Name:        "action_jira_service_desk_param_jira_priority",
					Description: `(Optional) Jira priority of issue`,
				},
				resource.Attribute{
					Name:        "action_jira_service_desk_param_jira_description",
					Description: `(Optional) Jira issue description`,
				},
				resource.Attribute{
					Name:        "action_webhook_param_url",
					Description: `(Optional) URL to send the HTTP POST request to. Must be accessible from the Splunk server`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `(Optional) A comma-separated list of actions to enable. For example: rss,email`,
				},
				resource.Attribute{
					Name:        "alert_comparator",
					Description: `(Optional) One of the following strings: greater than, less than, equal to, rises by, drops by, rises by perc, drops by percUsed with alert_threshold to trigger alert actions.`,
				},
				resource.Attribute{
					Name:        "alert_condition",
					Description: `(Optional) Contains a conditional search that is evaluated against the results of the saved search. Defaults to an empty string.`,
				},
				resource.Attribute{
					Name:        "alert_digest_mode",
					Description: `(Optional) Specifies whether alert actions are applied to the entire result set or on each individual result.Defaults to 1 (true).`,
				},
				resource.Attribute{
					Name:        "alert_expires",
					Description: `(Optional) Valid values: [number][time-unit]Sets the period of time to show the alert in the dashboard. Defaults to 24h.Use [number][time-unit] to specify a time. For example: 60 = 60 seconds, 1m = 1 minute, 1h = 60 minutes = 1 hour.`,
				},
				resource.Attribute{
					Name:        "alert_severity",
					Description: `(Optional) Valid values: (1 | 2 | 3 | 4 | 5 | 6) Sets the alert severity level.Valid values are:1 DEBUG 2 INFO 3 WARN 4 ERROR 5 SEVERE 6 FATAL Defaults to 3.`,
				},
				resource.Attribute{
					Name:        "alert_suppress",
					Description: `(Optional) Indicates whether alert suppression is enabled for this scheduled search.`,
				},
				resource.Attribute{
					Name:        "alert_suppress_fields",
					Description: `(Optional) Comma delimited list of fields to use for suppression when doing per result alerting. Required if suppression is turned on and per result alerting is enabled.`,
				},
				resource.Attribute{
					Name:        "alert_suppress_period",
					Description: `(Optional) Valid values: [number][time-unit] Specifies the suppresion period. Only valid if alert.supress is enabled.Use [number][time-unit] to specify a time. For example: 60 = 60 seconds, 1m = 1 minute, 1h = 60 minutes = 1 hour.`,
				},
				resource.Attribute{
					Name:        "alert_threshold",
					Description: `(Optional) Valid values are: Integer[%]Specifies the value to compare (see alert_comparator) before triggering the alert actions. If expressed as a percentage, indicates value to use when alert_comparator is set to rises by perc or drops by perc.`,
				},
				resource.Attribute{
					Name:        "alert_track",
					Description: `(Optional) Valid values: (true | false | auto) Specifies whether to track the actions triggered by this scheduled search.auto - determine whether to track or not based on the tracking setting of each action, do not track scheduled searches that always trigger actions. Default value true - force alert tracking.false - disable alert tracking for this search.`,
				},
				resource.Attribute{
					Name:        "alert_type",
					Description: `(Optional) What to base the alert on, overriden by alert_condition if it is specified. Valid values are: always, custom, number of events, number of hosts, number of sources.`,
				},
				resource.Attribute{
					Name:        "allow_skew",
					Description: `(Optional) Allows the search scheduler to distribute scheduled searches randomly and more evenly over their specified search periods.`,
				},
				resource.Attribute{
					Name:        "auto_summarize",
					Description: `(Optional) Indicates whether the scheduler should ensure that the data for this search is automatically summarized. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_command",
					Description: `(Optional) An auto summarization template for this search. See auto summarization options in savedsearches.conf for more details.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_cron_schedule",
					Description: `(Optional) Cron schedule that probes and generates the summaries for this saved search.The default value is`,
				},
				resource.Attribute{
					Name:        "auto_summarize_dispatch_earliest_time",
					Description: `(Optional) A time string that specifies the earliest time for summarizing this search. Can be a relative or absolute time.If this value is an absolute time, use the dispatch.time_format to format the value.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_dispatch_latest_time",
					Description: `(Optional) A time string that specifies the latest time for summarizing this saved search. Can be a relative or absolute time.If this value is an absolute time, use the dispatch.time_format to format the value.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_dispatch_time_format",
					Description: `(Optional) Defines the time format that Splunk software uses to specify the earliest and latest time. Defaults to %FT%T.%Q%:z`,
				},
				resource.Attribute{
					Name:        "auto_summarize_dispatch_ttl",
					Description: `(Optional) Valid values: Integer[p]. Defaults to 60.Indicates the time to live (in seconds) for the artifacts of the summarization of the scheduled search.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_max_disabled_buckets",
					Description: `(Optional) The maximum number of buckets with the suspended summarization before the summarization search is completely stopped, and the summarization of the search is suspended for auto_summarize.suspend_period. Defaults to 2.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_max_summary_ratio",
					Description: `(Optional) The maximum ratio of summary_size/bucket_size, which specifies when to stop summarization and deem it unhelpful for a bucket. Defaults to 0.1 Note: The test is only performed if the summary size is larger than auto_summarize.max_summary_size.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_max_summary_size",
					Description: `(Optional) The minimum summary size, in bytes, before testing whether the summarization is helpful.The default value is 52428800 and is equivalent to 5MB.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_max_time",
					Description: `(Optional) Maximum time (in seconds) that the summary search is allowed to run. Defaults to 3600.Note: This is an approximate time. The summary search stops at clean bucket boundaries.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_suspend_period",
					Description: `(Optional) Time specfier indicating when to suspend summarization of this search if the summarization is deemed unhelpful.Defaults to 24h.`,
				},
				resource.Attribute{
					Name:        "auto_summarize_timespan",
					Description: `(Optional) The list of time ranges that each summarized chunk should span. This comprises the list of available granularity levels for which summaries would be available. Specify a comma delimited list of time specifiers.For example a timechart over the last month whose granuality is at the day level should set this to 1d. If you need the same data summarized at the hour level for weekly charts, use: 1h,1d.`,
				},
				resource.Attribute{
					Name:        "cron_schedule",
					Description: `(Optional) Valid values: cron stringThe cron schedule to execute this search. For example:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of this saved search. Defaults to empty string.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Indicates if the saved search is enabled. Defaults to 0.Disabled saved searches are not visible in Splunk Web.`,
				},
				resource.Attribute{
					Name:        "dispatch_buckets",
					Description: `(Optional) The maximum number of timeline buckets. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "dispatch_earliest_time",
					Description: `(Optional) A time string that specifies the earliest time for this search. Can be a relative or absolute time. If this value is an absolute time, use the dispatch.time_format to format the value.`,
				},
				resource.Attribute{
					Name:        "dispatch_index_earliest",
					Description: `(Optional) A time string that specifies the earliest index time for this search. Can be a relative or absolute time. If this value is an absolute time, use the dispatch.time_format to format the value.`,
				},
				resource.Attribute{
					Name:        "dispatch_index_latest",
					Description: `(Optional) A time string that specifies the latest index time for this search. Can be a relative or absolute time. If this value is an absolute time, use the dispatch.time_format to format the value.`,
				},
				resource.Attribute{
					Name:        "dispatch_indexed_realtime",
					Description: `(Optional) A time string that specifies the earliest time for this search. Can be a relative or absolute time. If this value is an absolute time, use the dispatch.time_format to format the value.`,
				},
				resource.Attribute{
					Name:        "dispatch_indexed_realtime_minspan",
					Description: `(Optional) Allows for a per-job override of the [search] indexed_realtime_disk_sync_delay setting in limits.conf.`,
				},
				resource.Attribute{
					Name:        "dispatch_indexed_realtime_offset",
					Description: `(Optional) Allows for a per-job override of the [search] indexed_realtime_disk_sync_delay setting in limits.conf.`,
				},
				resource.Attribute{
					Name:        "dispatch_latest_time",
					Description: `(Optional) A time string that specifies the latest time for this saved search. Can be a relative or absolute time.If this value is an absolute time, use the dispatch.time_format to format the value.`,
				},
				resource.Attribute{
					Name:        "dispatch_lookups",
					Description: `(Optional) Enables or disables the lookups for this search. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "dispatch_max_count",
					Description: `(Optional) The maximum number of results before finalizing the search. Defaults to 500000.`,
				},
				resource.Attribute{
					Name:        "dispatch_max_time",
					Description: `(Optional) Indicates the maximum amount of time (in seconds) before finalizing the search. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "dispatch_reduce_freq",
					Description: `(Optional) Specifies, in seconds, how frequently the MapReduce reduce phase runs on accumulated map values. Defaults to 10.`,
				},
				resource.Attribute{
					Name:        "dispatch_rt_backfill",
					Description: `(Optional) Whether to back fill the real time window for this search. Parameter valid only if this is a real time search. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "dispatch_rt_maximum_span",
					Description: `(Optional) Allows for a per-job override of the [search] indexed_realtime_maximum_span setting in limits.conf.`,
				},
				resource.Attribute{
					Name:        "dispatch_spawn_process",
					Description: `(Optional) Specifies whether a new search process spawns when this saved search is executed. Defaults to 1. Searches against indexes must run in a separate process.`,
				},
				resource.Attribute{
					Name:        "dispatch_time_format",
					Description: `(Optional) A time format string that defines the time format for specifying the earliest and latest time. Defaults to %FT%T.%Q%:z`,
				},
				resource.Attribute{
					Name:        "dispatch_ttl",
					Description: `(Optional) Valid values: Integer[p]. Defaults to 2p.Indicates the time to live (in seconds) for the artifacts of the scheduled search, if no actions are triggered.`,
				},
				resource.Attribute{
					Name:        "display_view",
					Description: `(Optional) Defines the default UI view name (not label) in which to load the results. Accessibility is subject to the user having sufficient permissions.`,
				},
				resource.Attribute{
					Name:        "is_scheduled",
					Description: `(Optional) Whether this search is to be run on a schedule`,
				},
				resource.Attribute{
					Name:        "is_visible",
					Description: `(Optional) Specifies whether this saved search should be listed in the visible saved search list. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `(Optional) The maximum number of concurrent instances of this search the scheduler is allowed to run. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "realtime_schedule",
					Description: `(Optional) Defaults to 1. Controls the way the scheduler computes the next execution time of a scheduled search. If this value is set to 1, the scheduler bases its determination of the next scheduled search execution time on the current time. If this value is set to 0, the scheduler bases its determination of the next scheduled search on the last search execution time. This is called continuous scheduling. If set to 0, the scheduler never skips scheduled execution periods. However, the execution of the saved search might fall behind depending on the scheduler load. Use continuous scheduling whenever you enable the summary index option.`,
				},
				resource.Attribute{
					Name:        "request_ui_dispatch_app",
					Description: `(Optional) Specifies a field used by Splunk Web to denote the app this search should be dispatched in.`,
				},
				resource.Attribute{
					Name:        "request_ui_dispatch_view",
					Description: `(Optional) Specifies a field used by Splunk Web to denote the view this search should be displayed in.`,
				},
				resource.Attribute{
					Name:        "restart_on_searchpeer_add",
					Description: `(Optional) Specifies whether to restart a real-time search managed by the scheduler when a search peer becomes available for this saved search. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "run_on_startup",
					Description: `(Optional) Indicates whether this search runs at startup. If it does not run on startup, it runs at the next scheduled time. Defaults to 0. Set to 1 for scheduled searches that populate lookup tables.`,
				},
				resource.Attribute{
					Name:        "schedule_priority",
					Description: `(Optional) Raises the scheduling priority of the named search. Defaults to Default`,
				},
				resource.Attribute{
					Name:        "schedule_window",
					Description: `(Optional) Time window (in minutes) during which the search has lower priority. Defaults to 0. The scheduler can give higher priority to more critical searches during this window. The window must be smaller than the search period.Set to auto to let the scheduler determine the optimal window value automatically. Requires the edit_search_schedule_window capability to override auto.`,
				},
				resource.Attribute{
					Name:        "vsid",
					Description: `(Optional) Defines the viewstate id associated with the UI view listed in 'displayview'.`,
				},
				resource.Attribute{
					Name:        "workload_pool",
					Description: `(Optional) Specifies the new workload pool where the existing running search will be placed.` + "`" + ``,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The app/user context that is the namespace for the resource ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the saved search resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the saved search resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunk_sh_indexes_manager",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Create indexes on Splunk Cloud instances. [BETA]

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the index to create.`,
				},
				resource.Attribute{
					Name:        "datatype",
					Description: `(Optional) Valid values: (event | metric). Specifies the type of index.`,
				},
				resource.Attribute{
					Name:        "frozen_time_period_in_secs",
					Description: `(Optional) Number of seconds after which indexed data rolls to frozen. Defaults to 94608000 (3 years).Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation.`,
				},
				resource.Attribute{
					Name:        "max_global_raw_data_size_mb",
					Description: `(Optional) The maximum size of an index (in MB). If an index grows larger than the maximum size, the oldest data is frozen. Defaults to 100 MB. ## Attribute Reference In addition to all arguments above, This resource block exports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the splunk_sh_indexes_manager resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the splunk_sh_indexes_manager resource`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"splunk_acl":                         0,
		"splunk_admin_saml_groups":           1,
		"splunk_apps_local":                  2,
		"splunk_authentication_users":        3,
		"splunk_authorization_roles":         4,
		"splunk_configs_conf":                5,
		"splunk_data_ui_views":               6,
		"splunk_generic_acl":                 7,
		"splunk_global_http_event_collector": 8,
		"splunk_indexes":                     9,
		"splunk_inputs_http_event_collector": 10,
		"splunk_inputs_monitor":              11,
		"splunk_inputs_script":               12,
		"splunk_inputs_tcp_cooked":           13,
		"splunk_inputs_tcp_raw":              14,
		"splunk_inputs_tcp_splunktcptoken":   15,
		"splunk_inputs_tcp_ssl":              16,
		"splunk_inputs_udp":                  17,
		"splunk_outputs_tcp_default":         18,
		"splunk_outputs_tcp_group":           19,
		"splunk_outputs_tcp_server":          20,
		"splunk_outputs_tcp_syslog":          21,
		"splunk_saved_searches":              22,
		"splunk_sh_indexes_manager":          23,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
