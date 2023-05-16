package fortiflexvm

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fortiflexvm_config",
			Category:         "Configs",
			ShortDescription: `Create a new Flex VM Configuration under a Flex VM Program.`,
			Description:      ``,
			Keywords: []string{
				"configs",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "program_serial_number",
					Description: `(Required/String) The serial number of your Flex VM Program. This serial number should start with ` + "`" + `"ELAVMR"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required/String) The name of your Flex VM Configuration.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional/String) Configuration status. If you don't specify, the configuration status keeps unchanged. The default status is ` + "`" + `ACTIVE` + "`" + ` once you create a Flex VM Configuration. It must be one of the following options:`,
				},
				resource.Attribute{
					Name:        "faz_vm",
					Description: `(Block List) You must fill in this block if your ` + "`" + `product_type` + "`" + ` is ` + "`" + `"FAZ_VM"` + "`" + `. The structure of [` + "`" + `faz_vm` + "`" + ` block](#nestedobjatt--faz_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "fgt_vm_bundle",
					Description: `(Block List) You must fill in this block if your ` + "`" + `product_type` + "`" + ` is ` + "`" + `"FGT_VM_Bundle"` + "`" + `. The structure of [` + "`" + `fgt_vm_bundle` + "`" + ` block](#nestedatt--fgt_vm_bundle) is documented below.`,
				},
				resource.Attribute{
					Name:        "fgt_vm_lcs",
					Description: `(Block List) You must fill in this block if your ` + "`" + `product_type` + "`" + ` is ` + "`" + `"FGT_VM_LCS"` + "`" + `. The structure of [` + "`" + `fgt_vm_lcs` + "`" + ` block](#nestedatt--fgt_vm_lcs) is documented below.`,
				},
				resource.Attribute{
					Name:        "fmg_vm",
					Description: `(Block List) You must fill in this block if your ` + "`" + `product_type` + "`" + ` is ` + "`" + `"FMG_VM"` + "`" + `. The structure of [` + "`" + `fmg_vm` + "`" + ` block](#nestedatt--fmg_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "fpc_vm",
					Description: `(Block List) You must fill in this block if your ` + "`" + `product_type` + "`" + ` is ` + "`" + `"FPC_VM"` + "`" + `. The structure of [` + "`" + `fpc_vm` + "`" + ` block](#nestedobjatt--fpc_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "fwb_vm",
					Description: `(Block List) You must fill in this block if your ` + "`" + `product_type` + "`" + ` is ` + "`" + `"FWB_VM"` + "`" + `. The structure of [` + "`" + `fwb_vm` + "`" + ` block](#nestedatt--fwb_vm) is documented below. <a id="nestedblock--faz_vm"></a> The ` + "`" + `faz_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "adom_num",
					Description: `(Required if ` + "`" + `product_type = "FAZ_VM"` + "`" + `/Number) Number of ADOMs. A number between 0 and 1200 (inclusive).`,
				},
				resource.Attribute{
					Name:        "daily_storage",
					Description: `(Required if ` + "`" + `product_type = "FAZ_VM"` + "`" + `/Number) Daily Storage (GB). A number between 5 and 8300 (inclusive).`,
				},
				resource.Attribute{
					Name:        "support_service",
					Description: `(Required if ` + "`" + `product_type = "FAZ_VM"` + "`" + `/String) Support Service. Currently, the only available option is ` + "`" + `"FAZFC247"` + "`" + ` (FortiCare Premium). The default value is ` + "`" + `"FAZFC247"` + "`" + `. <a id="nestedblock--fgt_vm_bundle"></a> The ` + "`" + `fgt_vm_bundle` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "cpu_size",
					Description: `(Required if ` + "`" + `product_type = "FGT_VM_Bundle"` + "`" + `/String) The number of CPUs. The value of this attribute is one of ` + "`" + `"1"` + "`" + `, ` + "`" + `"2"` + "`" + `, ` + "`" + `"4"` + "`" + `, ` + "`" + `"8"` + "`" + `, ` + "`" + `"16"` + "`" + `, ` + "`" + `"32"` + "`" + ` or ` + "`" + `"2147483647"` + "`" + ` (unlimited).`,
				},
				resource.Attribute{
					Name:        "service_pkg",
					Description: `(Required if ` + "`" + `product_type = "FGT_VM_Bundle"` + "`" + `/String) The value of this attribute is one of ` + "`" + `"FC"` + "`" + ` (FortiCare), ` + "`" + `"UTM"` + "`" + `, ` + "`" + `"ENT"` + "`" + ` (Enterprise) or ` + "`" + `"ATP"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vdom_num",
					Description: `(Optional/Number) Number of VDOMs. A number between 0 and 500 (inclusive). The default number is 0. <a id="nestedblock--fgt_vm_lcs"></a> The ` + "`" + `fgt_vm_lcs` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "cloud_services",
					Description: `(Optional/List of String) The cloud services this FortiGate Virtual Machine supports. The default value is an empty list. It should be a combination of:`,
				},
				resource.Attribute{
					Name:        "cpu_size",
					Description: `(Required if ` + "`" + `product_type = "FGT_VM_LCS"` + "`" + `/String) The number of CPUs. A number between 1 and 96 (inclusive).`,
				},
				resource.Attribute{
					Name:        "fortiguard_services",
					Description: `(Optional/List of String) The fortiguard services this FortiGate Virtual Machine supports. The default value is an empty list. It should be a combination of:`,
				},
				resource.Attribute{
					Name:        "support_service",
					Description: `(Required if ` + "`" + `product_type = "FGT_VM_LCS"` + "`" + `/String) ` + "`" + `"FC247"` + "`" + ` (FortiCare 24x7) or ` + "`" + `"ASET"` + "`" + ` (FortiCare Elite).`,
				},
				resource.Attribute{
					Name:        "vdom_num",
					Description: `(Optional/Number) Number of VDOMs. A number between 0 and 500 (inclusive). The default number is 0. <a id="nestedblock--fmg_vm"></a> The ` + "`" + `fmg_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "adom_num",
					Description: `(Optional/Number) Number of ADOMs. A number between 1 and 100000 (inclusive). The default value is 0.`,
				},
				resource.Attribute{
					Name:        "managed_dev",
					Description: `(Optional/Number) Number of managed devices. A number between 1 and 100000 (inclusive). The default value is 1. <a id="nestedblock--fpc_vm"></a> The ` + "`" + `fpc_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "managed_dev",
					Description: `(Required if ` + "`" + `product_type = "FPC_VM"` + "`" + `/Number) Number of managed devices. A number between 0 and 100000 (inclusive). <a id="nestedblock--fwb_vm"></a> The ` + "`" + `fwb_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "cpu_size",
					Description: `(Required if ` + "`" + `product_type = "FWB_VM"` + "`" + `/String) Number of CPUs. The value of this attribute is one of ` + "`" + `"1"` + "`" + `, ` + "`" + `"2"` + "`" + `, ` + "`" + `"4"` + "`" + `, ` + "`" + `"8"` + "`" + ` or ` + "`" + `"16"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_pkg",
					Description: `(Required if ` + "`" + `product_type = "FWB_VM"` + "`" + `/String) Service Package. Valid values: ` + "`" + `"FWBSTD"` + "`" + ` (Standard) or ` + "`" + `"FWBADV"` + "`" + ` (Advanced). ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource. ## Import FlexVM Configuration can be imported by using the following steps: First, specify the ` + "`" + `program_serial_number` + "`" + ` when you configure the provider. ` + "`" + `` + "`" + `` + "`" + ` provider "fortiflexvm" { username = "ABCDEFG" password = "HIJKLMN" import_options= toset(["program_serial_number=ELAVMS000000XXXX"]) } ` + "`" + `` + "`" + `` + "`" + ` Then, use the following command to import the FlexVM Configuration. ` + "`" + `` + "`" + `` + "`" + ` terraform import fortiflexvm_config.labelname {{id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource. ## Import FlexVM Configuration can be imported by using the following steps: First, specify the ` + "`" + `program_serial_number` + "`" + ` when you configure the provider. ` + "`" + `` + "`" + `` + "`" + ` provider "fortiflexvm" { username = "ABCDEFG" password = "HIJKLMN" import_options= toset(["program_serial_number=ELAVMS000000XXXX"]) } ` + "`" + `` + "`" + `` + "`" + ` Then, use the following command to import the FlexVM Configuration. ` + "`" + `` + "`" + `` + "`" + ` terraform import fortiflexvm_config.labelname {{id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiflexvm_vms_create",
			Category:         "VMS",
			ShortDescription: `Create one or more VMs based on a Flex VM Configuration.`,
			Description:      ``,
			Keywords: []string{
				"vms",
				"create",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required/Number) The ID of a Flex VM Configuration.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional/String) The description of VM(s).`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional/String) VM(s) end date. It can not be before today's date or after the program's end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: ` + "`" + `YYYY-MM-DDThh:mm:ss` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "folder_path",
					Description: `(Optional/String) The folder path of the VM(s).`,
				},
				resource.Attribute{
					Name:        "vm_count",
					Description: `(Optional/Number) The number of VM(s) to be created. The default value is 1. ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource. ## Import !> You can not import this resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource. ## Import !> You can not import this resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiflexvm_vms_update",
			Category:         "VMS",
			ShortDescription: `Update a VM, stop or reactivate a VM, or regenerate a VM token.`,
			Description:      ``,
			Keywords: []string{
				"vms",
				"update",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "serial_number",
					Description: `(Required/String) The unique serial number of the VM to be updated.`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required if you want to update the ` + "`" + `config_id` + "`" + `, ` + "`" + `description` + "`" + ` or ` + "`" + `end_date` + "`" + `/Number) Set a new Flex VM Configuration.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional/String) Set a new description.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional/String) Set a new end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: ` + "`" + `YYYY-MM-DDThh:mm:ss` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "regenerate_token",
					Description: `(Optional/Boolean) Whether to regenerate a new token. If this argument is ` + "`" + `true` + "`" + `, every time you run ` + "`" + `terraform apply` + "`" + `, the system will generate a new token for your VM. Please remember to set it as ` + "`" + `false` + "`" + ` if you don't want to regenerate the token anymore.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional/String) Valid values: ` + "`" + `ACTIVE` + "`" + ` or ` + "`" + `STOPPED` + "`" + `. If you want to use this argument, it is highly recommended to also specify the argument ` + "`" + `config_id` + "`" + `. ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource. ## Import ~> Currently, our Provider only supports importing one VM profile. VM profile can be imported by using the following steps: First, specify the ` + "`" + `config_id` + "`" + ` when you configure the provider. ` + "`" + `` + "`" + `` + "`" + ` provider "fortiflexvm" { username = "ABCDEFG" password = "HIJKLMN" import_options= toset(["config_id=42"]) } ` + "`" + `` + "`" + `` + "`" + ` Then, use the following command to import the VM profile. ` + "`" + `` + "`" + `` + "`" + ` terraform import fortiflexvm_vms_update.labelname {{serial_number}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource. ## Import ~> Currently, our Provider only supports importing one VM profile. VM profile can be imported by using the following steps: First, specify the ` + "`" + `config_id` + "`" + ` when you configure the provider. ` + "`" + `` + "`" + `` + "`" + ` provider "fortiflexvm" { username = "ABCDEFG" password = "HIJKLMN" import_options= toset(["config_id=42"]) } ` + "`" + `` + "`" + `` + "`" + ` Then, use the following command to import the VM profile. ` + "`" + `` + "`" + `` + "`" + ` terraform import fortiflexvm_vms_update.labelname {{serial_number}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"fortiflexvm_config":     0,
		"fortiflexvm_vms_create": 1,
		"fortiflexvm_vms_update": 2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
