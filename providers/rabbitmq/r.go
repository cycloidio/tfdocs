package rabbitmq

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rabbitmq_binding",
			Category:         "Resources",
			ShortDescription: `Creates and manages a binding on a RabbitMQ server.`,
			Description:      ``,
			Keywords: []string{
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source exchange.`,
				},
				resource.Attribute{
					Name:        "vhost",
					Description: `(Required) The vhost to create the resource in.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The destination queue or exchange.`,
				},
				resource.Attribute{
					Name:        "destination_type",
					Description: `(Required) The type of destination (queue or exchange).`,
				},
				resource.Attribute{
					Name:        "routing_key",
					Description: `(Optional) A routing key for the binding.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `(Optional) Additional key/value arguments for the binding. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "properties_key",
					Description: `A unique key to refer to the binding. ## Import Bindings can be imported using the ` + "`" + `id` + "`" + ` which is composed of ` + "`" + `vhost/source/destination/destination_type/properties_key` + "`" + `. E.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rabbitmq_binding.test test/test/test/queue/%23 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "properties_key",
					Description: `A unique key to refer to the binding. ## Import Bindings can be imported using the ` + "`" + `id` + "`" + ` which is composed of ` + "`" + `vhost/source/destination/destination_type/properties_key` + "`" + `. E.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rabbitmq_binding.test test/test/test/queue/%23 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rabbitmq_exchange",
			Category:         "Resources",
			ShortDescription: `Creates and manages an exchange on a RabbitMQ server.`,
			Description:      ``,
			Keywords: []string{
				"exchange",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the exchange.`,
				},
				resource.Attribute{
					Name:        "vhost",
					Description: `(Required) The vhost to create the resource in.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Required) The settings of the exchange. The structure is described below. The ` + "`" + `settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of exchange.`,
				},
				resource.Attribute{
					Name:        "durable",
					Description: `(Optional) Whether the exchange survives server restarts. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional) Whether the exchange will self-delete when all queues have finished using it.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `(Optional) Additional key/value settings for the exchange. ## Attributes Reference No further attributes are exported. ## Import Exchanges can be imported using the ` + "`" + `id` + "`" + ` which is composed of ` + "`" + `name@vhost` + "`" + `. E.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import rabbitmq_exchange.test test@vhost ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rabbitmq_permissions",
			Category:         "Resources",
			ShortDescription: `Creates and manages a user's permissions on a RabbitMQ server.`,
			Description:      ``,
			Keywords: []string{
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `(Required) The user to apply the permissions to.`,
				},
				resource.Attribute{
					Name:        "vhost",
					Description: `(Required) The vhost to create the resource in.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) The settings of the permissions. The structure is described below. The ` + "`" + `permissions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "configure",
					Description: `(Required) The "configure" ACL.`,
				},
				resource.Attribute{
					Name:        "write",
					Description: `(Required) The "write" ACL.`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Required) The "read" ACL. ## Attributes Reference No further attributes are exported. ## Import Permissions can be imported using the ` + "`" + `id` + "`" + ` which is composed of ` + "`" + `user@vhost` + "`" + `. E.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import rabbitmq_permissions.test user@vhost ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rabbitmq_policy",
			Category:         "Resources",
			ShortDescription: `Creates and manages a policy on a RabbitMQ server.`,
			Description:      ``,
			Keywords: []string{
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy.`,
				},
				resource.Attribute{
					Name:        "vhost",
					Description: `(Required) The vhost to create the resource in.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The settings of the policy. The structure is described below. The ` + "`" + `policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) A pattern to match an exchange or queue name.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) The policy with the greater priority is applied first.`,
				},
				resource.Attribute{
					Name:        "apply_to",
					Description: `(Required) Can either be "exchanges", "queues", or "all".`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) Key/value pairs of the policy definition. See the RabbitMQ documentation for definition references and examples. ## Attributes Reference No further attributes are exported. ## Import Policies can be imported using the ` + "`" + `id` + "`" + ` which is composed of ` + "`" + `name@vhost` + "`" + `. E.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import rabbitmq_policy.test name@vhost ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rabbitmq_queue",
			Category:         "Resources",
			ShortDescription: `Creates and manages a queue on a RabbitMQ server.`,
			Description:      ``,
			Keywords: []string{
				"queue",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the queue.`,
				},
				resource.Attribute{
					Name:        "vhost",
					Description: `(Required) The vhost to create the resource in.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Required) The settings of the queue. The structure is described below. The ` + "`" + `settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "durable",
					Description: `(Optional) Whether the queue survives server restarts. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional) Whether the queue will self-delete when all consumers have unsubscribed.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `(Optional) Additional key/value settings for the queue. All values will be sent to RabbitMQ as a string. If you require non-string values, use ` + "`" + `arguments_json` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "arguments_json",
					Description: `(Optional) A nested JSON string which contains additional settings for the queue. This is useful for when the arguments contain non-string values. ## Attributes Reference No further attributes are exported. ## Import Queues can be imported using the ` + "`" + `id` + "`" + ` which is composed of ` + "`" + `name@vhost` + "`" + `. E.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import rabbitmq_queue.test name@vhost ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rabbitmq_user",
			Category:         "Resources",
			ShortDescription: `Creates and manages a user on a RabbitMQ server.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the user. The value of this argument is plain-text so make sure to secure where this is defined.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Which permission model to apply to the user. Valid options are: management, policymaker, monitoring, and administrator. ## Attributes Reference No further attributes are exported. ## Import Users can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import rabbitmq_user.test mctest ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rabbitmq_vhost",
			Category:         "Resources",
			ShortDescription: `Creates and manages a vhost on a RabbitMQ server.`,
			Description:      ``,
			Keywords: []string{
				"vhost",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the vhost. ## Attributes Reference No further attributes are exported. ## Import Vhosts can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import rabbitmq_vhost.my_vhost my_vhost ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"rabbitmq_binding":     0,
		"rabbitmq_exchange":    1,
		"rabbitmq_permissions": 2,
		"rabbitmq_policy":      3,
		"rabbitmq_queue":       4,
		"rabbitmq_user":        5,
		"rabbitmq_vhost":       6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
