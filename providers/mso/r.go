package mso

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mso_label",
			Category:         "Resources",
			ShortDescription: `Manages MSO Resource Label`,
			Description:      ``,
			Keywords: []string{
				"label",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) name of the label.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) type of the label. ## Attribute Reference ## No Attributes are Exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_rest",
			Category:         "Resources",
			ShortDescription: `MSO Rest Resource to manage the MSO objects via REST API.`,
			Description:      ``,
			Keywords: []string{
				"rest",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) MSO REST endpoint, where the data is being sent.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) HTTP method, allowed values are POST, PATCH, GET, DELETE and PUT.`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `(Required) JSON encoded payload data. NOTE: This resource will not work well in the case of Terraform destroy if there is a change in the terraform configuration required to destroy the object from the MSO, as Destroy only has the access to the data in the state file. To destroy the objects created via mso_rest in such cases modify the payload and use the Terraform apply instead. ## Attribute Reference ## No Attributes are Exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema`,
			Description:      ``,
			Keywords: []string{
				"schema",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the schema.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) name of template attached to this schema.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) tenant_id for this schema. ## Attribute Reference ## The only Attribute exposed for this resource is ` + "`" + `id` + "`" + `. Which is set to the id of schema created.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp",
			Category:         "Resources",
			ShortDescription: `MSO Schema Site Application Network Profile(ANP) Resource`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"anp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Site Anp.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Site Anp to be created.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Anp.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Site Anp. The name of the ANP should be present in the ANP list of the given ` + "`" + `schema_id` + "`" + ` and ` + "`" + `template_name` + "`" + ` ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Application Network Profiles Endpoint Groups.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"anp",
				"epg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Anp Epg to be created.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Anp Epg.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group to manage. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg_domain",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Application Network Profiles Endpoint Groups Domain.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"anp",
				"epg",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg Domain.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Anp Epg Domain to be created.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Anp Epg Domain.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group to manage.`,
				},
				resource.Attribute{
					Name:        "dn",
					Description: `(Required) The domain profile name.`,
				},
				resource.Attribute{
					Name:        "deploy_immediacy",
					Description: `(Required) The deployment immediacy of the domain. choices: [ immediate, lazy ]`,
				},
				resource.Attribute{
					Name:        "domain_type",
					Description: `(Required) The type of domain to associate. choices: [ vmmDomain, l3ExtDomain, l2ExtDomain, physicalDomain, fibreChannelDomain ]`,
				},
				resource.Attribute{
					Name:        "resolution_immediacy",
					Description: `(Required) Determines when the policies should be resolved and available. choices: [ immediate, lazy, pre-provision ]`,
				},
				resource.Attribute{
					Name:        "vlan_encap_mode",
					Description: `(Optional) Which VLAN enacap mode to use. This attribute can only be used with vmmDomain domain association. choices: [ static, dynamic ]`,
				},
				resource.Attribute{
					Name:        "allow_micro_segmentation",
					Description: `(Optional) Specifies microsegmentation is enabled or not. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) Which switching mode to use with this domain association. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "switch_type",
					Description: `(Optional) Which switch type to use with this domain association. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "micro_seg_vlan_type",
					Description: `(Optional) Virtual LAN type for microsegmentation. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "micro_seg_vlan",
					Description: `(Optional) Virtual LAN for microsegmentation. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "port_encap_vlan_type",
					Description: `(Optional) Virtual LAN type for port encap. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "port_encap_vlan",
					Description: `(Optional) Virtual LAN for port encap. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "enhanced_lagpolicy_name",
					Description: `(Optional) EPG enhanced lagpolicy name. This attribute can only be used with vmmDomain domain association.`,
				},
				resource.Attribute{
					Name:        "enhanced_lagpolicy_dn",
					Description: `(Optional) Distinguished name of EPG lagpolicy. This attribute can only be used with vmmDomain domain association. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg_static_port",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template Application Network Profiles Endpoint Groups Static Port.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"anp",
				"epg",
				"static",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Static Port.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Static Port.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name under which you want to deploy Static Port.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) ANP name under which you want to deploy Static Port.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) EPG name under which you want to deploy Static Port.`,
				},
				resource.Attribute{
					Name:        "path_type",
					Description: `(Required) The type of the static port. Allowed value is ` + "`" + `port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pod",
					Description: `(Required) The pod of the static port.`,
				},
				resource.Attribute{
					Name:        "leaf",
					Description: `(Required) The leaf of the static port.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of the static port.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The mode of the static port. Allowed values are ` + "`" + `native` + "`" + `, ` + "`" + `regular` + "`" + ` and ` + "`" + `untagged` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deployment_immediacy",
					Description: `(Required) The deployment immediacy of the static port. Allowed values are ` + "`" + `immediate` + "`" + ` and ` + "`" + `lazy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) The port encap VLAN id of the static port.`,
				},
				resource.Attribute{
					Name:        "micro_segvlan",
					Description: `(Optional) The microsegmentation VLAN id of the static port. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg_staticleaf",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Application Network Profiles Endpoint Groups StaticLeaf.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"anp",
				"epg",
				"staticleaf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg StaticLeaf.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Anp Epg StaticLeaf to be created.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Anp Epg StaticLeaf.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group to manage.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path Given to the StaticLeaf. ForceNew set to true.`,
				},
				resource.Attribute{
					Name:        "port_encap_vlan",
					Description: `(Required) The VLAN id of the static leaf. ForceNew set to true. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_anp_epg_subnet",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Application Network Profiles Endpoint Groups Subnet.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"anp",
				"epg",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) ANP name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) EPG name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP range in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of this subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) The scope of the subnet. Allowed values are ` + "`" + `private` + "`" + ` and ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Required) Whether this subnet is shared between VRFs.`,
				},
				resource.Attribute{
					Name:        "querier",
					Description: `(Optional) Whether this subnet is an IGMP querier.`,
				},
				resource.Attribute{
					Name:        "no_default_gateway",
					Description: `(Optional) Whether this subnet has a default gateway. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_bd",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Bridge Domain(BD).`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"bd",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Site Bd.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Site Bd to be created.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Bd.`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Required) Name of Site Bd. The name of the Bd should be present in the Bd list of the given ` + "`" + `schema_id` + "`" + ` and ` + "`" + `template_name` + "`" + ``,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Value to check whether host-based routing is enabled. Default value is ` + "`" + `false` + "`" + `. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_bd_l3out",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Bridge Domain L3out.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"bd",
				"l3out",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Bd L3out.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Bd L3out.`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Required) Name of Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "l3out_name",
					Description: `(Required) Name of L3out to manage.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Bd L3out to be created. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_bd_subnet",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Bridge Domain(BD) Subnet.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"bd",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Required) Bd name under which you want to deploy Subnet. The Bd Name Reference should have ` + "`" + `l2Stretch` + "`" + ` set to ` + "`" + `false` + "`" + ` to be able to add a subnet.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP of the Subnet.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name under which you want to deploy Subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet. Allowed values are ` + "`" + `private` + "`" + ` and ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this subnet is shared between VRFs.`,
				},
				resource.Attribute{
					Name:        "querier",
					Description: `(Optional) Whether this subnet is an IGMP querier.`,
				},
				resource.Attribute{
					Name:        "no_default_gateway",
					Description: `(Optional) Whether this subnet has a default gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of this subnet. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_vrf",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site VRF.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"vrf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Site Vrf.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Vrf.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Site Vrf to be created.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of Site Vrf. The name of the VRF should be present in the VRF list of the given ` + "`" + `schema_id` + "`" + ` and ` + "`" + `template_name` + "`" + ` ## Attribute Reference ## No attributes are exported`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_vrf_region",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Vrf Region.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"vrf",
				"region",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Vrf Region.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Vrf Region.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of Vrf.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Required) Name of Region to manage.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Vrf Region to be created. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_vrf_region_cidr",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Vrf Region Cidr.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"vrf",
				"region",
				"cidr",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Vrf Region.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Vrf Region.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of Vrf.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Required) Name of Region to manage.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The name of the region CIDR to manage.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Vrf Region to be created.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Required) Whether this is the primary CIDR. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site_vrf_region_cidr_subnet",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site Vrf Region Cidr Subnet.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
				"vrf",
				"region",
				"cidr",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Vrf Region Cidr Subnet.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) SiteID under which you want to deploy Vrf Region Cidr Subnet.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of Vrf.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Required) Name of Region to manage.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Vrf Region Cidr Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Required) The IP range of for the region CIDR where Vrf Region Cidr Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP subnet of this region CIDR.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the zone for the region CIDR subnet.`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) The usage for the region CIDR Subnet. ## Attribute Reference ## No attributes are exported. ## Note ## Multiple Subnets with same Ip are allowed, but the operations will take place on the first found Subnet with the given Ip.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) name of the schema.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Tenant-id to associate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the template.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the Template to be deployed on the site. ## Attribute Reference ## The only attribute exported with this resource is ` + "`" + `id` + "`" + `. Which is set to the id of schema template associated.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp",
			Category:         "Resources",
			ShortDescription: `Manages MSO Resource Schema Template Anp`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"anp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) The schema-id where anp is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the anp to add.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) template associated with the anp.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name as displayed on the MSO web interface. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp_epg",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template Application Network Profiles Endpoint Groups.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"anp",
				"epg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Anp Epg to be created.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Endpoint Group to manage.`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Required) Name of Bridge Domain to associate with.`,
				},
				resource.Attribute{
					Name:        "bd_schema_id",
					Description: `(Opional) The schemaID that defines the referenced BD.`,
				},
				resource.Attribute{
					Name:        "bd_template_name",
					Description: `(Optional) The template that defines the referenced BD.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of Vrf.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) The schemaID that defines the referenced VRF.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) The template that defines the referenced VRF.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name as displayed on the MSO web interface.`,
				},
				resource.Attribute{
					Name:        "useg_epg",
					Description: `(Optional) Boolean flag to enable or disable whether this is a USEG EPG. Default value is set to false.`,
				},
				resource.Attribute{
					Name:        "intra_epg",
					Description: `(Optional) Whether intra EPG isolation is enforced. choices: [ enforced, unenforced ]`,
				},
				resource.Attribute{
					Name:        "intersite_multicaste_source",
					Description: `(Optional) Whether intersite multicast source is enabled.`,
				},
				resource.Attribute{
					Name:        "preferred_group",
					Description: `(Optional) Boolean flag to enable or disable whether this EPG is added to preferred group. Default value is set to false. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp_epg_contract",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template Application Network Profile(ANP) Endpoint Group(EPG) Contract.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"anp",
				"epg",
				"contract",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy ANP EPG Contract.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name under which you want to deploy ANP EPG Contract.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) ANP name under which you want to deploy ANP EPG Contract.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) EPG name under which you want to deploy ANP EPG Contract.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) The contract name which you want to associate with.`,
				},
				resource.Attribute{
					Name:        "relationship_type",
					Description: `(Required) The type of the contract i.e. provider or consumer.`,
				},
				resource.Attribute{
					Name:        "contract_schema_id",
					Description: `(Optional) SchemaID of Contract. schema_id of ANP EPG will be used if not provided. Should use this parameter when Contract is in different schema than ANP EPG.`,
				},
				resource.Attribute{
					Name:        "contract_template_name",
					Description: `(Optional) Template Name associated with Contract. template_name of ANP EPG will be used if not provided. Should use this parameter when Contract is in different schema than ANP EPG. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_anp_epg_subnet",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template Application Network Profiles Endpoint Groups Subnets.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"anp",
				"epg",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Anp Epg Subnet.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Anp Epg Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "anp_name",
					Description: `(Required) Name of Application Network Profiles.`,
				},
				resource.Attribute{
					Name:        "epg_name",
					Description: `(Required) Name of Endpoint Group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) Ip Address of Subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Scope of Subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether the subnet should be shared or not. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_bd_subnet",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template Bridge Domain Subnet.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"bd",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Bridge Domain Subnet.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Bridge Domain Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "bd_name",
					Description: `(Required) Name of Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP range in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) The scope of the subnet. Allowed values are ` + "`" + `private` + "`" + `, ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the subnet.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Required) Whether this subnet is shared between VRFs.`,
				},
				resource.Attribute{
					Name:        "no_default_gateway",
					Description: `(Optional) Whether this subnet has a default gateway.`,
				},
				resource.Attribute{
					Name:        "querier",
					Description: `(Optional) Whether this subnet is an IGMP querier. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_contract",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template Bridge Domain.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"contract",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Contract.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Contract to be created.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) The name of the contract to manage.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display Name of the contract on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `(Optional) The type of filters defined in this contract. Allowed values are ` + "`" + `bothWay` + "`" + ` and ` + "`" + `oneWay` + "`" + `. Default to ` + "`" + `bothWay` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the contract.`,
				},
				resource.Attribute{
					Name:        "filter_relationships",
					Description: `(Required) Map to provide Filter Relationships.`,
				},
				resource.Attribute{
					Name:        "filter_relationships.filter_schema_id",
					Description: `(Optional) The schemaId in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationships.filter_template_name",
					Description: `(Optional) The template name in which the filter is located.`,
				},
				resource.Attribute{
					Name:        "filter_relationships.filter_name",
					Description: `(Required) The filter to associate with this contract.`,
				},
				resource.Attribute{
					Name:        "directives",
					Description: `(Required) A list of filter directives. Allowed values are ` + "`" + `log` + "`" + ` and ` + "`" + `none` + "`" + `. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_contract_filter",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template Contract Filter.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"contract",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Contract Filter.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Contract Filter to be created.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) The name of the contract to manage. There should be an existing contract with this name.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `(Required) The type of filters defined in this contract. Allowed values are ` + "`" + `bothWay` + "`" + `, ` + "`" + `provider_to_consumer` + "`" + ` and ` + "`" + `consumer_to_provider` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter_schema_id",
					Description: `(Optional) The schemaId in which the filter is located. Default is ` + "`" + `schema_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter_template_name",
					Description: `(Optional) The template name in which the filter is located. Default is ` + "`" + `template_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter_name",
					Description: `(Required) The filter name to associate with this contract. Filter must exist with the given ` + "`" + `filter_name` + "`" + `, ` + "`" + `filter_schema_id` + "`" + ` and ` + "`" + `filter_template_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directives",
					Description: `(Required) A list of filter directives. Allowed values are ` + "`" + `log` + "`" + ` and ` + "`" + `none` + "`" + `. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_deploy",
			Category:         "Resources",
			ShortDescription: `Manages deploy/undeploy operations for schema template on sites.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"deploy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) The schema-id of template.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) name of the template to deploy/undeploy.`,
				},
				resource.Attribute{
					Name:        "undeploy",
					Description: `(Optional) Boolean flag indicating whether to undeploy the template or not. Default is false.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Optional) Site-id from where the template is to be undeployed. It is required if you set undeploy = true. NOTE: This resource is intentionally created non-idempotent so that it deploys the template in every run, it will not fail if there is no change and we deploy the template again. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_externalepg",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template External Endpoint Group.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"externalepg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy External-epg.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where External-epg to be created.`,
				},
				resource.Attribute{
					Name:        "externalepg_name",
					Description: `(Required) Name of External-epg.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display Name of the External-epg on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) The VRF associated to this External-epg. VRF must exist.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) SchemaID of VRF. schema_id of External-epg will be used if not provided. Should use this parameter when VRF is in different schema than external-epg.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) Template Name of VRF. template_name of External-epg will be used if not provided. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_externalepg_contract",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template External Endpoint Group Contract.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"externalepg",
				"contract",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy External-epg.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where External-epg to be created.`,
				},
				resource.Attribute{
					Name:        "external_epg_name",
					Description: `(Required) Name of External-epg.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) Name of Contract.`,
				},
				resource.Attribute{
					Name:        "relationship_type",
					Description: `(Required) RelationType of the Contract. Values that can be used is provider or consumer`,
				},
				resource.Attribute{
					Name:        "contract_schema_id",
					Description: `(Optional) SchemaID of Contract. schema_id of External-epg will be used if not provided. Should use this parameter when Contract is in different schema than external-epg.`,
				},
				resource.Attribute{
					Name:        "contract_template_name",
					Description: `(Optional) Template Name of Contract. template_name of External-epg will be used if not provided. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_externalepg_subnet",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template External EPG Subnet.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"externalepg",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy External EPG Subnet.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where External EPG Subnet to be created.`,
				},
				resource.Attribute{
					Name:        "externalepg_name",
					Description: `(Required) Name of External EPG.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP range in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of Subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the subnet. Allowed values are ` + "`" + `shared-rtctrl` + "`" + `, ` + "`" + `export-rtctrl` + "`" + `, ` + "`" + `shared-security` + "`" + `, ` + "`" + `import-rtctrl` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) The aggregate of the subnet. Allowed values are ` + "`" + `shared-rtctrl` + "`" + `, ` + "`" + `export-rtctrl` + "`" + `, ` + "`" + `shared-security` + "`" + `, ` + "`" + `import-rtctrl` + "`" + `. Aggregate should be enabled only if shared-rtctrl is enabled in Scope. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_filter_entry",
			Category:         "Resources",
			ShortDescription: `Manages MSO Resource Schema Template Filter Entry`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"filter",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) The schema-id where Filter entry is associated.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) The template associated with the filter entry.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Filter associated with the filter entry.`,
				},
				resource.Attribute{
					Name:        "entry_name",
					Description: `(Required) The name of the entry.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name of the filter as displayed on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "entry_display_name",
					Description: `(Required) The name of the entry as displayed on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "entry_description",
					Description: `(Optional) Description of the entry created.`,
				},
				resource.Attribute{
					Name:        "ether_type",
					Description: `(Optional) The ethernet type to use for this filter entry. Allowed Values: arp, fcoe, ip, ipv4, ipv6, mac-security, mpls-unicast, trill, unspecified`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol to use for this filter entry. Allowed Values: eigrp, egp, icmp, icmpv6, igmp, igp, l2tp, ospfigp, pim, tcp, udp, unspecified`,
				},
				resource.Attribute{
					Name:        "tcp_session_rules",
					Description: `(Optional) A list of TCP session rules. Allowed Values : acknowledgement, established, finish, synchronize, reset, unspecified`,
				},
				resource.Attribute{
					Name:        "source_from",
					Description: `(Optional) The source port range from.`,
				},
				resource.Attribute{
					Name:        "source_to",
					Description: `(Optional) The source port range to.`,
				},
				resource.Attribute{
					Name:        "destination_from",
					Description: `(Optional) The destination port range from.`,
				},
				resource.Attribute{
					Name:        "destination_to",
					Description: `(Optional) The destination port range to.`,
				},
				resource.Attribute{
					Name:        "arp_flag",
					Description: `(Optional) The ARP flag to use for this filter entry. Allowed Values: reply, request, unspecified`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) Whether this filter entry is stateful. Allowed Values: true or false.`,
				},
				resource.Attribute{
					Name:        "match_only_fragments",
					Description: `(Optional) Whether this filter entry only matches fragments. Allowed Values: true or false. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_l3out",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template L3Out.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"l3out",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy L3Out.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where L3Out to be created.`,
				},
				resource.Attribute{
					Name:        "l3out_name",
					Description: `(Required) Name of L3Out.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display Name of the L3Out on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) The VRF associated to this L3out. VRF must exist.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) SchemaID of VRF. schema_id of L3Out will be used if not provided. Should use this parameter when VRF is in different schema than l3out.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) Template Name of VRF. template_name of L3Out will be used if not provided. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_site",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Site`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"site",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) name of the schema.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Site-id to associate.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template to be deployed on the site. ## Attribute Reference ## The only attribute exported with this resource is ` + "`" + `id` + "`" + `. Which is set to the id of schema site associated.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_vrf",
			Category:         "Resources",
			ShortDescription: `Manages MSO Resource Schema Template Vrf`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"vrf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) The schema-id where vrf is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the vrf to add.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) template associated with the vrf.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name as displayed on the MSO web interface.`,
				},
				resource.Attribute{
					Name:        "layer3_multicast",
					Description: `(Optional) Whether to enable L3 multicast. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_site",
			Category:         "Resources",
			ShortDescription: `Manages MSO Site`,
			Description:      ``,
			Keywords: []string{
				"site",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the site.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username for the APICs.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password for the APICs.`,
				},
				resource.Attribute{
					Name:        "apic_site_id",
					Description: `(Required) The site ID of the APICs.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(Required) A list of URLs to reference the APICs.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The labels for this site.`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Optional) Location of the site. ## Attribute Reference ## No Attributes are Exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_schema_template_bd",
			Category:         "Resources",
			ShortDescription: `Manages MSO Schema Template Bridge Domain.`,
			Description:      ``,
			Keywords: []string{
				"schema",
				"template",
				"bd",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Required) SchemaID under which you want to deploy Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template where Bridge Domain to be created.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display Name of the Bridge Domain on the MSO UI.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) Name of VRF to attach with Bridge Domain. VRF must exist.`,
				},
				resource.Attribute{
					Name:        "vrf_schema_id",
					Description: `(Optional) SchemaID of VRF. schema_id of Bridge Domain will be used if not provided. Should use this parameter when VRF is in different schema than BD.`,
				},
				resource.Attribute{
					Name:        "vrf_template_name",
					Description: `(Optional) Template Name of VRF. template_name of Bridge Domain will be used if not provided. Should use this parameter when VRF is in different schema than BD.`,
				},
				resource.Attribute{
					Name:        "layer2_unknown_unicast",
					Description: `(Optional) Type of layer 2 unknown unicast. Allowed values are ` + "`" + `flood` + "`" + ` and ` + "`" + `proxy` + "`" + `. Default to ` + "`" + `flood` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "intersite_bum_traffic",
					Description: `(Optional) Boolean Flag to enable or disable intersite bum traffic. Default to false.`,
				},
				resource.Attribute{
					Name:        "optimize_wan_bandwidth",
					Description: `(Optional) Boolean flag to enable or disable the wan bandwidth optimization. Default to false.`,
				},
				resource.Attribute{
					Name:        "layer2_stretch",
					Description: `(Optional) Boolean flag to enable or disable the layer-2 stretch. Default to false. Should enable this flag if you want to create subnets under this Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "layer3_multicast",
					Description: `(Optional) Boolean flag to enable or disable layer 3 multicast traffic. Default to false.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy",
					Description: `(Optional) Map to provide dhcp_policy configurations.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.name",
					Description: `(Optional) dhcp_policy name. Required if you specify the dhcp_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.version",
					Description: `(Optional) Version of dhcp_policy. Required if you specify the dhcp_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.dhcp_option_policy_name",
					Description: `(Optional) Name of dhcp_option_policy.`,
				},
				resource.Attribute{
					Name:        "dhcp_policy.dhcp_option_policy_version",
					Description: `(Optional) Version of dhcp_option_policy. Require if you specify the ` + "`" + `dhcp_policy.dhcp_option_policy_name` + "`" + `. ## Attribute Reference ## No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_tenant",
			Category:         "Resources",
			ShortDescription: `Manages MSO Tenant`,
			Description:      ``,
			Keywords: []string{
				"tenant",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tenant.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name of the tenant to be displayed in the web UI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for this tenant.`,
				},
				resource.Attribute{
					Name:        "user_associations",
					Description: `(Optional) A list of associated users for this tenant.`,
				},
				resource.Attribute{
					Name:        "site_association",
					Description: `(Optional) A list of associated sites for this tenant. ## Attribute Reference ## The only Attribute exposed for this resource is ` + "`" + `id` + "`" + `. Which is set to the id of tenant created.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mso_user",
			Category:         "Resources",
			ShortDescription: `Manages MSO User`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) username of the user. It must contain at least 1 character in length.`,
				},
				resource.Attribute{
					Name:        "user_password",
					Description: `(Required) password of the user. It must contain at least 8 characters in length.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) roles given to the user.`,
				},
				resource.Attribute{
					Name:        "roles.roleid",
					Description: `(Required) id of roles given to the user.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Optional) firstname of the user.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Optional) lastname of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email of the user.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `(Optional) phone of the user.`,
				},
				resource.Attribute{
					Name:        "account-status",
					Description: `(Optional) account status of the user.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) domain status of the user.`,
				},
				resource.Attribute{
					Name:        "roles.access_type",
					Description: `(Optional) access_type of roles given to the user. ## Attribute Reference ## The only attribute exported with this resource is ` + "`" + `id` + "`" + `. Which is set to the id of user associated.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"mso_label":                                0,
		"mso_rest":                                 1,
		"mso_schema":                               2,
		"mso_schema_site_anp":                      3,
		"mso_schema_site_anp_epg":                  4,
		"mso_schema_site_anp_epg_domain":           5,
		"mso_schema_site_anp_epg_static_port":      6,
		"mso_schema_site_anp_epg_staticleaf":       7,
		"mso_schema_site_anp_epg_subnet":           8,
		"mso_schema_site_bd":                       9,
		"mso_schema_site_bd_l3out":                 10,
		"mso_schema_site_bd_subnet":                11,
		"mso_schema_site_vrf":                      12,
		"mso_schema_site_vrf_region":               13,
		"mso_schema_site_vrf_region_cidr":          14,
		"mso_schema_site_vrf_region_cidr_subnet":   15,
		"mso_schema_template":                      16,
		"mso_schema_template_anp":                  17,
		"mso_schema_template_anp_epg":              18,
		"mso_schema_template_anp_epg_contract":     19,
		"mso_schema_template_anp_epg_subnet":       20,
		"mso_schema_template_bd_subnet":            21,
		"mso_schema_template_contract":             22,
		"mso_schema_template_contract_filter":      23,
		"mso_schema_template_deploy":               24,
		"mso_schema_template_externalepg":          25,
		"mso_schema_template_externalepg_contract": 26,
		"mso_schema_template_externalepg_subnet":   27,
		"mso_schema_template_filter_entry":         28,
		"mso_schema_template_l3out":                29,
		"mso_schema_site":                          30,
		"mso_schema_template_vrf":                  31,
		"mso_site":                                 32,
		"mso_schema_template_bd":                   33,
		"mso_tenant":                               34,
		"mso_user":                                 35,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
