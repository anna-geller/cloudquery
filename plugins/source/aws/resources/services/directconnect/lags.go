// Code generated by codegen; DO NOT EDIT.

package directconnect

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Lags() *schema.Table {
	return &schema.Table{
		Name:        "aws_directconnect_lags",
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Lag.html`,
		Resolver:    fetchDirectconnectLags,
		Multiplex:   client.ServiceAccountRegionMultiplexer("directconnect"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveLagARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LagId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "allows_hosted_connections",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllowsHostedConnections"),
			},
			{
				Name:     "aws_device",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsDevice"),
			},
			{
				Name:     "aws_device_v2",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsDeviceV2"),
			},
			{
				Name:     "aws_logical_device_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsLogicalDeviceId"),
			},
			{
				Name:     "connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Connections"),
			},
			{
				Name:     "connections_bandwidth",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionsBandwidth"),
			},
			{
				Name:     "encryption_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionMode"),
			},
			{
				Name:     "has_logical_redundancy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HasLogicalRedundancy"),
			},
			{
				Name:     "jumbo_frame_capable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("JumboFrameCapable"),
			},
			{
				Name:     "lag_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LagName"),
			},
			{
				Name:     "lag_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LagState"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "mac_sec_capable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MacSecCapable"),
			},
			{
				Name:     "mac_sec_keys",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MacSecKeys"),
			},
			{
				Name:     "minimum_links",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinimumLinks"),
			},
			{
				Name:     "number_of_connections",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumberOfConnections"),
			},
			{
				Name:     "owner_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerAccount"),
			},
			{
				Name:     "provider_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProviderName"),
			},
		},
	}
}
