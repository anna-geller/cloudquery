// Code generated by codegen; DO NOT EDIT.

package waf

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WebAcls() *schema.Table {
	return &schema.Table{
		Name:        "aws_waf_web_acls",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_WebACLSummary.html`,
		Resolver:    fetchWafWebAcls,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebACLArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafWebACLTags,
			},
			{
				Name:        "logging_configuration",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("LoggingConfiguration"),
				Description: `The LoggingConfiguration for the specified web ACL.`,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "web_acl_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebACLId"),
			},
		},
	}
}
