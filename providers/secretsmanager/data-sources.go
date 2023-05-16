package secretsmanager

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_address",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `A list containing address information: - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `A list containing address information: - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_bank_account",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "bank_account",
					Description: `A list containing bank account information: - ` + "`" + `account_type` + "`" + ` - Account type - ` + "`" + `other_type` + "`" + ` - Other type (if specified) - ` + "`" + `routing_number` + "`" + ` - City - ` + "`" + `account_number` + "`" + ` - State`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Account login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Account password.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Account URL.`,
				},
				resource.Attribute{
					Name:        "card_ref",
					Description: `A list containing card reference information. - ` + "`" + `uid` + "`" + ` - Card reference UID - ` + "`" + `payment_card` + "`" + ` - A list containing payment card information: - ` + "`" + `card_number` + "`" + ` - Card number - ` + "`" + `card_expiration_date` + "`" + ` - Card expiration date - ` + "`" + `card_security_code` + "`" + ` - Card security code - ` + "`" + `cardholder_name` + "`" + ` - Cardholder name - ` + "`" + `pin_code` + "`" + ` - PIN code`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
				resource.Attribute{
					Name:        "totp",
					Description: `A list containing Time-based One-time password information: - ` + "`" + `url` + "`" + ` - TOTP URL - ` + "`" + `token` + "`" + ` - Current TOTP password - ` + "`" + `ttl` + "`" + ` - Time to live in seconds for current token`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "bank_account",
					Description: `A list containing bank account information: - ` + "`" + `account_type` + "`" + ` - Account type - ` + "`" + `other_type` + "`" + ` - Other type (if specified) - ` + "`" + `routing_number` + "`" + ` - City - ` + "`" + `account_number` + "`" + ` - State`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Account login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Account password.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Account URL.`,
				},
				resource.Attribute{
					Name:        "card_ref",
					Description: `A list containing card reference information. - ` + "`" + `uid` + "`" + ` - Card reference UID - ` + "`" + `payment_card` + "`" + ` - A list containing payment card information: - ` + "`" + `card_number` + "`" + ` - Card number - ` + "`" + `card_expiration_date` + "`" + ` - Card expiration date - ` + "`" + `card_security_code` + "`" + ` - Card security code - ` + "`" + `cardholder_name` + "`" + ` - Cardholder name - ` + "`" + `pin_code` + "`" + ` - PIN code`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
				resource.Attribute{
					Name:        "totp",
					Description: `A list containing Time-based One-time password information: - ` + "`" + `url` + "`" + ` - TOTP URL - ` + "`" + `token` + "`" + ` - Current TOTP password - ` + "`" + `ttl` + "`" + ` - Time to live in seconds for current token`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_bank_card",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "payment_card",
					Description: `A list containing payment card information: - ` + "`" + `card_number` + "`" + ` - Card number - ` + "`" + `card_expiration_date` + "`" + ` - Card expiration date - ` + "`" + `card_security_code` + "`" + ` - Card security code`,
				},
				resource.Attribute{
					Name:        "cardholder_name",
					Description: `Cardholder name.`,
				},
				resource.Attribute{
					Name:        "pin_code",
					Description: `PIN code.`,
				},
				resource.Attribute{
					Name:        "address_ref",
					Description: `A list containing address information: - ` + "`" + `uid` + "`" + ` - The address reference record UID - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "payment_card",
					Description: `A list containing payment card information: - ` + "`" + `card_number` + "`" + ` - Card number - ` + "`" + `card_expiration_date` + "`" + ` - Card expiration date - ` + "`" + `card_security_code` + "`" + ` - Card security code`,
				},
				resource.Attribute{
					Name:        "cardholder_name",
					Description: `Cardholder name.`,
				},
				resource.Attribute{
					Name:        "pin_code",
					Description: `PIN code.`,
				},
				resource.Attribute{
					Name:        "address_ref",
					Description: `A list containing address information: - ` + "`" + `uid` + "`" + ` - The address reference record UID - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_birth_certificate",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "birth_date",
					Description: `Date of birth.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "birth_date",
					Description: `Date of birth.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_contact",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `Company name.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Contact's e-mail.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `A list containing phone information: - ` + "`" + `region` + "`" + ` - 2 letter country code (ISO 3166-1 alpha-2) - ` + "`" + `number` + "`" + ` - Phone number - ` + "`" + `ext` + "`" + ` - Phone extension - ` + "`" + `type` + "`" + ` - Phone type: Mobile, Home or Work`,
				},
				resource.Attribute{
					Name:        "address_ref",
					Description: `A list containing address information: - ` + "`" + `uid` + "`" + ` - The address reference record UID - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `Company name.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Contact's e-mail.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `A list containing phone information: - ` + "`" + `region` + "`" + ` - 2 letter country code (ISO 3166-1 alpha-2) - ` + "`" + `number` + "`" + ` - Phone number - ` + "`" + `ext` + "`" + ` - Phone extension - ` + "`" + `type` + "`" + ` - Phone type: Mobile, Home or Work`,
				},
				resource.Attribute{
					Name:        "address_ref",
					Description: `A list containing address information: - ` + "`" + `uid` + "`" + ` - The address reference record UID - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_database_credentials",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Login name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Login password.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A list containing hostname and port information: - ` + "`" + `host_name` + "`" + ` - Database server hostname - ` + "`" + `port` + "`" + ` - Port number`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Login name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Login password.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A list containing hostname and port information: - ` + "`" + `host_name` + "`" + ` - Database server hostname - ` + "`" + `port` + "`" + ` - Port number`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_driver_license",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "driver_license_number",
					Description: `Driver's License Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "birth_date",
					Description: `Date of birth.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Date of expiration.`,
				},
				resource.Attribute{
					Name:        "address_ref",
					Description: `A list containing address information: - ` + "`" + `uid` + "`" + ` - The address reference record UID - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "driver_license_number",
					Description: `Driver's License Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "birth_date",
					Description: `Date of birth.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Date of expiration.`,
				},
				resource.Attribute{
					Name:        "address_ref",
					Description: `A list containing address information: - ` + "`" + `uid` + "`" + ` - The address reference record UID - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_encrypted_notes",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `Encrypted note.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Date of the note.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `Encrypted note.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Date of the note.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_field",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path to a field of a secret stored in existing record in Keeper Vault. Provide full path to the field - regular fields are accessible by field type and custom fields are accessible by field label: ex. ` + "`" + `<record UID>/field/login` + "`" + `, ex. ` + "`" + `<record UID>/custom_field/custom1` + "`" + `, ex. ` + "`" + `<record UID>/custom_field/custom2` + "`" + `. Use ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Optional) The title of a secret stored in existing record in Keeper Vault. If there's a need to find record by title - use ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the selected field.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `The value of the selected field.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_file",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_general",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The secret login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The secret url.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
				resource.Attribute{
					Name:        "totp",
					Description: `A list containing Time-based One-time password information: - ` + "`" + `url` + "`" + ` - TOTP URL - ` + "`" + `token` + "`" + ` - Current TOTP password - ` + "`" + `ttl` + "`" + ` - Time to live in seconds for current token`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The secret login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The secret url.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
				resource.Attribute{
					Name:        "totp",
					Description: `A list containing Time-based One-time password information: - ` + "`" + `url` + "`" + ` - TOTP URL - ` + "`" + `token` + "`" + ` - Current TOTP password - ` + "`" + `ttl` + "`" + ` - Time to live in seconds for current token`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_health_insurance",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `Account Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Account login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Account password.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Account URL.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `Account Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Account login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Account password.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Account URL.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_login",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The secret login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The secret url.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
				resource.Attribute{
					Name:        "totp",
					Description: `A list containing Time-based One-time password information: - ` + "`" + `url` + "`" + ` - TOTP URL - ` + "`" + `token` + "`" + ` - Current TOTP password - ` + "`" + `ttl` + "`" + ` - Time to live in seconds for current token`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The secret login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The secret url.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
				resource.Attribute{
					Name:        "totp",
					Description: `A list containing Time-based One-time password information: - ` + "`" + `url` + "`" + ` - TOTP URL - ` + "`" + `token` + "`" + ` - Current TOTP password - ` + "`" + `ttl` + "`" + ` - Time to live in seconds for current token`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_membership",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `Account Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `Account Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_passport",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "passport_number",
					Description: `Passport Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "birth_date",
					Description: `Date of birth.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Expiration date.`,
				},
				resource.Attribute{
					Name:        "date_issued",
					Description: `Date issued.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "address_ref",
					Description: `A list containing address information: - ` + "`" + `uid` + "`" + ` - The address reference record UID - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "passport_number",
					Description: `Passport Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "birth_date",
					Description: `Date of birth.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Expiration date.`,
				},
				resource.Attribute{
					Name:        "date_issued",
					Description: `Date issued.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "address_ref",
					Description: `A list containing address information: - ` + "`" + `uid` + "`" + ` - The address reference record UID - ` + "`" + `street1` + "`" + ` - Street line 1 - ` + "`" + `street2` + "`" + ` - Street line 2 - ` + "`" + `city` + "`" + ` - City - ` + "`" + `state` + "`" + ` - State - ` + "`" + `zip` + "`" + ` - Zip - ` + "`" + `country` + "`" + ` - Country`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_photo",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_server_credentials",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A list containing hostname and port information: - ` + "`" + `host_name` + "`" + ` - Database server hostname - ` + "`" + `port` + "`" + ` - Port number`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The secret login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A list containing hostname and port information: - ` + "`" + `host_name` + "`" + ` - Database server hostname - ` + "`" + `port` + "`" + ` - Port number`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The secret login.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The secret password.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_software_license",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "license_number",
					Description: `License Number.`,
				},
				resource.Attribute{
					Name:        "activation_date",
					Description: `Date of activation.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Date of expiration.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "license_number",
					Description: `License Number.`,
				},
				resource.Attribute{
					Name:        "activation_date",
					Description: `Date of activation.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Date of expiration.`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_ssh_keys",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The secret login.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `A list containing public and private key pair information: - ` + "`" + `public_key` + "`" + ` - The public key. - ` + "`" + `private_key` + "`" + ` - The private key.`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `The passphrase to unlock the key.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A list containing hostname and port information: - ` + "`" + `host_name` + "`" + ` - Database server hostname - ` + "`" + `port` + "`" + ` - Port number`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The secret login.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `A list containing public and private key pair information: - ` + "`" + `public_key` + "`" + ` - The public key. - ` + "`" + `private_key` + "`" + ` - The private key.`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `The passphrase to unlock the key.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A list containing hostname and port information: - ` + "`" + `host_name` + "`" + ` - Database server hostname - ` + "`" + `port` + "`" + ` - Port number`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "secretsmanager_ssn_card",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The UID of existing record in Keeper Vault. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "identity_number",
					Description: `Identity Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Record title.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Record notes.`,
				},
				resource.Attribute{
					Name:        "identity_number",
					Description: `Identity Number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A list containing name information: - ` + "`" + `first` + "`" + ` - First name - ` + "`" + `middle` + "`" + ` - Middle name - ` + "`" + `last` + "`" + ` - Last name`,
				},
				resource.Attribute{
					Name:        "file_ref",
					Description: `A list containing file reference information: - ` + "`" + `uid` + "`" + ` - File UID - ` + "`" + `title` + "`" + ` - File title - ` + "`" + `name` + "`" + ` - File name - ` + "`" + `type` + "`" + ` - File content type - ` + "`" + `size` + "`" + ` - File size - ` + "`" + `last_modified` + "`" + ` - File last modification timestamp - ` + "`" + `content_base64` + "`" + ` - File content base64 encoded`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"secretsmanager_address":              0,
		"secretsmanager_bank_account":         1,
		"secretsmanager_bank_card":            2,
		"secretsmanager_birth_certificate":    3,
		"secretsmanager_contact":              4,
		"secretsmanager_database_credentials": 5,
		"secretsmanager_driver_license":       6,
		"secretsmanager_encrypted_notes":      7,
		"secretsmanager_field":                8,
		"secretsmanager_file":                 9,
		"secretsmanager_general":              10,
		"secretsmanager_health_insurance":     11,
		"secretsmanager_login":                12,
		"secretsmanager_membership":           13,
		"secretsmanager_passport":             14,
		"secretsmanager_photo":                15,
		"secretsmanager_server_credentials":   16,
		"secretsmanager_software_license":     17,
		"secretsmanager_ssh_keys":             18,
		"secretsmanager_ssn_card":             19,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
