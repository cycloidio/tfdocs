package ec

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ec_deployment",
			Category:         "Resources",
			ShortDescription: `Provides an Elastic Cloud deployment resource, which allows deployments to be created, updated, and deleted.`,
			Description: `

Provides an Elastic Cloud deployment resource, which allows deployments to be created, updated, and deleted.

~> **Note on traffic filters** If you use ` + "`" + `traffic_filter` + "`" + ` on an ` + "`" + `ec_deployment` + "`" + `, Terraform will manage the full set of traffic rules for the deployment, and treat additional traffic filters as drift. For this reason, ` + "`" + `traffic_filter` + "`" + ` cannot be mixed with the ` + "`" + `ec_deployment_traffic_filter_association` + "`" + ` resource for a given deployment.

~> **Note on Elastic Stack versions** Using a version prior to ` + "`" + `6.6.0` + "`" + ` is not supported.

~> **Note on regions and deployment templates** Before you start, you might want to read about [Elastic Cloud deployments](https://www.elastic.co/guide/en/cloud/current/ec-create-deployment.html) and check the [full list](https://www.elastic.co/guide/en/cloud/current/ec-regions-templates-instances.html) of regions and deployment templates available in Elasticsearch Service (ESS).

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_deployment_elasticsearch_keystore",
			Category:         "Resources",
			ShortDescription: `Provides an Elastic Cloud Deployment Elasticsearch keystore resource, which allows you to create and update Elasticsearch keystore settings. Elasticsearch keystore settings can be created and updated through this resource, **each resource represents a single Elasticsearch Keystore setting**. After adding a key and its secret value to the keystore, you can use the key in place of the secret value when you configure sensitive settings. ~> **Note on Elastic keystore settings** This resource offers weaker consistency guarantees and will not detect and update keystore setting values that have been modified outside of the scope of Terraform, usually referred to as _drift_. For example, consider the following scenario: 1. A keystore setting is created using this resource. 2. The keystore setting's value is modified to a different value using the Elasticsearch Service API. 3. Running <code>terraform apply</code> fails to detect the changes and does not update the keystore setting to the value defined in the terraform configuration. To force the keystore setting to the value it is configured to hold, you may want to taint the resource and force its recreation. Before you create Elasticsearch keystore settings, check the [official Elasticsearch keystore documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/elasticsearch-keystore.html) and the [Elastic Cloud specific documentation](https://www.elastic.co/guide/en/cloud/current/ec-configuring-keystore.html).`,
			Description: `

Provides an Elastic Cloud Deployment Elasticsearch keystore resource, which allows you to create and update Elasticsearch keystore settings.

  Elasticsearch keystore settings can be created and updated through this resource, **each resource represents a single Elasticsearch Keystore setting**. After adding a key and its secret value to the keystore, you can use the key in place of the secret value when you configure sensitive settings.

  ~> **Note on Elastic keystore settings** This resource offers weaker consistency guarantees and will not detect and update keystore setting values that have been modified outside of the scope of Terraform, usually referred to as _drift_. For example, consider the following scenario:
    1. A keystore setting is created using this resource.
    2. The keystore setting's value is modified to a different value using the Elasticsearch Service API.
    3. Running <code>terraform apply</code> fails to detect the changes and does not update the keystore setting to the value defined in the terraform configuration.
    To force the keystore setting to the value it is configured to hold, you may want to taint the resource and force its recreation.

  Before you create Elasticsearch keystore settings, check the [official Elasticsearch keystore documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/elasticsearch-keystore.html) and the [Elastic Cloud specific documentation](https://www.elastic.co/guide/en/cloud/current/ec-configuring-keystore.html).

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_deployment_extension",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `


  Provides an Elastic Cloud extension resource, which allows extensions to be created, updated, and deleted.

  Extensions allow users of Elastic Cloud to use custom plugins, scripts, or dictionaries to enhance the core functionality of Elasticsearch. Before you install an extension, be sure to check out the supported and official [Elasticsearch plugins](https://www.elastic.co/guide/en/elasticsearch/plugins/current/index.html) already available.

  **Tip :** If you experience timeouts when uploading an extension through a slow network, you might need to increase the [timeout setting](https://registry.terraform.io/providers/elastic/ec/latest/docs#timeout).


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_deployment_traffic_filter",
			Category:         "Resources",
			ShortDescription: `Provides an Elastic Cloud traffic filter resource, which allows traffic filter rules to be created, updated, and deleted. Traffic filter rules are used to limit inbound traffic to deployment resources. ~> **Note on traffic filters** If you use traffic_filter on an ec_deployment, Terraform will manage the full set of traffic rules for the deployment, and treat additional traffic filters as drift. For this reason, traffic_filter cannot be mixed with the ec_deployment_traffic_filter_association resource for a given deployment.`,
			Description: `

Provides an Elastic Cloud traffic filter resource, which allows traffic filter rules to be created, updated, and deleted. Traffic filter rules are used to limit inbound traffic to deployment resources.

  ~> **Note on traffic filters** If you use traffic_filter on an ec_deployment, Terraform will manage the full set of traffic rules for the deployment, and treat additional traffic filters as drift. For this reason, traffic_filter cannot be mixed with the ec_deployment_traffic_filter_association resource for a given deployment.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_deployment_traffic_filter_association",
			Category:         "Resources",
			ShortDescription: `Provides an Elastic Cloud traffic filter association resource, which allows traffic filter rules to be associated to an Elastic Cloud deployment outside of the control of Terraform. Associations can be created and deleted.`,
			Description: `

Provides an Elastic Cloud traffic filter association resource, which allows traffic filter rules to be associated to an Elastic Cloud deployment outside of the control of Terraform. Associations can be created and deleted.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_snapshot_repository",
			Category:         "Resources",
			ShortDescription: `Manages Elastic Cloud Enterprise snapshot repositories. ~> **This resource can only be used with Elastic Cloud Enterprise** For Elastic Cloud SaaS please use the [elasticstack_elasticsearch_snapshot_repository](https://registry.terraform.io/providers/elastic/elasticstack/latest/docs/resources/elasticsearch_snapshot_repository) resource from the [Elastic Stack terraform provider](https://registry.terraform.io/providers/elastic/elasticstack/latest).`,
			Description: `

Manages Elastic Cloud Enterprise snapshot repositories.

  ~> **This resource can only be used with Elastic Cloud Enterprise** For Elastic Cloud SaaS please use the [elasticstack_elasticsearch_snapshot_repository](https://registry.terraform.io/providers/elastic/elasticstack/latest/docs/resources/elasticsearch_snapshot_repository) resource from the [Elastic Stack terraform provider](https://registry.terraform.io/providers/elastic/elasticstack/latest).

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"ec_deployment":                            0,
		"ec_deployment_elasticsearch_keystore":     1,
		"ec_deployment_extension":                  2,
		"ec_deployment_traffic_filter":             3,
		"ec_deployment_traffic_filter_association": 4,
		"ec_snapshot_repository":                   5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
