package ovh

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_region",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of a region associated with a public cloud project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the region associated with the public cloud project. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project concatenated with the name of the region. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `(Deprecated) Use ` + "`" + `continent_code` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "datacenterLocation",
					Description: `(Deprecated) Use ` + "`" + `datacenter_location` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `(Deprecated) Use ` + "`" + `continent_code` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "datacenterLocation",
					Description: `(Deprecated) Use ` + "`" + `datacenter_location` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_regions",
			Category:         "Data Sources",
			ShortDescription: `Get the list of regions associated with a public cloud project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_domain_zone",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of a domain zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the domain zone. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the domain zone name. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update date of the DNS zone`,
				},
				resource.Attribute{
					Name:        "has_dns_anycast",
					Description: `hasDnsAnycast flag of the DNS zone`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Name servers that host the DNS zone`,
				},
				resource.Attribute{
					Name:        "dnssec_supported",
					Description: `Is DNSSEC supported by this zone`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update date of the DNS zone`,
				},
				resource.Attribute{
					Name:        "has_dns_anycast",
					Description: `hasDnsAnycast flag of the DNS zone`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Name servers that host the DNS zone`,
				},
				resource.Attribute{
					Name:        "dnssec_supported",
					Description: `Is DNSSEC supported by this zone`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of an IP Load Balancing product.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6",
					Description: `The IPV6 associated to your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The IPV4 associated to your IP load balancing`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Location where your service is. This takes an array of values.`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `The offer of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ip_loadbalancing",
					Description: `Your IP load balancing`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of your IP. Can take any of the following value: "blacklisted", "deleted", "free", "ok", "quarantined", "suspended"`,
				},
				resource.Attribute{
					Name:        "vrack_eligibility",
					Description: `Vrack eligibility. Takes a boolean value.`,
				},
				resource.Attribute{
					Name:        "vrack_name",
					Description: `Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product`,
				},
				resource.Attribute{
					Name:        "ssl_configuration",
					Description: `Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Can take any of the following value: "intermediate", "modern"`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `the name displayed in ManagerV6 for your iplb (max 50 chars) ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the service_name of your IP load balancing In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "metrics_token",
					Description: `The metrics token associated with your IP load balancing This attribute is sensitive.`,
				},
				resource.Attribute{
					Name:        "orderable_zone",
					Description: `Available additional zone for your Load Balancer`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The zone three letter code`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `The billing planCode for this zone`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metrics_token",
					Description: `The metrics token associated with your IP load balancing This attribute is sensitive.`,
				},
				resource.Attribute{
					Name:        "orderable_zone",
					Description: `Available additional zone for your Load Balancer`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The zone three letter code`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `The billing planCode for this zone`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_paymentmean_bankaccount",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of an ovh bank account payment mean`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regexp",
					Description: `(Optional) a regexp used to filter bank accounts on their ` + "`" + `description` + "`" + ` attributes.`,
				},
				resource.Attribute{
					Name:        "use_default",
					Description: `(Optional) Retrieve bank account marked as default payment mean.`,
				},
				resource.Attribute{
					Name:        "use_oldest",
					Description: `(Optional) Retrieve oldest bank account. project.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Filter bank accounts on their ` + "`" + `state` + "`" + ` attribute. Can be "blockedForIncidents", "valid", "pendingValidation" ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the bank account payment mean`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the bank account`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved bank account is marked as the default payment mean`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the bank account`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved bank account is marked as the default payment mean`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_paymentmean_creditcard",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of an ovh credit card payment mean`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regexp",
					Description: `(Optional) a regexp used to filter credit cards on their ` + "`" + `description` + "`" + ` attributes.`,
				},
				resource.Attribute{
					Name:        "use_default",
					Description: `(Optional) Retrieve credit card marked as default payment mean.`,
				},
				resource.Attribute{
					Name:        "use_last_to_expire",
					Description: `(Optional) Retrieve the credit card that will be the last to expire according to its expiration date.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) Filter credit cards on their ` + "`" + `state` + "`" + ` attribute. Can be "expired", "valid", "tooManyFailures" ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the credit card payment mean`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved credit card is marked as the default payment mean`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved credit card is marked as the default payment mean`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_publiccloud_region",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of a region associated with a public cloud project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The name of the region associated with the public cloud project. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project concatenated with the name of the region. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `(Deprecated) Use ` + "`" + `continent_code` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "datacenterLocation",
					Description: `(Deprecated) Use ` + "`" + `datacenter_location` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `(Deprecated) Use ` + "`" + `continent_code` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "datacenterLocation",
					Description: `(Deprecated) Use ` + "`" + `datacenter_location` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_publiccloud_regions",
			Category:         "Data Sources",
			ShortDescription: `Get the list of regions associated with a public cloud project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ovh_cloud_region":               0,
		"ovh_cloud_regions":              1,
		"ovh_domain_zone":                2,
		"ovh_iploadbalancing":            3,
		"ovh_me_paymentmean_bankaccount": 4,
		"ovh_me_paymentmean_creditcard":  5,
		"ovh_publiccloud_region":         6,
		"ovh_publiccloud_regions":        7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
