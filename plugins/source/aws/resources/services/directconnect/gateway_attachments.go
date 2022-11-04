// Code generated by codegen; DO NOT EDIT.

package directconnect

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GatewayAttachments() *schema.Table {
	return &schema.Table{
		Name:        "aws_directconnect_gateway_attachments",
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAttachment.html`,
		Resolver:    fetchDirectconnectGatewayAttachments,
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
				Name:     "gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "attachment_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AttachmentState"),
			},
			{
				Name:     "attachment_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AttachmentType"),
			},
			{
				Name:     "direct_connect_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectConnectGatewayId"),
			},
			{
				Name:     "state_change_error",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateChangeError"),
			},
			{
				Name:     "virtual_interface_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualInterfaceId"),
			},
			{
				Name:     "virtual_interface_owner_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualInterfaceOwnerAccount"),
			},
			{
				Name:     "virtual_interface_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualInterfaceRegion"),
			},
		},
	}
}
