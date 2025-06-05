package signoz

import (
	"github.com/ezeslucky/monitrix/pkg/modules/apdex"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/apdex/implapdex"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/dashboard"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/dashboard/impldashboard"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/organization"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/organization/implorganization"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/preference"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/preference/implpreference"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/savedview"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/savedview/implsavedview"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/user"
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
