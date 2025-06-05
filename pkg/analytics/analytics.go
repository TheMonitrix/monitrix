package analytics

import (
	"context"

	"github.com/ezeslucky/monitrix/pkg/factory"
	"github.com/ezeslucky/monitrix/pkg/types/analyticstypes"
)

type Analytics interface {
	factory.Service

	// Sends analytics messages to an analytics backend.
	Send(context.Context, ...analyticstypes.Message)
}
