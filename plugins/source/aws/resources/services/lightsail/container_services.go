// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ContainerServices() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_container_services",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerService.html`,
		Resolver:    fetchLightsailContainerServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "container_service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContainerServiceName"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "current_deployment",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CurrentDeployment"),
			},
			{
				Name:     "is_disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsDisabled"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "next_deployment",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NextDeployment"),
			},
			{
				Name:     "power",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Power"),
			},
			{
				Name:     "power_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PowerId"),
			},
			{
				Name:     "principal_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrincipalArn"),
			},
			{
				Name:     "private_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateDomainName"),
			},
			{
				Name:     "private_registry_access",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateRegistryAccess"),
			},
			{
				Name:     "public_domain_names",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PublicDomainNames"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "scale",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Scale"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_detail",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StateDetail"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Url"),
			},
		},

		Relations: []*schema.Table{
			ContainerServiceDeployments(),
			ContainerServiceImages(),
		},
	}
}
