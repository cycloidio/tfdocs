package mysql

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mysql_database",
			Category:         "Resources",
			ShortDescription: `Creates and manages a database on a MySQL server.`,
			Description:      ``,
			Keywords: []string{
				"database",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database. This must be unique within a given MySQL server and may or may not be case-sensitive depending on the operating system on which the MySQL server is running.`,
				},
				resource.Attribute{
					Name:        "default_character_set",
					Description: `(Optional) The default character set to use when a table is created without specifying an explicit character set. Defaults to "utf8".`,
				},
				resource.Attribute{
					Name:        "default_collation",
					Description: `(Optional) The default collation to use when a table is created without specifying an explicit collation. Defaults to ` + "`" + `` + "`" + `utf8_general_ci` + "`" + `` + "`" + `. Each character set has its own set of collations, so changing the character set requires also changing the collation. Note that the defaults for character set and collation above do not respect any defaults set on the MySQL server, so that the configuration can be set appropriately even though Terraform cannot see the server-level defaults. If you wish to use the server's defaults you must consult the server's configuration and then set the ` + "`" + `` + "`" + `default_character_set` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `default_collation` + "`" + `` + "`" + ` to match. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the database.`,
				},
				resource.Attribute{
					Name:        "default_character_set",
					Description: `The default_character_set of the database.`,
				},
				resource.Attribute{
					Name:        "default_collation",
					Description: `The default_collation of the database. ## Import Databases can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mysql_database.example my-example-database ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the database.`,
				},
				resource.Attribute{
					Name:        "default_character_set",
					Description: `The default_character_set of the database.`,
				},
				resource.Attribute{
					Name:        "default_collation",
					Description: `The default_collation of the database. ## Import Databases can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mysql_database.example my-example-database ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mysql_grant",
			Category:         "Resources",
			ShortDescription: `Creates and manages privileges given to a user on a MySQL server`,
			Description:      ``,
			Keywords: []string{
				"grant",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) The name of the user. Conflicts with ` + "`" + `role` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The source host of the user. Defaults to "localhost". Conflicts with ` + "`" + `role` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role to grant ` + "`" + `privileges` + "`" + ` to. Conflicts with ` + "`" + `user` + "`" + ` and ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) The database to grant privileges on.`,
				},
				resource.Attribute{
					Name:        "table",
					Description: `(Optional) Which table to grant ` + "`" + `privileges` + "`" + ` on. Defaults to ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Optional) A list of privileges to grant to the user. Refer to a list of privileges (such as [here](https://dev.mysql.com/doc/refman/5.5/en/grant.html)) for applicable privileges. Conflicts with ` + "`" + `roles` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) A list of rols to grant to the user. Conflicts with ` + "`" + `privileges` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_option",
					Description: `(Optional) An TLS-Option for the ` + "`" + `GRANT` + "`" + ` statement. The value is suffixed to ` + "`" + `REQUIRE` + "`" + `. A value of 'SSL' will generate a ` + "`" + `GRANT ... REQUIRE SSL` + "`" + ` statement. See the [MYSQL ` + "`" + `GRANT` + "`" + ` documentation](https://dev.mysql.com/doc/refman/5.7/en/grant.html) for more. Ignored if MySQL version is under 5.7.0.`,
				},
				resource.Attribute{
					Name:        "grant",
					Description: `(Optional) Whether to also give the user privileges to grant the same privileges to other users. ## Attributes Reference No further attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mysql_role",
			Category:         "Resources",
			ShortDescription: `Creates and manages a role on a MySQL server.`,
			Description:      ``,
			Keywords: []string{
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role. ## Attributes Reference No further attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mysql_user",
			Category:         "Resources",
			ShortDescription: `Creates and manages a user on a MySQL server.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The source host of the user. Defaults to "localhost".`,
				},
				resource.Attribute{
					Name:        "plaintext_password",
					Description: `(Optional) The password for the user. This must be provided in plain text, so the data source for it must be secured. An _unsalted_ hash of the provided password is stored in state. Conflicts with ` + "`" + `auth_plugin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Deprecated alias of ` + "`" + `plaintext_password` + "`" + `, whose value is`,
				},
				resource.Attribute{
					Name:        "auth_plugin",
					Description: `(Optional) Use an [authentication plugin][ref-auth-plugins] to authenticate the user instead of using password authentication. Description of the fields allowed in the block below. Conflicts with ` + "`" + `password` + "`" + ` and ` + "`" + `plaintext_password` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_option",
					Description: `(Optional) An TLS-Option for the ` + "`" + `CREATE USER` + "`" + ` or ` + "`" + `ALTER USER` + "`" + ` statement. The value is suffixed to ` + "`" + `REQUIRE` + "`" + `. A value of 'SSL' will generate a ` + "`" + `CREATE USER ... REQUIRE SSL` + "`" + ` statement. See the [MYSQL ` + "`" + `CREATE USER` + "`" + ` documentation](https://dev.mysql.com/doc/refman/5.7/en/create-user.html) for more. Ignored if MySQL version is under 5.7.0. [ref-auth-plugins]: https://dev.mysql.com/doc/refman/5.7/en/authentication-plugins.html The ` + "`" + `auth_plugin` + "`" + ` value supports:`,
				},
				resource.Attribute{
					Name:        "AWSAuthenticationPlugin",
					Description: `Allows the use of IAM authentication with [Amazon Aurora][ref-amazon-aurora]. For more details on how to use IAM auth with Aurora, see [here][ref-aurora-using-iam]. [ref-amazon-aurora]: https://aws.amazon.com/rds/aurora/ [ref-aurora-using-iam]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html#UsingWithRDS.IAMDBAuth.Creating`,
				},
				resource.Attribute{
					Name:        "mysql_no_login",
					Description: `Uses the MySQL No-Login Authentication Plugin. The No-Login Authentication Plugin must be active in MySQL. For more information, see [here][ref-mysql-no-login]. [ref-mysql-no-login]: https://dev.mysql.com/doc/refman/5.7/en/no-login-pluggable-authentication.html ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the user created, composed as "username@host".`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host where the user was created. ## Attributes Reference No further attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the user created, composed as "username@host".`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host where the user was created. ## Attributes Reference No further attributes are exported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mysql_user_password",
			Category:         "Resources",
			ShortDescription: `Creates and manages the password for a user on a MySQL server.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `(Required) The IAM user to associate with this access key.`,
				},
				resource.Attribute{
					Name:        "pgp_key",
					Description: `(Required) Either a base-64 encoded PGP public key, or a keybase username in the form ` + "`" + `keybase:some_person_that_exists` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The source host of the user. Defaults to ` + "`" + `localhost` + "`" + `. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the password`,
				},
				resource.Attribute{
					Name:        "encrypted_password",
					Description: `The encrypted password, base64 encoded. ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the password`,
				},
				resource.Attribute{
					Name:        "encrypted_password",
					Description: `The encrypted password, base64 encoded. ~>`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"mysql_database":      0,
		"mysql_grant":         1,
		"mysql_role":          2,
		"mysql_user":          3,
		"mysql_user_password": 4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
