// Code generated by codegen; DO NOT EDIT.

package docdb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EventSubscriptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_event_subscriptions",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_EventSubscription.html`,
		Resolver:    fetchDocdbEventSubscriptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "cust_subscription_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustSubscriptionId"),
			},
			{
				Name:     "customer_aws_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerAwsId"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "event_categories_list",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EventCategoriesList"),
			},
			{
				Name:     "event_subscription_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventSubscriptionArn"),
			},
			{
				Name:     "sns_topic_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnsTopicArn"),
			},
			{
				Name:     "source_ids_list",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SourceIdsList"),
			},
			{
				Name:     "source_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceType"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subscription_creation_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionCreationTime"),
			},
		},
	}
}
