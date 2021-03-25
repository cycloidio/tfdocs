package linode

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "linode_account",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode account.`,
			Description: `\_account

Provides information about a Linode account.

This data source should not be used in conjuction with the ` + "`" + `LINODE_DEBUG` + "`" + ` option.  See the [debugging notes](/providers/linode/linode/latest/docs#debugging) for more details.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The email address for this Account, for account management communications, and may be used for other communications as configured.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The first name of the person associated with this Account.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The last name of the person associated with this Account.`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `The company name associated with this Account.`,
				},
				resource.Attribute{
					Name:        "address_1",
					Description: `First line of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "address_2",
					Description: `Second line of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `The phone number associated with this Account.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city for this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `If billing address is in the United States, this is the State portion of the Account's billing address. If the address is outside the US, this is the Province associated with the Account's billing address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The two-letter country code of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "zip",
					Description: `The zip code of this Account's billing address.`,
				},
				resource.Attribute{
					Name:        "balance",
					Description: `This Account's balance, in US dollars.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_domain",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode domain.`,
			Description: `\_domain

Provides information about a Linode domain.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique numeric ID of the Domain record to query.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The unique domain name of the Domain record to query. ## Attributes The Linode Domain resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this Domain.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave)`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The group this Domain belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Used to control whether this Domain is currently being rendered.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for this Domain.`,
				},
				resource.Attribute{
					Name:        "master_ips",
					Description: `The IP addresses representing the master DNS for this Domain.`,
				},
				resource.Attribute{
					Name:        "axfr_ips",
					Description: `The list of IPs that may perform a zone transfer for this Domain.`,
				},
				resource.Attribute{
					Name:        "ttl_sec",
					Description: `'Time to Live'-the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers.`,
				},
				resource.Attribute{
					Name:        "retry_sec",
					Description: `The interval, in seconds, at which a failed refresh should be retried.`,
				},
				resource.Attribute{
					Name:        "expire_sec",
					Description: `The amount of time in seconds that may pass before this Domain is no longer authoritative.`,
				},
				resource.Attribute{
					Name:        "refresh_sec",
					Description: `The amount of time in seconds before this Domain should be refreshed.`,
				},
				resource.Attribute{
					Name:        "soa_email",
					Description: `Start of Authority email address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `An array of tags applied to this object.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_domain_record",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode Domain Record.`,
			Description: `

Provides information about a Linode Domain Record.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_firewall",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Firewall.`,
			Description: `\_firewall

Provides details about an LKE Cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The Firewall's ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label for the firewall.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags applied to the firewall.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `If true, the firewall is inactive.`,
				},
				resource.Attribute{
					Name:        "inbound_policy",
					Description: `The default behavior for inbound traffic.`,
				},
				resource.Attribute{
					Name:        "outbound_policy",
					Description: `The default behavior for outbound traffic. This setting can be overridden by updating the outbound.action property for an individual Firewall Rule.`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `The IDs of Linodes to apply this firewall to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the firewall.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Used to identify this rule. For display purposes only.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Controls whether traffic is accepted or dropped by this rule. Overrides the Firewall’s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The network protocol this rule controls.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `A list of IPv4 addresses or networks. Must be in IP/mask format.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `A list of IPv6 addresses or networks. Must be in IP/mask format. ### devices The following attributes are available on devices:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Device.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID of the underlying entity this device references (i.e. the Linode's ID).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of Firewall Device.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the underlying entity this device references.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `The label for the firewall.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags applied to the firewall.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `If true, the firewall is inactive.`,
				},
				resource.Attribute{
					Name:        "inbound_policy",
					Description: `The default behavior for inbound traffic.`,
				},
				resource.Attribute{
					Name:        "outbound_policy",
					Description: `The default behavior for outbound traffic. This setting can be overridden by updating the outbound.action property for an individual Firewall Rule.`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `The IDs of Linodes to apply this firewall to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the firewall.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Used to identify this rule. For display purposes only.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Controls whether traffic is accepted or dropped by this rule. Overrides the Firewall’s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The network protocol this rule controls.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `A list of IPv4 addresses or networks. Must be in IP/mask format.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `A list of IPv6 addresses or networks. Must be in IP/mask format. ### devices The following attributes are available on devices:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Device.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID of the underlying entity this device references (i.e. the Linode's ID).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of Firewall Device.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the underlying entity this device references.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_image",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode image.`,
			Description: `\_image

Provides information about a Linode image

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique ID of this Image. The ID of private images begin with ` + "`" + `private/` + "`" + ` followed by the numeric identifier of the private image, for example ` + "`" + `private/12345` + "`" + `. ## Attributes The Linode Image resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A short description of the Image.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Image was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The name of the User who created this Image, or "linode" for official Images.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Whether or not this Image is deprecated. Will only be true for deprecated public Images.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A detailed description of this Image.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `True if the Image is public.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The minimum size this Image needs to deploy. Size is in MB. example: 2500`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `How the Image was created. Manual Images can be created at any time. image"Automatic" Images are created automatically from a deleted Linode.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The upstream distribution vendor. ` + "`" + `None` + "`" + ` for private Images.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_instance_type",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode instance type.`,
			Description: `\_instance\_type

Provides information about a Linode instance type

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Label used to identify instance type ## Attributes The Linode Instance Type resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID representing the Linode Type`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Linode Type's label is for display purposes only`,
				},
				resource.Attribute{
					Name:        "class",
					Description: `The class of the Linode Type`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The Disk size, in MB, of the Linode Type`,
				},
				resource.Attribute{
					Name:        "price.0.hourly",
					Description: `Cost (in US dollars) per hour.`,
				},
				resource.Attribute{
					Name:        "price.0.monthly",
					Description: `Cost (in US dollars) per month.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.hourly",
					Description: `The cost (in US dollars) per hour to add Backups service.`,
				},
				resource.Attribute{
					Name:        "addons.0.backups.0.price.0.monthly",
					Description: `The cost (in US dollars) per month to add Backups service.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_lke_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides details about an LKE Cluster.`,
			Description: `\_lke_cluster

Provides details about an LKE Cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The LKE Cluster's ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `The Kubernetes version for this Kubernetes cluster in the format of ` + "`" + `major.minor` + "`" + ` (e.g. ` + "`" + `1.17` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `This Kubernetes cluster's location.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags applied to the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the cluster.`,
				},
				resource.Attribute{
					Name:        "api_endpoints",
					Description: `The endpoints for the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The base64 encoded kubeconfig for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "pools",
					Description: `Node pools associated with this cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Node Pool.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The linode type for all of the nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `The nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the node.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the underlying Linode instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "k8s_version",
					Description: `The Kubernetes version for this Kubernetes cluster in the format of ` + "`" + `major.minor` + "`" + ` (e.g. ` + "`" + `1.17` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `This Kubernetes cluster's location.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags applied to the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the cluster.`,
				},
				resource.Attribute{
					Name:        "api_endpoints",
					Description: `The endpoints for the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The base64 encoded kubeconfig for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "pools",
					Description: `Node pools associated with this cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Node Pool.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The linode type for all of the nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `The nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the node.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the underlying Linode instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_networking_ip",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode Networking IP Address.`,
			Description: `\_network\_ip

Provides information about a Linode Networking IP Address

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The IP Address to access. The address must be associated with the account and a resource that the user has access to view. ## Attributes The Linode Network IP Address resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The default gateway for this address.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The mask that separates host bits from network bits for this address.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The number of bits set in the subnet mask.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of address this is (ipv4, ipv6, ipv6/pool, ipv6/range).`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether this is a public or private IP address.`,
				},
				resource.Attribute{
					Name:        "rdns",
					Description: `The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.`,
				},
				resource.Attribute{
					Name:        "linode_id",
					Description: `The ID of the Linode this address currently belongs to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Region this IP address resides in.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_object_storage_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode Object Storage Cluster.`,
			Description: `\_object\_storage\_cluster

Provides information about a Linode Object Storage Cluster

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique ID of this cluster. ## Attributes The Linode Object Storage Cluster resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The base URL for this cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `This cluster's status.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this cluster is located in.`,
				},
				resource.Attribute{
					Name:        "static_site_domain",
					Description: `The base URL for this cluster used when hosting static sites.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_profile",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode profile.`,
			Description: `\_profile

Provides information about a Linode profile.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The profile email address. This address will be used for communication with Linode as necessary.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The profile's preferred timezone. This is not used by the API, and is for the benefit of clients only. All times the API returns are in UTC.`,
				},
				resource.Attribute{
					Name:        "email_notifications",
					Description: `If true, email notifications will be sent about account activity. If false, when false business-critical communications may still be sent through email.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username for logging in to Linode services.`,
				},
				resource.Attribute{
					Name:        "ip_whitelist_enabled",
					Description: `If true, logins for the user will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled.`,
				},
				resource.Attribute{
					Name:        "lish_auth_method",
					Description: `The methods of authentication allowed when connecting via Lish. 'keys_only' is the most secure with the intent to use Lish, and 'disabled' is recommended for users that will not use Lish at all.`,
				},
				resource.Attribute{
					Name:        "authorized_keys",
					Description: `The list of SSH Keys authorized to use Lish for this user. This value is ignored if lish_auth_method is 'disabled'.`,
				},
				resource.Attribute{
					Name:        "two_factor_auth",
					Description: `If true, logins from untrusted computers will require Two Factor Authentication.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `If true, the user has restrictions on what can be accessed on the Account.`,
				},
				resource.Attribute{
					Name:        "referrals",
					Description: `Credit Card information associated with this Account.`,
				},
				resource.Attribute{
					Name:        "referrals.0.total",
					Description: `The number of users who have signed up with the referral code.`,
				},
				resource.Attribute{
					Name:        "referrals.0.credit",
					Description: `The amount of account credit in US Dollars issued to the account through the referral program.`,
				},
				resource.Attribute{
					Name:        "referrals.0.completed",
					Description: `The number of completed signups with the referral code.`,
				},
				resource.Attribute{
					Name:        "referrals.0.pending",
					Description: `The number of pending signups for the referral code. To receive credit the signups must be completed.`,
				},
				resource.Attribute{
					Name:        "referrals.0.code",
					Description: `The Profile referral code. If new accounts use this when signing up for Linode, referring account will receive credit.`,
				},
				resource.Attribute{
					Name:        "referrals.0.url",
					Description: `The referral URL.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_region",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific service region`,
			Description: `\_region

` + "`" + `linode_region` + "`" + ` provides details about a specific Linode region. See all regions [here](https://api.linode.com/v4/regions).

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_sshkey",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a profile SSH Key`,
			Description: `\_sshkey

` + "`" + `linode_sshkey` + "`" + ` provides access to a specifically labeled SSH Key in the Profile of the User identified by the access token.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_stackscript",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode StackScript.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique numeric ID of the StackScript to query. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The StackScript's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The script to execute when provisioning a new Linode with this StackScript.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the StackScript.`,
				},
				resource.Attribute{
					Name:        "rev_note",
					Description: `This field allows you to add notes for the set of revisions made to this StackScript.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `An array of Image IDs representing the Images that this StackScript is compatible for deploying with.`,
				},
				resource.Attribute{
					Name:        "deployments_active",
					Description: `Count of currently active, deployed Linodes created from this StackScript.`,
				},
				resource.Attribute{
					Name:        "user_gravatar_id",
					Description: `The Gravatar ID for the User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "deployments_total",
					Description: `The total number of times this StackScript has been deployed.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this StackScript was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date this StackScript was updated.`,
				},
				resource.Attribute{
					Name:        "user_defined_fields",
					Description: `This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A human-readable label for the field that will serve as the input prompt for entering the value during deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the field.`,
				},
				resource.Attribute{
					Name:        "example",
					Description: `An example value for the field.`,
				},
				resource.Attribute{
					Name:        "one_of",
					Description: `A list of acceptable single values for the field.`,
				},
				resource.Attribute{
					Name:        "many_of",
					Description: `A list of acceptable values for the field in any quantity, combination or order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `The default value. If not specified, this value will be used. ## Import Linodes StackScripts can be imported using the Linode StackScript ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_stackscript.mystackscript 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_user",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode user.`,
			Description: `\_user

Provides information about a Linode user

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The unique username of this User. ## Attributes The Linode User resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the authorized_users field of a create Linode, rebuild Linode, or create Disk request.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address for this User, for account management communications, and may be used for other communications as configured.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `If true, this User must be granted access to perform actions or access entities on this Account.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_volume",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a Linode Volume.`,
			Description: `

Provides information about a Linode Volume.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"linode_account":                0,
		"linode_domain":                 1,
		"linode_domain_record":          2,
		"linode_firewall":               3,
		"linode_image":                  4,
		"linode_instance_type":          5,
		"linode_lke_cluster":            6,
		"linode_networking_ip":          7,
		"linode_object_storage_cluster": 8,
		"linode_profile":                9,
		"linode_region":                 10,
		"linode_sshkey":                 11,
		"linode_stackscript":            12,
		"linode_user":                   13,
		"linode_volume":                 14,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
