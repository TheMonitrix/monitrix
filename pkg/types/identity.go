package types

import (
	"github.com/ezeslucky/monitrix/pkg/valuer"
)

type Identifiable struct {
	ID valuer.UUID `json:"id" bun:"id,pk,type:text"`
}
