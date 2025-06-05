package implorganization

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ezeslucky/monitrix/pkg/errors"
	"github.com/ezeslucky/monitrix/pkg/http/render"
	"github.com/ezeslucky/monitrix/pkg/modules/organization"
	"github.com/ezeslucky/monitrix/pkg/types"
	"github.com/ezeslucky/monitrix/pkg/types/authtypes"
	"github.com/ezeslucky/monitrix/pkg/valuer"
)

type handler struct {
	module organization.Module
}

func NewHandler(module organization.Module) organization.Handler {
	return &handler{module: module}
}

func (handler *handler) Get(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	claims, err := authtypes.ClaimsFromContext(ctx)
	if err != nil {
		render.Error(rw, err)
		return
	}

	orgID, err := valuer.NewUUID(claims.OrgID)
	if err != nil {
		render.Error(rw, errors.Newf(errors.TypeInvalidInput, errors.CodeInvalidInput, "orgId is invalid"))
		return
	}

	organization, err := handler.module.Get(ctx, orgID)
	if err != nil {
		render.Error(rw, err)
		return
	}

	render.Success(rw, http.StatusOK, organization)
}

func (handler *handler) Update(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	claims, err := authtypes.ClaimsFromContext(ctx)
	if err != nil {
		render.Error(rw, err)
		return
	}

	orgID, err := valuer.NewUUID(claims.OrgID)
	if err != nil {
		render.Error(rw, errors.Newf(errors.TypeInvalidInput, errors.CodeInvalidInput, "invalid org id"))
		return
	}

	var req *types.Organization
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Error(rw, err)
	}

	req.ID = orgID
	err = handler.module.Update(ctx, req)
	if err != nil {
		render.Error(rw, err)
		return
	}

	render.Success(rw, http.StatusNoContent, nil)
}
