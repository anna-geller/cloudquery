// Code generated by codegen; DO NOT EDIT.

package docdb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PendingMaintenanceActions() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_pending_maintenance_actions",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_PendingMaintenanceAction.html`,
		Resolver:    fetchDocdbPendingMaintenanceActions,
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
				Name:     "action",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Action"),
			},
			{
				Name:     "auto_applied_after_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AutoAppliedAfterDate"),
			},
			{
				Name:     "current_apply_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CurrentApplyDate"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "forced_apply_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ForcedApplyDate"),
			},
			{
				Name:     "opt_in_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OptInStatus"),
			},
		},
	}
}
