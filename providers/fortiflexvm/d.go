package fortiflexvm

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fortiflexvm_configs_list",
			Category:         "Data Sources",
			ShortDescription: `Get list of Flex VM Configurations for a Program.`,
			Description: `
Get list of Flex VM Configurations for a Program.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "program_serial_number",
					Description: `(Required/String) The unique serial number of the Flex VM Program. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "configs",
					Description: `(List of Object) The list of Flex VM Configurations for a Program. The structure of [` + "`" + `configs` + "`" + ` block](#nestedatt--configs) is documented below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "program_serial_number",
					Description: `(String) The unique serial number of this Flex VM Program. <a id="nestedatt--configs"></a> The ` + "`" + `configs` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "faz_vm",
					Description: `(List of Object) List of FortiAnalyzer Virtual Machines this Flex VM Project contains. The structure of [` + "`" + `configs.faz_vm` + "`" + ` block](#nestedobjatt--configs--faz_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "fgt_vm_bundle",
					Description: `(List of Object) List of FortiGate Virtual Machines (Service Bundle) this Flex VM Project contains. The structure of [` + "`" + `configs.fgt_vm_bundle` + "`" + ` block](#nestedobjatt--configs--fgt_vm_bundle) is documented below.`,
				},
				resource.Attribute{
					Name:        "fgt_vm_lcs",
					Description: `(List of Object) List of FortiGate Virtual Machines (A La Carte Services) this Flex VM Project contains. The structure of [` + "`" + `configs.fgt_vm_lcs` + "`" + ` block](#nestedobjatt--configs--fgt_vm_lcs) is documented below.`,
				},
				resource.Attribute{
					Name:        "fmg_vm",
					Description: `(List of Object) List of FortiManager Virtual Machines this Flex VM Project contains. The structure of [` + "`" + `configs.fmg_vm` + "`" + ` block](#nestedobjatt--configs--fmg_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "fpc_vm",
					Description: `(List of Object) List of FortiPortal Virtual Machines this Flex VM Project contains. The structure of [` + "`" + `configs.fpc_vm` + "`" + ` block](#nestedobjatt--configs--fpc_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "fwb_vm",
					Description: `(List of Object) List of FortiWeb Virtual Machines this Flex VM Project contains. The structure of [` + "`" + `configs.fwb_vm` + "`" + ` block](#nestedobjatt--configs--fwb_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) Flex VM Configuration name.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `(String) Flex VM Configuration type. Possible values:`,
				},
				resource.Attribute{
					Name:        "program_serial_number",
					Description: `(String) The unique serial number of the Flex VM Program this Flex VM Configuration belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) The status of this Flex VM Configuration. ` + "`" + `ACTIVATE` + "`" + ` or ` + "`" + `DISABLED` + "`" + `. <a id="nestedobjatt--configs--faz_vm"></a> The ` + "`" + `configs.faz_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "adom_num",
					Description: `(Number) Number of ADOMs. A number between 0 and 1200 (inclusive).`,
				},
				resource.Attribute{
					Name:        "daily_storage",
					Description: `(Number) Daily Storage (GB). A number between 5 and 8300 (inclusive).`,
				},
				resource.Attribute{
					Name:        "support_service",
					Description: `(String) Support Service. ` + "`" + `"FAZFC247"` + "`" + ` (FortiCare Premium). <a id="nestedobjatt--configs--fgt_vm_bundle"></a> The ` + "`" + `configs.fgt_vm_bundle` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "cpu_size",
					Description: `(String) The number of CPUs. The value of this attribute is one of ` + "`" + `"1"` + "`" + `, ` + "`" + `"2"` + "`" + `, ` + "`" + `"4"` + "`" + `, ` + "`" + `"8"` + "`" + `, ` + "`" + `"16"` + "`" + `, ` + "`" + `"32"` + "`" + ` or ` + "`" + `"2147483647"` + "`" + ` (unlimited).`,
				},
				resource.Attribute{
					Name:        "service_pkg",
					Description: `(String) The value of this attribute is one of ` + "`" + `"FC"` + "`" + ` (FortiCare), ` + "`" + `"UTM"` + "`" + `, ` + "`" + `"ENT"` + "`" + ` (Enterprise) or ` + "`" + `"ATP"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vdom_num",
					Description: `(Number) Number of VDOMs. A number between 0 and 500 (inclusive). <a id="nestedobjatt--configs--fgt_vm_lcs"></a> The ` + "`" + `configs.fgt_vm_lcs` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "cloud_services",
					Description: `(List of String) The cloud services this FortiGate Virtual Machine supports. The combination of:`,
				},
				resource.Attribute{
					Name:        "cpu_size",
					Description: `(String) The number of CPUs. A number between 1 and 96 (inclusive).`,
				},
				resource.Attribute{
					Name:        "fortiguard_services",
					Description: `(List of String) The FortiGuard services this FortiGate Virtual Machine supports. The combination of:`,
				},
				resource.Attribute{
					Name:        "support_service",
					Description: `(String) ` + "`" + `"FC247"` + "`" + ` (FortiCare 24x7) or ` + "`" + `"ASET"` + "`" + ` (FortiCare Elite).`,
				},
				resource.Attribute{
					Name:        "vdom_num",
					Description: `(Number) Number of VDOMs. A number between 0 and 500 (inclusive). <a id="nestedobjatt--configs--fmg_vm"></a> The ` + "`" + `configs.fmg_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "adom_num",
					Description: `(Number) Number of ADOMs. A number between 1 and 100000 (inclusive).`,
				},
				resource.Attribute{
					Name:        "managed_dev",
					Description: `(Number) Number of managed devices. A number between 1 and 100000 (inclusive). <a id="nestedobjatt--configs--fpc_vm"></a> The ` + "`" + `configs.fpc_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "managed_dev",
					Description: `(Number) Number of managed devices. A number between 0 and 100000 (inclusive). <a id="nestedobjatt--configs--fwb_vm"></a> The ` + "`" + `configs.fwb_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "cpu_size",
					Description: `(String) Number of CPUs. The value of this attribute is one of ` + "`" + `"1"` + "`" + `, ` + "`" + `"2"` + "`" + `, ` + "`" + `"4"` + "`" + `, ` + "`" + `"8"` + "`" + ` or ` + "`" + `"16"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_pkg",
					Description: `(String) Service Package. ` + "`" + `"FWBSTD"` + "`" + ` (Standard) or ` + "`" + `"FWBADV"` + "`" + ` (Advanced).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "configs",
					Description: `(List of Object) The list of Flex VM Configurations for a Program. The structure of [` + "`" + `configs` + "`" + ` block](#nestedatt--configs) is documented below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "program_serial_number",
					Description: `(String) The unique serial number of this Flex VM Program. <a id="nestedatt--configs"></a> The ` + "`" + `configs` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "faz_vm",
					Description: `(List of Object) List of FortiAnalyzer Virtual Machines this Flex VM Project contains. The structure of [` + "`" + `configs.faz_vm` + "`" + ` block](#nestedobjatt--configs--faz_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "fgt_vm_bundle",
					Description: `(List of Object) List of FortiGate Virtual Machines (Service Bundle) this Flex VM Project contains. The structure of [` + "`" + `configs.fgt_vm_bundle` + "`" + ` block](#nestedobjatt--configs--fgt_vm_bundle) is documented below.`,
				},
				resource.Attribute{
					Name:        "fgt_vm_lcs",
					Description: `(List of Object) List of FortiGate Virtual Machines (A La Carte Services) this Flex VM Project contains. The structure of [` + "`" + `configs.fgt_vm_lcs` + "`" + ` block](#nestedobjatt--configs--fgt_vm_lcs) is documented below.`,
				},
				resource.Attribute{
					Name:        "fmg_vm",
					Description: `(List of Object) List of FortiManager Virtual Machines this Flex VM Project contains. The structure of [` + "`" + `configs.fmg_vm` + "`" + ` block](#nestedobjatt--configs--fmg_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "fpc_vm",
					Description: `(List of Object) List of FortiPortal Virtual Machines this Flex VM Project contains. The structure of [` + "`" + `configs.fpc_vm` + "`" + ` block](#nestedobjatt--configs--fpc_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "fwb_vm",
					Description: `(List of Object) List of FortiWeb Virtual Machines this Flex VM Project contains. The structure of [` + "`" + `configs.fwb_vm` + "`" + ` block](#nestedobjatt--configs--fwb_vm) is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) Flex VM Configuration name.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `(String) Flex VM Configuration type. Possible values:`,
				},
				resource.Attribute{
					Name:        "program_serial_number",
					Description: `(String) The unique serial number of the Flex VM Program this Flex VM Configuration belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) The status of this Flex VM Configuration. ` + "`" + `ACTIVATE` + "`" + ` or ` + "`" + `DISABLED` + "`" + `. <a id="nestedobjatt--configs--faz_vm"></a> The ` + "`" + `configs.faz_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "adom_num",
					Description: `(Number) Number of ADOMs. A number between 0 and 1200 (inclusive).`,
				},
				resource.Attribute{
					Name:        "daily_storage",
					Description: `(Number) Daily Storage (GB). A number between 5 and 8300 (inclusive).`,
				},
				resource.Attribute{
					Name:        "support_service",
					Description: `(String) Support Service. ` + "`" + `"FAZFC247"` + "`" + ` (FortiCare Premium). <a id="nestedobjatt--configs--fgt_vm_bundle"></a> The ` + "`" + `configs.fgt_vm_bundle` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "cpu_size",
					Description: `(String) The number of CPUs. The value of this attribute is one of ` + "`" + `"1"` + "`" + `, ` + "`" + `"2"` + "`" + `, ` + "`" + `"4"` + "`" + `, ` + "`" + `"8"` + "`" + `, ` + "`" + `"16"` + "`" + `, ` + "`" + `"32"` + "`" + ` or ` + "`" + `"2147483647"` + "`" + ` (unlimited).`,
				},
				resource.Attribute{
					Name:        "service_pkg",
					Description: `(String) The value of this attribute is one of ` + "`" + `"FC"` + "`" + ` (FortiCare), ` + "`" + `"UTM"` + "`" + `, ` + "`" + `"ENT"` + "`" + ` (Enterprise) or ` + "`" + `"ATP"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vdom_num",
					Description: `(Number) Number of VDOMs. A number between 0 and 500 (inclusive). <a id="nestedobjatt--configs--fgt_vm_lcs"></a> The ` + "`" + `configs.fgt_vm_lcs` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "cloud_services",
					Description: `(List of String) The cloud services this FortiGate Virtual Machine supports. The combination of:`,
				},
				resource.Attribute{
					Name:        "cpu_size",
					Description: `(String) The number of CPUs. A number between 1 and 96 (inclusive).`,
				},
				resource.Attribute{
					Name:        "fortiguard_services",
					Description: `(List of String) The FortiGuard services this FortiGate Virtual Machine supports. The combination of:`,
				},
				resource.Attribute{
					Name:        "support_service",
					Description: `(String) ` + "`" + `"FC247"` + "`" + ` (FortiCare 24x7) or ` + "`" + `"ASET"` + "`" + ` (FortiCare Elite).`,
				},
				resource.Attribute{
					Name:        "vdom_num",
					Description: `(Number) Number of VDOMs. A number between 0 and 500 (inclusive). <a id="nestedobjatt--configs--fmg_vm"></a> The ` + "`" + `configs.fmg_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "adom_num",
					Description: `(Number) Number of ADOMs. A number between 1 and 100000 (inclusive).`,
				},
				resource.Attribute{
					Name:        "managed_dev",
					Description: `(Number) Number of managed devices. A number between 1 and 100000 (inclusive). <a id="nestedobjatt--configs--fpc_vm"></a> The ` + "`" + `configs.fpc_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "managed_dev",
					Description: `(Number) Number of managed devices. A number between 0 and 100000 (inclusive). <a id="nestedobjatt--configs--fwb_vm"></a> The ` + "`" + `configs.fwb_vm` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "cpu_size",
					Description: `(String) Number of CPUs. The value of this attribute is one of ` + "`" + `"1"` + "`" + `, ` + "`" + `"2"` + "`" + `, ` + "`" + `"4"` + "`" + `, ` + "`" + `"8"` + "`" + ` or ` + "`" + `"16"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_pkg",
					Description: `(String) Service Package. ` + "`" + `"FWBSTD"` + "`" + ` (Standard) or ` + "`" + `"FWBADV"` + "`" + ` (Advanced).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiflexvm_groups_list",
			Category:         "Data Sources",
			ShortDescription: `Get list of FlexVM groups (asset folders).`,
			Description: `
Get list of FlexVM groups (asset folders).

Returns list of FlexVM groups (asset folders that have FlexVM products in them).


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "groups",
					Description: `(List of Object) List of groups. The structure of [` + "`" + `groups` + "`" + ` block](#nestedatt--groups) is documented below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource. <a id="nestedatt--groups"></a> The ` + "`" + `groups` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "available_tokens",
					Description: `(Number) The number of available tokens in the group.`,
				},
				resource.Attribute{
					Name:        "folder_path",
					Description: `(String) The folder path of the group.`,
				},
				resource.Attribute{
					Name:        "used_tokens",
					Description: `(Number) The number of tokens used in the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "groups",
					Description: `(List of Object) List of groups. The structure of [` + "`" + `groups` + "`" + ` block](#nestedatt--groups) is documented below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource. <a id="nestedatt--groups"></a> The ` + "`" + `groups` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "available_tokens",
					Description: `(Number) The number of available tokens in the group.`,
				},
				resource.Attribute{
					Name:        "folder_path",
					Description: `(String) The folder path of the group.`,
				},
				resource.Attribute{
					Name:        "used_tokens",
					Description: `(Number) The number of tokens used in the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiflexvm_groups_nexttoken",
			Category:         "Data Sources",
			ShortDescription: `Get next available (unused) token.`,
			Description: `
Get next available (unused) token.

Returns first available token by asset folder or Configuration id (or both can be specified in the request).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `(Number) The ID of the Flex VM Configuration this VM used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The description of the VM.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(String) VM end date.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `(String) The unique serial number of the VM.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(String) VM creation date.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) VM status. Possible values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `STOPPED` + "`" + ` or ` + "`" + `EXPIRED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(String) VM token.`,
				},
				resource.Attribute{
					Name:        "token_status",
					Description: `(String) The status of the VM token. Possible values: ` + "`" + `NOTUSED` + "`" + ` or ` + "`" + `USED` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `(Number) The ID of the Flex VM Configuration this VM used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The description of the VM.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(String) VM end date.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `(String) The unique serial number of the VM.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(String) VM creation date.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) VM status. Possible values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `STOPPED` + "`" + ` or ` + "`" + `EXPIRED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(String) VM token.`,
				},
				resource.Attribute{
					Name:        "token_status",
					Description: `(String) The status of the VM token. Possible values: ` + "`" + `NOTUSED` + "`" + ` or ` + "`" + `USED` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiflexvm_programs_list",
			Category:         "Data Sources",
			ShortDescription: `Get list of Flex VM Programs for the account.`,
			Description: `
Get list of Flex VM Programs for the account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "programs",
					Description: `(List of Object) List of Flex VM Programs for the account. The structure of [` + "`" + `programs` + "`" + ` block](#nestedatt--programs) is documented below. <a id="nestedatt--programs"></a> The ` + "`" + `programs` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Number) Your account ID.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(String) Flex VM Program end date.`,
				},
				resource.Attribute{
					Name:        "has_support_coverage",
					Description: `(Boolean) <!-- Whether the current date is between the start_date and end_date. -->`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `(String) The unique serial number of this Flex VM Program.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(String) Flex VM Program creation date.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "programs",
					Description: `(List of Object) List of Flex VM Programs for the account. The structure of [` + "`" + `programs` + "`" + ` block](#nestedatt--programs) is documented below. <a id="nestedatt--programs"></a> The ` + "`" + `programs` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Number) Your account ID.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(String) Flex VM Program end date.`,
				},
				resource.Attribute{
					Name:        "has_support_coverage",
					Description: `(Boolean) <!-- Whether the current date is between the start_date and end_date. -->`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `(String) The unique serial number of this Flex VM Program.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(String) Flex VM Program creation date.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiflexvm_vms_list",
			Category:         "Data Sources",
			ShortDescription: `Get list of existing VMs for a Flex VM Configuration.`,
			Description: `
Get list of existing VMs for a Flex VM Configuration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required/Number) The ID of a Flex VM Configuration. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `(List of Object) List of existing VMs using the specified Flex VM Configuration. The structure of [` + "`" + `vms` + "`" + ` block](#nestedatt--vms) is documented below. <a id="nestedatt--vms"></a> The ` + "`" + `vms` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `(Number) The ID of the Flex VM Configuration this VM used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The description of VM.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(String) VM end date.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `(String) The unique serial number of the VM.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(String) VM creation date.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) VM status. Possible values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `STOPPED` + "`" + ` or ` + "`" + `EXPIRED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(String) VM token.`,
				},
				resource.Attribute{
					Name:        "token_status",
					Description: `(String) The status of the VM token. Possible values: ` + "`" + `NOTUSED` + "`" + ` or ` + "`" + `USED` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `(List of Object) List of existing VMs using the specified Flex VM Configuration. The structure of [` + "`" + `vms` + "`" + ` block](#nestedatt--vms) is documented below. <a id="nestedatt--vms"></a> The ` + "`" + `vms` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `(Number) The ID of the Flex VM Configuration this VM used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The description of VM.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(String) VM end date.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `(String) The unique serial number of the VM.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(String) VM creation date.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) VM status. Possible values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `STOPPED` + "`" + ` or ` + "`" + `EXPIRED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(String) VM token.`,
				},
				resource.Attribute{
					Name:        "token_status",
					Description: `(String) The status of the VM token. Possible values: ` + "`" + `NOTUSED` + "`" + ` or ` + "`" + `USED` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiflexvm_vms_points",
			Category:         "Data Sources",
			ShortDescription: `Get point usage for VMs.`,
			Description: `
Get point usage for VMs.

Returns total points consumed by one or more virtual machines in a date range.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required/Number) The ID of a Flex VM Configuration.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Required/String) Specify an end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: ` + "`" + `YYYY-MM-DD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Required/String) Specify a start date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: ` + "`" + `YYYY-MM-DD` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `(List of Object) List of existing VMs using the specified Flex VM Configuration. The structure of [` + "`" + `vms` + "`" + ` block](#nestedatt--vms) is documented below. <a id="nestedatt--vms"></a> The ` + "`" + `vms` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "points",
					Description: `(Number) The points consumed by this VM in the date range.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `(String) The unique serial number of the VM.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) An ID for the resource.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `(List of Object) List of existing VMs using the specified Flex VM Configuration. The structure of [` + "`" + `vms` + "`" + ` block](#nestedatt--vms) is documented below. <a id="nestedatt--vms"></a> The ` + "`" + `vms` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "points",
					Description: `(Number) The points consumed by this VM in the date range.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `(String) The unique serial number of the VM.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"fortiflexvm_configs_list":     0,
		"fortiflexvm_groups_list":      1,
		"fortiflexvm_groups_nexttoken": 2,
		"fortiflexvm_programs_list":    3,
		"fortiflexvm_vms_list":         4,
		"fortiflexvm_vms_points":       5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
