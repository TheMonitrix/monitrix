package quickfilter

import (
	"context"
	v3 "github.com/ezeslucky/monitrix/pkg/query-service/model/v3"
	"github.com/ezeslucky/monitrixitrixitrixitrixitrixitrix/pkg/types/quickfiltertypes"
	"github.com/ezeslucky/monitrixitrixitrixitrixitrixitrix/pkg/valuer"
)

type Usecase interface {
	GetQuickFilters(ctx context.Context, orgID valuer.UUID) ([]*quickfiltertypes.SignalFilters, error)
	UpdateQuickFilters(ctx context.Context, orgID valuer.UUID, signal quickfiltertypes.Signal, filters []v3.AttributeKey) error
	GetSignalFilters(ctx context.Context, orgID valuer.UUID, signal quickfiltertypes.Signal) (*quickfiltertypes.SignalFilters, error)
	SetDefaultConfig(ctx context.Context, orgID valuer.UUID) error
}
