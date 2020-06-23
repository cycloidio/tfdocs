package infoblox

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "infoblox_a_record",
			Category:         "Resources",
			ShortDescription: `Creates an A record in NIOS.`,
			Description:      ``,
			Keywords: []string{
				"a",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_view_name",
					Description: `(Optional) Unless specified, the providers tries to update IP properties in default network view`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Required) A name you want to associate with the IP address.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Optional) Updates the VM id of the vm used to provision`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The network block in cidr format`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Links the network to a tenant`,
				},
				resource.Attribute{
					Name:        "dns_view",
					Description: `(Optional) The view which contains the details of the zone. If not provided , record will be created under default view`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone in which you want to update a host record`,
				},
				resource.Attribute{
					Name:        "ip_addr",
					Description: `(Required) - The IP address you want to update in NIOS. Use the Same IP you have passed during IP allocation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infoblox_cname_record",
			Category:         "Resources",
			ShortDescription: `Creates a cname record in NIOS.`,
			Description:      ``,
			Keywords: []string{
				"cname",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "canonical",
					Description: `(Required) A name you want to associate with the IP address.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Optional) Updates the VM id of the vm used to provision`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Links the network to a tenant`,
				},
				resource.Attribute{
					Name:        "dns_view",
					Description: `(Optional) The view which contains the details of the zone. If not provided , record will be created under default view`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone in which you want to update a host record`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infoblox_ip_allocation",
			Category:         "Resources",
			ShortDescription: `Reserves an IP from a network in NIOS.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"allocation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_view_name",
					Description: `(Optional) Unless specified the resource Reserves the IP under default network view`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Required) A name you want to associate with the IP address.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The network block in cidr format`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Links the network to a tenant`,
				},
				resource.Attribute{
					Name:        "dns_view",
					Description: `(Optional) The view which contains the details of the zone.If not provided , record will be created under default view`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone in which you want to create a host record`,
				},
				resource.Attribute{
					Name:        "enable_dns",
					Description: `(optional) A boolean value which either creates or not creates for DNS purposes`,
				},
				resource.Attribute{
					Name:        "ip_addr",
					Description: `(Optional) If set , a record will be created in NIOS using a passed IP address value. Takes in a string. If no value is given, a next available IP address will be allocated in NIOS`,
				},
				resource.Attribute{
					Name:        "mac_addr",
					Description: `(Optional) If not set , a reservation will be created in NIOS. ## Additional Note Dont set the mac address if you are integrating with cloud providers to deploy a Vm and use Infoblox to give the IP address.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infoblox_ip_association",
			Category:         "Resources",
			ShortDescription: `Updates the properties of an IP address in NIOS.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_view_name",
					Description: `(Optional) Unless specified, the providers tries to update IP properties in default network view`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Required) A name you want to associate with the IP address.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Required) Updates the VM id of the vm used to provision`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The network block in cidr format`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Links the network to a tenant`,
				},
				resource.Attribute{
					Name:        "dns_view",
					Description: `(Optional) The view which contains the details of the zone. If not provided , record will be created under default view`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone in which you want to update a host record`,
				},
				resource.Attribute{
					Name:        "ip_addr",
					Description: `(Required) - The IP address you want to update in NIOS. Use the Same IP you have passed during IP allocation.`,
				},
				resource.Attribute{
					Name:        "mac_addr",
					Description: `(Optional) - Updates the actual mac adress when used with another provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infoblox_network",
			Category:         "Resources",
			ShortDescription: `Creates a network on NIOS.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_view_name",
					Description: `(Optional) Unless specified the resource creates network under default network view`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `(optional) Unless specified the resource does not associate any name to the network`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The network block in cidr format`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Links the network to a tenant`,
				},
				resource.Attribute{
					Name:        "reserve_ip",
					Description: `(optional) reserves the number of Ip's for later use. Takes an ` + "`" + `int` + "`" + ` value`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional) give the IP you want to reserve for gateway, by default the first IP gets reserved for gateway ## Note While linking the provider with azure , give ` + "`" + `reserve_ip =3` + "`" + ` because azure reserves first 4 IP's in it's cloud`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infoblox_network_view",
			Category:         "Resources",
			ShortDescription: `Creates a network view in NIOS.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"view",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Links the network view to a tenant`,
				},
				resource.Attribute{
					Name:        "network_view_name",
					Description: `(Required) Create a network view with a given name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infoblox_ptr_record",
			Category:         "Resources",
			ShortDescription: `Creates a PTR record in NIOS.`,
			Description:      ``,
			Keywords: []string{
				"ptr",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_view_name",
					Description: `(Optional) Unless specified, the providers tries to update IP properties in default network view`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Required) A name you want to associate with the IP address.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Optional) Updates the VM id of the vm used to provision`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The network block in cidr format`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Links the network to a tenant`,
				},
				resource.Attribute{
					Name:        "dns_view",
					Description: `(Optional) The view which contains the details of the zone. If not provided , record will be created under default view`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone in which you want to update a host record`,
				},
				resource.Attribute{
					Name:        "ip_addr",
					Description: `(Required) - The IP address you want to update in NIOS. Use the Same IP you have passed during IP allocation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"infoblox_a_record":       0,
		"infoblox_cname_record":   1,
		"infoblox_ip_allocation":  2,
		"infoblox_ip_association": 3,
		"infoblox_network":        4,
		"infoblox_network_view":   5,
		"infoblox_ptr_record":     6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
