package fmc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_access_policies",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_access_rules",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_fqdn_objects",
			Category:         "Resources",
			ShortDescription: `Resource for FQDN Objects in FMC Example An example is shown below: hcl resource "fmc_fqdn_objects" "new" { name = "Cisco" value = "cisco.com" description = "Cisco domain" dns_resolution = "IPV4_ONLY" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_ftd_autonat_rules",
			Category:         "Resources",
			ShortDescription: `Resource for Auto NAT Rules in FMC Example An example is shown below: ` + "`" + `` + "`" + `` + "`" + `hcl resource "fmcftdautonatrules" "newrule" { natpolicy = fmcftdnatpolicies.natpolicy.id description = "Testing Auto NAT priv-pub" nattype = "static" sourceinterface { id = data.fmcsecurityzones.inside.id type = data.fmcsecurityzones.inside.type } destinationinterface { id = data.fmcsecurityzones.outside.id type = data.fmcsecurityzones.outside.type } originalnetwork { id = data.fmcnetworkobjects.private.id type = data.fmcnetworkobjects.private.type } translatednetwork { id = data.fmcnetworkobjects.public.id type = data.fmcnetworkobjects.public.type } translatednetworkisdestinationinterface = false originalport { port = 53 protocol = "udp" } translatedport = 5353 ipv6 = true } resource "fmcftdautonatrules" "newrule2" { natpolicy = fmcftdnatpolicies.natpolicy.id description = "Testing Auto NAT pub-priv" nattype = "dynamic" sourceinterface { id = data.fmcsecurityzones.inside.id type = data.fmcsecurityzones.inside.type } destinationinterface { id = data.fmcsecurityzones.outside.id type = data.fmcsecurityzones.outside.type } originalnetwork { id = data.fmchostobjects.CUCMPub.id type = data.fmchostobjects.CUCMPub.type } translatednetworkisdestinationinterface = false patoptions { patpooladdress { id = data.fmcnetworkobjects.private.id type = data.fmcnetworkobjects.private.type } extendedpattable = true roundrobin = true } ipv6 = true } ` + "`" + `` + "`" + ` **Note** If creating multiple rules during a singleterraform apply, remember to usedepends_on` + "`" + ` to chain the rules so that terraform creates it in the same order that you intended.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_ftd_deploy",
			Category:         "Resources",
			ShortDescription: `Resource for deploying changes to FTD in FMC Example An example is shown below: hcl resource "fmc_ftd_deploy" "ftd" { device = data.fmc_devices.ftd.id ignore_warning = false force_deploy = false }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_ftd_manualnat_rules",
			Category:         "Resources",
			ShortDescription: `Resource for Manual NAT Rules in FMC Example An example is shown below: ` + "`" + `` + "`" + `` + "`" + `hcl resource "fmcftdmanualnatrules" "newrule" { natpolicy = fmcftdnatpolicies.natpolicy.id description = "Testing Manual NAT priv-pub" nattype = "static" sourceinterface { id = data.fmcsecurityzones.inside.id type = data.fmcsecurityzones.inside.type } destinationinterface { id = data.fmcsecurityzones.outside.id type = data.fmcsecurityzones.outside.type } originalsource { id = data.fmcnetworkobjects.public.id type = data.fmcnetworkobjects.public.type } translateddestination { id = data.fmcnetworkobjects.public.id type = data.fmcnetworkobjects.public.type } interfaceinoriginaldestination = true interfaceintranslatedsource = true ipv6 = true } resource "fmcftdmanualnatrules" "newruleafter" { natpolicy = fmcftdnatpolicies.natpolicy.id description = "Testing Manual NAT priv-pub" nattype = "static" section = "afterauto" sourceinterface { id = data.fmcsecurityzones.inside.id type = data.fmcsecurityzones.inside.type } destinationinterface { id = data.fmcsecurityzones.outside.id type = data.fmcsecurityzones.outside.type } originalsource { id = data.fmcnetworkobjects.public.id type = data.fmcnetworkobjects.public.type } translateddestination { id = data.fmcnetworkobjects.public.id type = data.fmcnetworkobjects.public.type } interfaceinoriginaldestination = true interfaceintranslatedsource = true ipv6 = true } resource "fmcftdmanualnatrules" "newrulebefore1" { natpolicy = fmcftdnatpolicies.natpolicy.id description = "Testing Manual NAT before priv-pub" nattype = "static" section = "beforeauto" targetindex = 1 sourceinterface { id = data.fmcsecurityzones.inside.id type = data.fmcsecurityzones.inside.type } destinationinterface { id = data.fmcsecurityzones.outside.id type = data.fmcsecurityzones.outside.type } originalsource { id = data.fmcnetworkobjects.public.id type = data.fmcnetworkobjects.public.type } translateddestination { id = data.fmchostobjects.CUCMPub.id type = data.fmchostobjects.CUCMPub.type } interfaceinoriginaldestination = true interfaceintranslatedsource = true ipv6 = true } ` + "`" + `` + "`" + ` **Note** If creating multiple rules during a singleterraform apply, remember to usedepends_on` + "`" + ` to chain the rules so that terraform creates it in the same order that you intended.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_ftd_nat_policies",
			Category:         "Resources",
			ShortDescription: `Resource for NAT Policies in FMC Example An example is shown below: hcl resource "fmc_ftd_nat_policies" "nat_policy" { name = "Terraform NAT Policy" description = "New NAT policy!" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_host_objects",
			Category:         "Resources",
			ShortDescription: `Resource for Host Objects in FMC Example An example is shown below: hcl resource "fmc_host_objects" "host" { name = "Web Server" value = "10.10.10.10" description = "K8s primary" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_icmpv4_objects",
			Category:         "Resources",
			ShortDescription: `Resource for ICMPv4 Objects in FMC Example An example is shown below: hcl resource "fmc_icmpv4_objects" "wrong-proto" { name = "wrong-proto" icmp_type = "3" code = 2 }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_network_group_objects",
			Category:         "Resources",
			ShortDescription: `Resource for Network Group Objects in FMC Example An example is shown below: hcl resource "fmc_network_group_objects" "PrivateGroup" { name = "PrivateGroup" description = "Terraform private group" objects { id = data.fmc_network_objects.PrivateVLAN.id type = data.fmc_network_objects.PrivateVLAN.type } objects { id = fmc_network_objects.PrivateVLANDR.id type = fmc_network_objects.PrivateVLANDR.type } literals { value = "10.10.10.10" type = "Host" } }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_network_objects",
			Category:         "Resources",
			ShortDescription: `Resource for Network Objects in FMC Example An example is shown below: hcl resource "fmc_network_objects" "PrivateVLANDR" { name = "VLAN-Private-DRsite" value = "10.10.10.0/24" description = "Terraform DR network object" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_policy_devices_assignments",
			Category:         "Resources",
			ShortDescription: `Resource for Policy Device Assignments in FMC Example An example is shown below: hcl resource "fmc_policy_devices_assignments" "policy_assignment" { policy { id = fmc_access_policies.access_policy.id type = fmc_access_policies.access_policy.type } target_devices { id = data.fmc_devices.device.id type = data.fmc_devices.device.type } } Note You cannot delete a policy assignment, only reassign the devices to another policy. So, the delete operation on terraform does nothing, but the assignment is not deleted until you have manually moved the devices to another policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_port_group_objects",
			Category:         "Resources",
			ShortDescription: `Resource for Port Group Objects in FMC An example is shown below: ` + "`" + `` + "`" + `` + "`" + `hcl Example resource "fmcportgroupobjects" "port-group" { name = "TCP-ICMP" description = "Combo ports" objects { id = fmcportobjects.http.id type = fmcportobjects.http.type } objects { id = fmcicmpv4objects.wrong-proto.id type = fmcicmpv4_objects.wrong-proto.type } } ` + "`" + `` + "`" + `` + "`" + ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_port_objects",
			Category:         "Resources",
			ShortDescription: `Resource for Port Objects in FMC Example An example is shown below: hcl resource "fmc_port_objects" "http" { name = "HTTP" port = "80" protocol = "TCP" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_range_objects",
			Category:         "Resources",
			ShortDescription: `Resource for Network Range Objects in FMC Example An example is shown below: hcl resource "fmc_range_objects" "servers" { name = "k8s-cluster" value = "10.10.10.10-10.10.10.16" description = "K8s Prod Cluster" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_url_object_group",
			Category:         "Resources",
			ShortDescription: `Resource for URL Object Groups in FMC Example An example is shown below: hcl resource "fmc_url_object_group" "URLGroup" { name = "URLGroup" description = "Websites" objects { id = data.fmc_url_objects.cisco-home.id type = data.fmc_url_objects.cisco-home.type } objects { id = fmc_url_objects.cisco-tac.id type = fmc_url_objects.cisco-tac.type } literals { url = "https://www.terraform.io/" type = "Url" } }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fmc_fmc_url_objects",
			Category:         "Resources",
			ShortDescription: `Resource for URL Objects in FMC Example An example is shown below: hcl resource "fmc_url_objects" "cisco-home" { name = "cisco-home" url = "https://www.cisco.com/" description = "Cisco home page" }`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"fmc_fmc_access_policies":            0,
		"fmc_fmc_access_rules":               1,
		"fmc_fmc_fqdn_objects":               2,
		"fmc_fmc_ftd_autonat_rules":          3,
		"fmc_fmc_ftd_deploy":                 4,
		"fmc_fmc_ftd_manualnat_rules":        5,
		"fmc_fmc_ftd_nat_policies":           6,
		"fmc_fmc_host_objects":               7,
		"fmc_fmc_icmpv4_objects":             8,
		"fmc_fmc_network_group_objects":      9,
		"fmc_fmc_network_objects":            10,
		"fmc_fmc_policy_devices_assignments": 11,
		"fmc_fmc_port_group_objects":         12,
		"fmc_fmc_port_objects":               13,
		"fmc_fmc_range_objects":              14,
		"fmc_fmc_url_object_group":           15,
		"fmc_fmc_url_objects":                16,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
