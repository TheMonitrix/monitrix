package signoz

import (
	"github.com/ezeslucky/monitrix/pkg/modules/apdex"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/apdex/implapdex"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/dashboard"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/dashboard/impldashboard"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/organization"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/organization/implorganization"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/preference"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/preference/implpreference"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/savedview"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/savedview/implsavedview"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/user"
)

type Handlers struct {
	Organization organization.Handler
	Preference   preference.Handler
	User         user.Handler
	SavedView    savedview.Handler
	Apdex        apdex.Handler
	Dashboard    dashboard.Handler
}

func NewHandlers(modules Modules, user user.Handler) Handlers {
	return Handlers{
		Organization: implorganization.NewHandler(modules.Organization),
		Preference:   implpreference.NewHandler(modules.Preference),
		User:         user,
		SavedView:    implsavedview.NewHandler(modules.SavedView),
		Apdex:        implapdex.NewHandler(modules.Apdex),
		Dashboard:    impldashboard.NewHandler(modules.Dashboard),
	}
}
