package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/logging"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	ServicesManager ServicesManager
	logger          zerolog.Logger
	// this is set by table clientList
	AccountID            string
	Region               string
	AutoscalingNamespace string
	WAFScope             wafv2types.Scope
	Partition            string
	LanguageCode         string
	stateClient          state.Client
	specificRegions      bool
	Spec                 *spec.Spec
}

type AwsLogger struct {
	l zerolog.Logger
}

type AssumeRoleAPIClient interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

const (
	defaultRegion              = "us-east-1"
	defaultVar                 = "default"
	awsCloudfrontScopeRegion   = defaultRegion
	awsCnCloudfrontScopeRegion = "cn-north-1"
)

var errInvalidRegion = errors.New("region wildcard \"*\" is only supported as first argument")
var errUnknownRegion = func(region string) error {
	return fmt.Errorf("unknown region: %q", region)
}
var errRetrievingCredentials = errors.New("error retrieving AWS credentials (see logs for details). Please verify your credentials and try again")

var ErrPaidAPIsNotEnabled = errors.New("not fetching resource because `use_paid_apis` is set to false")

func NewAwsClient(logger zerolog.Logger, s *spec.Spec) Client {
	return Client{
		ServicesManager: make(ServicesManager),
		logger:          logger,
		Spec:            s,
		stateClient:     new(state.NoOpClient),
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	idStrings := []string{
		c.AccountID,
		c.Region,
		c.AutoscalingNamespace,
		string(c.WAFScope),
		c.LanguageCode,
	}
	return strings.TrimRight(strings.Join(idStrings, ":"), ":")
}

func (c *Client) Services(serviceNames ...AWSServiceName) *Services {
	svc := c.ServicesManager.ServicesByPartitionAccount(c.Partition, c.AccountID)
	for _, service := range serviceNames {
		svc.InitService(service)
	}
	return svc
}

func (c *Client) StateClient() state.Client {
	return c.stateClient
}

// SetStateClient will set state.Client value (or state.NopClient if the param is nil) to the current Client state backend.
func (c *Client) SetStateClient(client state.Client) {
	if client == nil {
		client = new(state.NoOpClient)
	}
	c.stateClient = client
}

func (c *Client) Duplicate() *Client {
	duplicateClient := *c
	return &duplicateClient
}

func (c *Client) withPartitionAccountIDAndRegion(partition, accountID, region string) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: c.AutoscalingNamespace,
		WAFScope:             c.WAFScope,
		stateClient:          c.stateClient,
		Spec:                 c.Spec,
	}
}

func (c *Client) withPartitionAccountIDRegionAndNamespace(partition, accountID, region, namespace string) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Str("autoscaling_namespace", namespace).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: namespace,
		WAFScope:             c.WAFScope,
		stateClient:          c.stateClient,
		Spec:                 c.Spec,
	}
}

func (c *Client) withPartitionAccountIDRegionAndScope(partition, accountID, region string, scope wafv2types.Scope) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Str("waf_scope", string(scope)).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: c.AutoscalingNamespace,
		WAFScope:             scope,
		stateClient:          c.stateClient,
		Spec:                 c.Spec,
	}
}

func (c *Client) withLanguageCode(code string) *Client {
	newC := *c
	newC.LanguageCode = code
	newC.logger = newC.logger.With().Str("language_code", code).Logger()
	return &newC
}

// Configure is the entrypoint into configuring the AWS plugin. It is called by the plugin initialization in resources/plugin/aws.go
func Configure(ctx context.Context, logger zerolog.Logger, s spec.Spec) (schema.ClientMeta, error) {
	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("spec validation failed: %w", err)
	}
	s.SetDefaults()

	client := NewAwsClient(logger, &s)

	var adminAccountSts AssumeRoleAPIClient

	if client.Spec.Organization != nil {
		var err error
		client.Spec.Accounts, adminAccountSts, err = loadOrgAccounts(ctx, logger, client.Spec)
		if err != nil {
			logger.Error().Err(err).Msg("error getting child accounts")
			return nil, err
		}
	}
	if len(client.Spec.Accounts) == 0 {
		client.Spec.Accounts = []spec.Account{{ID: defaultVar}}
	}

	errorGroup, gtx := errgroup.WithContext(ctx)
	errorGroup.SetLimit(client.Spec.InitializationConcurrency)
	for _, account := range client.Spec.Accounts {
		account := account
		errorGroup.Go(func() error {
			svcsDetail, err := client.setupAWSAccount(gtx, logger, client.Spec, adminAccountSts, account)
			if err != nil {
				return err
			}
			if svcsDetail == nil {
				return nil
			}
			client.ServicesManager.InitServices(*svcsDetail)

			return nil
		})
	}
	if err := errorGroup.Wait(); err != nil {
		return nil, err
	}

	if len(client.ServicesManager) == 0 {
		// This is a special error case where we found active accounts, but just weren't able to assume a role in any of them
		if client.Spec.Organization != nil && len(client.Spec.Accounts) > 0 && client.Spec.Organization.MemberCredentials == nil {
			return nil, fmt.Errorf("discovered %d accounts in the AWS Organization, but the credentials specified in 'admin_account' were unable to assume a role in the member accounts. Verify that the role you are trying to assume (arn:aws:iam::<account_id>:role/%s) exists. If you need to use a different set of credentials to do the role assumption use 'member_trusted_principal'", len(client.Spec.Accounts), client.Spec.Organization.ChildAccountRoleName)
		}
		return nil, fmt.Errorf("no AWS accounts were successfully configured. See warning messages in the logs for details")
	}
	return &client, nil
}

func (a AwsLogger) Logf(classification logging.Classification, format string, v ...any) {
	if classification == logging.Warn {
		a.l.Warn().Msg(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug().Msg(fmt.Sprintf(format, v...))
	}
}
