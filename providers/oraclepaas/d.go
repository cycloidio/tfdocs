package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "oraclepaas_database_service_instance",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the configuration of an Oracle Database Cloud Service instance on the Oracle Cloud Platform.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Database Service Instance ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "apex_url",
					Description: `The URL to use to connect to Oracle Application Express on the service instance.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Name of the availability domain within the region where the Oracle Database Cloud Service instance is provisioned.`,
				},
				resource.Attribute{
					Name:        "character_set",
					Description: `The database character set of the database.`,
				},
				resource.Attribute{
					Name:        "cloud_storage_container",
					Description: `The Oracle Storage Cloud container for backups.`,
				},
				resource.Attribute{
					Name:        "compute_site_name",
					Description: `The Oracle Cloud location housing the service instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the service instance.`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `The software edition of the service instance.`,
				},
				resource.Attribute{
					Name:        "enterprise_manager_url",
					Description: `The URL to use to connect to Enterprise Manager on the service instance.`,
				},
				resource.Attribute{
					Name:        "failover_database",
					Description: `Indicates whether the service instance hosts an Oracle Data Guard configuration.`,
				},
				resource.Attribute{
					Name:        "glassfish_url",
					Description: `The URL to use to connect to the Oracle GlassFish Server Administration Console on the service instance.`,
				},
				resource.Attribute{
					Name:        "hybrid_disaster_recovery_ip",
					Description: `Data Guard Role of the on-premise instance in Oracle Hybrid Disaster Recovery configuration.`,
				},
				resource.Attribute{
					Name:        "identity_domain",
					Description: `The identity domain housing the service instance.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `The three-part name of an IP network to which the service instance is added.`,
				},
				resource.Attribute{
					Name:        "ip_reservations",
					Description: `Groups one or more IP reservations in use on this service instance.`,
				},
				resource.Attribute{
					Name:        "bring_your_own_license",
					Description: `Indicates whether service instance was provisioned with the 'Bring Your Own License' option.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `The service level of the service instance.`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `The listener port for Oracle Net Services (SQL`,
				},
				resource.Attribute{
					Name:        "monitor_url",
					Description: `The URL to use to connect to Oracle DBaaS Monitor on the service instance.`,
				},
				resource.Attribute{
					Name:        "national_character_set",
					Description: `The national character set of the database.`,
				},
				resource.Attribute{
					Name:        "pluggable_database_name",
					Description: `The name of the default PDB (pluggable database) created when the service instance was created.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Location where the service instance is provisioned.`,
				},
				resource.Attribute{
					Name:        "shape",
					Description: `The Oracle Compute Cloud shape of the service instance.`,
				},
				resource.Attribute{
					Name:        "high_performance_storage",
					Description: `Indicates whether the service instance was provisioned with high performance storage.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The REST endpoint URI of the service instance.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Oracle Database version on the service instance.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "apex_url",
					Description: `The URL to use to connect to Oracle Application Express on the service instance.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Name of the availability domain within the region where the Oracle Database Cloud Service instance is provisioned.`,
				},
				resource.Attribute{
					Name:        "character_set",
					Description: `The database character set of the database.`,
				},
				resource.Attribute{
					Name:        "cloud_storage_container",
					Description: `The Oracle Storage Cloud container for backups.`,
				},
				resource.Attribute{
					Name:        "compute_site_name",
					Description: `The Oracle Cloud location housing the service instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the service instance.`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `The software edition of the service instance.`,
				},
				resource.Attribute{
					Name:        "enterprise_manager_url",
					Description: `The URL to use to connect to Enterprise Manager on the service instance.`,
				},
				resource.Attribute{
					Name:        "failover_database",
					Description: `Indicates whether the service instance hosts an Oracle Data Guard configuration.`,
				},
				resource.Attribute{
					Name:        "glassfish_url",
					Description: `The URL to use to connect to the Oracle GlassFish Server Administration Console on the service instance.`,
				},
				resource.Attribute{
					Name:        "hybrid_disaster_recovery_ip",
					Description: `Data Guard Role of the on-premise instance in Oracle Hybrid Disaster Recovery configuration.`,
				},
				resource.Attribute{
					Name:        "identity_domain",
					Description: `The identity domain housing the service instance.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `The three-part name of an IP network to which the service instance is added.`,
				},
				resource.Attribute{
					Name:        "ip_reservations",
					Description: `Groups one or more IP reservations in use on this service instance.`,
				},
				resource.Attribute{
					Name:        "bring_your_own_license",
					Description: `Indicates whether service instance was provisioned with the 'Bring Your Own License' option.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `The service level of the service instance.`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `The listener port for Oracle Net Services (SQL`,
				},
				resource.Attribute{
					Name:        "monitor_url",
					Description: `The URL to use to connect to Oracle DBaaS Monitor on the service instance.`,
				},
				resource.Attribute{
					Name:        "national_character_set",
					Description: `The national character set of the database.`,
				},
				resource.Attribute{
					Name:        "pluggable_database_name",
					Description: `The name of the default PDB (pluggable database) created when the service instance was created.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Location where the service instance is provisioned.`,
				},
				resource.Attribute{
					Name:        "shape",
					Description: `The Oracle Compute Cloud shape of the service instance.`,
				},
				resource.Attribute{
					Name:        "high_performance_storage",
					Description: `Indicates whether the service instance was provisioned with high performance storage.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The REST endpoint URI of the service instance.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Oracle Database version on the service instance.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]Resource{

		"oraclepaas_database_service_instance": 0,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
