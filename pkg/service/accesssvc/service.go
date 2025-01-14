package accesssvc

import (
	"context"

	"github.com/benbjohnson/clock"

	"github.com/common-fate/ddb"
	ahTypes "github.com/common-fate/granted-approvals/accesshandler/pkg/types"
	"github.com/common-fate/granted-approvals/pkg/access"
	"github.com/common-fate/granted-approvals/pkg/cache"
	"github.com/common-fate/granted-approvals/pkg/gevent"
	"github.com/common-fate/granted-approvals/pkg/rule"
	"github.com/common-fate/granted-approvals/pkg/service/grantsvc"
	"github.com/common-fate/granted-approvals/pkg/types"
)

// Service holds business logic relating to Access Requests.
type Service struct {
	Clock       clock.Clock
	DB          ddb.Storage
	Granter     Granter
	EventPutter EventPutter
	Cache       CacheService
	AHClient    AHClient
	Rules       AccessRuleService
}

//go:generate go run github.com/golang/mock/mockgen -destination=mocks/granter.go -package=mocks . Granter

// Granter creates Grants in the Access Handler.
type Granter interface {
	CreateGrant(ctx context.Context, opts grantsvc.CreateGrantOpts) (*access.Request, error)
	RevokeGrant(ctx context.Context, opts grantsvc.RevokeGrantOpts) (*access.Request, error)
	ValidateGrant(ctx context.Context, opts grantsvc.CreateGrantOpts) error
}

//go:generate go run github.com/golang/mock/mockgen -destination=mocks/eventputter.go -package=mocks . EventPutter
type EventPutter interface {
	Put(ctx context.Context, detail gevent.EventTyper) error
}

//go:generate go run github.com/golang/mock/mockgen -destination=mocks/cache.go -package=mocks . CacheService
type CacheService interface {
	RefreshCachedProviderArgOptions(ctx context.Context, providerId string, argId string) (bool, []cache.ProviderOption, []cache.ProviderArgGroupOption, error)
	LoadCachedProviderArgOptions(ctx context.Context, providerId string, argId string) (bool, []cache.ProviderOption, []cache.ProviderArgGroupOption, error)
}

//go:generate go run github.com/golang/mock/mockgen -destination=mocks/mock_accessrule_service.go -package=mocks . AccessRuleService

// AccessRuleService can create and get rules
type AccessRuleService interface {
	RequestArguments(ctx context.Context, accessRuleTarget rule.Target) (map[string]types.RequestArgument, error)
}

type AHClient interface {
	ahTypes.ClientWithResponsesInterface
}
