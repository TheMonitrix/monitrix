package api

import (
	"net/http"

	baseapp "github.com/ezeslucky/monitrix/pkg/query-service/app"
	basemodel "github.com/ezeslucky/monitrix/pkg/query-service/model"
)

func RespondError(w http.ResponseWriter, apiErr basemodel.BaseApiError, data interface{}) {
	baseapp.RespondError(w, apiErr, data)
}
