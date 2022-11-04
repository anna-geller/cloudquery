// Code generated by codegen; DO NOT EDIT.

package athena

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkGroupPreparedStatements() *schema.Table {
	return &schema.Table{
		Name:        "aws_athena_work_group_prepared_statements",
		Description: `https://docs.aws.amazon.com/athena/latest/APIReference/API_PreparedStatement.html`,
		Resolver:    fetchAthenaWorkGroupPreparedStatements,
		Multiplex:   client.ServiceAccountRegionMultiplexer("athena"),
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
				Name:     "work_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "query_statement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QueryStatement"),
			},
			{
				Name:     "statement_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatementName"),
			},
			{
				Name:     "work_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkGroupName"),
			},
		},
	}
}
