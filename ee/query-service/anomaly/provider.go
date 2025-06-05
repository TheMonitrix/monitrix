package anomaly

import (
	"context"

	"github.com/ezeslucky/monitrix/pkg/valuer"
)

type Provider interface {
	GetAnomalies(ctx context.Context, orgID valuer.UUID, req *GetAnomaliesRequest) (*GetAnomaliesResponse, error)
}
