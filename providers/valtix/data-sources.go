package valtix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "valtix_address_object",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Address Object resource`,
				},
				resource.Attribute{
					Name:        "address_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Address Object resource`,
				},
				resource.Attribute{
					Name:        "address_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_alert_profile",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Alert Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Alert Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Alert Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_alert_rule",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Alert Rule resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Alert Rule resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Alert Rule resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_certificate",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Certificate Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Certificate Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Certificate Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_cloud_account",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Cloud Account resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Cloud Account resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Cloud Account resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_gateway",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Gateway resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Gateway resource`,
				},
				resource.Attribute{
					Name:        "gateway_gwlb_endpoints",
					Description: `(AWS only) AWS Gateway Load Balancer endpoints created in each of the AZs displayed in the format as follows: ` + "`" + `` + "`" + `` + "`" + `hcl gateway_gwlb_endpoints { endpoint_id = "vpce-047c749fc6f7e0c0d" network_interface_id = "eni-017eacdb23d2ebaf4" subnet_id = "subnet-0d61750e97caafd9d" } gateway_gwlb_endpoints { endpoint_id = "vpce-0707fa3f03c5064a7" network_interface_id = "eni-020464bd838461bca" subnet_id = "subnet-0fd61e07f200224f1" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gwlb_service_name",
					Description: `(AWS only) VPC Endpoint Service Name associated with the AWS Gateway Load Balancer. This name can be used by the AWS Terraform Provider for establishing a GWLB Endpoint connection.`,
				},
				resource.Attribute{
					Name:        "gwlb_service_id",
					Description: `(AWS only) VPC Endpoint Service ID associated with the AWS Gateway Load Balancer. This ID can be used by the AWS Terraform Provider for accepting GWLB Endpoint connections and assigning Principals.`,
				},
				resource.Attribute{
					Name:        "gateway_endpoint",
					Description: `For Gateways of ` + "`" + `security_type = INGRESS` + "`" + `, this represents the NLB endpoint (FQDN, IP Address) to be used as the target for the client communicating with any application protected by the Valtix Ingress Gateway. This information is most often used in a DNS A record (IP Address) or CNAME record (FQDN) to resolve the application FQDN to the Valtix Ingress Gateway endpoint. Valtix will receive traffic on this endpoint and proxy the traffic to the appropriate backend application based on the configured policy. For the Ingress Gateway, this attribute is populated for Gateways deployed in all CSPs (AWS, Azure, GCP, OCI). For Gateways of ` + "`" + `security_type = EGRESS` + "`" + `, this represents the NLB endpoint (IP Address) to be used as a target for routing traffic from the Spoke VPC/VNet/VCN to the Valtix Egress / East-West Gateway. Valtix will receive traffic from clients, and forward or proxy the traffic to the appropriate destination based on the configured policy. For the Egress / East-West Gateway, this attribute is only populated for non-AWS Gateways (Azure, GCP, OCI). For the AWS Gateways, traffic is routed to the AWS Transit Gateway (TGW) or Gateway Load Balancer (GWLB) endpoints.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Gateway resource`,
				},
				resource.Attribute{
					Name:        "gateway_gwlb_endpoints",
					Description: `(AWS only) AWS Gateway Load Balancer endpoints created in each of the AZs displayed in the format as follows: ` + "`" + `` + "`" + `` + "`" + `hcl gateway_gwlb_endpoints { endpoint_id = "vpce-047c749fc6f7e0c0d" network_interface_id = "eni-017eacdb23d2ebaf4" subnet_id = "subnet-0d61750e97caafd9d" } gateway_gwlb_endpoints { endpoint_id = "vpce-0707fa3f03c5064a7" network_interface_id = "eni-020464bd838461bca" subnet_id = "subnet-0fd61e07f200224f1" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gwlb_service_name",
					Description: `(AWS only) VPC Endpoint Service Name associated with the AWS Gateway Load Balancer. This name can be used by the AWS Terraform Provider for establishing a GWLB Endpoint connection.`,
				},
				resource.Attribute{
					Name:        "gwlb_service_id",
					Description: `(AWS only) VPC Endpoint Service ID associated with the AWS Gateway Load Balancer. This ID can be used by the AWS Terraform Provider for accepting GWLB Endpoint connections and assigning Principals.`,
				},
				resource.Attribute{
					Name:        "gateway_endpoint",
					Description: `For Gateways of ` + "`" + `security_type = INGRESS` + "`" + `, this represents the NLB endpoint (FQDN, IP Address) to be used as the target for the client communicating with any application protected by the Valtix Ingress Gateway. This information is most often used in a DNS A record (IP Address) or CNAME record (FQDN) to resolve the application FQDN to the Valtix Ingress Gateway endpoint. Valtix will receive traffic on this endpoint and proxy the traffic to the appropriate backend application based on the configured policy. For the Ingress Gateway, this attribute is populated for Gateways deployed in all CSPs (AWS, Azure, GCP, OCI). For Gateways of ` + "`" + `security_type = EGRESS` + "`" + `, this represents the NLB endpoint (IP Address) to be used as a target for routing traffic from the Spoke VPC/VNet/VCN to the Valtix Egress / East-West Gateway. Valtix will receive traffic from clients, and forward or proxy the traffic to the appropriate destination based on the configured policy. For the Egress / East-West Gateway, this attribute is only populated for non-AWS Gateways (Azure, GCP, OCI). For the AWS Gateways, traffic is routed to the AWS Transit Gateway (TGW) or Gateway Load Balancer (GWLB) endpoints.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_policy_rule_set",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Policy Ruleset resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy Ruleset resource`,
				},
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy Ruleset resource`,
				},
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_anti_malware",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Anti-Malware Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Network Anti-Malware Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Network Anti-Malware Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_application_threat",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Application Threat (WAF) Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Application Threat (WAF) Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Application Threat (WAF) Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_decryption",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Decryption Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Decryption Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Decryption Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_diagnostic",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Diagnostic Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Diagnostic Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Diagnostic Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_dlp",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Data Loss Prevention (DLP) IP Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Data Loss Prevention (DLP) Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Data Loss Prevention (DLP) Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_fqdn",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the FQDN Filtering Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the FQDN Filtering Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the FQDN Filtering Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_l7dos",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the L7DOS Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Network L7DOS Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Network L7DOS Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_log_forwarding",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Log Forwarding Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Log Forwarding Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Log Forwarding Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_malicious_ip",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Malicious IP Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Malicious IP Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Malicious IP Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_network_intrusion",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Network Intrusion (IDS/IPS) IP Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Network Intrusion (IDS/IPS) Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Network Intrusion (IDS/IPS) Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_ntp",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the NTP Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NTP Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NTP Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_packet_capture",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Packet Capture Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Packet Capture Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Packet Capture Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_urlfilter",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the URL Filtering Profile resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the URL Filtering Profile resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the URL Filtering Profile resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_service_object",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Service Object resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Service Object resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Service Object resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_service_vpc",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) Name of the onboarded CSP Account where the Service VPC/VNet resource exists`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the Service VPC/VNet resource exists`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC/VNet ID of the Service VPC/VNet resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Service VPC/VNet resource`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Service VPC/VNet resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Service VPC/VNet resource`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Service VPC/VNet resource`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"valtix_address_object":             0,
		"valtix_alert_profile":              1,
		"valtix_alert_rule":                 2,
		"valtix_certificate":                3,
		"valtix_cloud_account":              4,
		"valtix_gateway":                    5,
		"valtix_policy_rule_set":            6,
		"valtix_profile_anti_malware":       7,
		"valtix_profile_application_threat": 8,
		"valtix_profile_decryption":         9,
		"valtix_profile_diagnostic":         10,
		"valtix_profile_dlp":                11,
		"valtix_profile_fqdn":               12,
		"valtix_profile_l7dos":              13,
		"valtix_profile_log_forwarding":     14,
		"valtix_profile_malicious_ip":       15,
		"valtix_profile_network_intrusion":  16,
		"valtix_profile_ntp":                17,
		"valtix_profile_packet_capture":     18,
		"valtix_profile_urlfilter":          19,
		"valtix_service_object":             20,
		"valtix_service_vpc":                21,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
