package quickfilter

import (
	"encoding/json"
	"github.com/ezeslucky/monitrix/pkg/http/render"
	"github.com/ezeslucky/monitrixitrixitrixitrixitrixitrix/pkg/types/authtypes"
	"github.com/ezeslucky/monitrixitrixitrixitrixitrixitrix/pkg/types/quickfiltertypes"
	"github.com/ezeslucky/monitrixitrixitrixitrixitrixitrix/pkg/valuer"
	"github.com/gorilla/mux"
	"net/http"
)

type API interface {
	GetQuickFilters(http.ResponseWriter, *http.Request)
	UpdateQuickFilters(http.ResponseWriter, *http.Request)
	GetSignalFilters(http.ResponseWriter, *http.Request)
}

type quickFiltersAPI struct {
	usecase Usecase
}

func NewAPI(usecase Usecase) API {
	return &quickFiltersAPI{usecase: usecase}
}

func (q *quickFiltersAPI) GetQuickFilters(rw http.ResponseWriter, r *http.Request) {
	claims, err := authtypes.ClaimsFromContext(r.Context())
	if err != nil {
		render.Error(rw, err)
		return
	}

	filters, err := q.usecase.GetQuickFilters(r.Context(), valuer.MustNewUUID(claims.OrgID))
	if err != nil {
		render.Error(rw, err)
		return
	}

	render.Success(rw, http.StatusOK, filters)
}

func (q *quickFiltersAPI) UpdateQuickFilters(rw http.ResponseWriter, r *http.Request) {
	claims, err := authtypes.ClaimsFromContext(r.Context())
	if err != nil {
		render.Error(rw, err)
		return
	}

	var req quickfiltertypes.UpdatableQuickFilters
	decodeErr := json.NewDecoder(r.Body).Decode(&req)
	if decodeErr != nil {
		render.Error(rw, decodeErr)
		return
	}

	err = q.usecase.UpdateQuickFilters(r.Context(), valuer.MustNewUUID(claims.OrgID), req.Signal, req.Filters)
	if err != nil {
		render.Error(rw, err)
		return
	}

	render.Success(rw, http.StatusNoContent, nil)
}

func (q *quickFiltersAPI) GetSignalFilters(rw http.ResponseWriter, r *http.Request) {
	claims, err := authtypes.ClaimsFromContext(r.Context())
	if err != nil {
		render.Error(rw, err)
		return
	}

	signal := mux.Vars(r)["signal"]
	validatedSignal, err := quickfiltertypes.NewSignal(signal)
	if err != nil {
		render.Error(rw, err)
		return
	}

	filters, err := q.usecase.GetSignalFilters(r.Context(), valuer.MustNewUUID(claims.OrgID), validatedSignal)
	if err != nil {
		render.Error(rw, err)
		return
	}

	render.Success(rw, http.StatusOK, filters)
}
