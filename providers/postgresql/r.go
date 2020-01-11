package postgresql

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "postgresql_database",
			Category:         "Resources",
			ShortDescription: `Creates and manages a database on a PostgreSQL server.`,
			Description:      ``,
			Keywords: []string{
				"database",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database. Must be unique on the PostgreSQL server instance where it is configured.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) The role name of the user who will own the database, or ` + "`" + `DEFAULT` + "`" + ` to use the default (namely, the user executing the command). To create a database owned by another role or to change the owner of an existing database, you must be a direct or indirect member of the specified role, or the username in the provider is a superuser.`,
				},
				resource.Attribute{
					Name:        "tablespace_name",
					Description: `(Optional) The name of the tablespace that will be associated with the database, or ` + "`" + `DEFAULT` + "`" + ` to use the template database's tablespace. This tablespace will be the default tablespace used for objects created in this database.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional) How many concurrent connections can be established to this database. ` + "`" + `-1` + "`" + ` (the default) means no limit.`,
				},
				resource.Attribute{
					Name:        "allow_connections",
					Description: `(Optional) If ` + "`" + `false` + "`" + ` then no one can connect to this database. The default is ` + "`" + `true` + "`" + `, allowing connections (except as restricted by other mechanisms, such as ` + "`" + `GRANT` + "`" + ` or ` + "`" + `REVOKE CONNECT` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "is_template",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, then this database can be cloned by any user with ` + "`" + `CREATEDB` + "`" + ` privileges; if ` + "`" + `false` + "`" + ` (the default), then only superusers or the owner of the database can clone it.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) The name of the template database from which to create the database, or ` + "`" + `DEFAULT` + "`" + ` to use the default template (` + "`" + `template0` + "`" + `). NOTE: the default in Terraform is ` + "`" + `template0` + "`" + `, not ` + "`" + `template1` + "`" + `. Changing this value will force the creation of a new resource as this value can only be changed when a database is created.`,
				},
				resource.Attribute{
					Name:        "encoding",
					Description: `(Optional) Character set encoding to use in the database. Specify a string constant (e.g. ` + "`" + `UTF8` + "`" + ` or ` + "`" + `SQL_ASCII` + "`" + `), or an integer encoding number. If unset or set to an empty string the default encoding is set to ` + "`" + `UTF8` + "`" + `. If set to ` + "`" + `DEFAULT` + "`" + ` Terraform will use the same encoding as the template database. Changing this value will force the creation of a new resource as this value can only be changed when a database is created.`,
				},
				resource.Attribute{
					Name:        "lc_collate",
					Description: `(Optional) Collation order (` + "`" + `LC_COLLATE` + "`" + `) to use in the database. This affects the sort order applied to strings, e.g. in queries with ` + "`" + `ORDER BY` + "`" + `, as well as the order used in indexes on text columns. If unset or set to an empty string the default collation is set to ` + "`" + `C` + "`" + `. If set to ` + "`" + `DEFAULT` + "`" + ` Terraform will use the same collation order as the specified ` + "`" + `template` + "`" + ` database. Changing this value will force the creation of a new resource as this value can only be changed when a database is created.`,
				},
				resource.Attribute{
					Name:        "lc_ctype",
					Description: `(Optional) Character classification (` + "`" + `LC_CTYPE` + "`" + `) to use in the database. This affects the categorization of characters, e.g. lower, upper and digit. If unset or set to an empty string the default character classification is set to ` + "`" + `C` + "`" + `. If set to ` + "`" + `DEFAULT` + "`" + ` Terraform will use the character classification of the specified ` + "`" + `template` + "`" + ` database. Changing this value will force the creation of a new resource as this value can only be changed when a database is created. ## Import Example ` + "`" + `postgresql_database` + "`" + ` supports importing resources. Supposing the following Terraform: ` + "`" + `` + "`" + `` + "`" + `hcl provider "postgresql" { alias = "admindb" } resource "postgresql_database" "db1" { provider = "postgresql.admindb" name = "testdb1" } ` + "`" + `` + "`" + `` + "`" + ` It is possible to import a ` + "`" + `postgresql_database` + "`" + ` resource with the following command: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import postgresql_database.db1 testdb1 ` + "`" + `` + "`" + `` + "`" + ` Where ` + "`" + `testdb1` + "`" + ` is the name of the database to import and ` + "`" + `postgresql_database.db1` + "`" + ` is the name of the resource whose state will be populated as a result of the command.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "postgresql_default_privileges",
			Category:         "Resources",
			ShortDescription: `Creates and manages default privileges given to a user for a database schema.`,
			Description:      ``,
			Keywords: []string{
				"default",
				"privileges",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the role to which grant default privileges on.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) The database to grant default privileges for this role.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `(Required) The database schema to set default privileges for this role.`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `(Required) The PostgreSQL object type to set the default privileges on (one of: table, sequence).`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Required) The list of privileges to apply as default privileges.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "postgresql_extension",
			Category:         "Resources",
			ShortDescription: `Creates and manages an extension on a PostgreSQL server.`,
			Description:      ``,
			Keywords: []string{
				"extension",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the extension.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `(Optional) Sets the schema of an extension.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Sets the version number of the extension.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Optional) Which database to create the extension on. Defaults to provider database.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "postgresql_grant",
			Category:         "Resources",
			ShortDescription: `Creates and manages privileges given to a user for a database schema.`,
			Description:      ``,
			Keywords: []string{
				"grant",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the role to grant privileges on.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) The database to grant privileges on for this role.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `(Required) The database schema to grant privileges on for this role.`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `(Required) The PostgreSQL object type to grant the privileges on (one of: table, sequence).`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Required) The list of privileges to grant.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "postgresql_role",
			Category:         "Resources",
			ShortDescription: `Creates and manages a role on a PostgreSQL server.`,
			Description:      ``,
			Keywords: []string{
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role. Must be unique on the PostgreSQL server instance where it is configured.`,
				},
				resource.Attribute{
					Name:        "superuser",
					Description: `(Optional) Defines whether the role is a "superuser", and therefore can override all access restrictions within the database. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_database",
					Description: `(Optional) Defines a role's ability to execute ` + "`" + `CREATE DATABASE` + "`" + `. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_role",
					Description: `(Optional) Defines a role's ability to execute ` + "`" + `CREATE ROLE` + "`" + `. A role with this privilege can also alter and drop other roles. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "inherit",
					Description: `(Optional) Defines whether a role "inherits" the privileges of roles it is a member of. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(Optional) Defines whether role is allowed to log in. Roles without this attribute are useful for managing database privileges, but are not users in the usual sense of the word. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `(Optional) Defines whether a role is allowed to initiate streaming replication or put the system in and out of backup mode. Default value is ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bypass_row_level_security",
					Description: `(Optional) Defines whether a role bypasses every row-level security (RLS) policy. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional) If this role can log in, this specifies how many concurrent connections the role can establish. ` + "`" + `-1` + "`" + ` (the default) means no limit.`,
				},
				resource.Attribute{
					Name:        "encrypted_password",
					Description: `(Optional) Defines whether the password is stored encrypted in the system catalogs. Default value is ` + "`" + `true` + "`" + `. NOTE: this value is always set (to the conservative and safe value), but may interfere with the behavior of [PostgreSQL's ` + "`" + `password_encryption` + "`" + ` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Sets the role's password. A password is only of use for roles having the ` + "`" + `login` + "`" + ` attribute set to true.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) Defines list of roles which will be granted to this new role.`,
				},
				resource.Attribute{
					Name:        "search_path",
					Description: `(Optional) Alters the search path of this new role. Note that due to limitations in the implementation, values cannot contain the substring ` + "`" + `", "` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "valid_until",
					Description: `(Optional) Defines the date and time after which the role's password is no longer valid. Established connections past this ` + "`" + `valid_time` + "`" + ` will have to be manually terminated. This value corresponds to a PostgreSQL datetime. If omitted or the magic value ` + "`" + `NULL` + "`" + ` is used, ` + "`" + `valid_until` + "`" + ` will be set to ` + "`" + `infinity` + "`" + `. Default is ` + "`" + `NULL` + "`" + `, therefore ` + "`" + `infinity` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "skip_drop_role",
					Description: `(Optional) When a PostgreSQL ROLE exists in multiple databases and the ROLE is dropped, the [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html) in each of the respective databases must occur before the ROLE can be dropped from the catalog. Set this option to true when there are multiple databases in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership. This is the third and final step taken when removing a ROLE from a database.`,
				},
				resource.Attribute{
					Name:        "skip_reassign_owned",
					Description: `(Optional) When a PostgreSQL ROLE exists in multiple databases and the ROLE is dropped, a [` + "`" + `REASSIGN OWNED` + "`" + `](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in must be executed on each of the respective databases before the ` + "`" + `DROP ROLE` + "`" + ` can be executed to dropped the ROLE from the catalog. This is the first and second steps taken when removing a ROLE from a database (the second step being an implicit [` + "`" + `DROP OWNED` + "`" + `](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)). ## Import Example ` + "`" + `postgresql_role` + "`" + ` supports importing resources. Supposing the following Terraform: ` + "`" + `` + "`" + `` + "`" + `hcl provider "postgresql" { alias = "admindb" } resource "postgresql_role" "replication_role" { provider = "postgresql.admindb" name = "replication_name" } ` + "`" + `` + "`" + `` + "`" + ` It is possible to import a ` + "`" + `postgresql_role` + "`" + ` resource with the following command: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import postgresql_role.replication_role replication_name ` + "`" + `` + "`" + `` + "`" + ` Where ` + "`" + `replication_name` + "`" + ` is the name of the role to import and ` + "`" + `postgresql_role.replication_role` + "`" + ` is the name of the resource whose state will be populated as a result of the command.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "postgresql_schema",
			Category:         "Resources",
			ShortDescription: `Creates and manages a schema within a PostgreSQL database.`,
			Description:      ``,
			Keywords: []string{
				"schema",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the schema. Must be unique in the PostgreSQL database instance where it is configured.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) The ROLE who owns the schema.`,
				},
				resource.Attribute{
					Name:        "if_not_exists",
					Description: `(Optional) When true, use the existing schema if it exists. (Default: true)`,
				},
				resource.Attribute{
					Name:        "drop_cascade",
					Description: `(Optional) When true, will also drop all the objects that are contained in the schema. (Default: false)`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) Can be specified multiple times for each policy. Each policy block supports fields documented below. The ` + "`" + `policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Optional) Should the specified ROLE have CREATE privileges to the specified SCHEMA.`,
				},
				resource.Attribute{
					Name:        "create_with_grant",
					Description: `(Optional) Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The ROLE who is receiving the policy. If this value is empty or not specified it implies the policy is referring to the [` + "`" + `PUBLIC` + "`" + ` role](https://www.postgresql.org/docs/current/static/sql-grant.html).`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) Should the specified ROLE have USAGE privileges to the specified SCHEMA.`,
				},
				resource.Attribute{
					Name:        "usage_with_grant",
					Description: `(Optional) Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"postgresql_database":           0,
		"postgresql_default_privileges": 1,
		"postgresql_extension":          2,
		"postgresql_grant":              3,
		"postgresql_role":               4,
		"postgresql_schema":             5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
